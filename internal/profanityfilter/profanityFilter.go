package profanityfilter

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"slices"
	"strings"
)

var rxBlacklistedWords *regexp.Regexp
var rxWhitelistedWords *regexp.Regexp

func SanitizeMessage(message string) string {
	findAllWhitelist := rxWhitelistedWords.FindAllString(message, -1)
	findAllBlacklist := rxBlacklistedWords.FindAllString(message, -1)
	for _, blackListedWord := range findAllBlacklist {
		if slices.Contains(findAllWhitelist, blackListedWord) {
			continue
		}
		var censorBuilder strings.Builder
		for i := 0; i < len(blackListedWord); i++ {
			censorBuilder.WriteString("*")
		}
		message = strings.ReplaceAll(message, blackListedWord, censorBuilder.String())
	}
	return message
}

func InitProfanityFilter() {
	blacklistFile, err := os.Open("chatwatcher-blacklist.txt")
	if err != nil {
		blacklistFile, err = os.Create("chatwatcher-blacklist.txt")
		if err != nil {
			panic(err)
		}
	}
	rxBlacklistedWords = compilePattern(blacklistFile)
	whitelistFile, err := os.Open("chatwatcher-whitelist.txt")
	if err != nil {
		whitelistFile, err = os.Create("chatwatcher-whitelist.txt")
		if err != nil {
			panic(err)
		}
	}
	rxWhitelistedWords = compilePattern(whitelistFile)
}

func compilePattern(reader io.Reader) *regexp.Regexp {
	var stringBuilder strings.Builder
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "#") {
			continue
		}
		if stringBuilder.Len() <= 0 {
			stringBuilder.WriteString("(" + text)
		} else {
			stringBuilder.WriteString(")|(" + text)
		}
	}
	if stringBuilder.Len() > 0 {
		stringBuilder.WriteString(")")
	}
	return regexp.MustCompile("(?i)" + stringBuilder.String())
}
