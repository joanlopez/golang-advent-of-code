package intcode

type IOChan interface {
	Send(data int)
	Receive() int
	Halt()
}
