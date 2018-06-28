package fpgo

type person struct {
	name string
}

func createPerson(name string) person {
	return person{name: name}
}

func changeName(p *person, name string) {
	p.name = name
}

func cantChangeName(p person, name string) {
	p.name = name
}
