package cliapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mateuszmidor/FlightFinder/pkg/application"
	"github.com/mateuszmidor/FlightFinder/pkg/infrastructure/csv"
)

func Run(flights_data_dir string) {
	const promptMsg = "Try: krk gdn. For exit: exit"
	const maxSegmentCount = 2

	repo := csv.NewFlightsDataRepoCSV(flights_data_dir)
	finder := application.NewConnectionFinder(repo)
	renderer := NewPathRendererAsText(os.Stdout, "\n")

	fmt.Println(promptMsg)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if line == "exit" {
			fmt.Println("exiting now.")
			break
		}

		if from, to, ok := parseFromTo(line); ok {
			fmt.Println("working...")
			finder.Find(strings.ToUpper(from), strings.ToUpper(to), maxSegmentCount, renderer)
			fmt.Println("\ndone.")
		} else {
			fmt.Println(promptMsg)
		}
	}
}

func parseFromTo(line string) (from string, to string, ok bool) {
	_, err := fmt.Sscanf(line, "%s %s\n", &from, &to)
	if err != nil {
		return "", "", false
	}
	return from, to, true
}
