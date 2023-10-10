package main

import (
    "fmt"
	"sync"
	"net/http"
	"time"
	
   )
   

  
  func fetchWiki(city string, ch chan<- string, wg *sync.WaitGroup) {
  
       url := fmt.Sprintf("https://en.wikipedia.org/wiki/")
	   resp, err := http.Get(url)
	   if err != nil{
			fmt.Printf("error fetching page for %s: %s", city, err)
	   }
	   fmt.Println(city)
	   
	   defer resp.Body.Close()
	   
	   ch <- fmt. Sprintf("City: %s", city)
	   
       
   }
   
   func main() {
       cities := []string{"Manila", "Seoul", "Paris", "Detroit"}
	   ch := make(chan string)
	   var wg sync.WaitGroup
	   start := time.Now()
	   
	   for _, city := range cities{
			wg.Add(1)
			fetchWiki(city, ch, &wg)
	   }
	   go func() {
			wg.Wait()
			close(ch)
	   }()
	   
	   
	   for result := range ch{
			fmt.Println(result)
	   }
	   
	   
	   fmt.Println("Time: ", time.Since(start))
       
}      
