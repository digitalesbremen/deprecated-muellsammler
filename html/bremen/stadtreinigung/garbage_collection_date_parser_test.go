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

	assertCollectionDate(t, collectionDates, 0, `05.07.2018`)
	assertCollectionType(t, collectionDates, 0, YellowBag, PaperWaste)
	assertCollectionDate(t, collectionDates, 1, `12.07.2018`)
	assertCollectionType(t, collectionDates, 1, ResidualWaste, BioWaste)
	assertCollectionDate(t, collectionDates, 2, `01.06.2019`)
	assertCollectionType(t, collectionDates, 2, ResidualWaste, BioWaste)
	assertCollectionDate(t, collectionDates, 3, `06.06.2019`)
	assertCollectionType(t, collectionDates, 3, PaperWaste, YellowBag)
	assertCollectionDate(t, collectionDates, 4, `11.01.2020`)
	assertCollectionType(t, collectionDates, 4, ChristmasTree)
	assertCollectionDate(t, collectionDates, 5, `23.05.2020`)
	assertCollectionType(t, collectionDates, 5, PaperWaste, YellowBag)
	assertCollectionDate(t, collectionDates, 6, `28.05.2020`)
	assertCollectionType(t, collectionDates, 6, ResidualWaste, BioWaste)
}

func assertCollectionDate(t *testing.T, collectionDates []GarageCollection, index int, want string) {
	if collectionDates[index].Date != want {
		t.Errorf(`ParseGarbageCollectionDates(sample)[%d].WasteType = %s; want %s`, index, collectionDates[index].Date, want)
	}
}

func assertCollectionType(t *testing.T, collectionDates []GarageCollection, index int, want ...WasteType) {
	if len(collectionDates[index].Type) != len(want) {
		t.Errorf(`len(ParseGarbageCollectionDates(sample)[%d].WasteType) = %d; want %d`, index, len(collectionDates[index].Type), len(want))
	}
	for _, wantType := range want {
		if !contains(collectionDates[index].Type, wantType) {
			t.Errorf(`ParseGarbageCollectionDates(sample)[%d].WasteType does not contain %d`, index, int(wantType))
		}
	}
}

func contains(s []WasteType, e WasteType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
