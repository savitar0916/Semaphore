package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	FeatureA_Semaphore = make(chan struct{}, 1)  // 功能A的信號量，最多允許2個併發
	FeatureB_Semaphore = make(chan struct{}, 10) // 功能B的信號量，最多允許3個併發
)

// 模擬功能A，執行5秒的任務
func featureA(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	FeatureA_Semaphore <- struct{}{}        // 獲取功能A的信號量
	defer func() { <-FeatureA_Semaphore }() // 釋放功能A的信號量
	fmt.Printf("[%v] Feature A Task %d is starting, will take 3 seconds\n", time.Now().Format("15:04:05"), id)
	time.Sleep(3 * time.Second) // 模擬功能A的執行時間
	fmt.Printf("[%v] Feature A Task %d is done\n", time.Now().Format("15:04:05"), id)
}

// 模擬功能B，執行3秒的任務
func featureB(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	FeatureB_Semaphore <- struct{}{}        // 獲取功能B的信號量
	defer func() { <-FeatureB_Semaphore }() // 釋放功能B的信號量
	fmt.Printf("[%v] Feature B Task %d is starting, will take 3 seconds\n", time.Now().Format("15:04:05"), id)
	time.Sleep(3 * time.Second) // 模擬功能B的執行時間
	fmt.Printf("[%v] Feature B Task %d is done\n", time.Now().Format("15:04:05"), id)
}

func main() {
	var wg sync.WaitGroup

	// 這裡我們可以同時啟動功能A和功能B
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go featureA(i, &wg) // 同時啟動功能A的任務
		wg.Add(1)
		go featureB(i, &wg) // 同時啟動功能B的任務
	}

	// 等待所有任務完成
	wg.Wait()
	fmt.Println("All tasks completed.")
}
