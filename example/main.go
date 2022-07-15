package main

import (
	"encoding/json"
	"fmt"
)

type Email struct {
	Name    string
	Address string
}

func (e *Email) GetName() string {
	return e.Name
}

type Address struct {
	Street string
	Number int
}

type Person struct {
	ID      int64
	Name    string
	Address *Address
	Emails  []*Email
}

type User struct {
	ID   int64
	Name string
}

func main() {
	person := useNewBuilder()

	useToBuilder(person)

	useNewBuilderFrom(person)
}

func useNewBuilder() *Person {
	person1 := NewPersonBuilder().
		ID(10).
		Name("person").
		Address(
			NewAddressBuilder().
				Street("street").
				Number(99).
				Build(),
		).
		AddEmail(NewEmailBuilder().
			Name("email").
			Address("email@example.com").
			Build(),
		).
		Build()

	j, _ := json.MarshalIndent(person1, "", "  ")
	fmt.Printf("\n%s\n", string(j))

	return person1
}

func useToBuilder(person1 *Person) {
	person2 := person1.ToBuilder().
		ID(5).
		Name("changed").
		AddressFunc(func(a *Address) *Address {
			return a.ToBuilder().
				Street("changed").
				Number(55).
				Build()
		}).
		ForEachEmail(func(email *Email) *Email {
			return email.ToBuilder().
				Name("changed").
				Address("changed@example.com").
				Build()
		}).
		Build()

	j, _ := json.MarshalIndent([]*Person{person1, person2}, "", "  ")
	fmt.Printf("\n%s\n", string(j))
}

func useNewBuilderFrom(person *Person) {
	user := NewUserBuilderFrom(person).Build()

	j, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("\n%s\n", string(j))
}
