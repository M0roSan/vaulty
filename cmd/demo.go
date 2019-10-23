package main

import (
	"fmt"

	secret "github.com/masahiroyoshida/vaulty"
)

func main() {
	v := secret.Memory("my-fake-key")
	err := v.Set("demo_key", "some value")
	if err != nil {
		panic(err)
	}
	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println(plain)
}
