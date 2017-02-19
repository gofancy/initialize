package initialize_test

import (
	"testing"

	"github.com/gofancy/initialize"
)

var (
	results  []int
	expected = []int{1, 2}
)

type myInit struct{}

func (v myInit) Initialize01Log() {
	results = append(results, 1)
}

func (v myInit) Initialize02DB() {
	results = append(results, 2)
}

func TestFromAll(t *testing.T) {
	results = make([]int, 0)
	initialize.AllFrom(myInit{})
	if len(results) != 2 || results[0] != 1 || results[1] != 2 {
		t.Error(
			"For", "All",
			"expected", expected,
			"got", results,
		)
	}
}
