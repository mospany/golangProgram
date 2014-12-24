package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi, pandaman!"

	Md5Init := md5.New()
	Md5Init.Write([]byte(TestString))
	Result := Md5Init.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Init := sha1.New()
	Sha1Init.Write([]byte(TestString))
	Result = Sha1Init.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
