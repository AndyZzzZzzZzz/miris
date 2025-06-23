package main

// Main driver for MIRIS
/*
1. Loads the query plan for disk
2. Loads the relevant dataset + model config
3. Builds the output path
4. Runs the full MIRIS pipeline via exec.Exec
*/
import (
	"github.com/mitroadmaps/miris/data"
	"github.com/mitroadmaps/miris/exec"
	"github.com/mitroadmaps/miris/miris"

	"fmt"
	"os"
)

func main() {
	predName := os.Args[1]
	planFname := os.Args[2]

	ppCfg, modelCfg := data.Get(predName)
	detectionPath, framePath := data.GetExec(predName)
	var plan miris.PlannerConfig
	miris.ReadJSON(planFname, &plan)
	execCfg := miris.ExecConfig{
		DetectionPath: detectionPath,
		FramePath: framePath,
		TrackOutput: fmt.Sprintf("logs/%s/%d/%v/track.json", predName, plan.Freq, plan.Bound),
		FilterOutput: fmt.Sprintf("logs/%s/%d/%v/filter.json", predName, plan.Freq, plan.Bound),
		UncertaintyOutput: fmt.Sprintf("logs/%s/%d/%v/uncertainty.json", predName, plan.Freq, plan.Bound),
		RefineOutput: fmt.Sprintf("logs/%s/%d/%v/refine.json", predName, plan.Freq, plan.Bound),
		OutPath: fmt.Sprintf("logs/%s/%d/%v/final.json", predName, plan.Freq, plan.Bound),
	}
	exec.Exec(ppCfg, modelCfg, plan, execCfg)
}

