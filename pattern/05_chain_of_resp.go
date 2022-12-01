package main

import "fmt"

type department interface {
	execute(*client)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *client) {
	if p.registrationDone {
		fmt.Println("Registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering client")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type servicemanager struct {
	next department
}

func (d *servicemanager) execute(p *client) {
	if p.serviceGetDone {
		fmt.Println("Servicealready done")
		d.next.execute(p)
		return
	}
	fmt.Println("Manager is serving client")
	p.serviceGetDone = true
	d.next.execute(p)
}

func (d *servicemanager) setNext(next department) {
	d.next = next
}

type passport struct {
	next department
}

func (m *passport) execute(p *client) {
	if p.paperWorkDone {
		fmt.Println("Document already given to client")
		m.next.execute(p)
		return
	}
	fmt.Println("Passport handing to client")
	p.paperWorkDone = true
	m.next.execute(p)
}

func (m *passport) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *client) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from client client")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type client struct {
	name             string
	registrationDone bool
	serviceGetDone   bool
	paperWorkDone    bool
	paymentDone      bool
}

func main() {
	cashier := &cashier{}

	//Set next for passport department
	passport := &passport{}
	passport.setNext(cashier)

	//Set next for servicemanager department
	servicemanager := &servicemanager{}
	servicemanager.setNext(passport)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(servicemanager)

	client := &client{name: "qwerty"}
	//Patient visiting
	reception.execute(client)
}
