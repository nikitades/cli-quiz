package start

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nikitades/cli-quiz/internal/csv"
	"github.com/nikitades/cli-quiz/internal/fake"
	"github.com/nikitades/cli-quiz/internal/game"
)

type Mode int

const (
	ModeCsv = iota
	ModeHardcode
)

type StartOptions struct {
	Mode
	CsvPath string
}

func Start(opts StartOptions) {

	var Reader game.QuestionReader
	switch opts.Mode {
	case ModeHardcode:
		Reader = &fake.FakeReader{}
	case ModeCsv:
		csvReader, err := csv.NewReader(opts.CsvPath)
		if err != nil {
			panic(err)
		}
		Reader = csvReader
	}

	game := &game.Game{Reader}

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	game.Run(screen)
}
