### Meine Meinung / Wissen zu GO

##### Version 1.0.2

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

### Das hässliche an Go

Variablen können auf mehrere Weisen deklariert werden..

Keine Objekt Orientierte Sprache, nichts mit Vererbung, kein Casting

> **<small>,,Ich habe heute leider kein Foto für dich´´</small>**

**<small>(Generell hat es den immer den Vorteil sich auf eine Convention festzulegen ob  alleine oder im Team, da umso mehr.)</small>**

___

### ist gleichzeitig auch das schöne an Go

Variablen können auf mehrere Weisen deklariert werden..

Keine Objekt Orientierte Sprache, nichts mit Vererbung, kein Casting.

Läuft nicht auf einer virtuellen Maschine, das sind die tatsächlichen Adressen die da benutzt werden.

Direkt sehen woran man ist, durch Integrität sogar gezwungen, sauber

und nicht verschwenderisch zu arbeiten. In meinen Augen ein Pluspunkt.



Der Fokus richtet sich hierbei auf ein Programm, dass gut stukturiert ist und seine Arbeit schnell erledigt. Zu viel Kreative möglichkeiten können einem selbst oftmals nur in ein tiefes Loch reißen und man findet kein Ende beim Programmieren. **"Die Qual der Wahl"** fällt hier fast weg.

___

### Nice To Know

Soll eine Variable zu Begin keinen Wert enthalten eignet sich das Schlüsselwort **var** besten

um einen 100% garantierten **Null-Pointer** Verweis zu erhalten, somit lassen sich viele Speicherkosten von vornerein schon vermeiden.

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
