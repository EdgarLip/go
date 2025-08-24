package examplesOnly

import (
	//"fmt"
	"fmt"
	"log"
	"os"
	"time"
)


func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i
		_log(fmt.Sprintf("generate - sent i=%v to ch\n", i ))
	}
	_log(fmt.Sprintf("generate - closing  ch\n"))
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	// 1: mySieve = [3, 4, 5, .., 99], prime=2 output=[3, 5, 7, 9, .., 99]
	// 2: mySieve = [5, .., 99], prime=3 output=[5, 7, 11, ..]
	// 3: mySieve = [7, 11, ..], prime=5, output=[7, 11, 13, ..]
	// 4: mySieve = [11, 13, ..], prime=7
 
	for i := range src {
		_log(fmt.Sprintf("filter - cheking i=%v, prime=%v\n", i, prime))
		if i%prime != 0 {
			dst <- i
			_log(fmt.Sprintf("filter - sent i=%v to dst\n", i))
		}
	}
	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)

	go generate(limit, ch)

	for {
		
		prime, ok := <-ch
		_log(fmt.Sprintf("sieve-1 in for loop got prime = %v from ch ---------\n", prime))
		if !ok {
			break
		}

		ch1 := make(chan int)

		go filter(ch, ch1, prime)

		ch = ch1

		_log(fmt.Sprintf("sieve-2 in for loop prime = %v ---------\n", prime))
	}
}

func _log(txt string) {
	now := time.Now()
    timestamp := now.Format("2006/01/02 15:04:05") + fmt.Sprintf(".%010d", now.Nanosecond())
	log.Printf("%v %v", timestamp,  txt)
}

func RunPrimeNumbers() {
	//log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
    log.SetFlags(0) // Disable default timestamp

	
	sieve(100) // 2 3 5 7 
}