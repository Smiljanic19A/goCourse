package main

//requirements
// create deck done
// deal hand
// write to file
// read from file

func main() {
	d := newDeck()

	d.shuffle(12)

	d.print()

}
