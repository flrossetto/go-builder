// Code generated by go-builder. DO NOT EDIT.
// Source: main.go
// Timestamp: 2022-07-08T10:01:29-03:00

package main

import "github.com/jinzhu/copier"

type AddressBuilder struct {
	xStreet string
	xNumber int
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{}
}

func NewAddressBuilderFrom(a interface{}) *AddressBuilder {
	var b AddressBuilder
	_ = copier.Copy(&b, a)
	return &b
}

func (b *AddressBuilder) LoadValues(a interface{}) *AddressBuilder {
	_ = copier.Copy(b, a)
	return b
}

func (b *AddressBuilder) Street(v string) *AddressBuilder {
	b.xStreet = v
	return b
}

func (b *AddressBuilder) Number(v int) *AddressBuilder {
	b.xNumber = v
	return b
}

func (b *AddressBuilder) Build() *Address {
	return &Address{
		Street: b.xStreet,
		Number: b.xNumber,
	}
}

func (a *Address) ToBuilder() *AddressBuilder {
	return &AddressBuilder{
		xStreet: a.Street,
		xNumber: a.Number,
	}
}

func (a *Address) GetStreet() string {
	if a == nil {
		return ""
	}
	return a.Street
}

func (a *Address) GetNumber() int {
	if a == nil {
		return 0
	}
	return a.Number
}

type PersonBuilder struct {
	xID      int64
	xName    string
	xAddress *Address
	xEmails  []*Email
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{}
}

func NewPersonBuilderFrom(a interface{}) *PersonBuilder {
	var b PersonBuilder
	_ = copier.Copy(&b, a)
	return &b
}

func (b *PersonBuilder) LoadValues(a interface{}) *PersonBuilder {
	_ = copier.Copy(b, a)
	return b
}

func (b *PersonBuilder) ID(v int64) *PersonBuilder {
	b.xID = v
	return b
}

func (b *PersonBuilder) Name(v string) *PersonBuilder {
	b.xName = v
	return b
}

func (b *PersonBuilder) Address(v *Address) *PersonBuilder {
	b.xAddress = v
	return b
}

func (b *PersonBuilder) AddressFunc(f func(*Address) *Address) *PersonBuilder {
	b.xAddress = f(b.xAddress)
	return b
}

func (b *PersonBuilder) Emails(v []*Email) *PersonBuilder {
	b.xEmails = v
	return b
}

func (b *PersonBuilder) ClearEmails() *PersonBuilder {
	b.xEmails = nil
	return b
}

func (b *PersonBuilder) AddEmails(value *Email) *PersonBuilder {
	b.xEmails = append(b.xEmails, value)
	return b
}

func (b *PersonBuilder) ForEachEmail(f func(*Email) *Email) *PersonBuilder {
	list := make([]*Email, 0, len(b.xEmails))
	for _, item := range b.xEmails {
		list = append(list, f(item))
	}
	b.xEmails = list
	return b
}

func (b *PersonBuilder) Build() *Person {
	return &Person{
		ID:      b.xID,
		Name:    b.xName,
		Address: b.xAddress,
		Emails:  b.xEmails,
	}
}

func (a *Person) ToBuilder() *PersonBuilder {
	return &PersonBuilder{
		xID:      a.ID,
		xName:    a.Name,
		xAddress: a.Address,
		xEmails:  a.Emails,
	}
}

func (a *Person) GetID() int64 {
	if a == nil {
		return 0
	}
	return a.ID
}

func (a *Person) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *Person) GetAddress() *Address {
	if a == nil {
		return nil
	}
	return a.Address
}

func (a *Person) GetEmails() []*Email {
	if a == nil {
		return nil
	}
	return a.Emails
}

type UserBuilder struct {
	xID   int64
	xName string
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func NewUserBuilderFrom(a interface{}) *UserBuilder {
	var b UserBuilder
	_ = copier.Copy(&b, a)
	return &b
}

func (b *UserBuilder) LoadValues(a interface{}) *UserBuilder {
	_ = copier.Copy(b, a)
	return b
}

func (b *UserBuilder) ID(v int64) *UserBuilder {
	b.xID = v
	return b
}

func (b *UserBuilder) Name(v string) *UserBuilder {
	b.xName = v
	return b
}

func (b *UserBuilder) Build() *User {
	return &User{
		ID:   b.xID,
		Name: b.xName,
	}
}

func (a *User) ToBuilder() *UserBuilder {
	return &UserBuilder{
		xID:   a.ID,
		xName: a.Name,
	}
}

func (a *User) GetID() int64 {
	if a == nil {
		return 0
	}
	return a.ID
}

func (a *User) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

type EmailBuilder struct {
	xName    string
	xAddress string
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{}
}

func NewEmailBuilderFrom(a interface{}) *EmailBuilder {
	var b EmailBuilder
	_ = copier.Copy(&b, a)
	return &b
}

func (b *EmailBuilder) LoadValues(a interface{}) *EmailBuilder {
	_ = copier.Copy(b, a)
	return b
}

func (b *EmailBuilder) Name(v string) *EmailBuilder {
	b.xName = v
	return b
}

func (b *EmailBuilder) Address(v string) *EmailBuilder {
	b.xAddress = v
	return b
}

func (b *EmailBuilder) Build() *Email {
	return &Email{
		Name:    b.xName,
		Address: b.xAddress,
	}
}

func (a *Email) ToBuilder() *EmailBuilder {
	return &EmailBuilder{
		xName:    a.Name,
		xAddress: a.Address,
	}
}

func (a *Email) GetName() string {
	if a == nil {
		return ""
	}
	return a.Name
}

func (a *Email) GetAddress() string {
	if a == nil {
		return ""
	}
	return a.Address
}
