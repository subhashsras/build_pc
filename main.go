package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"practice.com/build_pc/components"
	"practice.com/build_pc/components/harddisk"
	"practice.com/build_pc/components/processor"
)

type PC struct {
	processor.Processable
	components.Ram
	harddisk.Storable
}

func (p PC) BootUp() {
	fmt.Println("Booting up...")
}

func BuildAMDPC() {
	startTime := time.Now()
	amdProcessor := processor.NewAMDProcessor()
	wdHardDisk := harddisk.NewWDHardDisk()
	amdPC := PC{amdProcessor, components.Ram{}, wdHardDisk}
	time.Sleep(3 * time.Second)
	amdPC.BootUp()
	amdPC.ReadFromDisk()
	amdPC.SaveToDisk()
	amdPC.LoadFromDisk()
	amdPC.SendToProcessor()
	amdPC.Compute()
	fmt.Println("AMD BuildTime:", time.Since(startTime))
}

func BuildIntelPC() {
	startTime := time.Now()
	intelProcessor := processor.NewIntelProcessor()
	seagateHardDisk := harddisk.NewSeagateHardDisk()
	intelPC := PC{intelProcessor, components.Ram{}, seagateHardDisk}
	intelPC.BootUp()
	intelPC.ReadFromDisk()
	intelPC.SaveToDisk()
	intelPC.LoadFromDisk()
	intelPC.SendToProcessor()
	intelPC.Compute()
	fmt.Println("Intel BuildTime:", time.Since(startTime))
}

func main() {
	start := time.Now()

	var wg sync.WaitGroup

	var ch = make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	totalRequests := 0
	go func(ctx context.Context) {
		defer wg.Done()

		fmt.Println("PC Billing started...")
		for {

			fmt.Println("PC Biller Waiting")
			var input string
			fmt.Println("Enter a for AMD PC or i for Intel PC or q for exit")
			fmt.Scanln(&input)
			if input == "q" {
				cancel()
				fmt.Println("PC Billing stopped.")
				return
			}
			ch <- input
			totalRequests += 1
		}

	}(ctx)

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		fmt.Println("PC Building started...")
		for {
			fmt.Println("PC Builder Waiting")

			select {

			case <-ctx.Done():
				fmt.Println("PC Building stopped.")
				return

			case r := <-ch:

				switch r {
				case "a":
					{
						go BuildAMDPC()
					}
				case "i":
					{
						BuildIntelPC()
					}

				default:
					{
						fmt.Println("Invalid input", r)
					}
				}
			}

		}

	}(ctx)

	wg.Wait()
	fmt.Println("\nMain function ended: ", time.Since(start))
	fmt.Println("\nTotal Requests processed: ", totalRequests)
	// var input string
	// fmt.Scan(&input)

	// switch input {
	// case "a":
	// 	{
	// 		amdProcessor := processor.NewAMDProcessor()
	// 		wdHardDisk := harddisk.NewWDHardDisk()
	// 		amdPC := PC{amdProcessor, components.Ram{}, wdHardDisk}
	// 		amdPC.BootUp()
	// 		amdPC.ReadFromDisk()
	// 		amdPC.SaveToDisk()
	// 		amdPC.LoadFromDisk()
	// 		amdPC.SendToProcessor()
	// 		amdPC.Compute()
	// 	}
	// case "i":
	// 	{
	// 		intelProcessor := processor.NewIntelProcessor()
	// 		seagateHardDisk := harddisk.NewSeagateHardDisk()
	// 		intelPC := PC{intelProcessor, components.Ram{}, seagateHardDisk}
	// 		intelPC.BootUp()
	// 		intelPC.ReadFromDisk()
	// 		intelPC.SaveToDisk()
	// 		intelPC.LoadFromDisk()
	// 		intelPC.SendToProcessor()
	// 		intelPC.Compute()
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Invalid input")
	// 	}
	// }

}
