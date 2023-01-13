package main

import (
	"webpage-download/src"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.POST("/pagesource/:retrylimit", src.DownloadPage)
	r.Run(":3000")
}

func Route() {

	// var wg sync.WaitGroup
	// var nums []int

	// for i := 1; i <= 10000; i++ {
	// 	nums = append(nums, i)
	// }
	// activeusers := 50

	// wg.Add(activeusers)
	// go Route()

	// wg.Wait()

}
