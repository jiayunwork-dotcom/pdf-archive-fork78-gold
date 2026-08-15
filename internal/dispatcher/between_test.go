package dispatcher

import "testing"

func TestEvaluateCondition_BetweenInclusive(t *testing.T) {
	if !EvaluateCondition("between", []interface{}{10, 20}, "10") {
		t.Fatal(`EvaluateCondition(between, [10,20], "10") = false, want true`)
	}
}
