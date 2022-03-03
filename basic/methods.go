package basic

type PersonRule interface{}

func (p *Person) AddPerson(name string, age uint) Person {
	person := Person{}
	person.Name = name
	person.Age = age
	return person
}
