package Person

import "strconv"

type Person struct {
	name string
	id   int
}

func (p *Person) Set_name(name string) {
	p.name = name
}

func (p *Person) Set_id(id int) {
	p.id = id
}

func (p Person) Display() string {
	return p.name + " " + strconv.Itoa(p.id)
}
