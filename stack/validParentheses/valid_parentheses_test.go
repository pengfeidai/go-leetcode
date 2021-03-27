package stack

import (
	"testing"
)

type question struct {
	param  string
	answer bool
}

func TestValidParentheses(t *testing.T) {
	qs := []question{
		{
			param:  "([])",
			answer: true,
		},
		{
			param:  "([{]])",
			answer: false,
		},
		{
			param:  "(())]]",
			answer: false,
		},
		{
			param:  "(",
			answer: false,
		},
		{
			param:  "(){[({[]})]}",
			answer: true,
		},
	}
	for _, q := range qs {
		t.Logf("【input】:%v       【output】:%v\n", q.answer, isValid(q.param))
	}
}
