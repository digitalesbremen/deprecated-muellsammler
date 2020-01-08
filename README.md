# Bremer Abfallkalender

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://travis-ci.org/larmic/bremen_abfall_kalender.svg?branch=master)](https://travis-ci.org/larmic/bremen_abfall_kalender)

## Hintergrund

Mich stört es jede Woche, dass ich vergessen habe, welche Tonnen ich letzte Woche an die Straße gestellt habe. Um zu Wissen,
welche Tonnen diese Woche raus müssen, muss ich mich mühseelig durch `http://213.168.213.236/bremereb/bify/index.jsp` klicken.

Ich habe bisher wenig Erfahrung mit der Sprache Golang machen dürfen. Ich versuche hier etwas zu entwickeln, was mir jede 
Woche direkt sagen kann, welche Tonnen heute an die Straße müssen

#### Hinweis

Das Projekt ist in einem sehr frühen Stadium und noch in Arbeit. 

## Projekt

Dieses Projekt ist in Golang geschrieben und dient zunächst als reines
Backend. Ob zukünftig ein Web- oder App-Client hinzukommt, wird die Zeit zeigen.

### Anforderungen

* Golang 1.13

### Project bauen und starten

```ssh
$ git clone https://github.com/larmic/bremen_abfall_kalender
$ go build
$ ./bremen_trash
```
