package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {
	p1 := person{
		First: "Mateus",
	}

	p2 := person{
		First: "Marcos",
	}

	xp := []person{p1, p2}

	bs, err := json.Marshal(xp)
	
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bs))
}