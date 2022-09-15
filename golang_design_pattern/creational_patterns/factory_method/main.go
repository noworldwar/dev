package main

import "fmt"

//
// Product Interface
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

//
// Concrete Product
type gun struct {
	name  string
	power int
}

type ak47 struct {
	gun
}

type musket struct {
	gun
}

//
// Entry Point
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

//
// Print The Product
func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

//
// Factory
func getGun(gunType string) (iGun, error) {

	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("Wrong Gun Type Passed")
	}

}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 Gun",
			power: 4,
		},
	}
}

func newMusket() iGun {
	return &musket{
		gun: gun{
			name:  "Musket Gun",
			power: 2,
		},
	}
}

//
// Properties
func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}
