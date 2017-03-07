/*
	mathgo 包是一个数学工具--脚本解释器。语法类似 Scilab 。
*/
package mathgo

type Interpreter interface {
	Run(exp string) bool
}

func NewInterpreter() Interpreter {
	return newRatInterpreter()
}
