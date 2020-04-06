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

	/*  Fazit eine Struktur kann named typed oder anonymous typed sein.
	Wenn es auf einem Namenstyp basiert, müssen wir explizit sein.
	Wir müssen unsere Absicht zeigen, wenn beide kompatibel sind also die selben Felder haben
	ist es dann auch kein Problem.

	Handelt es sich nicht um anonyme Strukturen	in die man Konvertieren will,
	dann kümmert sich der Kompiler darum.
	*/

	// Durch Integrität ist man gezwungen Variablen auch zu benutzen, das ist sehr sinnvoll.
	// Aber um hier erstmal mögliche Mimimis zu vermeiden, einfach sinnlos Printen :)
	fmt.Println(b, a, e)
}
