package xml

import "regexp"

var (
	regex = regexp.MustCompile(`\s([a-zA-Z]+)=([\w.+%?\-=]+)`)
)

func RepairInvalidHtml(html string) string {
	return regex.ReplaceAllString(html, ` $1="$2"`)
}
