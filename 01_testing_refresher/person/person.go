package person

type PersonParams struct {
	FirstName string
	LastName  string
	Age       int
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func NewPerson(params PersonParams) *Person {
	return &Person{
		firstName: params.FirstName,
		lastName:  params.LastName,
		age:       params.Age,
	}
}

func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}
func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}
func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetFirstName() string {
	return p.firstName
}
func (p *Person) GetLastName() string {
	return p.lastName
}
func (p *Person) GetAge() int {
	return p.age
}
