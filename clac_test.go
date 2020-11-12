package go_modt

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Fatal("sum err")
	}
}

func TestSub(t *testing.T) {
	if Sub(2, 3) != -1 {
		t.Fatal("sub err")
	}
}

//add sth


//add sth 2