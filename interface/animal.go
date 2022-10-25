package _interface

type Animal interface {
	Eat()
	Run()
}
type Cat struct {
	name string
}

func (c Cat) Run() {
	println(c.name, "跑")
}

type Dog struct {
	name string
}

func (c Cat) Eat() {
	println(c.name, "吃")
}

func (d Dog) Eat() {
	println(d.name, "吃")
}

func TestInterface() {
	//var a Animal
	//a = Cat{name: "猫咪"}
	//a.Eat()
	//a.Run()
	//
	//AbstractFun(a)
}

func AbstractFun(a Animal) {
	println("AbstractFun")
	a.Run()
	a.Eat()
}
