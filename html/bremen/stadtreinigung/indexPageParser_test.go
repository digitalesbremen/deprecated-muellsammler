package stadtreinigung

import (
	"testing"

	"bremen_trash/html/bremen/stadtreinigung/parser"
)

var (
	sample = `<tr>
				<td id="A" class = "BAKChr">
					<a href="strasse.jsp?strasse=A">A</a>
				</td>
    		  	<td id="Tag" class = "BAKChr">
					<a href="strasse.jsp?strasse=Tag">Tag</a>
				</td>
				<td id="" class = "BAKChr"></td>
    		  	<td id="AE" class = "BAKChr" align="center">
					<a href="strasse.jsp?strasse=%C4">&Auml;</a>
				</td>
    		  	<td id="OE" class = "BAKChr">
					<a href="strasse.jsp?strasse=%DC">&Uuml;</a>
				</td>
    		</tr>
    		<tr>
    		  	<td id="N" class = "BAKChr">
					<a href="strasse.jsp?strasse=N">N</a>
				</td>
    		  	<td id="O" class = "BAKChr">
					<a href="strasse.jsp?strasse=O">O</a>
				</td>
    		  	<td id="" class = "BAKChr"></td>
    		</tr> `
)

func TestParseIndexPage(t *testing.T) {
	firstLetters := ParseIndexPage(sample, "www.test.url/")

	assertFirstLetterValue(t, firstLetters, 0, `A`)
	assertFirstLetterUrlUrl(t, firstLetters, 0, `www.test.url/strasse.jsp?strasse=A`)
	assertFirstLetterValue(t, firstLetters, 1, `Tag`)
	assertFirstLetterUrlUrl(t, firstLetters, 1, `www.test.url/strasse.jsp?strasse=Tag`)
	assertFirstLetterValue(t, firstLetters, 2, `Ä`)
	assertFirstLetterUrlUrl(t, firstLetters, 2, `www.test.url/strasse.jsp?strasse=%C4`)
	assertFirstLetterValue(t, firstLetters, 3, `Ü`)
	assertFirstLetterUrlUrl(t, firstLetters, 3, `www.test.url/strasse.jsp?strasse=%DC`)
	assertFirstLetterValue(t, firstLetters, 4, `N`)
	assertFirstLetterUrlUrl(t, firstLetters, 4, `www.test.url/strasse.jsp?strasse=N`)
	assertFirstLetterValue(t, firstLetters, 5, `O`)
	assertFirstLetterUrlUrl(t, firstLetters, 5, `www.test.url/strasse.jsp?strasse=O`)
}

func assertFirstLetterValue(t *testing.T, firstLetters []parser.Dto, index int, want string) {
	if firstLetters[index].Value != want {
		t.Errorf(`ParseIndexPage(sample, test-root-url)[%d].FirstLetter = %s; want %s`, index, firstLetters[index].Value, want)
	}
}

func assertFirstLetterUrlUrl(t *testing.T, firstLetters []parser.Dto, index int, want string) {
	if firstLetters[index].Url != want {
		t.Errorf(`ParseIndexPage(sample, test-root-url)[%d].Url = %s; want %s`, index, firstLetters[index].Url, want)
	}
}
