package src

import "log"

func Concc() {

	mychannel := make(chan string, 3)
	char := []string{"a", "b", "c"}

	for _, s := range char {
		mychannel <- s
	}
	close(mychannel)
	for result := range mychannel {
		log.Println(result)
	}
}

func Pipeline() {

	nums := []int{1, 2, 3, 4}
	a := Numstochannel(nums)
	b := Sq(a)
	for c := range b {
		log.Println(c)
	}

}
func Numstochannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, s := range nums {
			out <- s
		}
		close(out)
	}()
	return out
}

func Sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for s := range in {
			out <- s * s
		}
		close(out)
	}()
	return out
}

func HitUser() <-chan int {
	var nums []int
	requestbuffer := make(chan int, 10)
	for i := 0; i < 1000; i++ {
		nums = append(nums, i)
	}

	go func() {
		for _, s := range nums {
			requestbuffer <- s
		}
		close(requestbuffer)
	}()

	return requestbuffer
}
