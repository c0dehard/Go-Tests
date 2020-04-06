### Meine Meinung / Wissen zu GO

##### Version 1.1.3

<small>Version = (Haupt Version. Erweitert. Fehler behoben)</small>

(englische Variante folgt sofern diese hier Fertig ist)

```go
package main
import "fmt"

func getElement()string{
	return "Donner"
}

func main(){
	name, attacke := "Pikachu",getElement()
	fmt.Printf("%s's Element ist %d",name,power)
}
```

### Das hässliche an Go..

Variablen können auf mehrere Weisen deklariert werden..

Keine Objekt Orientierte Sprache, nichts mit Vererbung, kein Casting

> **<small>,,Ich habe heute leider kein Foto für dich´´</small>**

**<small>(Generell hat es immer einen Vorteil, sich auf eine Convention festzulegen. Egal ob alleine oder im Team, da sogar umso mehr.)</small>**

___

### ..ist gleichzeitig auch das schöne an Go

Variablen können auf mehrere Weisen deklariert werden..

Keine Objekt Orientierte Sprache, nichts mit Vererbung, kein Casting.

Läuft nicht auf einer virtuellen Maschine, das sind die tatsächlichen Adressen die da benutzt werden.

Direkt sehen woran man ist, durch Integrität sogar gezwungen, sauber

und nicht verschwenderisch zu arbeiten. In meinen Augen ein Pluspunkt.



Der Fokus richtet sich hierbei auf ein Programm, dass gut stukturiert ist und seine Arbeit schnell erledigt. Zu viel Kreative möglichkeiten können einem selbst oftmals nur in ein tiefes Loch reißen und man findet kein Ende beim Programmieren. **"Die Qual der Wahl"** fällt hier fast weg.

___

### Nice To Know

Soll eine Variable zu Begin keinen Wert enthalten eignet sich das Schlüsselwort **var** bestens, um einen 100% garantierten **Null-Pointer** Verweis zu erhalten, somit lassen sich viele Speicherkosten von vornherein schon vermeiden.

Will man einen bei Begin schon festen Wert zuweisen, empfiehlt sich die **shorthand notation**.

```go
	// Diesr Integer ist gleich nil
	// Garantiert NUll-Pointer Verweis zu 100%
	var hatSicherNull int
	
	// Dieser String ist gleich "" 
	// Vorsicht, empty ist nicht gleich nil!
  	// Ein String besteht aus 2 Words in Go
	// Das erste ist ein Pointer auf ein Array im Hintergrund
	// und das Zweite die Länge des Arrays.
	var isSicherEmpty string 


	// Kurznotation aka shorthand notation
	// Deklaration und Initialisierung durch :=
	// immer gut wenn man schon einen Wert zuweisen will
	// Go erkennt den Tyen automatisch, ist praktisch!
	eins,zwei := 1,2
	drei := 3
	vier := 4

	// Es gibt sogar platzhalter der sogenannte leerer Bezeichner.
	// Erlaubt Verwurft einzelner Werte bei multiplen Rückgaben
	// Beispiel: zu Testzwecken wird nur ein bestimmter Wert benötigt,
	// dann schreibt man einach wie folgt um nur den 2ten zu bekommen
	_,fuenf := methodeMitZweiReturnsZb()

```

___

### Strkturen Typen, Konvertierung und was man wissen sollte

```go
package main

import "fmt"

// Aufgrund von Padding-Bytes am besten die Felder der Größe nach notieren
// Optimierung von Performance

type albert struct {
	pi      float32
	counter int16
	flag    bool
}

type berta struct {
	pi      float32
	counter int16
	flag    bool
}

func main() {

	// Anonyme Struktur mit der shorthand notation
	e := struct {
		pi      float32
		counter int16
		flag    bool
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Ladies first, erstmal 2 Null-Pointer Variablen zum testen:
	// Obwohl berta und albert kompatibel sind..
	var b berta
	var a albert

	// ..sagt der Kompiler: ,,Ich macht keine implizierte Umwandlung, INTEGRITY first..!"
	// b = a
	// Er will uns nicht ärgern, sondern Errors vermeiden. Gibt man Ihm aber die klare
	// Absicht was man machen will, sprich in diesem Fall Konvertierung von passenden
	// Strukturen und direkt sagt:,, Ja, das ist ein Typ vom anderen!" Dann ist es OK.
	b = berta(a)

	// e ist keine Struktur mit Name sondern anonym, da meckert der Kompiler nicht rum
	// und übernimmt das Konvertieren sofern die Felder der Typen kompatibel sind.
	e = b

	fmt.Println(b, a, e)
```

### Fazit 

- Eine Struktur kann named typed oder anonymous typed sein.

- Wenn es auf einem Namenstyp basiert, müssen wir explizit sein.

- Wir müssen unsere Absicht zeigen, wenn beide kompatibel sind also die selben Felder haben ist es dann auch kein Problem.

- Handelt es sich nicht um anonyme Strukturen in die man Konvertieren will, dann kümmert sich der Kompiler darum.

- Durch Integrität ist man gezwungen Variablen auch zu benutzen, das ist sehr sinnvoll.

Aber um hier erstmal mögliche Mimimis zu vermeiden, habe man auch einfach vorerst sinnlos Printen :)
