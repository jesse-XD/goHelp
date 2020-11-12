package help

import (
	"fmt"
	"testing"
)

func TestGetStrCn(t *testing.T) {
	str := "扣12你412吉512瓦"
	s := GetStrCn(str)
	fmt.Println(s)
}
