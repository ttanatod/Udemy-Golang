package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "is even")
// 		} else {
// 			fmt.Println(i, "is odd")
// 		}
// 	}
// }
