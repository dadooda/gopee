
//
// Originally from: https://github.com/dadooda/gopee.
//

package gopee_test

import . "dadooda/gopee"

func ExampleP() {
	P("Joe")
	P(25)

	type person struct {
		firstName string
		gender string
	}

	P(person{"Joe", "m"})

	// Output:
	// "Joe"
	// 25
	// gopee_test.person{firstName:"Joe", gender:"m"}
}

func ExampleP_labelsBasic() {
	P("name:", "Joe", "age:", 25)

	// Output:
	// name:"Joe"
	// age:25
}

func ExampleP_labelsCompact() {
	P(1.1, " name:", "Joe", " age:", 25)

	// Output:
	// 1.1 name:"Joe" age:25
}
