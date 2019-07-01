package main 

import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup


func for2() {
	defer wg.Done()
	for j := 0; j < 15; j++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("=")
	}
	wg.Wait()
}

 func for1() {
	defer wg.Done()
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go for2()
		fmt.Println()
	}
	
}

func main(){
	fmt.Println()
	
	wg.Add(1)
	
	go for1()
	wg.Wait()

	fmt.Println()
}