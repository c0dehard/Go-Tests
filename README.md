### Meine Meinung / Wissen zu GO

##### Version 1

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

**<small>(Generell hat es den Vorteil sich auf eine Convention festzulegen alleine und im Team umso mehr)</small>**

___

### Das schöne an Go

Variablen können auf mehrere Weisen deklariert werden..

Keine Objekt Orientierte Sprache, nichts mit Vererbung, kein Casting.

Läuft nicht auf einer virtuellen Maschine, das sind die tatsächlichen Adressen

die da benutzt werden.



Daher sieht direkt woran man ist, durch Integrität sogar gezwungen, sauber

und nicht verschwenderisch zu arbeiten.



Der Fokus wird hierbei auf ein Programm gelegt das gut durchdacht/stukturiert ist und seine Arbeit macht.

Zu viel Kreative möglichkeiten können einem selbst oftmals nur in ein tiefes Loch reißen.

**"Die Qual der Wahl"**

___

### Beispiele

Soll eine Variable noch nichts enthalten, dann am besten **var**  nehmen für einen garantierten **Null-Pointer** verweis, somit lassen sich viele Speicherkosten von vornerein vermeiden.

Will man einen bei Begin schon festen Wert zuweisen, empfiehlt sich die **shorthand notation**.

```go
	// Will man 100 % den 0Pointer Verweis
	// so sollte man besser so schreiben:
	var hatSicherNull int
	// der int hat dann gleich nil 
	var isSicherEmpty string 
	// der string hat dann "" also empty nicht gleich nil
  // ein string besteht aus 2 Words, der erste ist ein Pointer
  // auf ein Array im Hintergrund, das Zweite die Länge

	// Kurznotation
	eins,zwei := 1,2
	drei := 3
	vier := 4

	// Garantiert NUll-Pointer verweis zu 100%
	var keineAhnungLeer string

	// Es gibt auch platzhalter, kann ja sein das man bei mehreren Rückgaben
	// nur einen Wert benötigt, dann verweist man  mit _ auf man verzichtet
	// auf den Wert an der Stelle? muss googeln
	_,fuenf := methodeMitZweiReturnsZb()

```

