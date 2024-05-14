package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano int `json:"-"`
}

func (c Carro) Andar() {
	fmt.Println("O carro, " + c.Nome + " está andando")
}

func (c Carro) Parar() {
	fmt.Println("O carro, " + c.Nome + " está parando")
}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Olá, Mundo!")
}

func main() {
  http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}

func mainOld01() {
	print("Hello, World!")
}

func mainOld02() {
	http.ListenAndServe(":3333", nil);
}

func mainOld03() {
	a := 1
	a = 2
	s := "hello"
	x := true

	print(a)
	print("\n")
	print(s)
	print("\n")
	print(x)
	print("\n")
}

func mainOld04() {
	carro1 := Carro{Nome: "Toyota Corolla", Modelo: "Sedan", Ano: 2022}
	carro2 := Carro{Nome: "Ford Mustang", Modelo: "GT", Ano: 2023}
	
	fmt.Println(carro1.Nome)
	fmt.Println(carro2.Nome)

	carro1.Andar()
	carro1.Parar()
	carro2.Andar()
	carro2.Parar()
}

func mainOld05() {
	carro1 := Carro{Nome: "Toyota Corolla", Modelo: "Sedan", Ano: 2022}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(carro1)
	})

	http.ListenAndServe(":8080", nil)
}
