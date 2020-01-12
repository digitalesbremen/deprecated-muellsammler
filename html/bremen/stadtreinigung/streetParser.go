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
	var err error = nil

	if strings.Contains(content, `<!-- BEGIN: Keine Strassen gefunden:-->`) {
		err = fmt.Errorf("Page does not contains streets")
	} else if strings.Contains(content, `bitte w&auml;hlen Sie die Hausnummer:`) {
		regex := regexp.MustCompile(`<!-- BEGIN: Strassen gefunden:-->\s*<h[0-9]>([` + firstLetter.FirstLetter + `][a-zA-Z-]*)<\/h[0-9]>`)
		submatch := regex.FindStringSubmatch(content)
		streetName := submatch[1]
		streetUrl := firstLetter.Url
		street := Street{streetName, streetUrl}
		streets = append(streets, street)
	} else {

	}

	return streets, err
}
