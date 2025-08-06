package utils

import (
	"fmt"
	"strings"
	"time"
)

type Person struct {
	name string
	age  int
}

func (p Person) GetPersonName() string {
	// reciver value baiknya untuk read-only operasi

	return p.name
}

func (p *Person) SetPersonName(name string) {
	// reciver pointer baik untuk setter atau mengubah nilai dari reciver itu
	p.name = name
}

func CallMethodSection() {
	fmt.Println("Hellow world methods")

	person := Person{name: "Gusti Bisman Taka", age: 20}

	nameOfperson := person.GetPersonName()

	fmt.Println("Hai nama aku adalah ", nameOfperson)

	person.SetPersonName("BismenBoy")

	fmt.Println("Hai aku berubah nama sekarang nama aku adalah ", person.GetPersonName())
	fmt.Println()
	chainingMethodExampel()

	fmt.Println()

	callUserMethod()
}

type Calcullator struct {
	result float64
}

// Chining method
func (c *Calcullator) Add(n float64) *Calcullator {
	c.result += n
	return c
}

func (c *Calcullator) Substract(n float64) *Calcullator {
	c.result -= n
	return c
}

func (c *Calcullator) Multiply(n float64) *Calcullator {
	c.result *= n

	return c
}

func (c Calcullator) GetResult() float64 {
	return c.result
}

func chainingMethodExampel() {
	cal := Calcullator{}

	cal.Add(2).Add(4).Substract(100)

	valResult := cal.GetResult()

	fmt.Println("value chaining: ", valResult)
}

type User struct {
	Id        int
	Username  string
	Email     string
	CreatedAt time.Time
	IsActive  bool
}

func NewUser(username string, email string) *User {
	return &User{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		IsActive:  true,
	}
}

func (u *User) ValidateEmail() bool {
	return len(u.Email) > 3 && strings.Contains(u.Email, "@")
}

func (u *User) ValidateUsername() bool {
	return len(u.Username) > 3
}

func (u User) GetFullInfo() string {
	status := "Active"

	if !u.IsActive {
		status = "Inactive"
	}

	return fmt.Sprintf("User : %s (%s) - Status %s - CreatedAt - %s", u.Email, u.Username, status, u.CreatedAt.Format("23-08-2001"))
}

func (u User) IsValidUser() bool {
	return u.ValidateEmail() && u.ValidateUsername()
}

func callUserMethod() {
	newUser := NewUser("bisma_taka", "bism@mail.com")

	if newUser.IsValidUser() {
		getInfoUser := newUser.GetFullInfo()

		fmt.Println(getInfoUser)
	} else {
		fmt.Println("Error user method")
	}
}
