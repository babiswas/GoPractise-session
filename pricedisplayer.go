package app

import (
	"fmt"
	"math/rand"
	"strconv"
)

func pricePoller(priceArr [10]chan int) {
	for {
		select {
		case price := <-priceArr[0]:
			fmt.Println("Price of stock:", "STCK1", " :", price)
		case price := <-priceArr[1]:
			fmt.Println("Price of stock:", "STCK2", " :", price)
		case price := <-priceArr[2]:
			fmt.Println("Price of stock:", "STCK3", " :", price)
		case price := <-priceArr[3]:
			fmt.Println("Price of stock:", "STCK4", " :", price)
		case price := <-priceArr[4]:
			fmt.Println("Price of stock:", "STCK5", " :", price)
		case price := <-priceArr[5]:
			fmt.Println("Price of stock:", "STCK6", " :", price)
		case price := <-priceArr[6]:
			fmt.Println("Price of stock:", "STCK7", " :", price)
		case price := <-priceArr[7]:
			fmt.Println("Price of stock:", "STCK8", " :", price)
		case price := <-priceArr[8]:
			fmt.Println("Price of stock:", "STCK9", " :", price)
		case price := <-priceArr[9]:
			fmt.Println("Price of stock:", "STCK10", " :", price)
		}
	}
}

func priceGenerator(stck <-chan string, priceArr [10]chan int) {
	for stk := range stck {
		switch stk {
		case "STCK1":
			price := rand.Intn(100)
			priceArr[0] <- price
		case "STCK2":
			price := rand.Intn(100)
			priceArr[1] <- price
		case "STCK3":
			price := rand.Intn(100)
			priceArr[2] <- price
		case "STCK4":
			price := rand.Intn(100)
			priceArr[3] <- price
		case "STCK5":
			price := rand.Intn(100)
			priceArr[4] <- price
		case "STCK6":
			price := rand.Intn(100)
			priceArr[5] <- price
		case "STCK7":
			price := rand.Intn(100)
			priceArr[6] <- price
		case "STCK8":
			price := rand.Intn(100)
			priceArr[7] <- price
		case "STCK9":
			price := rand.Intn(100)
			priceArr[8] <- price
		case "STCK10":
			price := rand.Intn(100)
			priceArr[9] <- price
		}

	}
}

func StockPriceDisplay() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	stocks := make(chan string, 10)
	priceArr := [10]chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int),
	}

	go priceGenerator(stocks, priceArr)

	go pricePoller(priceArr)

	for {
		index := rand.Intn(len(numbers))
		stck_name := "STCK" + strconv.Itoa(numbers[index])
		stocks <- stck_name
	}

}
