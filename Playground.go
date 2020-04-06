// Nothing serious yet
package main

import (
	"fmt"
	"os"
)

type World struct {
	Parent Planet
	Name   string
}

type Planet struct {
	Size   int
	Radius int
	Name   string
}

type Person struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

type Jedi struct {
	Name   string
	Force  bool
	Phrase string
}

func (j *Jedi) addPhrase(newPhrase string) {
	j.Phrase = newPhrase
}


func (p Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func (p Planet) NtimesRadius() int {
	return p.Size * p.Radius
}

func (p *Planet) SetSize(size int) {
	p.Size = size
}

type DidYouKnow struct {
	name string
	age int
}

func NewDidYouKnow(name string,age int) *DidYouKnow{
	p := DidYouKnow{name:name,age:age}
	return &p
}

func whoIsFileOwner(yamlFile string) string {
	// nonsence to see how an anonymous function could be used
	// I know it ain't right 
	file, _ := os.Open(yamlFile)
	infos, _ := file.Stat()
	anonymous := struct {
		owner string
	}{
		owner: infos.Name(),
	}
	return anonymous.owner
}

func main() {
	var test Planet = Planet{Size: 4, Name: "Yavin"}
	fmt.Print(test.Name)

	print("\n- - - - \n")

	jedis := []string{"Obi-Wan", "Qui-Gon Jinn", "Yoda", "Luke"}
	for _, jedis := range jedis {
		fmt.Print(jedis, "\n")
	}
	
	y := Jedi{Name: "Yoda", Force: true}
	y.addPhrase("Patience you must have")
	fmt.Printf("%s: ,,%s, %t this is!''\n", y.Name, y.Phrase, y.Force)

	print("\n- - - - \n")

	p := Person{
		Name:    "c0deN",
		Surname: "c0dehard",
		Hobbies: []string{"Tech", "Computer Science", "Cartoons", "Games"},
		id:      "0123456789ABCDEF",
		// why Comma? A
	}

	fmt.Printf("%s likes %s, %s, %s and %s\n", p.GetFullName(), p.Hobbies[0], p.Hobbies[1], p.Hobbies[2], p.Hobbies[3])
	print("\n- - - - \n")
	dataOne := NewDidYouKnow("Unixtime",1970)
	dataTwo := "was born in"
	fmt.Printf("%s %s %d ",dataOne.name,dataTwo,dataOne.age)
	xyz := whoIsFileOwner("newPOST.yaml")
	print(xyz)
}
