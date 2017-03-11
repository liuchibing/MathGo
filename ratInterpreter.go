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

/*
执行表达式
Parameter: exp: 要执行的表达式
Return: result: 表达式的结果
        needNextLine: 是否需要下一行语句才能得到结果
*/
func (ri ratInterpreter) Run(exp string) (result MthExp, needNextLine bool) {
	//TrimSpace
	exp = strings.TrimSpace(exp)
	//判断是否有这个变量
	item, hasVar := ri.vars[exp]
	//判断是否可以转为数字
	num, error := strconv.ParseFloat(exp, 64)
	//解析表达式
	switch {
	case error == nil: //表达式是数值
		return num, false
	case hasVar: //表达式是变量名
		return item, false
	case ri.regSetVar.MatchString(exp): //变量赋值
		matches := ri.regSetVar.FindStringSubmatch(exp)
		ri.vars[matches[1]], _ = ri.Run(matches[2])
		return ri.vars[matches[1]], false
	case ri.regCalc.MatchString(exp): //四则运算
		matches := ri.regCalc.FindStringSubmatch(exp)
		return matches, false
	case ri.regRunFunc.MatchString(exp): //执行函数
		matches := ri.regRunFunc.FindStringSubmatch(exp)
		result := ri.handleFunc(matches)
		return result, false
	default:
		return "error", false
	}
}

func newRatInterpreter() ratInterpreter {
	var ri ratInterpreter
	ri.vars = make(map[string]MthExp)

	ri.regSetVar = regexp.MustCompile("([^0-9][a-zA-Z0-9_]+)=([^=]+);?")
	ri.regRunFunc = regexp.MustCompile("([^0-9][a-zA-Z0-9_]+)\\((.+)\\);?")
	ri.regCalc = regexp.MustCompile("(.+)([+\\-*^%])(.+);?")

	return ri
}

func (ri ratInterpreter) handleFunc(sign []string) MthExp {
	//处理参数列表
	var args []string
	if strings.Contains(sign[2], ",") {
		args = strings.Split(sign[2], ",")
	}
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
