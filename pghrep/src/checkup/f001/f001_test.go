package f001

import (
	"fmt"
	"testing"

	checkup ".."
)

var TestLastNodesJson checkup.ReportLastNodes = checkup.ReportLastNodes{
	Hosts: checkup.ReportHosts{
		"test-host": {
			Role: "master",
		},
	},
}

func TestF001Successful(t *testing.T) {
	fmt.Println(t.Name())
	var report F001Report
	var hostResult F001ReportHostResult
	hostResult.Data = F001ReportHostResultData{
		Settings: F001Settings{
			GlobalSettings: map[string]F001GlobalSetting{
				"autovacuum_vacuum_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_vacuum_scale_factor",
					Setting: "0.01",
				},
				"autovacuum_vacuum_threshold": F001GlobalSetting{
					Name:    "autovacuum_vacuum_threshold",
					Setting: "50",
				},
				"autovacuum_analyze_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_analyze_scale_factor",
					Setting: "0.05",
				},
				"autovacuum_analyze_threshold": F001GlobalSetting{
					Name:    "autovacuum_analyze_threshold",
					Setting: "50",
				},
				"autovacuum_vacuum_cost_delay": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_delay",
					Setting: "20",
				},
				"autovacuum_vacuum_cost_limit": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_limit",
					Setting: "0",
				},
			},
			TableSettings: map[string]F001TableSetting{},
		},
	}

	report.Results = F001ReportHostsResults{"test-host": hostResult}
	report.LastNodesJson = TestLastNodesJson
	result, err := F001Process(report)

	if err != nil || result.P1 || result.P2 || result.P3 {
		t.Fatal()
	}

	checkup.PrintResultConclusions(result)
	checkup.PrintResultRecommendations(result)
}

func TestF001AllCases(t *testing.T) {
	fmt.Println(t.Name())
	var report F001Report
	var hostResult F001ReportHostResult
	hostResult.Data = F001ReportHostResultData{
		Settings: F001Settings{
			GlobalSettings: map[string]F001GlobalSetting{
				"autovacuum_vacuum_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_vacuum_scale_factor",
					Setting: "0.2",
				},
				"autovacuum_vacuum_threshold": F001GlobalSetting{
					Name:    "autovacuum_vacuum_threshold",
					Setting: "50",
				},
				"autovacuum_analyze_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_analyze_scale_factor",
					Setting: "0.1",
				},
				"autovacuum_analyze_threshold": F001GlobalSetting{
					Name:    "autovacuum_analyze_threshold",
					Setting: "50",
				},
				"autovacuum_vacuum_cost_delay": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_delay",
					Setting: "20",
				},
				"autovacuum_vacuum_cost_limit": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_limit",
					Setting: "-1",
				},
			},
			TableSettings: map[string]F001TableSetting{},
		},
	}

	report.Results = F001ReportHostsResults{"test-host": hostResult}
	report.LastNodesJson = TestLastNodesJson
	result, err := F001Process(report)

	if err != nil || !result.P1 || result.P2 || result.P3 ||
		!checkup.ResultInList(result.Conclusions, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_TUNE_RECOMMENDATION) {
		t.Fatal()
	}

	checkup.PrintResultConclusions(result)
	checkup.PrintResultRecommendations(result)
}

func TestF001FirstCase(t *testing.T) {
	fmt.Println(t.Name())
	var report F001Report
	var hostResult F001ReportHostResult
	hostResult.Data = F001ReportHostResultData{
		Settings: F001Settings{
			GlobalSettings: map[string]F001GlobalSetting{
				"autovacuum_vacuum_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_vacuum_scale_factor",
					Setting: "0.2",
				},
				"autovacuum_vacuum_threshold": F001GlobalSetting{
					Name:    "autovacuum_vacuum_threshold",
					Setting: "50",
				},
				"autovacuum_analyze_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_analyze_scale_factor",
					Setting: "0.5",
				},
				"autovacuum_analyze_threshold": F001GlobalSetting{
					Name:    "autovacuum_analyze_threshold",
					Setting: "100",
				},
				"autovacuum_vacuum_cost_delay": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_delay",
					Setting: "20",
				},
				"autovacuum_vacuum_cost_limit": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_limit",
					Setting: "-1",
				},
			},
			TableSettings: map[string]F001TableSetting{},
		},
	}
	report.Results = F001ReportHostsResults{"test-host": hostResult}
	report.LastNodesJson = TestLastNodesJson
	result, err := F001Process(report)

	if err != nil || !result.P1 || result.P2 || result.P3 ||
		!checkup.ResultInList(result.Conclusions, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_TUNE_RECOMMENDATION) {
		t.Fatal()
	}

	checkup.PrintResultConclusions(result)
	checkup.PrintResultRecommendations(result)
}

