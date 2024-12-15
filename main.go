package main

import (
	"bitcoin-lottery/bitcoin_address"
	"bitcoin-lottery/validater"
	"context"
	"fmt"
	"runtime"
	"sync"
)

var numWorkers = runtime.NumCPU()
var jobQueueSize = numWorkers * 10

func worker(ctx context.Context, cancel context.CancelFunc, id int, jobs <-chan int, results chan<- bitcoin_address.AddressBlock, wg *sync.WaitGroup) {
	defer wg.Done()
	vl := validater.NewValidater("bloom_filter.bf", "address_after_clean_only_address_sorted.txt")
	defer vl.Close()

	for iteration := range jobs {
		select {
		case <-ctx.Done():
			return
		default:
		}
		newAddressBlock := bitcoin_address.NewAddressBlock()
		match := vl.ValidateAddressBlock(newAddressBlock)
		if match {
			cancel()
			fmt.Printf("Worker %d found a match at iteration: %d\n", id, iteration)
			results <- newAddressBlock
			return
		}
	}
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	jobs := make(chan int, jobQueueSize)
	results := make(chan bitcoin_address.AddressBlock)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, cancel, w, jobs, results, &wg)
	}

	go func() {
		for iteration := 1; ; iteration++ {
			if iteration%1_000_000 == 0 {
				fmt.Println(iteration)
				//fmt.Println("Number of goroutines: ", runtime.NumGoroutine())
				fmt.Println("job queue: ", len(jobs), "/", jobQueueSize)
			}
			jobs <- iteration
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Match Address Block: ", result.String())
		//cancel()
		break
	}
}
