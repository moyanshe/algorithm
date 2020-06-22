package classics

import (
	"reflect"
	"testing"
)

func TestBubbleAsort(t *testing.T)  {
	tests := []struct{
		input []int
		ans []int
	}{
		{[]int{2,4,5,3}, []int{2,3,4,5}},
		{[]int{2,2,5,0}, []int{0,2,2,5}},
		{[]int{1,0,0,0}, []int{0,0,0,1}},
	}

	for i := 0; i < len(tests); i++ {
		result := BubbleAsort(tests[i].input)
		if !reflect.DeepEqual(result, tests[i].ans) {
			t.Errorf("got %v for input %v, expected %v", tests[i].input, tests[i].ans, result)
		}
	}
}