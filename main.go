package main

import(
	"fmt"
	"sync"
)
func makeHotesleservation(wg *sync.WaitGroup){
	fmt.Println("Hotel reservation is done")
	wg.Done()
}
func flightTicketBookings(wg *sync.WaitGroup){
	fmt.Println("Flight booking is done")
	wg.Done()
}
func orderDress(wg *sync.WaitGroup){
	fmt.Println("Dress is ordered")
	wg.Done()
}
func payCreditCardBills(wg *sync.WaitGroup){
	fmt.Println("Credit card bills are paid")
	wg.Done()
}
func writeMail(wg *sync.WaitGroup){
	fmt.Println("Writing mail first 1/3 is done")
	go continueWritingMail(wg)
}
func continueWritingMail(wg *sync.WaitGroup){
	fmt.Println("Writing mail second 1/3 is done")
	go continueWritingMail1(wg)
}
func continueWritingMail1(wg *sync.WaitGroup){
	fmt.Println("Writing mail last 1/3 is done")
	fmt.Println("mail completed")
	wg.Done()
}

func listenAudioBook(wg *sync.WaitGroup){
	fmt.Println("Listen audio book 1/2 is done")
	go listenAudioBook1(wg)
}
func listenAudioBook1(wg *sync.WaitGroup){
	fmt.Println("Listen audio book 1/2 is done")
	fmt.Println("audio book is completed")
	wg.Done()
}
var allTask = []func(*sync.WaitGroup){
	makeHotesleservation,payCreditCardBills,flightTicketBookings,orderDress,
	writeMail,listenAudioBook,
}
func main(){
	fmt.Println("JSMPJ Corporation")
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(allTask))
	for _,task := range allTask{
       go task(&waitGroup)
	}
	waitGroup.Wait()
}