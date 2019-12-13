package arcade

const (
	halt   = 0
	input  = 1
	output = 2
)

type IOMsg struct {
	msgType int
	C       chan int
}

func (io IOMsg) IsHalt() bool {
	return io.msgType == halt
}

func (io IOMsg) IsInput() bool {
	return io.msgType == input
}

func (io IOMsg) IsOutput() bool {
	return io.msgType == output
}

type IO struct {
	C chan IOMsg
}

func NewIO() IO {
	return IO{C: make(chan IOMsg)}
}

func (io IO) Send(data int) {
	ch := make(chan int)
	io.C <- IOMsg{msgType: output, C: ch}
	ch <- data
}

func (io IO) Receive() int {
	ch := make(chan int)
	io.C <- IOMsg{msgType: input, C: ch}
	return <-ch
}

func (io IO) Halt() {
	io.C <- IOMsg{msgType: halt}
}
