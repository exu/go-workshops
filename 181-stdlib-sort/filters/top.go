package filters

import (
	"sort"

	"ioki.com.pl/perf/gatling-report-compare/comparator"
)

func Top(n int, decorator comparator.Sortable) comparator.ContentsArray {
	sort.Sort(decorator)

	if decorator.Len() < n {
		return decorator.ToArray()
	}

	return decorator.Slice(0, n)
}
