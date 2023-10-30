package util

import (
	"strings"
)

func FindMatchingStrings(targets []string, q string) []string {
	q = strings.ReplaceAll(q, "　", " ")
	qWords := strings.Fields(q)

	if q == "" || len(qWords) == 0 {
		return targets
	}

	matchingStrings := []string{}

	for _, target := range targets {
		matchesAllWords := true
		for _, word := range qWords {
			if !strings.Contains(target, word) {
				matchesAllWords = false
				break
			}
		}
		if matchesAllWords {
			matchingStrings = append(matchingStrings, target)
		}
	}

	return matchingStrings
}