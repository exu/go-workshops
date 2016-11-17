package filters

import (
	. "ioki.com.pl/perf/gatling-report-compare/comparator"
	"testing"
)

func getContents() Contents {
	return Contents{
		"six":   Content{Stats: Stats{Percentiles1: ReqStat{Total: 7}, StandardDeviation: ReqStat{Total: 3}}},
		"five":  Content{Stats: Stats{Percentiles1: ReqStat{Total: 6}, StandardDeviation: ReqStat{Total: 2}}},
		"four":  Content{Stats: Stats{Percentiles1: ReqStat{Total: 5}, StandardDeviation: ReqStat{Total: 1}}},
		"three": Content{Stats: Stats{Percentiles1: ReqStat{Total: 4}, StandardDeviation: ReqStat{Total: 5}}},
		"two":   Content{Stats: Stats{Percentiles1: ReqStat{Total: 3}, StandardDeviation: ReqStat{Total: 12}}},
		"one":   Content{Stats: Stats{Percentiles1: ReqStat{Total: 2}, StandardDeviation: ReqStat{Total: 11}}},
		"zero":  Content{Stats: Stats{Percentiles1: ReqStat{Total: 1}, StandardDeviation: ReqStat{Total: 10}}},
	}
}

func TestTop1_ReturnsOneResult(t *testing.T) {
	contents := getContents()

	filtered := Top(1, ByPercentiles1{contents.ToArray()})
	assertLen(t, 1, filtered)
}

func TestTop5ByPercentiles1_ReturnsSortedResults(t *testing.T) {
	contents := Contents{
		"one":   Content{Stats: Stats{Percentiles1: ReqStat{Ko: 1, Ok: 1, Total: 2}, StandardDeviation: ReqStat{Ko: 1, Ok: 1, Total: 3}}},
		"two":   Content{Stats: Stats{Percentiles1: ReqStat{Ko: 1, Ok: 1, Total: 1}, StandardDeviation: ReqStat{Ko: 1, Ok: 1, Total: 2}}},
		"three": Content{Stats: Stats{Percentiles1: ReqStat{Ko: 1, Ok: 1, Total: 3}, StandardDeviation: ReqStat{Ko: 1, Ok: 1, Total: 1}}},
	}

	filtered := Top(3, ByPercentiles1{contents.ToArray()})

	assertPosition(t, "three", contents, 0, filtered)
	assertPosition(t, "one", contents, 1, filtered)
	assertPosition(t, "two", contents, 2, filtered)
}

func TestTop5ByPercentiles1_ReturnsTop5Ressults_SortedByPercentiles1(t *testing.T) {
	contents := getContents()

	filtered := Top(5, ByPercentiles1{contents.ToArray()})

	assertLen(t, 5, filtered)

	assertPosition(t, "six", contents, 0, filtered)
	assertPosition(t, "five", contents, 1, filtered)
	assertPosition(t, "four", contents, 2, filtered)
	assertPosition(t, "three", contents, 3, filtered)
	assertPosition(t, "two", contents, 4, filtered)
}

func TestTop3ByByStandardDeviation_ReturnsTop5Ressults_SortedByStandardDeviation(t *testing.T) {
	contents := getContents()

	filtered := Top(3, ByStandardDeviation{contents.ToArray()})

	assertLen(t, 3, filtered)

	assertPosition(t, "two", contents, 0, filtered)
	assertPosition(t, "one", contents, 1, filtered)
	assertPosition(t, "zero", contents, 2, filtered)
}
