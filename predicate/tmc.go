package predicate

// turning movement count (TMC) queries
/*
	Represent pre-registered query configurations that MIRIS can reference by name during planning and execution
	Not full queries but define the core logic of what the query is looking for, directly used when a named query is invoked.
*/
import (
	"github.com/mitroadmaps/gomapinfer/common"
)

/*
init() function is automatically called when the package is imported. Here, it is used to register query predicates into the
global predicates map.
*/
func init() {
	// object starts in lower-left region and ends in an upper-right section
	predicates["uav"] = StartEndPredicate(
		common.Rect(362, 446, 706, 1080).ToPolygon(),
		common.Rect(784, 176, 1920, 642).ToPolygon(),
	)
	
	// left to right
	// object must pass through left region and right region
	predicates["warsawlr"] = WaypointPredicate([]common.Polygon{
		common.Rect(0, 610, 930, 1080).ToPolygon(),
		common.Rect(1190, 700, 1920, 1080).ToPolygon(),
	})

	// object must pass through both in order
	predicates["warsawtb"] = WaypointPredicate([]common.Polygon{
		{
			common.Point{978, 337},
			common.Point{1150, 680},
			common.Point{1450, 680},
			common.Point{1580, 590},
			common.Point{1107, 325},
		},
		{
			common.Point{1920, 685},
			common.Point{1645, 669},
			common.Point{1400, 780},
			common.Point{1573, 1080},
			common.Point{1920, 1080},
		},
	})

	// highway movement
	// describe distinct movement pattern
	predicates["warsawhw"] = WaypointPredicate([]common.Polygon{
		{
			common.Point{1314, 403},
			common.Point{901, 253},
			common.Point{978, 202},
			common.Point{1390, 333},
		},
		{
			common.Point{1558, 393},
			common.Point{1491, 466},
			common.Point{1920, 680},
			common.Point{1920, 550},
		},
	})
	
	// matches if any of the three movement patterns are observed
	predicates["warsaw"] = Or(predicates["warsawlr"], predicates["warsawtb"], predicates["warsawhw"])
	

	shibuyaPolys := map[string]common.Polygon{
		"right": {
			common.Point{1332, 0},
			common.Point{1332, 440},
			common.Point{1614, 550},
			common.Point{1920, 550},
			common.Point{1920, 0},
		},
		"left": {
			common.Point{0, 525},
			common.Point{500, 525},
			common.Point{800, 1080},
			common.Point{0, 1080},
		},
		"top": {
			common.Point{0, 525},
			common.Point{550, 525},
			common.Point{1200, 420},
			common.Point{1200, 0},
			common.Point{0, 0},
		},
		"bottom": {
			common.Point{1920, 630},
			common.Point{1640, 630},
			common.Point{1040, 1080},
			common.Point{1920, 1080},
		},
	}

	predicates["shibuyabt"] = StartEndPredicate(shibuyaPolys["bottom"], shibuyaPolys["top"])
	predicates["shibuyabl"] = StartEndPredicate(shibuyaPolys["bottom"], shibuyaPolys["left"])
	predicates["shibuyarl"] = StartEndPredicate(shibuyaPolys["right"], shibuyaPolys["left"])
	predicates["shibuyart"] = StartEndPredicate(shibuyaPolys["right"], shibuyaPolys["top"])
	predicates["shibuyarb"] = StartEndPredicate(shibuyaPolys["right"], shibuyaPolys["bottom"])

	predicates["shibuya"] = Or(predicates["shibuyabt"], predicates["shibuyabl"], predicates["shibuyarl"], predicates["shibuyart"], predicates["shibuyarb"])
}
