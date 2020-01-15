package stadtreinigung

import "testing"

var (
	collectionDateSample = `
				<!-- Start Titel Jahr 2018-->
				<tr>
				  <td colspan="2" class="bakY">
				    <b>2018</b>
				  </td>
				</tr>
				<!-- End Titel Jahr 2018-->
				<!-- Start Inhalt Termine Jahr 2018 -->
				<tr>
					<td valign='top'>
						<b>Juli 2018</b>
						<br>
						<nobr>05.07.&nbsp;Papier / Gelber Sack</nobr>
						<br>
						<nobr>12.07.&nbsp;Restmüll / Bioabfall</nobr>
						<br>
					</td>
					<td valign='top'>
						<b>Juni 2019</b>
						<br>
						<nobr>(Sa) 01.06.&nbsp;Restm. / Bioabf.</nobr>
						<br>
						<nobr>06.06.&nbsp;Papier / Gelber Sack</nobr>
						<br>
					</td>
				</tr>
				<!-- Start Titel Jahr 2020-->
				<tr>
				  <td colspan="2" class="bakY">
				    <b>2020</b>
				  </td>
				</tr>
				<!-- End Titel Jahr 2020-->
				<!-- Start Inhalt Termine Jahr 2020 -->
				<tr>
					<td valign='top'>
						<b>Januar 2020</b>
						<br>
						<nobr>11.01.&nbsp;Tannenbaumabfuhr</nobr>
						<br>
					</td>
					<td valign='top'>
						<b>Mai 2020</b>
						<br>
						<nobr>(Sa) 23.05.&nbsp;Papier / G.Sack</nobr>
						<br>
						<nobr>28.05.&nbsp;Restmüll / Bioabfall</nobr>
						<br>
					</td>
				</tr>`
)

func TestParseGarbageCollectionDates(t *testing.T) {
	collectionDates := ParseGarbageCollectionDates(collectionDateSample)

	if collectionDates[0].Date != `05.07.` {
		t.Errorf(`ParseGarbageCollectionDates(sample)[%d].Type = %s; want %s`, 0, collectionDates[0].Date, `05.07.`)
	}
	if collectionDates[0].Type != `Papier / Gelber Sack` {
		t.Errorf(`ParseGarbageCollectionDates(sample)[%d].Type = %s; want %s`, 0, collectionDates[0].Type, `Papier / Gelber Sack`)
	}
}
