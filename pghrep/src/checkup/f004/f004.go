package f004

import (
	"encoding/json"
	"fmt"
	"strings"

	"../../log"

	checkup ".."
	"../../fmtutils"
	"../../pyraconv"
)

const WARNING_BLOAT_RATIO float32 = 40.0
const CRITICAL_BLOAT_RATIO float32 = 90.0
const CRITICAL_TOTAL_BLOAT_RATIO float32 = 20.0
const MIN_INDEX_SIZE_TO_ANALYZE int64 = 1024 * 1024

func appendTable(list []string, tableBloatData F004HeapBloat) []string {
	return append(list, fmt.Sprintf(TABLE_DETAILS, tableBloatData.TableName,
		fmtutils.ByteFormat(float64(tableBloatData.RealSizeBytes), 2),
		tableBloatData.BloatRatio, fmtutils.ByteFormat(float64(tableBloatData.ExtraSizeBytes), 2),
		tableBloatData.BloatRatioPercent))
}

// Generate conclusions and recommendatons
func F004Process(report F004Report) checkup.ReportOutcome {
	var result checkup.ReportOutcome
	// check total values
	var criticalTables []string
	var warningTables []string
	totalBloatIsCritical := false
	var totalData F004HeapBloatTotal
	var databaseSize int64
	for _, hostData := range report.Results {
		databaseSize = hostData.Data.DatabaseSizeBytes
		totalData = hostData.Data.HeapBloatTotal
		if hostData.Data.HeapBloatTotal.BloatRatioPercentAvg > CRITICAL_TOTAL_BLOAT_RATIO {
			totalBloatIsCritical = true
			result.P1 = true
		}
		for _, heapBloatData := range hostData.Data.HeapBloat {
			if (heapBloatData.RealSizeBytes > MIN_INDEX_SIZE_TO_ANALYZE) && (heapBloatData.BloatRatioPercent >= WARNING_BLOAT_RATIO) &&
				(heapBloatData.BloatRatioPercent < CRITICAL_BLOAT_RATIO) {
				warningTables = appendTable(warningTables, heapBloatData)
			}
			if heapBloatData.RealSizeBytes > MIN_INDEX_SIZE_TO_ANALYZE && heapBloatData.BloatRatioPercent >= CRITICAL_BLOAT_RATIO {
				criticalTables = appendTable(criticalTables, heapBloatData)
			}
		}
	}
	if totalBloatIsCritical {
		result.AppendConclusion(MSG_TOTAL_BLOAT_EXCESS_CONCLUSION,
			fmtutils.ByteFormat(float64(totalData.BloatSizeBytesSum), 2),
			totalData.BloatRatioPercentAvg,
			float64(float64(totalData.BloatSizeBytesSum)/float64(databaseSize)*100),
			fmtutils.ByteFormat(float64(databaseSize-totalData.BloatSizeBytesSum), 2),
			fmtutils.ByteFormat(float64(totalData.BloatSizeBytesSum), 2),
			totalData.BloatRatioAvg)
		result.P1 = true
	} else {
		result.AppendConclusion(MSG_TOTAL_BLOAT_LOW_CONCLUSION, totalData.BloatRatioPercentAvg,
			fmtutils.ByteFormat(float64(totalData.BloatSizeBytesSum), 2))
	}
	if len(criticalTables) > 0 {
		result.AppendConclusion(MSG_BLOAT_CRITICAL_CONCLUSION, len(criticalTables), CRITICAL_BLOAT_RATIO,
			strings.Join(checkup.LimitList(criticalTables), ""))
		result.AppendRecommendation(MSG_BLOAT_CRITICAL_RECOMMENDATION)
		result.P1 = true
	}
	if len(warningTables) > 0 {
		result.AppendConclusion(MSG_BLOAT_WARNING_CONCLUSION, len(warningTables), WARNING_BLOAT_RATIO, CRITICAL_BLOAT_RATIO,
			strings.Join(checkup.LimitList(warningTables), ""))
		if !result.P1 {
			result.AppendRecommendation(MSG_BLOAT_WARNING_RECOMMENDATION)
		}
		result.P2 = true
	}
	if len(result.Recommendations) > 0 {
		result.AppendRecommendation(MSG_BLOAT_GENERAL_RECOMMENDATION_1)
		result.AppendRecommendation(MSG_BLOAT_GENERAL_RECOMMENDATION_2)
	}
	if result.P1 || result.P2 {
		result.AppendRecommendation(MSG_BLOAT_PX_RECOMMENDATION)
	}
	return result
}

func F004PreprocessReportData(data map[string]interface{}) {
	filePath := pyraconv.ToString(data["source_path_full"])
	jsonRaw := checkup.LoadRawJsonReport(filePath)
	var report F004Report
	err := json.Unmarshal(jsonRaw, &report)
	if err != nil {
		log.Err("Cannot load json report to process")
		return
	}
	result := F004Process(report)
	if len(result.Recommendations) == 0 {
		result.AppendRecommendation(MSG_NO_RECOMMENDATIONS)
	}
	checkup.SaveConclusionsRecommendations(data, result)
}
