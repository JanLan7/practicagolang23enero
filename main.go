package main

import "fmt"

//creamos la estructura
type Usuario struct {
	nombre string
	CI int
	barrio string
	telefono int
}

// creamos el slice de usuarios
func main(){
	usuarios:= []Usuario{
		{nombre:"Victor Morel", CI: 333555, barrio: "Nazaret", telefono: 981771871},
		{nombre:"Matias Insaurralde", CI: 333555, barrio: "Kure-luque", telefono: 981771871},
		{nombre:"Giovani Primerano", CI: 2267676, barrio: "Remanso", telefono: 891891891},
		{nombre:"Bruno Muñoz", CI: 28982989, barrio: "Barrio o", telefono: 18981981},
		{nombre:"Jose Villamayor", CI: 28989989, barrio: "Sajonia", telefono: 91891891},
	}

	//mostramos detalles de usuarios
	fmt.Println("Estos son los detalles de los usuarios: ")
	for i, Usuario:= range usuarios{
		fmt.Printf("%d. %s con CI: %d, del barrio %s y con numero de telefono numero %d\n", i +1, Usuario.nombre, Usuario.CI, Usuario.barrio, Usuario.telefono)
	}
}