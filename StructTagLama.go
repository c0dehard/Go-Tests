package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Lama struct {
	LamaName  string //!`example:"name"` it ignores annotations when using encoding/json
	Password  string //? can only get Exported fields for automatic annotation, YEP!
	CreatedAt time.Time
}

func (l *Lama) String() string {
	return fmt.Sprintf("Hi! I'm %s", l.LamaName)
}

func main() {
	l := &Lama{
		LamaName: "Giraffe Sheep",
		// Funfact: It would take a computer about 9 QUINQUAGINTILLION YEARS to crack this fictive password and its not pwned :D
		Password:  base64.StdEncoding.EncodeToString([]byte("HappyLamaSadLamaMentallyDisturbedLamaSuperLamaDramaLamaBigFatMamaLama!*1")),
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
