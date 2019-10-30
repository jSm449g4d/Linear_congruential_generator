package main

import "time"

func main() {
	//X=(A*X+C)%M
	var x = 0
	var a = 214013
	var c = 2531011
	var m = 0x80000000 - 1
	var j = 0
	var t = time.Now().UnixNano()
	x = c
	for x != 0 {
		x = (a*x + c) % m
		j++
	}
	print((time.Now().UnixNano()-t)/int64(time.Microsecond), "[μs]", j, "[1]")
}
