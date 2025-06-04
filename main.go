package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func sum(numbers []int, result chan int) {
	total := 0
	for _, num := range numbers {
		total += num
	}
	result <- total
}

var (
	dbURL         string
	concurency    = flag.Int("concurency", 10, "Number Of Concurent Workers")
	bathSize      = flag.Int("batch-size", 1000, "Number Of row per batch")
	numBatches    = flag.Int("batches", 5000, "Total Number Of Batches")
	noConcurrency = flag.Bool("no-concurrency", false, "Disable concurrency (no goroutines)")
)

func init() {
	flag.Parse()
	dbURL = os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Missing required variable: DB_URL")
	}
}
func main() {

}

//Chanel Implementation
// values := []int{1, 3, 6, 5, -2, -3}
// resultChanel := make(chan int)
// go sum(values[:len(values)/2], resultChanel)
// go sum(values[len(values)/2:], resultChanel)

// leftSum, rightSum := <-resultChanel, <-resultChanel //menerima datadari chanel
// fmt.Println(leftSum, rightSum, leftSum+rightSum)
// var wg sync.WaitGroup
// totaly := 5
// for i := 0; i < totaly; i++ {
// 	wg.Add(1)
// 	go func(id int) {
// 		defer wg.Done()
// 		fmt.Printf("Goroutine %d done\n", id)
// 	}(i)
// }
// wg.Wait()
// fmt.Println("All Goroutine finish running")

func firstPrint() {
	fmt.Println("Cetak nama sebagai pertama")
}

func secondPrint() {
	fmt.Println("cetak nama dengan kedua")
}
