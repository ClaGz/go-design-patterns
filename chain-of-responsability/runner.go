package chainofresponsability

func Run() {

	cashier := &cashier{}

	//Set next for medical department
	medical := &medicineRoom{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	reception.do()
}
