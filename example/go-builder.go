// Code generated by go-builder. DO NOT EDIT.
// Source: main.go

package main

import "github.com/jinzhu/copier"

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
	xAddress *PersonAddress
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

func (b *PersonBuilder) Address(v *PersonAddress) *PersonBuilder {
	b.xAddress = v
	return b
}

func (b *PersonBuilder) AddressFunc(f func(*PersonAddress) *PersonAddress) *PersonBuilder {
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

func (b *PersonBuilder) AddEmail(value *Email) *PersonBuilder {
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

func (a *Person) GetAddress() *PersonAddress {
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

type EmailBuilder struct {
	xName      string
	xAddress   AddressType
	xEmailType EmailType
	xTypes     []EmailType
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

func (b *EmailBuilder) Address(v AddressType) *EmailBuilder {
	b.xAddress = v
	return b
}

func (b *EmailBuilder) EmailType(v EmailType) *EmailBuilder {
	b.xEmailType = v
	return b
}

func (b *EmailBuilder) Types(v []EmailType) *EmailBuilder {
	b.xTypes = v
	return b
}

func (b *EmailBuilder) ClearTypes() *EmailBuilder {
	b.xTypes = []EmailType{}
	return b
}

func (b *EmailBuilder) AddType(value EmailType) *EmailBuilder {
	b.xTypes = append(b.xTypes, value)
	return b
}

func (b *EmailBuilder) ForEachType(f func(EmailType) EmailType) *EmailBuilder {
	list := make([]EmailType, 0, len(b.xTypes))
	for _, item := range b.xTypes {
		list = append(list, f(item))
	}
	b.xTypes = list
	return b
}

func (b *EmailBuilder) Build() *Email {
	return &Email{
		Name:      b.xName,
		Address:   b.xAddress,
		EmailType: b.xEmailType,
		Types:     b.xTypes,
	}
}

func (a *Email) ToBuilder() *EmailBuilder {
	return &EmailBuilder{
		xName:      a.Name,
		xAddress:   a.Address,
		xEmailType: a.EmailType,
		xTypes:     a.Types,
	}
}

func (a *Email) GetAddress() AddressType {
	if a == nil {
		return AddressType("")
	}
	return a.Address
}

func (a *Email) GetEmailType() EmailType {
	if a == nil {
		return EmailType(0)
	}
	return a.EmailType
}

func (a *Email) GetTypes() []EmailType {
	if a == nil {
		return []EmailType{}
	}
	return a.Types
}

type PersonAddressBuilder struct {
	xStreet string
	xNumber int
}

func NewPersonAddressBuilder() *PersonAddressBuilder {
	return &PersonAddressBuilder{}
}

func NewPersonAddressBuilderFrom(a interface{}) *PersonAddressBuilder {
	var b PersonAddressBuilder
	_ = copier.Copy(&b, a)
	return &b
}

func (b *PersonAddressBuilder) LoadValues(a interface{}) *PersonAddressBuilder {
	_ = copier.Copy(b, a)
	return b
}

func (b *PersonAddressBuilder) Street(v string) *PersonAddressBuilder {
	b.xStreet = v
	return b
}

func (b *PersonAddressBuilder) Number(v int) *PersonAddressBuilder {
	b.xNumber = v
	return b
}

func (b *PersonAddressBuilder) Build() *PersonAddress {
	return &PersonAddress{
		Street: b.xStreet,
		Number: b.xNumber,
	}
}

func (a *PersonAddress) ToBuilder() *PersonAddressBuilder {
	return &PersonAddressBuilder{
		xStreet: a.Street,
		xNumber: a.Number,
	}
}

func (a *PersonAddress) GetStreet() string {
	if a == nil {
		return ""
	}
	return a.Street
}

func (a *PersonAddress) GetNumber() int {
	if a == nil {
		return 0
	}
	return a.Number
}
