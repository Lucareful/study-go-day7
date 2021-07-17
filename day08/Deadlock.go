package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 派出所证明 
	var psCertificate RecursiveMutex
	// 物业证明 
	var propertyCertificate RecursiveMutex

	var wg sync.WaitGroup 
	
	wg.Add(2) // 需要派出所和物业都处理
	
	// 派出所处理 goroutine
	go func ()  {
		defer wg.Done() // 派出所处理完成

		psCertificate.Lock()

		defer psCertificate.Unlock()

		// 检查材料 
		time.Sleep(5 * time.Second)
		
		// 请求物业的证明
		propertyCertificate.Lock() 
		
		propertyCertificate.Unlock()

	}()

	// 物业处理 goroutine
	go func ()  {
		defer wg.Done()

		propertyCertificate.Lock() 
		defer propertyCertificate.Unlock()

		// 检查材料 
		time.Sleep(5 * time.Second) 

		// 请求派出所的证明 
		psCertificate.Lock() 

		psCertificate.Unlock()
	}()
	
	wg.Wait() 
	fmt.Println("成功完成")

}