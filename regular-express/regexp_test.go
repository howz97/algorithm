package regexp

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	//fmt.Println(IsMatch(`abc`, "abc"))
	fmt.Println(IsMatch(`(1((\.|\()|c)*2)`, "1c((..2"))
}
