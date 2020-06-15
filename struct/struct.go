package main

import "fmt"

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}
 
type notifier interface {
	notify()
}

// value receiver method
func (u user) notify() {
	fmt.Printf("Sending user email to %s <%s>\n", u.name, u.email)
}

func (u admin) notify() {
	fmt.Printf("Sending admin email to %s <%s>\n", u.name, u.email)
}

// notifier

func sendNotification(n notifier) {
	n.notify()
}

// value receiver method
func (u *user) changeEmail(email string) {
	fmt.Printf("changing user email to %s\n", email)
	u.email = email
}

func (u *admin) changeEmail(email string) {
	fmt.Printf("changing user email to %s\n", email)
	u.email = email
}

func main() {
	lisa := user{"lisa","lisa@gmail.com"}
	lisa.notify()
	lisa.changeEmail("newlisa@gmail.com")
	

	sendNotification(lisa)


	bill := admin{"bill", "bill@gmail.com"}
	bill.notify()
	bill.changeEmail("newbill@gmail.com")
	sendNotification(bill)
}
