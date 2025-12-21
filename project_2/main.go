package main

func main() {
	philip := person{
		name:     "Philip",
		lastName: "Philip",
		age:      25,
		contact: contact{
			email: "a@mail.com",
			zip:   "123456",
		},
	}

	philip.print()

	philip.updateName("Keva tvoja")

	philip.print()
}
