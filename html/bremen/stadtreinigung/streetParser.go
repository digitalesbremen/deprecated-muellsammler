package stadtreinigung

import (
	"fmt"
	"regexp"
	"strings"
)

type Street struct {
	Name string
	Url  string
}

func ParseStreetPage(content string, firstLetter FirstLetter) ([]Street, error) {
	streets := make([]Street, 0)

	if strings.Contains(content, `<!-- BEGIN: Keine Strassen gefunden:-->`) {
		return streets, fmt.Errorf("Page does not contains streets")
	}

	if strings.Contains(content, `bitte w&auml;hlen Sie die Hausnummer:`) {
		regex := regexp.MustCompile(`<!-- BEGIN: Strassen gefunden:-->\s*<h3>([` + firstLetter.FirstLetter + `][a-zA-Z-]*)<\/h2>`)
		submatch := regex.FindStringSubmatch(content)
		streetName := submatch[1]
		streetUrl := firstLetter.Url
		street := Street{streetName, streetUrl}
		streets = append(streets, street)
		return streets, nil
	}

	return streets, nil
}
