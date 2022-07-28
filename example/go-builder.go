// Code generated by go-builder. DO NOT EDIT.
// Source: main.go

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

type EmailBuilder struct {
	xName      string
	xAddress   AddressType
	xEmailType EmailType
	xTypes     []EmailType
	xInterface interface{}
	xMap1      map[string]string
	xMap2      map[string]Address
	xMap3      map[string]*Address
	xMap4      map[*Address]*Address
	xMap5      map[interface{}]interface{}
	xMap6      *map[interface{}]interface{}
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

func (b *EmailBuilder) Interface(v interface{}) *EmailBuilder {
	b.xInterface = v
	return b
}

func (b *EmailBuilder) InterfaceFunc(f func(interface{}) interface{}) *EmailBuilder {
	b.xInterface = f(b.xInterface)
	return b
}

func (b *EmailBuilder) Map1(v map[string]string) *EmailBuilder {
	b.xMap1 = v
	return b
}

func (b *EmailBuilder) Map1Func(f func(map[string]string) map[string]string) *EmailBuilder {
	b.xMap1 = f(b.xMap1)
	return b
}

func (b *EmailBuilder) Map2(v map[string]Address) *EmailBuilder {
	b.xMap2 = v
	return b
}

func (b *EmailBuilder) Map2Func(f func(map[string]Address) map[string]Address) *EmailBuilder {
	b.xMap2 = f(b.xMap2)
	return b
}

func (b *EmailBuilder) Map3(v map[string]*Address) *EmailBuilder {
	b.xMap3 = v
	return b
}

func (b *EmailBuilder) Map3Func(f func(map[string]*Address) map[string]*Address) *EmailBuilder {
	b.xMap3 = f(b.xMap3)
	return b
}

func (b *EmailBuilder) Map4(v map[*Address]*Address) *EmailBuilder {
	b.xMap4 = v
	return b
}

func (b *EmailBuilder) Map4Func(f func(map[*Address]*Address) map[*Address]*Address) *EmailBuilder {
	b.xMap4 = f(b.xMap4)
	return b
}

func (b *EmailBuilder) Map5(v map[interface{}]interface{}) *EmailBuilder {
	b.xMap5 = v
	return b
}

func (b *EmailBuilder) Map5Func(f func(map[interface{}]interface{}) map[interface{}]interface{}) *EmailBuilder {
	b.xMap5 = f(b.xMap5)
	return b
}

func (b *EmailBuilder) Map6(v *map[interface{}]interface{}) *EmailBuilder {
	b.xMap6 = v
	return b
}

func (b *EmailBuilder) Map6Func(f func(*map[interface{}]interface{}) *map[interface{}]interface{}) *EmailBuilder {
	b.xMap6 = f(b.xMap6)
	return b
}

func (b *EmailBuilder) Build() *Email {
	return &Email{
		Name:      b.xName,
		Address:   b.xAddress,
		EmailType: b.xEmailType,
		Types:     b.xTypes,
		Interface: b.xInterface,
		Map1:      b.xMap1,
		Map2:      b.xMap2,
		Map3:      b.xMap3,
		Map4:      b.xMap4,
		Map5:      b.xMap5,
		Map6:      b.xMap6,
	}
}

func (a *Email) ToBuilder() *EmailBuilder {
	return &EmailBuilder{
		xName:      a.Name,
		xAddress:   a.Address,
		xEmailType: a.EmailType,
		xTypes:     a.Types,
		xInterface: a.Interface,
		xMap1:      a.Map1,
		xMap2:      a.Map2,
		xMap3:      a.Map3,
		xMap4:      a.Map4,
		xMap5:      a.Map5,
		xMap6:      a.Map6,
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

func (a *Email) GetInterface() interface{} {
	if a == nil {
		return nil
	}
	return a.Interface
}

func (a *Email) GetMap1() map[string]string {
	if a == nil {
		return map[string]string{}
	}
	return a.Map1
}

func (a *Email) GetMap2() map[string]Address {
	if a == nil {
		return map[string]Address{}
	}
	return a.Map2
}

func (a *Email) GetMap3() map[string]*Address {
	if a == nil {
		return map[string]*Address{}
	}
	return a.Map3
}

func (a *Email) GetMap4() map[*Address]*Address {
	if a == nil {
		return map[*Address]*Address{}
	}
	return a.Map4
}

func (a *Email) GetMap5() map[interface{}]interface{} {
	if a == nil {
		return map[interface{}]interface{}{}
	}
	return a.Map5
}

func (a *Email) GetMap6() *map[interface{}]interface{} {
	if a == nil {
		return nil
	}
	return a.Map6
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
