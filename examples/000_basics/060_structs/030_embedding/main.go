                
// Animal is embedded, if we add Animal Animal it will be only just a field
// Embedding
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Field named Animal with type Animal
type Cat struct {
	Animal   Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name) // we can access directly using field names of embedded struct

	// literal need to precise a type
	b2 := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		SpeedKPH: 48,
		CanFly:   false,
	}
	fmt.Println(b2)
	fmt.Println(b2.Name)

	c := Cat{}
	// c.Name = "Siamese" // c.Name undefined (type Cat has no field or method Name)
	c.Animal.Name = "Siamese"
	// c.Origin = "Thailand" // c.Origin undefined (type Cat has no field or method Origin)
	c.Animal.Origin = "Thailand"
	c.SpeedKPH = 100
	c.CanFly = false
	fmt.Println(c)
	fmt.Println(c.Animal.Name) // we can access directly using field names of embedded struct
}
