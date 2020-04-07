package main

import "fmt"

type operatingSystem interface {
	name() string
	ownedBy() string
	deri() string
}

type gnuLinux struct {
	distribution,
	derivat,
	owner string
}
type macOS struct {
	distribution,
	derivat,
	owner string
}

func (g gnuLinux) name() string {
	return g.distribution
}
func (g gnuLinux) ownedBy() string {
	return g.owner
}
func (g gnuLinux) deri() string {
	return g.derivat
}

func (m macOS) name() string {
	return m.distribution
}
func (m macOS) ownedBy() string {
	return m.owner
}
func (m macOS) deri() string {
	return m.derivat
}

// Excuse me, the only windows I have, are not operating stystems

//

func getFullOSInfo(os operatingSystem) {
	fmt.Printf("Owner: %s\nDistribution: %s\nDerivat: %s\n", os.ownedBy(), os.name(), os.deri())
}

func main() {

	u := gnuLinux{distribution: "Debian", derivat: "Ubuntu", owner: "Canonical"}
	m := macOS{distribution: "BSD", derivat: "macOS Catalina", owner: "Apple"}

	getFullOSInfo(u)
	println("==========")
	getFullOSInfo(m)

}
