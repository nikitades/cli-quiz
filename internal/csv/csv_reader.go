package csv

import (
	native_csv "encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/nikitades/cli-quiz/internal/game"
)

type CsvReader struct {
	data [][]string
}

// NewReader creates an instance of CsvReader. error is returned when the path is wrong
func NewReader(path string) (*CsvReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	nativeReader := native_csv.NewReader(f)
	records, err := nativeReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return &CsvReader{records}, nil
}

func (csvr *CsvReader) ReadAll() ([]game.Question, error) {
	output := make([]game.Question, 0, len(csvr.data))

	for _, line := range csvr.data {
		
		correctOptions := []int{}
		for _, opt := range strings.Split(line[2], ":") {
			intv, err := strconv.Atoi(opt)
			if err != nil {
				return []game.Question{}, err
			}
			correctOptions = append(correctOptions, intv)
		}
		
		output = append(output, game.Question{
			Title:          line[0],
			Options:        strings.Split(line[1], ":"),
			CorrectOptions: correctOptions,
		})
	}

	return output, nil
}
