package package1

import "testing"

func TestOut1(t *testing.T) {
	if Out1() != 41 {
		t.Fatal("ret !42 err\n")
	}
	t.Log("ok\n")
}
