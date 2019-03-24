package main

import (
	"regexp"
)

func subtractPatternsFromFails(patterns, fails []string) []string {
	for f, fail := range fails {
		for _, pattern := range patterns {
			r := regexp.MustCompile(`([^\.])\*$`)
			fixedPattern := "^" + r.ReplaceAllString(pattern, "$1.*") + "$"
			rf := regexp.MustCompile(fixedPattern)

			if rf.MatchString(fail) {
				fails[f] = ""
				break
			}
		}
	}

	final := make([]string, 0)

	for index := 0; index < len(fails); index++ {
		if len(fails[index]) > 0 {
			final = append(final, fails[index])
		}
	}

	return final
}
