// package main

// import "fmt"

// //aca definimos una estructura para representar un libro
// type Book struct {
// 	title string
// 	author string
// }

// //Creamos un slice de libros
// func main(){
// 	biblioteca:= []Book{
// 		{title: "Cien años de soledad", author: "Gabriel Garcia Marquez"},
// 		{title: "El aleph", author: "Jorge Luis Borges"},
	
// 	}
// 	// mostramos los detalles de libros
// 	fmt.Println("Estos son los libros que tenemos: ")
// 	for i, book := range biblioteca{
// 		fmt.Printf("%d. %s de %s\n",i+1, book.title, book.author)
// 	}
// 	//añadir un libro

// 	nuevoLibro := Book{title: "Jazz en Paraguay", author: "Ricardo Castellani/JoseVillamayor"}
// 	biblioteca = append(biblioteca, nuevoLibro)
// 	fmt.Println("Estos son los libros que tenemos: ")
// 	for i, book := range biblioteca{
// 		fmt.Printf("%d. %s de %s\n",i+1, book.title, book.author)
// 	}


// }