package main

type Animal interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d *Dog) Speak() string {
	return d.name + " barks"
}

func Act(animal Animal) {
	println(animal.Speak())
}

func main() {
	dog := Dog{"Bob"}
	Act(&dog)
}
