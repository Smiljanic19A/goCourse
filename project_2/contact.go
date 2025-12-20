package main

type contact struct {
	email string
	zip   string
}

func (c contact) toString() string {
	return "Email : " + c.email + "\nZip : " + c.zip
}
