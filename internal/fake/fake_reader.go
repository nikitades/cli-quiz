package fake

import "github.com/nikitades/cli-quiz/internal/game"

type FakeReader struct {
}

func (fr *FakeReader) ReadAll() ([]game.Question, error) {
	return []game.Question{
		{
			Title:          "First q",
			Options:        []string{"Right", "Wrong"},
			CorrectOptions: []int{1},
		},
		{
			Title:          "Second q",
			Options:        []string{"Wrong", "Right"},
			CorrectOptions: []int{2},
		},
		{
			Title:          "Third q",
			Options:        []string{"Wrong 1", "Wrong 2", "Right 1", "Right 2"},
			CorrectOptions: []int{3, 4},
		},
	}, nil
}
