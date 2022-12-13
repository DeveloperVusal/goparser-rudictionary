package types

type ParserInterface interface {
	Execute(*map[string]int)
	MorphyBuild(string) map[string][]map[string]string
}
