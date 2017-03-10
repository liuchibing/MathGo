/*
	mathgo 包是一个数学工具--脚本解释器。语法类似 Scilab 。
*/
package mathgo

type Interpreter interface {
	/*
		执行表达式
		Parameter: exp: 要执行的表达式
		Return: result: 表达式的结果
		        needNextLine: 是否需要下一行语句才能得到结果
	*/
	Run(exp string) (result MthExp, needNextLine bool)
}

func CreateInterpreter() Interpreter {
	return newRatInterpreter()
}
