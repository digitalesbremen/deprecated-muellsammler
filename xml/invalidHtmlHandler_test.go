package xml

import "testing"

var (
	//
	dataSets = [5][2]string{
		{` href=hausnummer.jsp?strasse=Landsberger+Stra%DFe`, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`},   // invalid -> must fixed
		{` href=hausnummer.jsp?strasse=Landsberger+Stra%DFe `, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe" `}, // invalid -> must fixed
		{` width=60%`, ` width="60%"`}, // invalid -> must fixed
		//{` content="text/html; charset=ISO-8859-1"`, ` content="text/html; charset=ISO-8859-1"`}, 							// valid -> won't fixed
		{` href='hausnummer.jsp?strasse=Landsberger+Stra%DFe'`, ` href='hausnummer.jsp?strasse=Landsberger+Stra%DFe'`}, // valid -> won't fixed
		{` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`, ` href="hausnummer.jsp?strasse=Landsberger+Stra%DFe"`}, // valid -> won't fixed
	}
)

func TestAddMissingDoubleQuoteIfNecessary(t *testing.T) {
	for index, row := range dataSets {
		input := row[0]
		output := row[1]
		got := RepairInvalidHtml(input)
		if got != output {
			t.Errorf(`Dataset: %d: RepairInvalidHtml(%s) = %s; want %s`, index, input, got, output)
		}
	}
}
