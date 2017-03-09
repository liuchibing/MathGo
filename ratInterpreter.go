package mathgo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
	实数运算解释器 -- 实现了Interpreter 接口
*/
type ratInterpreter struct {
	vars map[string]MthExp

	//用于判断命令的Regexp
	regSetVar, regRunFunc, regCalc *regexp.Regexp
}

func (ri ratInterpreter) Run(exp string) (result MthExp, needNextLine bool) {
	//判断是否有这个变量
	item, hasVar := ri.vars[exp]
	//判断是否可以转为数字
	num, error := strconv.ParseFloat(strings.TrimSpace(exp), 64)
	//解析表达式
	switch {
	case error == nil: //表达式是数值
		return num, false
	case hasVar: //表达式是变量名
		return item, false
	case ri.regSetVar.MatchString(exp): //创建变量
		matches := ri.regSetVar.FindStringSubmatch(exp)
		ri.vars[matches[1]], _ = ri.Run(matches[2])
		return ri.vars[matches[1]], false
	case ri.regRunFunc.MatchString(exp): //执行函数
		matches := ri.regRunFunc.FindStringSubmatch(exp)
		res := ri.handleFunc(matches)
		return res, false
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
	//处理参数列表
	args := strings.Split(sign[2], ",")
	//筛选处理
	switch sign[1] {
	case "print":
		for _, v := range args {
			r, _ := ri.Run(v)
			fmt.Print(r)
			fmt.Print(" ")
		}
		fmt.Print("\n")
		return nil
	}
	return nil
}
