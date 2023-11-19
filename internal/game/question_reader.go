package game

type QuestionReader interface {
	ReadAll() ([]Question, error)
}
