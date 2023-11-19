package start

import (
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/nikitades/cli-quiz/internal/fake"
	"github.com/nikitades/cli-quiz/internal/game"
	"github.com/stretchr/testify/assert"
)

func dumpString(cells []tcell.SimCell, x int, y int) string {
	str := []rune{}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			str = append(str, cells[x*i+j].Runes...)
		}
		str = append(str, rune('\n'))
	}

	return string(str)
}

func TestGameInitialScreen(t *testing.T) {
	game := &game.Game{&fake.FakeReader{}}

	fakeScreen := tcell.NewSimulationScreen("UTF-8")

	go game.Run(fakeScreen)
	time.Sleep(50 * time.Millisecond)
	content := dumpString(fakeScreen.GetContents())

	expected := `
┌──────────────────────────────────────────────────────────────────────────────┐
│     Quiz Game!                                                               │
│                                                                              │
│     - You will be offered a set of questions. Try and answer these, to       │
├──────────────────────────────────────────────────────────────────────────────┤
│┌────────────────────────────────────────────────────────────────────────────┐│
││Are you ready?                                                              ││
│├────────────────────────────────────────────────────────────────────────────┤│
││(y) Yes                                                                     ││
││    To start the game                                                       ││
││(n) No                                                                      ││
││    To decline the generous offer                                           ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
│└────────────────────────────────────────────────────────────────────────────┘│
└──────────────────────────────────────────────────────────────────────────────┘
                                                                                
                                                                                
                                                                                
                                                                                
`[1:]

	assert.Equal(t, expected, content)
}

func TestWhenUserAgreesGameStarts(t *testing.T) {
	game := &game.Game{&fake.FakeReader{}}

	fakeScreen := tcell.NewSimulationScreen("UTF-8")

	go game.Run(fakeScreen)
	time.Sleep(time.Millisecond * 50)
	fakeScreen.InjectKeyBytes([]byte("y"))
	time.Sleep(time.Millisecond * 50)
	fakeScreen.Show()

	content := dumpString(fakeScreen.GetContents())

	expected := `
┌──────────────────────────────────────────────────────────────────────────────┐
│     Quiz Game!                                                               │
│                                                                              │
│     - You will be offered a set of questions. Try and answer these, to       │
├──────────────────────────────────────────────────────────────────────────────┤
│┌────────────────────────────────────────────────────────────────────────────┐│
││First q                                                                     ││
│├────────────────────────────────────────────────────────────────────────────┤│
││(1) Right                                                                   ││
││    1                                                                       ││
││(2) Wrong                                                                   ││
││    2                                                                       ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
│└────────────────────────────────────────────────────────────────────────────┘│
└──────────────────────────────────────────────────────────────────────────────┘
                                                                                
                                                                                
                                                                                
                                                                                
`[1:]

	assert.Equal(t, expected, content)
}

func TestWhenUserAnswersWrongGameEnds(t *testing.T) {
	game := &game.Game{&fake.FakeReader{}}

	fakeScreen := tcell.NewSimulationScreen("UTF-8")

	go game.Run(fakeScreen)
	time.Sleep(time.Millisecond * 50)
	fakeScreen.InjectKeyBytes([]byte("y"))
	time.Sleep(time.Millisecond * 50)
	fakeScreen.InjectKeyBytes([]byte("2"))
	time.Sleep(time.Millisecond * 50)
	fakeScreen.Show()

	content := dumpString(fakeScreen.GetContents())

	expected := `
┌──────────────────────────────────────────────────────────────────────────────┐
│     Quiz Game!                                                               │
│                                                                              │
│     - You will be offered a set of questions. Try and answer these, to       │
├──────────────────────────────────────────────────────────────────────────────┤
│┌────────────────────────────────────────────────────────────────────────────┐│
││Faileeeed!!!!! Arrrrarrr Now you have to restart the quiz to try again!     ││
│├────────────────────────────────────────────────────────────────────────────┤│
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
││                                                                            ││
│└────────────────────────────────────────────────────────────────────────────┘│
└──────────────────────────────────────────────────────────────────────────────┘
                                                                                
                                                                                
                                                                                
                                                                                
`[1:]

	assert.Equal(t, expected, content)
}
