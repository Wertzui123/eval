package eval

import "testing"

func TestEval_Addition(t *testing.T) {
	input := "5 + 3 + 12 + -10"

	expected := 10.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_Subtraction(t *testing.T) {
	input := "24 - 12 - -2"

	expected := 14.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_Multiplication(t *testing.T) {
	input := "0.5 * 60 * 3"

	expected := 90.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_Division(t *testing.T) {
	input := "15 / 0.5 / 6"

	expected := 5.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_Power(t *testing.T) {
	input := "2 ^ 3"

	expected := 8.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_Parenthesis(t *testing.T) {
	input := "-5 + 5 * ( 7 - 2 )"

	expected := 20.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}

func TestEval_All(t *testing.T) {
	input := "( 6 - 2 * ( 6 / 3 ) ) ^ 3"

	expected := 8.0
	actual, err := Eval(input)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Got %f, expected %f", actual, expected)
	}
}
