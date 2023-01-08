package parsers

type Parser interface {
	Parse(data string) (string, error)
}
