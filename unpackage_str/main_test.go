package main

import "testing"

func TestSolve(t *testing.T) {
	s := "a4bc2d5e"
	ans := "aaaabccddddde"
	if res := Solve(s); res != ans {
		t.Fatalf("Wrong answer EXPECTED = %s, RESULT = %s", ans, res)
	}
}

func TestSolve2(t *testing.T) {
	s := "abcd"
	ans := "abcd"
	if res := Solve(s); res != ans {
		t.Fatalf("Wrong answer EXPECTED = %s, RESULT = %s", ans, res)
	}
}

func TestSolve3(t *testing.T) {
	s := "45"
	ans := ""
	if res := Solve(s); res != ans {
		t.Fatalf("Wrong answer EXPECTED = %s, RESULT = %s", ans, res)
	}
}

func TestSolve4(t *testing.T) {
	s := ""
	ans := ""
	if res := Solve(s); res != ans {
		t.Fatalf("Wrong answer EXPECTED = %s, RESULT = %s", ans, res)
	}
}

func TestSolve5(t *testing.T) {
	s := "ah5du6"
	ans := "ahhhhhduuuuuu"
	if res := Solve(s); res != ans {
		t.Fatalf("Wrong answer EXPECTED = %s, RESULT = %s", ans, res)
	}
}
