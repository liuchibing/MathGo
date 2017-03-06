package mathgo

import "regexp"

/*
	实数运算解释器 -- 实现了Interpreter 接口
 */
type ratInterpreter struct {
	vars map[string]MthExp
}

func (ri *ratInterpreter) Run(exp string) bool {
	regSetVar := regexp.MustCompile("(.+)=(.+);?")
	regRunFunc := regexp.MustCompile("(.+)\\((.+)\\)")
}
