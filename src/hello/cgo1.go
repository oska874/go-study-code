package hello

/*
#include <stdio.h>
void hello() {
printf("Hello, Cgo! -- From C world.\n");
}
*/
import "C"

func Hello() int {
	return int(C.hello())
}
