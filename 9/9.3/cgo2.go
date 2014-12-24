package main

/*
#include <stdio.h>
int hello() {
    printf("Hello. Cgo! -- From C world.\n");
    return 0;
}
*/
import "C"

func Hello() int {
	return int(C.hello())
}

func main() {
	Hello()
}
