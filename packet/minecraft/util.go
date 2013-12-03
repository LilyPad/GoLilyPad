package minecraft

import "strings"

func Colorize(input string) string {
	return strings.Replace(strings.Replace(input, "&", MAGIC, -1), MAGIC + MAGIC, "&", -1)
}