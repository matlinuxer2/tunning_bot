package main

import (
	"fmt"
	"net"
	"time"
	//"io"
	"bufio"
	//"strings"
)

type TimeRec struct {
	t_ns int64
	name string
}

type TimeRecs struct {
	recs []TimeRec
}

func (history *TimeRecs) record(name string) {
	t_ns := time.Now().UnixNano()
	item := TimeRec{t_ns, name}
	history.recs = append(history.recs, item)

	return
}

func (history *TimeRecs) show() {
	for i, curr := range history.recs {
		if i < 1 {
			continue
		}
		prev := history.recs[i-1]

		t_delta := float64(curr.t_ns-prev.t_ns) / 1000000 // in ms

		fmt.Printf("[%3d] %12.3f ms | [ %25s -> %-25s ] \n", i, float64(t_delta), prev.name, curr.name)
	}

	return
}

func (history *TimeRecs) calc() []float64 {
	result := []float64{}

	indeces := []int{2, 4}

	for _, idx := range indeces {
		curr := history.recs[idx]
		prev := history.recs[idx-1]
		t_delta := float64(curr.t_ns-prev.t_ns) / 1000000 // in ms
		result = append(result, t_delta)
	}

	return result
}

func main() {
	history := TimeRecs{}
	history.record("init")
	// history = append(history, { time.Now().UnixNano(),"beg" } )

	history.record("connecting")
	raddr, _ := net.ResolveTCPAddr("tcp", "216.58.200.36:80")
	ipconn, _ := net.DialTCP("tcp", nil, raddr)
	ipconn.SetNoDelay(true)
	//ipconn, _ := net.Dial("ip4:tcp", "192.168.1.1:80")

	history.record("sending")
	ret1, _ := ipconn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	history.record("receiving")
	//ret2, _ := ipconn.Read(buf)
	ret2, _ := bufio.NewReader(ipconn).ReadString('\n')
	buf := ret2
	//ret2 := io.ReadFull(ipconn,buf)
	fmt.Println(string(buf))

	history.record("closing")
	ipconn.Close()

	fmt.Println(ret1, ret2)

	history.record("exit")

	history.show()
	ret3 := history.calc()
	fmt.Println(ret3)
}