func TestF001SecondCase(t *testing.T) {
	fmt.Println(t.Name())
	var report F001Report
	var hostResult F001ReportHostResult
	hostResult.Data = F001ReportHostResultData{
		Settings: F001Settings{
			GlobalSettings: map[string]F001GlobalSetting{
				"autovacuum_vacuum_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_vacuum_scale_factor",
					Setting: "0.2",
				},
				"autovacuum_vacuum_threshold": F001GlobalSetting{
					Name:    "autovacuum_vacuum_threshold",
					Setting: "50",
				},
				"autovacuum_analyze_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_analyze_scale_factor",
					Setting: "0.5",
				},
				"autovacuum_analyze_threshold": F001GlobalSetting{
					Name:    "autovacuum_analyze_threshold",
					Setting: "70",
				},
				"autovacuum_vacuum_cost_delay": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_delay",
					Setting: "40",
				},
				"autovacuum_vacuum_cost_limit": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_limit",
					Setting: "0",
				},
			},
			TableSettings: map[string]F001TableSetting{},
		},
	}

	report.Results = F001ReportHostsResults{"test-host": hostResult}
	report.LastNodesJson = TestLastNodesJson
	result, err := F001Process(report)

	if err != nil || !result.P1 || result.P2 || result.P3 ||
		!checkup.ResultInList(result.Conclusions, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_TUNE_RECOMMENDATION) {
		t.Fatal()
	}

	checkup.PrintResultConclusions(result)
	checkup.PrintResultRecommendations(result)
}

func TestF001ThirdCase(t *testing.T) {
	fmt.Println(t.Name())
	var report F001Report
	var hostResult F001ReportHostResult
	hostResult.Data = F001ReportHostResultData{
		Settings: F001Settings{
			GlobalSettings: map[string]F001GlobalSetting{
				"autovacuum_vacuum_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_vacuum_scale_factor",
					Setting: "1",
				},
				"autovacuum_vacuum_threshold": F001GlobalSetting{
					Name:    "autovacuum_vacuum_threshold",
					Setting: "100",
				},
				"autovacuum_analyze_scale_factor": F001GlobalSetting{
					Name:    "autovacuum_analyze_scale_factor",
					Setting: "1",
				},
				"autovacuum_analyze_threshold": F001GlobalSetting{
					Name:    "autovacuum_analyze_threshold",
					Setting: "100",
				},
				"autovacuum_vacuum_cost_delay": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_delay",
					Setting: "20",
				},
				"autovacuum_vacuum_cost_limit": F001GlobalSetting{
					Name:    "autovacuum_vacuum_cost_limit",
					Setting: "-1",
				},
			},
			TableSettings: map[string]F001TableSetting{},
		},
	}

	report.Results = F001ReportHostsResults{"test-host": hostResult}
	report.LastNodesJson = TestLastNodesJson
	result, err := F001Process(report)

	if err != nil || !result.P1 || result.P2 || result.P3 ||
		!checkup.ResultInList(result.Conclusions, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_NOT_TUNED) ||
		!checkup.ResultInList(result.Recommendations, F001_AUTOVACUUM_TUNE_RECOMMENDATION) {
		t.Fatal()
	}

	checkup.PrintResultConclusions(result)
	checkup.PrintResultRecommendations(result)
}
