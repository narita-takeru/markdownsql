package sag

import (
	"bufio"
	"os"
	"strings"
)

type SagMarkDown struct {
	OnOneLines   map[string]func(line string) error
	OnMultiLines map[string]func(lines []string) error
	OnTable      func(columns map[string]string) error
}

func (sag SagMarkDown) tableEach(line string, scanner *bufio.Scanner) error {
	tokens := strings.Split(line, `|`)
	tableColumns := make([]string, len(tokens))
	for i, name := range tokens {
		tableColumns[i] = name
	}

	scanner.Scan() // Skip bar line.

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, `|`) {
			return nil
		}

		tokens := strings.Split(line, `|`)
		tableValues := make(map[string]string, len(tokens))
		for i, token := range tokens {
			tableValues[tableColumns[i]] = strings.Trim(token, ` `)
		}

		if err := sag.OnTable(tableValues); err != nil {
			return err
		}
	}

	return nil
}

func (sag SagMarkDown) Start(filePath string) error {
	fp, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, `|`) {
			if err := sag.tableEach(line, scanner); err != nil {
				return nil
			}

			continue
		}

		for search, onOneLine := range sag.OnOneLines {
			tokens := strings.Split(line, ` `)
			if tokens[0] == search {
				onOneLine(tokens[1])
			}
		}
	}

	return nil
}
