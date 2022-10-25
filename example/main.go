package main

import (
	"encoding/json"
	"fmt"

	"github.com/flrossetto/go-builder/example/types"
)

type AddressType string

type EmailType int

const (
	PrimaryEmail EmailType = iota
	PersonalEmail
)

type EmptyStruct struct{}

type AnyValue[T any] struct {
	Value T
}

type Email struct {
	Name      string
	Address   AddressType
	EmailType EmailType
	Types     []EmailType
	Interface interface{}
	Map1      map[string]string
	Map2      map[string]Address
	Map3      map[string]*Address
	Map4      map[*Address]*Address
	Map5      map[interface{}]interface{}
	Map6      *map[interface{}]interface{}
	Generic1  AnyValue[EmptyStruct]
	Generic2  []AnyValue[EmptyStruct]
	Generic3  AnyValue[[]EmptyStruct]
	Generic4  AnyValue[[]*map[interface{}]interface{}]
	Generic5  types.Values[types.Value]
}

func (e *Email) GetName() string {
	return e.Name
}

type Address struct {
	Street string
	Number int
}

type PersonAddress Address

type Person struct {
	ID      int64
	Name    string
	Address *PersonAddress
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
			NewPersonAddressBuilder().
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
		AddressFunc(func(a *PersonAddress) *PersonAddress {
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
