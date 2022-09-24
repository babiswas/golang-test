package Person

import "strconv"

type Person struct {
	name string
	id   int
}

type Job interface {
	Work() string
	Salary() int
}

func (p *Person) Work() string {
	return "working"
}

func (p *Person) Salary() int {
	return 15000
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
