package main

import (
	"regexp"
)

func GetCommitHashes(log string) []string {

	hashRegex := regexp.MustCompile(`commit ([^\s]+)`)

	matches := hashRegex.FindAllStringSubmatch(log, -1)

	var hashes []string
	for _, match := range matches {
		if len(match) > 1 {
			hashes = append(hashes, match[1])
		}
	}

	return hashes
}
