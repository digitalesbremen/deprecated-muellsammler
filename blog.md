# Müllsammler

## Disclaimer

Das Projekt ist in einem frühen Stadium und noch in Arbeit. Je nach meiner freien Zeit passiert hier etwas. Ich habe 
von Golang wenig Ahnung und freue mich daher über jede Art von Feedback. 

## Meine Herausforderungen und Erfahrungen

Hier eine kleine Berichterstattung meiner Herausforderungen bzw. gemachten Erfahrungen

* Golang grundsätzlich verstehen.
* Feststellen, dass der [Abfallkalender der Bremer Stadtreinigung](http://213.168.213.236/bremereb/bify/index.jsp) kein valides HTML liefert.
* Feststellen, dass der [Abfallkalender der Bremer Stadtreinigung](http://213.168.213.236/bremereb/bify/index.jsp) statt UTF-8 in Windows-1252 codiert ist.
* Pointer spielen in golang wieder eine Rolle.
* Feststellen, dass es offensichtlich ok ist, ein `<h3>` mit einem `</h2>` zu schließen.
* Feststellen, dass die Endung der URL `/strasse.jsp?strasse=...` einerseits alle Straßen mit einem gewissen Anfangsbuchstaben anzeigt, aber auch (sofern es nur eine
Straße mit diesem Anfangsbuchstaben gibt, nur die Hausnummern der Straße). Zudem gibt es `http://213.168.213.236/bremereb/bify/hausnummer.jsp?strasse=...` 
um Hausnummern anzuzeigen. Verwirrend und schwierig. 
* Ich muss mich in Regex immer erst wieder einarbeiten. Dann macht es aber Spaß.
* Golang kann auch Funktionen als Übergabeparameter.
* `go build`führt keine Tests aus. `go test` führt keine Tests in sub modules aus.
* Feststellen, dass es `<nobr>28.05.&nbsp;Restmüll / Bioabfall</nobr>` zwar nicht valide ist, aber geliefert wird.
* Warum hat golang kein LocalDate?
* `width=60%` ist ein Problem, dass generell per Regex lösbar ist. Jetzt gibt es aber auch `width="0`. Da wird es mit einem generellem Regex interessant.

## Stand

* Die Indexseite wird geparsed und alle Anfangsbuchstaben aller Straßen werden geladen.
* Zu jedem Anfangsbuchstaben werden alle (ca 3700) Straßen geladen.
* Zu jeder Straße werden alle (ca 130000) Hausnummern geladen.
* Eine Progressbar im Log gibt den Ladezustand an. 
* Die Abholzeiten werden geladen
* Ein erster Import ist durch. Knapp 4,5 Stunden hat es gedauert. Ergebnis ist ein JSON-File von ca 1GB Größe.
* Ein zweiter Import benötigt knapp 3,5 Stunden. Hier werden nur die zukünftigen Daten gesammelt. Das JSON-File ist entsprechend nur 200MB groß.