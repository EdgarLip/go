package cuncurrencyFunctions

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var start = time.Now()


func CuncurencyInAction() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 1000
	functionsToRun := 2

	const goRoutins = 10000
	var waitGroup1 sync.WaitGroup // create  wait group example.
	// var waitGroup2 sync.WaitGroup                                     - no need for this , i think this was a mistake creating 2 waiting groups.
	waitGroup1.Add(goRoutins * functionsToRun) //  set the amount of wait you need.
	// waitGroup2.Add(goRoutins)                                             - no need for this , i think this was a mistake creating 2 waiting groups.
	var mutex1 sync.Mutex // create a mutex

	go spendy(goRoutins, &counter, &waitGroup1, &mutex1)
	go spengy(goRoutins, &counter, &waitGroup1, &mutex1)

	waitGroup1.Wait()
	// waitGroup2.Wait()                             - no need for this , i think this was a mistake creating 2 waiting groups.
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count after cuncurrency happend:", counter)
}

func spendy(goRoutins int, counter *int, waitGroup *sync.WaitGroup, mutex1 *sync.Mutex) {
	for i := 0; i < goRoutins; i++ {
		mutex1.Lock() // i do not know why i dont need to put * before mutex1 ...
		*counter = *counter - 10
		mutex1.Unlock()
		//fmt.Printf("spendy's counter: %v\n", *counter)
		runtime.Gosched() // time.Sleep(time.Second)
		waitGroup.Done()  // each loop we will decrees the amount of wait groups that the "waitGroup1.Wait()" needs to wait.
	}
}

func spengy(goRoutins int, counter *int, waitGroup *sync.WaitGroup, mutex1 *sync.Mutex) {
	for i := 0; i < goRoutins; i++ {
		mutex1.Lock() // i do not know why i dont need to put * before mutex1 ...
		*counter = *counter + 10
		mutex1.Unlock()
		//fmt.Printf("spengy's counter: %v\n", *counter)
		runtime.Gosched() // time.Sleep(time.Second)
		waitGroup.Done()  // each loop we will decrees the amount of wait groups that the "waitGroup1.Wait()" needs to wait.
	}
}

func send(c chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)
	wg.Done()
}

func receive(c <-chan int, wg *sync.WaitGroup) {
	for v := range c {
		fmt.Println("Received from the channel:", v)
	}
	wg.Done()
}

func ChannelTestInAction() {
	c := make(chan int, 1)
	var waitGroup1 sync.WaitGroup
	waitGroup1.Add(2)

	go send(c, &waitGroup1)
	go receive(c, &waitGroup1)
	waitGroup1.Wait()

}
func PrintTime()  {
    elapsed := time.Since(start)
	fmt.Printf("%v - ", elapsed)
}

func FanInMain() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go FanInSend(even, odd)
	go FanInReceive(even, odd, fanin)

	// this is an endless loop that will keep running once the channel will have content,
	// it will pop that content
	// this loop will stop only in case the channel will get closed.
	for v := range fanin {
		PrintTime() 
		fmt.Printf("FAN in concept - fanin channel: %v\n", v)
	}
	PrintTime() 
	fmt.Println("about to exit")
}

func FanInSend(even, odd chan<- int) { // send channel
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			PrintTime() 
			fmt.Printf("FAN in concept sending: %v --> even \n", i)
			even <- i
		} else {
			PrintTime() 
			fmt.Printf("FAN in concept sending: %v --> odd \n", i)
			odd <- i
		}
	}
	
	close(even)
	PrintTime() 
	fmt.Printf("FAN in concept closed even \n")
	close(odd)
	PrintTime() 
	fmt.Printf("FAN in concept closed odd \n")
}

func FanInReceive(even, odd <-chan int, fanin chan<- int) {
	PrintTime() 
	fmt.Printf("FAN in concept FanInReceive sleeping \n")
	time.Sleep(time.Millisecond * 1000)
	
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			time.Sleep(time.Millisecond * 1000)
			PrintTime() 
			fmt.Printf("FAN in concept sending: %v from even --> fanin \n", v)
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			time.Sleep(time.Millisecond * 1000)
			PrintTime() 
			fmt.Printf("FAN in concept sending: %v from odd --> fanin \n", v)
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
