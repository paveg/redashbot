package match

import "regexp"

func IsTextMatch(text string) bool {
	r := regexp.MustCompile(`https://.+/queries`)
	return r.MatchString(text)
}
