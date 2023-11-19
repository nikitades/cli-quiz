package game

import (
	"fmt"
	"slices"
	"strconv"
	"time"

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

- You will be offered a set of questions. Try and answer these, to achieve [red]victory[white]!
- You will be given 30 seconds to answer each question, the game ends otherwise`

func newQuizViewModel(app *tview.Application) *quizViewModel {
	maingrid := tview.NewGrid().SetRows(5, 0, 3).SetBorders(true)

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
		qvm.ShowError(fmt.Errorf("The source contains no questions!"))
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

	enoughChan := make(chan int)
	timeChan := make(chan int)

	go func() {
		for {
			select {
			case <-enoughChan:
				return
			case timeChan <- 0:
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		timeLimit := 30
		for {
			select {
			case <-timeChan:
				qvm.currentQuestion.SetText(fmt.Sprintf("%v (%d)", question.Title, timeLimit))
				qvm.app.ForceDraw()
				timeLimit--
				if timeLimit <= 0 {
					enoughChan <- 0
					close(enoughChan)
					qvm.showQuitOption()
					qvm.currentQuestion.SetText("Time is up! 👺")
					qvm.app.ForceDraw()
				}
			case <-enoughChan:
				return
			}
		}
	}()

	for i, opt := range question.Options {
		qn := i
		qvm.options.AddItem(opt, strconv.Itoa(qn+1), rune(strconv.Itoa(qn + 1)[0]), func() {
			if slices.Contains(question.CorrectOptions, qn+1) {
				enoughChan <- 0
				qvm.showQuestion(questions, next+1)
			} else {
				enoughChan <- 0
				qvm.Fail()
			}
		})
	}
}

func (qvm *quizViewModel) Fail() {
	qvm.options.Clear()
	qvm.currentQuestion.SetText("[red]Faileeeed!!!!! [teal]Arrrrarrr [white]Now you have to restart the quiz to try again!")
}

func (qvm *quizViewModel) Success(questionsCount int) {
	qvm.currentQuestion.SetText(
		fmt.Sprintf("[green]CONGRATULATIONS! You did it! You solved %d question(s)!", questionsCount),
	)
	qvm.showQuitOption()
}

func (qvm *quizViewModel) ShowError(err error) {
	qvm.options.Clear()
	qvm.currentQuestion.SetText(fmt.Sprintf("[orange]Error! [white]<[red]%s[white]>", err))
}

func (qvm *quizViewModel) showQuitOption() {
	qvm.options.Clear().AddItem("Press Q to quit", "exits the app", 'q', qvm.app.Stop)
}
