package channel

import "fmt"

type ReceiverService struct {
	ch <-chan string
}

func (rs *ReceiverService) Init(chh <-chan string) {
	rs.ch = chh
	rs.ReceiveLoop()
}

func (rs *ReceiverService) ReceiveLoop() {
	for {
		msg, isOpen := <- rs.ch
		if !isOpen {
			break
		}
		fmt.Println(msg)
	}
}