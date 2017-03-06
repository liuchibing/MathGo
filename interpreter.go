package mathgo

type Interpreter interface {
	Run(exp string) bool
}
