package game

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/rivo/tview"
)

type quizViewModel struct {
	app             *tview.Application
	maingrid        *tview.Grid
	header          *tview.TextView
	quizGrid        *tview.Grid
	currentQuestion *tview.TextView
	options         *tview.List
}

const initialText = `[green]Quiz [white]Game!

- You will be offered a set of questions. Try and answer these, to achieve [red]victory[white]!`

func newQuizViewModel(app *tview.Application) *quizViewModel {
	maingrid := tview.NewGrid().SetRows(3, 0, 3).SetBorders(true)

	header := tview.NewTextView().SetTextAlign(tview.AlignLeft).SetDynamicColors(true)
	header.SetBorderPadding(0, 0, 5, 0)
	header.SetText(initialText)

	maingrid.AddItem(header, 0, 0, 1, 1, 0, 0, false)

	quizGrid := tview.NewGrid().SetRows(1, 0).SetBorders(true)

	currentQuestion := tview.NewTextView().SetDynamicColors(true)

	options := tview.NewList()

	quizGrid.AddItem(currentQuestion, 0, 0, 1, 1, 0, 0, false)
	quizGrid.AddItem(options, 1, 0, 1, 1, 0, 0, true)

	maingrid.AddItem(quizGrid, 1, 0, 1, 1, 0, 0, true)

	app.SetRoot(maingrid, true).SetFocus(maingrid)

	return &quizViewModel{
		app,
		maingrid,
		header,
		quizGrid,
		currentQuestion,
		options,
	}
}

func (qvm *quizViewModel) Start(questions []Question) {
	qvm.currentQuestion.SetText("Are you [red]ready[white]?")

	qvm.options.
		Clear().
		AddItem("[green]Yes", "To start the game", 'y', func() {
			qvm.ServeQuestions(questions)
		}).
		AddItem("[red]No", "To decline the generous offer", 'n', qvm.app.Stop)

	qvm.app.SetFocus(qvm.maingrid)
}

func (qvm *quizViewModel) ServeQuestions(questions []Question) {
	if len(questions) == 0 {
		qvm.NoQuestionsWarning()
		return
	}

	qvm.showQuestion(questions, 0)
}

func (qvm *quizViewModel) showQuestion(questions []Question, next int) {
	qvm.options.Clear()

	if next >= len(questions) {
		qvm.Success(len(questions))
		return
	}

	question := questions[next]

	qvm.currentQuestion.SetText(question.Title)

	for i, opt := range question.Options {
		qn := i
		qvm.options.AddItem(opt, strconv.Itoa(qn+1), rune(strconv.Itoa(qn+1)[0]), func() {
			if slices.Contains(question.CorrectOptions, qn+1) {
				qvm.showQuestion(questions, next+1)
			} else {
				qvm.Fail()
			}
		})
	}
}

func (qvm *quizViewModel) Fail() {
	qvm.options.Clear()
	qvm.currentQuestion.SetText("[red]Faileeeed!!!!! [teal]Arrrrarrr [white]Now you have to restart the quiz to try again!")
}

func (qvm *quizViewModel) NoQuestionsWarning() {
	qvm.options.Clear()
	qvm.currentQuestion.SetText("The source contains [black]no [white]questions!")
}

func (qvm *quizViewModel) Success(questionsCount int) {
	qvm.currentQuestion.SetText(
		fmt.Sprintf("[green]CONGRATULATIONS! You did it! You solved %d question(s)!", questionsCount),
	)
}

func (qvm *quizViewModel) ShowError(err error) {
	qvm.options.Clear()
	qvm.currentQuestion.SetText(fmt.Sprintf("[orange]Error! [white]<[red]%s[white]>", err))
}
