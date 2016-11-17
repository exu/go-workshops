package filters

import (
	. "ioki.com.pl/perf/gatling-report-compare/comparator"
	"testing"
)

func assertPosition(t *testing.T, key string, hashMap Contents, index int, array ContentsArray) {
	if hashMap[key] != *array[index] {
		t.Fatalf("Not equal \n\n%s:%#v\n\n%d:%#v", key, hashMap[key], index, *array[index])
	}
}

func assertLen(t *testing.T, expectedLen int, contents Sortable) {
	if expectedLen != contents.Len() {
		t.Fatalf("Expected %d but %d given, in: %+v", expectedLen, contents.Len(), contents)
	}
}
