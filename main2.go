// package main

// import "fmt"

// //definimos una estructura para representar una cancion
// type Song struct {
// 	Title string
// 	Artist string
// }

// //Creamos un slice para almacenar las canciones

// func main(){
// 	playlist := []Song{
// 		{Title: "Bohemian Rhapsody", Artist: "Queen"},
// 		{Title: "Starway to heaven", Artist: "Led Zeppelin"},
// 	}
// 	//mostramos la playlist
// 	fmt.Println("Tu playlist: ")
// 	for i, song := range playlist{
// 		fmt.Printf("%d. %s - %s\n", i+1, song.Title, song.Artist)
// 	}

// 	//Añadir una nueva cancion
// 	newSong := Song{Title: "Hotel California", Artist: "Eagles"}
// 	playlist = append(playlist, newSong)
// 	fmt.Printf("\nCancion añadida: %s - %s ", newSong.Title, newSong.Artist)

// 	//mostramos la playlist actualizada
// 	fmt.Println("\nLista actualizada: ")
// 	for i, song:= range playlist{
// 		fmt.Printf("%d. %s - %s\n", i+1, song.Title, song.Artist)
// 	}


// }

