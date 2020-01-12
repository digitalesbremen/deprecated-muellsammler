# Bremer Abfallkalender

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://travis-ci.org/larmic/bremen_abfall_kalender.svg?branch=master)](https://travis-ci.org/larmic/bremen_abfall_kalender)

## Hintergrund

Mich stört es jede Woche, dass ich vergessen habe, welche Tonnen ich letzte Woche an die Straße gestellt habe. Um zu Wissen,
welche Tonnen diese Woche raus müssen, muss ich mich mühselig durch den [Abfallkalender der Bremer Stadtreinigung](http://213.168.213.236/bremereb/bify/index.jsp) 
klicken.

Ich habe bisher wenig Erfahrung mit der Sprache Golang machen dürfen. Ich versuche hier etwas zu entwickeln, was mir jede 
Woche direkt sagen kann, welche Tonnen heute an die Straße können.

#### Hinweis

Das Projekt ist in einem frühen Stadium und noch in Arbeit. Je nach meiner freien Zeit passiert hier etwas. Ich habe 
von Golang wenig Ahnung und freue mich daher über jede Art von Feedback. 

#### Meine Herausforderungen und Erfahrungen

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
* `go build`führt keine Tests aus. `go test` führt keine Tests in submodules aus.

## Projekt

Dieses Projekt ist in Golang geschrieben und dient zunächst als reines
Backend. Ob zukünftig ein Web- oder App-Client hinzukommt, wird die Zeit zeigen.

### Stand

* Die Indexseite wird geparsed und alle Anfangsbuchstaben aller Straßen werden geladen.
* Zu jedem Anfangsbuchstaben werden alle (ca 3700) Straßen geladen.
* Zu jeder Straße werden alle (ca 130000) Hausnummern geladen.
* Eine Progressbar im Log gibt den Ladezustand an. Zur Zeit werden ungefährt 130.000 Request in 2m43s durchgeführt

### Anforderungen

* Golang 1.13

### Project bauen und starten

```ssh
$ git clone https://github.com/larmic/bremen_abfall_kalender
$ go build
$ ./bremen_trash
```
