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
	regSetVar, regRunFunc, regCalc *regexp.Regexp
}

func (ri ratInterpreter) Run(exp string) (MthExp, bool) {
	switch {
	case ri.regSetVar.MatchString(exp): //创建变量
		matches := ri.regSetVar.FindStringSubmatch(exp)
		fmt.Println(matches)
		ri.vars[matches[1]] = matches[2]
		return matches[2], false
	case ri.regRunFunc.MatchString(exp): //执行函数
		matches := ri.regRunFunc.FindStringSubmatch(exp)
		ri.handleFunc(matches)
		fmt.Println(matches)
	}
	return nil, false
}

func newRatInterpreter() ratInterpreter {
	var ri ratInterpreter
	ri.vars = make(map[string]MthExp)

	ri.regSetVar = regexp.MustCompile("([^0-9][a-zA-Z0-9_]+)=([^=\\n]+);?")
	ri.regRunFunc = regexp.MustCompile("(.+)\\((.+)\\);?")
	ri.regCalc = regexp.MustCompile("(.+)([\\+\\-\\*\\/\\^\\%])(.+);?")

	return ri
}

func (ri ratInterpreter) handleFunc(sign []string) MthExp {

}
