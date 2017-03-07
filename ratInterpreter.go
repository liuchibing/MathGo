package mathgo

import (
	"fmt"
	"regexp"
)

/*
	实数运算解释器 -- 实现了Interpreter 接口
*/
type ratInterpreter struct {
	vars map[string]MthExp

	//用于判断命令的Regexp
	regSetVar, regRunfunc, regCalc *regexp.Regexp
}

func (ri ratInterpreter) Run(exp string) bool {
	if ri.regSetVar.MatchString(exp) {
		result := ri.regSetVar.FindStringSubmatch(exp)
		fmt.Println(result)
	}
	return false
}

func newRatInterpreter() ratInterpreter {
	var ri ratInterpreter
	ri.vars = make(map[string]MthExp)

	ri.regSetVar = regexp.MustCompile("(.+)=(.+);?")
	ri.regRunfunc = regexp.MustCompile("(.+)\\((.+)\\);?")
	ri.regCalc = regexp.MustCompile("(.+)([\\+\\-\\*\\/\\^\\%])(.+);?")

	return ri
}
