package chainofresponsability

import "fmt"

type Department interface {
	do()
	setNext(Department)
}

type reception struct {
	next    Department
	hasNext bool
}

type doctor struct {
	next    Department
	hasNext bool
}

type medicineRoom struct {
	next    Department
	hasNext bool
}

type cashier struct {
	next    Department
	hasNext bool
}

func (r *reception) do() {
	fmt.Println("Doing reception")
	if r.hasNext {
		r.next.do()
	}
}

func (r *reception) setNext(department Department) {
	r.next = department
	r.hasNext = true
}

func (r *doctor) do() {
	fmt.Println("Doing doctor")
	if r.hasNext {
		r.next.do()
	}
}

func (r *doctor) setNext(department Department) {
	r.next = department
	r.hasNext = true
}

func (r *medicineRoom) do() {
	fmt.Println("Doing medicineRoom")
	if r.hasNext {
		r.next.do()
	}
}

func (r *medicineRoom) setNext(department Department) {
	r.next = department
	r.hasNext = true
}

func (r *cashier) do() {
	fmt.Println("Doing cashier")
	if r.hasNext {
		r.next.do()
	}
}

func (r *cashier) setNext(department Department) {
	r.next = department
	r.hasNext = true
}
