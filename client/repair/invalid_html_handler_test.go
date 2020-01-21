package repair

import "testing"

var (
	//
	dataSets = [13][2]string{
		{` href=hausnummer.jsp?strasse=Landsberger+Stra%DFe`, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`},   // invalid -> must fixed
		{` href=hausnummer.jsp?strasse=Landsberger+Stra%DFe `, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe" `}, // invalid -> must fixed
		{`<nobr>05.07.&nbsp;Papier / Gelber Sack<nobr>`, `<nobr>05.07.&nbsp;Papier / Gelber Sack</nobr>`},              // invalid -> must fixed
		{` width=60%`, ` width="60%"`}, // invalid -> must fixed
		{` content="text/html; charset=ISO-8859-1"`, ` content="text/html;charset=ISO-8859-1"`},                        // valid -> but trim within
		{` href='hausnummer.jsp?strasse=Landsberger+Stra%DFe'`, ` href='hausnummer.jsp?strasse=Landsberger+Stra%DFe'`}, // valid -> won't fixed
		{` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`}, // valid -> won't fixed
		{` name="hello world"`, ` name="hello world"`},                                                                 // valid -> won't fixed
		{`<br>`, ``},       // cannot handled ->  must fixed
		{`</br>`, ``},      // cannot handled ->  must fixed
		{`<h2>`, `<h3>`},   // cannot handled ->  must fixed
		{`</h2>`, `</h3>`}, // cannot handled ->  must fixed
		{` -- `, ` `},      // cannot handled ->  must fixed
	}
)

func TestAddMissingDoubleQuoteIfNecessary(t *testing.T) {
	for index, row := range dataSets {
		input := row[0]
		want := row[1]
		got := RepairInvalidHtml(input)
		if got != want {
			t.Errorf(`Dataset: %d: RepairInvalidHtml(%s) = %s; want %s`, index, input, got, want)
		}
	}
}
