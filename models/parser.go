package models

import (
	"regexp"
	"strings"
)

const (
	rowDelimiter = "----------------------------------------------------------------------------"
	lineCode     = "\n"
)

var patternRex = regexp.MustCompile(`(?s)r(?P<Revision>[0-9]+) \| (?P<User>.*) \| (?P<Date>.*?)\s\s(?P<Comment>.*)\s`)

func ParseChangeLog(s *string) (*[][]string, error) {
	rows := strings.Split(formatLineCode(*s), rowDelimiter)
	parsedRows := [][]string{}
	for _, row := range rows {
		if row == "" {
			continue
		}
		parsedText, err := parseRow(&row)
		if err != nil {
			return nil, err
		}
		parsedRows = append(parsedRows, *parsedText)
	}
	return &parsedRows, nil
}

func formatLineCode(s string) string {
	return strings.NewReplacer(
		"\r\n", lineCode,
		"\r", lineCode,
		"\n", lineCode,
	).Replace(s)
}

func parseRow(row *string) (*[]string, error) {
	group := getGroupRegex(patternRex, row)
	result := []string{}
	result = append(result, group["Revision"])
	result = append(result, group["User"])
	result = append(result, group["Date"])
	result = append(result, group["Comment"])
	return &result, nil
}

func getGroupRegex(rex *regexp.Regexp, s *string) map[string]string {
	match := rex.FindStringSubmatch(*s)
	result := make(map[string]string)
	for i, name := range rex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
	return result
}
