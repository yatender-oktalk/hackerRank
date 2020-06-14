package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	person user
	level  string
}

// value receiver method
func (u user) notify() {
	fmt.Printf("Sending user email to %s <%s>\n", u.name, u.email)
}

func (u admin) notify() {
	fmt.Printf("Sending admin email to %s <%s>\n", u.person.name, u.person.email)
}

// value receiver method
func (u *user) changeEmail(email string) {
	fmt.Printf("changing user email to %s\n", email)
	u.email = email
}

func main() {
	lisa := &user{"lisa","lisa@gmail.com"}
	lisa.notify()
	lisa.changeEmail("newlisa@gmail.com")
	lisa.notify()


	bill := user{"bill", "bill@gmail.com"}
	bill.notify()
	bill.changeEmail("newbill@gmail.com")
	bill.notify()
}
