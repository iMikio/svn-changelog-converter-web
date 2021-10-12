package models

import (
	"regexp"
	"strings"

	"github.com/beego/beego/v2/core/logs"
)

const (
	rowDelimiter = "----------------------------------------------------------------------------"
)

func ParseChangeLog(s *string) (*[][]string, error) {
	rows := strings.Split(*s, rowDelimiter)
	logs.Debug("rows:%v", rows)
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

func parseRow(row *string) (*[]string, error) {
	patternRex := regexp.MustCompile(`(?s)r(?P<Revision>[0-9]+) \| (?P<User>.*) \| (?P<Date>.*?)[\r\n]{2,}(?P<Comment>.*)\r\n`)
	group := getGroupRegex(patternRex, row)
	result := []string{}
	result = append(result, group["Revision"])
	result = append(result, group["User"])
	result = append(result, group["Date"])
	result = append(result, group["Comment"])
	return &result, nil
}

func getGroupRegex(rex *regexp.Regexp, s *string) map[string]string {
	logs.Debug("s:%s", *s)
	match := rex.FindStringSubmatch(*s)
	result := make(map[string]string)
	for i, name := range rex.SubexpNames() {
		logs.Debug("i:%v, name:%s", i, name)
		if i != 0 && name != "" {
			logs.Debug(match[i])
			result[name] = match[i]
		}
	}
	return result
}
