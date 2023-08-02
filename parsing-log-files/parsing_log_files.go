package parsinglogfiles

import (
	"fmt"
	"regexp"
)

const userTag string = "[USR]"

func IsValidLine(text string) bool {
	regex := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	regex := regexp.MustCompile(`<(~|\*|=|-)*>`)
	return regex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := regexp.MustCompile(`(?i)"(.|\s)*password"`)
	count := 0
	for _,line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(`end-of-line\d+`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	result := make([]string, 0, len(lines))
	regex := regexp.MustCompile(`User\s+(\w+)`)
	for _, line := range lines {
		if matches := regex.FindStringSubmatch(line); len(matches) > 1 {
			newLine := fmt.Sprintf("%s %s %s", userTag, matches[1], line)
			result = append(result, newLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}
