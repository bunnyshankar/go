package main

import (
	"fmt"
	"reflect"
)
type Doctor struct {
	number int
	actor string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actor: "shankar",
		companions: []string {
			"asdasads",
			"asdsadsasa",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actor) //to print actor 
	//other way of defining and initializing structs
	bDoc := struct{name string}{name: "Madhu"}
	anotherDoc := bDoc
	anotherDoc.name = "mmasdsa"
	fmt.Println(bDoc)
	fmt.Println(anotherDoc)
	aDoc := &bDoc //referencing to bDoc struct
	aDoc.name = "reference"
	fmt.Println(bDoc)
	fmt.Println(aDoc)
//Implementation of inheritence called Composition in go
	type animal struct {
		name string
		origin 	string
	}
	type bird struct {
		animal				//composition(similar to inheritence)
		speed 	float32
		canFly 	bool
	}
	b := bird{}
	b.name = "Emu"
	b.origin = "Australia"
	b.speed = 48
	b.canFly = false
	fmt.Println(b)
	fmt.Println(b.name)
	type animalTag struct {
		name string  `required max:"100"` //tagging name needs reflect package
		origin 	string
	}
	t := reflect.TypeOf(animalTag{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)
}

