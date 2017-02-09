package stringutil

type converter interface {
	Process(text string) string
}
