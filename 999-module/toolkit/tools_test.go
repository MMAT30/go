package toolkit

import "testing"


func TestAdder(t *testing.T) {
	var MathTest Math


	mathReturn := MathTest.Adder(1,2)

	if (mathReturn != 3) {
		t.Error("wrong addition of 1 + 2")
	}

}

