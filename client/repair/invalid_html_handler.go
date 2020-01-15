package repair

import (
	"regexp"
	"strings"
)

var (
	// A bit of hacking
	// -> missingQuotesRegex cannot handle content="text/html; charset=ISO-8859-1"
	// -> so replace it by removing white space within double quotes content="text/html;charset=ISO-8859-1"
	metaTagContentRegex = regexp.MustCompile(`(="[\w.+%?\-/;=]+)(;\s)([\w.+%?\-/;=]+")`)
	missingQuotesRegex  = regexp.MustCompile(`\s([a-zA-Z]+)=([\w.+%?\-=]+)`)
	closeNobrTagRegex   = regexp.MustCompile(`(<nobr>)(.*)(<nobr>)`)
)

func RepairInvalidHtml(html string) string {
	repairedHtml := metaTagContentRegex.ReplaceAllString(html, "$1;$3")
	repairedHtml = missingQuotesRegex.ReplaceAllString(repairedHtml, ` $1="$2"`)
	repairedHtml = closeNobrTagRegex.ReplaceAllString(repairedHtml, `$1$2</nobr>`)

	// Hack: Don't now why but parsing </br> does not work.
	repairedHtml = strings.ReplaceAll(repairedHtml, "<br>", "")
	repairedHtml = strings.ReplaceAll(repairedHtml, "</br>", "")

	// Hack: Fix <h3> ends with </h2>
	repairedHtml = strings.ReplaceAll(repairedHtml, "<h2>", "<h3>")
	repairedHtml = strings.ReplaceAll(repairedHtml, "</h2>", "</h3>")

	return repairedHtml
}
