package checkup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"

	"../log"
	"../orderedmap"
	"../pyraconv"
)

// General for all reports

const MSG_ALL_GOOD_CONCLUSION string = "Hooray, all good. Keep this up!"
const MSG_NO_RECOMMENDATION string = "No recommendations."
const MSG_ETC_ITEM string = "    - etc."
const RECOMMENDATION_ITEMS_LIMIT int = 5

type ReportHost struct {
	InternalAlias        string `json:"internal_alias"`
	Index                string `json:"index"`
	Role                 string `json:"role"`
	RoleChangeDetectedAt string `json:"role_change_detected_at"`
}

type ReportLastCheck struct {
	Epoch       string `json:"epoch"`
	Timestamptz string `json:"timestamptz"`
	Dir         string `json:"dir"`
}

type ReportHosts map[string]ReportHost

type ReportLastNodes struct {
	Hosts     ReportHosts     `json:"hosts"`
	LastCheck ReportLastCheck `json:"last_check"`
	//	LastCheck
}

type ReportResultItem struct {
	Id      string
	Message string
}

type ReportResult struct {
	P1              bool
	P2              bool
	P3              bool
	Conclusions     []ReportResultItem
	Recommendations []ReportResultItem
}

func (r *ReportResult) AppendConclusion(id string, conclusion string, a ...interface{}) {
	r.Conclusions = append(r.Conclusions, ReportResultItem{
		Id:      id,
		Message: fmt.Sprintf(conclusion, a...),
	})
}

func (r *ReportResult) AppendRecommendation(id string, reccomendation string,
	a ...interface{}) {
	if reccomendation != "" {
		r.Recommendations = append(r.Recommendations, ReportResultItem{
			Id:      id,
			Message: fmt.Sprintf(reccomendation, a...),
		})
	}
}

func LoadRawJsonReport(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}
	}
	defer file.Close()
	jsonRaw, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}
	}
	return jsonRaw
}

func CheckUnmarshalResult(err error) bool {
	if err != nil {
		log.Err("Cannot load json report to process")
		return false
	}
	return true
}

func SaveJsonReportResults(data map[string]interface{}, reportResult ReportResult) {
	filePath := pyraconv.ToString(data["source_path_full"])
	jsonData, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		return
	}
	orderedData := orderedmap.New()
	if err := json.Unmarshal([]byte(jsonData), &orderedData); err != nil {
		return
	} else {
		orderedData.Set("processed", true)
		orderedData.Set("conclusions", reportResult.Conclusions)
		orderedData.Set("recommendations", reportResult.Recommendations)
		orderedData.Set("p1", reportResult.P1)
		orderedData.Set("p2", reportResult.P2)
		orderedData.Set("p3", reportResult.P3)
		resultJson, merr := orderedData.MarshalJSON()
		if merr != nil {
			return
		}
		var out bytes.Buffer
		json.Indent(&out, resultJson, "", "  ")
		jfile, err := os.Create(filePath)
		if err != nil {
			return
		}
		defer jfile.Close()
		out.WriteTo(jfile)
	}
}

func SaveReportResult(data map[string]interface{},
	result ReportResult) map[string]interface{} {
	data["conclusions"] = result.Conclusions
	data["recommendations"] = result.Recommendations
	data["p1"] = result.P1
	data["p2"] = result.P2
	data["p3"] = result.P3
	data["processed"] = true
	SaveJsonReportResults(data, result)
	return data
}

func GetUniques(array []string) []string {
	items := map[string]bool{}
	for _, item := range array {
		items[item] = true
	}

	res := make([]string, len(items))
	i := 0
	for key, _ := range items {
		res[i] = key
		i++
	}
	return res
}

func LimitList(array []string) []string {
	if len(array) > (RECOMMENDATION_ITEMS_LIMIT + 1) {
		limitedArray := array[0:RECOMMENDATION_ITEMS_LIMIT]
		limitedArray = append(limitedArray, MSG_ETC_ITEM)
		return limitedArray
	} else {
		return array
	}
}

func InList(items []string, item string) bool {
	for _, itemValue := range items {
		if strings.Trim(itemValue, " \n") == strings.Trim(item, " \n") {
			return true
		}
	}
	return false
}

func InListPartial(items []string, item string) bool {
	for _, itemValue := range items {
		if strings.Contains(strings.Trim(itemValue, " \n"), strings.Trim(item, " \n")) {
			return true
		}
	}
	return false
}

func ResultInList(items []ReportResultItem, id string) bool {
	for _, itemValue := range items {
		if itemValue.Id == id {
			return true
		}
	}
	return false
}

func GetResultItem(items []ReportResultItem, id string) (*ReportResultItem, error) {
	for _, itemValue := range items {
		if itemValue.Id == id {
			return &itemValue, nil
		}
	}
	return nil, fmt.Errorf("Item not found")
}

func PrintResultConclusions(result ReportResult) {
	for _, conclusion := range result.Conclusions {
		fmt.Println("C:  ", conclusion.Message)
	}
}

func PrintResultRecommendations(result ReportResult) {
	for _, recommendation := range result.Recommendations {
		fmt.Println("R:  ", recommendation.Message)
	}
}

func GetMasterHostName(hosts ReportHosts) string {
	var firstHostName string = ""

	for hostName, host := range hosts {
		if firstHostName == "" {
			firstHostName = hostName
		}

		if host.Role == "master" {
			return hostName
		}
	}

	return firstHostName
}

// Get map keys sorted by defined int field inside struct
func SortItemsByInt(data interface{}, field string, reverse bool) []string {
	var result []string
	var numData map[int]string = map[int]string{}
	var keys []int

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Map {
		v2 := v.MapKeys()

		for _, itemData := range v2 {
			id := itemData.Interface()
			val := v.MapIndex(itemData)

			if val.Kind() != reflect.Struct {
				continue
			}

			valNum := val.FieldByName(field)

			if valNum.Kind() == reflect.Invalid {
				continue
			}

			num := valNum.Interface()
			intNum := num.(int)
			numData[intNum] = id.(string)
			keys = append(keys, intNum)
		}

		sort.Ints(keys)

		if reverse {
			sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		}

		for _, key := range keys {
			result = append(result, numData[key])
		}
	}

	return result
}

// Get map keys sorted by defined float64 field inside struct
func SortItemsByFloat64(data interface{}, field string, reverse bool) []string {
	var result []string
	var numData map[float64]string = map[float64]string{}
	var keys []float64

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Map {
		v2 := v.MapKeys()

		for _, itemData := range v2 {
			id := itemData.Interface()
			val := v.MapIndex(itemData)

			if val.Kind() != reflect.Struct {
				continue
			}

			valNum := val.FieldByName(field)

			if valNum.Kind() == reflect.Invalid {
				continue
			}

			num := valNum.Interface()
			floatNum := num.(float64)
			numData[floatNum] = id.(string)
			keys = append(keys, floatNum)
		}

		sort.Float64s(keys)

		if reverse {
			sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		}

		for _, key := range keys {
			result = append(result, numData[key])
		}
	}

	return result
}

// Get map keys sorted by field num inside struct
func GetItemsSortedByNum(data interface{}) []string {
	return SortItemsByInt(data, "Num", false)
}

// Check if the string exists in the array. If it is so, return its index
func StringInArray(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}
