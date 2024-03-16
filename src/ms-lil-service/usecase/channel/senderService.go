package channel

import "time"

type SenderService struct {
	ch chan<- string
}

func (rs *SenderService) Init(chh chan<- string) {
	rs.ch = chh
}

func (rs *SenderService) Send() {
	rs.ch <- "first mesig"
	time.Sleep(1 * time.Second)
	rs.ch <- "second mesig"
	rs.ch <- "third mesig"
	//time.Sleep(100 * time.Millisecond) // ne salje se zbog ovoga????
	rs.ch <- "I do be spammin" // OVO SE NE SALJE???? fren why
	close(rs.ch)
}