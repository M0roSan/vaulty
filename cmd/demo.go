package main

import (
	"fmt"

	secret "github.com/masahiroyoshida/vaulty"
)

func main() {
	v := secret.File("my-fake-key", ".secrets")
	err := v.Set("demo_key", "some value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo_key2", "some crazy value")
	if err != nil {
		panic(err)
	}
	err = v.Set("demo_key3", "some value 1234567898765432")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println(plain)
	plain, err = v.Get("demo_key2")
	if err != nil {
		panic(err)
	}
	fmt.Println(plain)
	plain, err = v.Get("demo_key3")
	if err != nil {
		panic(err)
	}
	fmt.Println(plain)
}
