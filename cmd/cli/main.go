package main

import (
	"flag"

	"github.com/nikitades/cli-quiz/internal/start"
)

func main() {
	mode := flag.String("mode", "h", "h for hardcode, csv for csv; default: h")
	csvPath := flag.String("csvpath", "", "absolute path to csv file; format: 1+1,2:2:7,1:1")

	flag.Parse()

	var startMode start.Mode

	if *mode == "h" {
		startMode = start.ModeHardcode
	}

	if *mode == "csv" {
		startMode = start.ModeCsv
	}

	start.Start(start.StartOptions{
		Mode:    startMode,
		CsvPath: *csvPath,
	})
}
