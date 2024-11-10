package main

import "fmt" //+
func main() {
	fmt.Println("Hello, world!")
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i*30+j*25 == 335 {
				fmt.Println("30->", i, "25->", j)
			}
		}
	}
}
