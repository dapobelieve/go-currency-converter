package main

import "fmt"

type engine struct {
	fuel   int
	thrust int
	hp     int
}

func (e *engine) start() {
	fmt.Print("Engine has started!\n")
}

type automobile struct {
	make  string
	model string
	engine
	electric bool
}

func (a *automobile) drive() {
	fmt.Printf("Automobile %s %s is on the road\n", a.make, a.model)
}

type plane struct {
	automobile
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p *plane) fly() {
	fmt.Printf("Aircraft %s %s is clear for takeoff!\n", p.make, p.model)
	fmt.Println()
}

func (p *plane) land() {
	fmt.Printf("Aircraft %s %s will be landing now!\n", p.make, p.model)
}

func main() {
	benz := &automobile{
		make:  "Mercedes Benz",
		model: "c300",
		engine: engine{
			fuel:   30,
			thrust: 4000,
			hp:     400,
		},
		electric: false,
	}

	benz.start()
	benz.drive()

	audi := &automobile{}

	audi.make = "BMW"
	audi.model = "x5 Series"
	audi.fuel = 45
	audi.hp = 600
	audi.thrust = 400
	audi.electric = true

	audi.start()
	audi.drive()

	//Engine has started!
	//	Automobile Mercedes Benz c300 is on the road
	//Engine has started!
	//	Automobile BMW x5 Series is on the road

	plane := &plane{
		automobile: automobile{
			make:  "Boeing",
			model: "737",
			engine: engine{
				30000000, 4000, 50000,
			},
			electric: false,
		},
		engineCount: 2,
		fixedWings:  true,
		maxAltitude: 1209904,
	}

	plane.start()
	plane.fly()

	//Engine has started!
	//	Automobile Mercedes Benz c300 is on the road
	//Engine has started!
	//	Automobile BMW x5 Series is on the road
	//Engine has started!
	//	Aircraft Boeing 737 is clear for takeoff!

}
