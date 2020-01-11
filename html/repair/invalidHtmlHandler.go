package repair

import (
	"regexp"
)

var (
	// A bit of hacking
	// -> missingQuotesRegex cannot handle content="text/html; charset=ISO-8859-1"
	// -> so replace it by removing white space within double quotes content="text/html;charset=ISO-8859-1"
	metaTagContentRegex = regexp.MustCompile(`(="[\w.+%?\-/;=]+)(;\s)([\w.+%?\-/;=]+")`)
	missingQuotesRegex  = regexp.MustCompile(`\s([a-zA-Z]+)=([\w.+%?\-=]+)`)
)

func RepairInvalidHtml(html string) string {
	return missingQuotesRegex.ReplaceAllString(metaTagContentRegex.ReplaceAllString(html, "$1;$3"), ` $1="$2"`)
}
