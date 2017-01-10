package types

type MethodWrapper struct {
	Package *string
	Service *string
	Method  *string
	Pattern *Pattern
}

type Pattern struct {
	Verb string
	Path string
	Body string
}
