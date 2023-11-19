package game

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Game struct {
	Reader QuestionReader
}

func (g *Game) Run(screen tcell.Screen) {
	app := tview.NewApplication().SetScreen(screen)
	quizViewmodel := newQuizViewModel(app)

	questions, err := g.Reader.ReadAll()
	if err != nil {
		quizViewmodel.ShowError(err)
	} else {
		quizViewmodel.Start(questions)
	}

	app.Run()
}
