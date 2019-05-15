package main

// 每个PipeData中持有一个channel
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

// 遍历传入的所有PipeData将自身的value经过handler处理之后写入自己的next对应的channel
func handler(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

func main() {

}
