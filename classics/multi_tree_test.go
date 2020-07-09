package classics

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tests := []struct {
		input []string
		ans   string
	}{
		{[]string{"a", "b", "c"}, "a,b,c"},
	}

	for i := 0; i < len(tests); i++ {
		fmt.Println(tests[i].ans)
	}
}