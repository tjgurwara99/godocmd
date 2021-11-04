package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

func prepareFileInjection(file io.Reader, inject fmt.Stringer) ([]string, error) {
	// create a slice to hold the lines
	lines := []string{}

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	beginRegexp := regexp.MustCompile(`GODOCMD BEGIN`)
	endRegexp := regexp.MustCompile(`GODOCMD END`)

	// loop through the lines
	for scanner.Scan() {
		// append the line to the slice
		line := scanner.Text()
		match := beginRegexp.MatchString(line)
		lines = append(lines, line)
		if match {
			break
		}
	}

	stringReader := strings.NewReader(inject.String())
	injectScanner := bufio.NewScanner(stringReader)
	for injectScanner.Scan() {
		line := injectScanner.Text()
		lines = append(lines, line)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !endRegexp.MatchString(line) {
			continue
		}
		lines = append(lines, line)
	}

	// return the slice of lines
	return lines, nil
}

func writeToFile(file io.Writer, lines []string) error {
	// create a writer to write the file
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	// flush the writer
	err := writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
