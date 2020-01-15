package stadtreinigung

import (
	"testing"
)

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

	assertCollectionDate(t, collectionDates, 0, `05.07.`)
	assertCollectionType(t, collectionDates, 0, `Papier / Gelber Sack`)
	assertCollectionDate(t, collectionDates, 1, `12.07.`)
	assertCollectionType(t, collectionDates, 1, `Restmüll / Bioabfall`)
	assertCollectionDate(t, collectionDates, 2, `01.06.`)
	assertCollectionType(t, collectionDates, 2, `Restm. / Bioabf.`)
	assertCollectionDate(t, collectionDates, 3, `06.06.`)
	assertCollectionType(t, collectionDates, 3, `Papier / Gelber Sack`)
	assertCollectionDate(t, collectionDates, 4, `11.01.`)
	assertCollectionType(t, collectionDates, 4, `Tannenbaumabfuhr`)
	assertCollectionDate(t, collectionDates, 5, `23.05.`)
	assertCollectionType(t, collectionDates, 5, `Papier / G.Sack`)
	assertCollectionDate(t, collectionDates, 6, `28.05.`)
	assertCollectionType(t, collectionDates, 6, `Restmüll / Bioabfall`)
}

func assertCollectionDate(t *testing.T, collectionDates []GarageCollection, index int, want string) {
	if collectionDates[index].Date != want {
		t.Errorf(`ParseGarbageCollectionDates(sample)[%d].Type = %s; want %s`, index, collectionDates[index].Date, want)
	}
}

func assertCollectionType(t *testing.T, collectionDates []GarageCollection, index int, want string) {
	if collectionDates[index].Type != want {
		t.Errorf(`ParseGarbageCollectionDates(sample)[%d].Type = %s; want %s`, index, collectionDates[index].Type, want)
	}
}
