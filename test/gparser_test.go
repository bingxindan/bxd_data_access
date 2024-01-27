package test

import (
	"fmt"
	"github.com/wangxin1248/gparser"
	"testing"
)

func TestParse(t *testing.T) {
	ruleStr := "!(a == 1 && b == 2 && c == \"test\" && d == false)"

	// 匹配变量
	params := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": "test",
		"d": true,
	}

	result, err := gparser.Match(ruleStr, params)

	fmt.Println(result, "err:", err)
}
