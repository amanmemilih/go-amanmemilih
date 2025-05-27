package interfaces

type WordClient interface {
	GetRandomWords(count int) ([]string, error)
}
