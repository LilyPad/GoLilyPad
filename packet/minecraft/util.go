package minecraft

import (
	"strings"
)

func Colorize(in string) (out string) {
	out = strings.Replace(strings.Replace(in, "&", MAGIC, -1), MAGIC+MAGIC, "&", -1)
	return
}
