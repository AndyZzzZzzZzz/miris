package main

/*
go run plan.go shibuya 16 0.9
1. query predicates: loads shibuya
2. sampling frequency: 16
3. precision/recall bound: 0.9

*/
import (
	"github.com/mitroadmaps/miris/data"
	"github.com/mitroadmaps/miris/miris"
	"github.com/mitroadmaps/miris/planner"

	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// predicate name
	predName := os.Args[1]
	// sampling frequency (1 frame out of every N)
	freq, _ := strconv.Atoi(os.Args[2])
	// target a value (desired precision/recall lower bound)
	bound, _ := strconv.ParseFloat(os.Args[3], 64)
	
	// optional: if a prev plan file is passed
	// this loads it and reuses the sample quality scores
	var existingPlan miris.PlannerConfig
	var qSamples map[int][]float64
	if len(os.Args) >= 5 {
		miris.ReadJSON(os.Args[4], &existingPlan)
		qSamples = existingPlan.QSamples
	}
	
	// loads meta data
	// ppCgf: preprocessing config (bounding boxes, track setup)
	// modelCfg: Model config
	ppCfg, modelCfg := data.Get(predName)
	
	/*
	If we didn't load an existing plan, generate new
	Quality curves: estimated tradeoffs between query accuracy and model effort, precomputed across different q values
	(thresholds for filtering/uncertainty)
	*/
	if qSamples == nil {
		qSamples = planner.GetQSamples(2*freq, ppCfg, modelCfg)
	}

	// Uses the quality samples to choose the best q value
	// Goal: maximize speed while keeping disired accuracy
	q := planner.PlanQ(qSamples, bound)
	log.Println("finished planning q", q)
	plan := miris.PlannerConfig{
		Freq: freq,
		Bound: bound,
		QSamples: qSamples,
		Q: q,
	}
	miris.WriteJSON(fmt.Sprintf("logs/%s/%d/%v/plan.json", predName, freq, bound), plan)
	filterPlan, refinePlan := planner.PlanFilterRefine(ppCfg, modelCfg, freq, bound, nil)
	plan.Filter = filterPlan
	plan.Refine = refinePlan
	log.Println(plan)
	miris.WriteJSON(fmt.Sprintf("logs/%s/%d/%v/plan.json", predName, freq, bound), plan)
}

