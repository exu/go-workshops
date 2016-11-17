package filters

import (
	"ioki.com.pl/perf/gatling-report-compare/comparator"
)

func NewSorter(name string, data comparator.ContentsArray) comparator.Sortable {

	switch name {

	case "group_1":
		return ByGroup1{data}

	case "group_2":
		return ByGroup2{data}

	case "group_3":
		return ByGroup3{data}

	case "group_4":
		return ByGroup4{data}

	case "max_response_time":
		return ByMaxResponseTime{data}

	case "mean_number_ofrequests_per_second":
		return ByMeanNumberOfRequestsPerSecond{data}

	case "mean_response_time":
		return ByMeanResponseTime{data}

	case "number_of_requests":
		return ByNumberOfRequests{data}

	case "percentiles_1":
		return ByPercentiles1{data}

	case "percentiles_2":
		return ByPercentiles2{data}

	case "percentiles_3":
		return ByPercentiles3{data}

	case "percentiles_4":
		return ByPercentiles4{data}

	case "standard_deviation":
		return ByStandardDeviation{data}

	}

	return data
}

type ByGroup1 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByGroup1) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Group1.Count >= f.ContentsArray[j].Stats.Group1.Count
}

type ByGroup2 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByGroup2) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Group2.Count >= f.ContentsArray[j].Stats.Group2.Count
}

type ByGroup3 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByGroup3) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Group3.Count >= f.ContentsArray[j].Stats.Group3.Count
}

type ByGroup4 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByGroup4) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Group4.Count >= f.ContentsArray[j].Stats.Group4.Count
}

type ByMaxResponseTime struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByMaxResponseTime) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.MaxResponseTime.Total >= f.ContentsArray[j].Stats.MaxResponseTime.Total
}

type ByMeanNumberOfRequestsPerSecond struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByMeanNumberOfRequestsPerSecond) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.MeanNumberOfRequestsPerSecond.Total >= f.ContentsArray[j].Stats.MeanNumberOfRequestsPerSecond.Total
}

type ByMeanResponseTime struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByMeanResponseTime) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.MeanResponseTime.Total >= f.ContentsArray[j].Stats.MeanResponseTime.Total
}

type ByNumberOfRequests struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByNumberOfRequests) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.NumberOfRequests.Total >= f.ContentsArray[j].Stats.NumberOfRequests.Total
}

type ByPercentiles1 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByPercentiles1) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Percentiles1.Total >= f.ContentsArray[j].Stats.Percentiles1.Total
}

type ByPercentiles2 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByPercentiles2) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Percentiles2.Total >= f.ContentsArray[j].Stats.Percentiles2.Total
}

type ByPercentiles3 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByPercentiles3) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Percentiles3.Total >= f.ContentsArray[j].Stats.Percentiles3.Total
}

type ByPercentiles4 struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByPercentiles4) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.Percentiles4.Total >= f.ContentsArray[j].Stats.Percentiles4.Total
}

type ByStandardDeviation struct{ comparator.ContentsArray }

// Less sorts in descending order
func (f ByStandardDeviation) Less(i, j int) bool {
	return f.ContentsArray[i].Stats.StandardDeviation.Total >= f.ContentsArray[j].Stats.StandardDeviation.Total
}
