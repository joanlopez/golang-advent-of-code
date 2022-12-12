package main

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

const input = `move 1 from 9 to 2
move 4 from 6 to 1
move 4 from 2 to 6
move 5 from 8 to 7
move 4 from 9 to 2
move 1 from 5 to 8
move 1 from 3 to 1
move 2 from 3 to 1
move 1 from 4 to 2
move 11 from 7 to 2
move 5 from 5 to 1
move 1 from 6 to 8
move 1 from 7 to 6
move 3 from 6 to 7
move 1 from 3 to 2
move 1 from 6 to 8
move 11 from 2 to 1
move 1 from 9 to 8
move 1 from 3 to 7
move 4 from 7 to 9
move 3 from 3 to 7
move 4 from 8 to 2
move 3 from 7 to 6
move 2 from 6 to 3
move 5 from 4 to 1
move 1 from 6 to 5
move 26 from 1 to 7
move 1 from 4 to 6
move 22 from 7 to 5
move 4 from 9 to 1
move 3 from 7 to 3
move 1 from 6 to 3
move 6 from 1 to 7
move 2 from 7 to 5
move 8 from 1 to 9
move 4 from 3 to 4
move 10 from 2 to 7
move 6 from 7 to 4
move 2 from 9 to 5
move 1 from 5 to 1
move 8 from 4 to 1
move 2 from 5 to 9
move 1 from 3 to 6
move 1 from 9 to 1
move 1 from 3 to 6
move 2 from 5 to 2
move 1 from 4 to 2
move 1 from 2 to 3
move 7 from 1 to 4
move 9 from 7 to 4
move 1 from 3 to 4
move 2 from 2 to 4
move 5 from 9 to 6
move 1 from 4 to 5
move 2 from 9 to 3
move 1 from 1 to 6
move 2 from 6 to 1
move 2 from 6 to 5
move 2 from 9 to 7
move 1 from 3 to 9
move 1 from 9 to 5
move 2 from 7 to 3
move 1 from 1 to 7
move 7 from 4 to 5
move 2 from 1 to 2
move 3 from 3 to 8
move 3 from 8 to 9
move 31 from 5 to 8
move 1 from 7 to 1
move 1 from 2 to 1
move 1 from 1 to 5
move 1 from 5 to 6
move 2 from 5 to 7
move 10 from 4 to 9
move 5 from 6 to 2
move 3 from 2 to 6
move 2 from 7 to 8
move 1 from 6 to 3
move 1 from 4 to 1
move 1 from 3 to 6
move 1 from 4 to 2
move 2 from 1 to 2
move 1 from 8 to 7
move 10 from 8 to 2
move 13 from 2 to 9
move 1 from 1 to 5
move 18 from 8 to 2
move 21 from 9 to 6
move 1 from 7 to 8
move 2 from 9 to 7
move 1 from 2 to 3
move 1 from 7 to 8
move 9 from 2 to 4
move 1 from 7 to 8
move 3 from 9 to 1
move 1 from 8 to 1
move 6 from 2 to 3
move 5 from 4 to 7
move 1 from 5 to 8
move 2 from 4 to 3
move 5 from 7 to 3
move 2 from 2 to 7
move 15 from 6 to 1
move 12 from 1 to 2
move 6 from 2 to 9
move 4 from 9 to 5
move 4 from 5 to 6
move 14 from 3 to 9
move 1 from 6 to 7
move 1 from 7 to 2
move 1 from 7 to 8
move 9 from 2 to 6
move 1 from 1 to 6
move 2 from 9 to 8
move 4 from 9 to 7
move 1 from 1 to 5
move 8 from 8 to 3
move 1 from 5 to 4
move 2 from 1 to 2
move 3 from 1 to 4
move 9 from 6 to 2
move 1 from 7 to 4
move 1 from 8 to 2
move 1 from 6 to 4
move 4 from 7 to 8
move 12 from 6 to 8
move 3 from 2 to 1
move 6 from 8 to 7
move 5 from 3 to 6
move 3 from 3 to 6
move 3 from 1 to 3
move 8 from 2 to 9
move 2 from 4 to 5
move 2 from 7 to 2
move 10 from 8 to 5
move 3 from 3 to 2
move 10 from 5 to 3
move 1 from 4 to 3
move 1 from 2 to 1
move 1 from 1 to 7
move 14 from 9 to 6
move 5 from 2 to 4
move 15 from 6 to 5
move 3 from 9 to 3
move 1 from 8 to 6
move 1 from 3 to 8
move 7 from 3 to 8
move 16 from 5 to 1
move 2 from 7 to 1
move 1 from 5 to 9
move 2 from 9 to 3
move 15 from 1 to 5
move 3 from 8 to 2
move 3 from 3 to 1
move 3 from 7 to 3
move 8 from 4 to 6
move 5 from 1 to 6
move 9 from 5 to 7
move 2 from 8 to 3
move 2 from 2 to 7
move 1 from 1 to 4
move 2 from 5 to 8
move 4 from 3 to 1
move 4 from 8 to 1
move 1 from 8 to 6
move 9 from 7 to 6
move 2 from 7 to 5
move 3 from 1 to 8
move 1 from 4 to 8
move 1 from 2 to 4
move 12 from 6 to 2
move 3 from 8 to 6
move 1 from 4 to 7
move 2 from 6 to 8
move 5 from 5 to 9
move 13 from 2 to 9
move 2 from 4 to 7
move 13 from 9 to 5
move 2 from 6 to 5
move 1 from 3 to 9
move 6 from 9 to 4
move 5 from 1 to 3
move 1 from 7 to 9
move 15 from 5 to 8
move 2 from 4 to 7
move 2 from 4 to 6
move 1 from 4 to 6
move 1 from 5 to 7
move 18 from 6 to 2
move 2 from 7 to 3
move 3 from 6 to 7
move 3 from 2 to 8
move 5 from 7 to 3
move 1 from 9 to 6
move 2 from 3 to 8
move 11 from 3 to 2
move 2 from 2 to 9
move 1 from 6 to 2
move 1 from 7 to 5
move 1 from 5 to 9
move 9 from 8 to 4
move 1 from 4 to 6
move 2 from 3 to 1
move 2 from 1 to 5
move 12 from 8 to 3
move 1 from 8 to 2
move 14 from 3 to 4
move 1 from 6 to 4
move 1 from 5 to 4
move 20 from 2 to 7
move 2 from 9 to 5
move 1 from 5 to 3
move 1 from 9 to 2
move 1 from 2 to 8
move 2 from 2 to 3
move 5 from 4 to 5
move 6 from 5 to 7
move 2 from 8 to 2
move 3 from 3 to 9
move 5 from 4 to 5
move 2 from 9 to 7
move 2 from 2 to 3
move 1 from 9 to 3
move 22 from 7 to 3
move 4 from 7 to 4
move 24 from 3 to 6
move 4 from 2 to 6
move 18 from 6 to 9
move 15 from 4 to 6
move 8 from 6 to 3
move 6 from 6 to 1
move 7 from 9 to 6
move 2 from 7 to 4
move 8 from 3 to 9
move 14 from 6 to 3
move 2 from 3 to 9
move 1 from 9 to 6
move 13 from 9 to 1
move 3 from 4 to 5
move 1 from 9 to 6
move 5 from 1 to 8
move 3 from 3 to 9
move 2 from 1 to 5
move 8 from 5 to 8
move 10 from 3 to 5
move 3 from 4 to 6
move 6 from 1 to 9
move 4 from 5 to 3
move 5 from 8 to 2
move 6 from 6 to 3
move 7 from 3 to 6
move 1 from 3 to 4
move 5 from 8 to 7
move 5 from 2 to 6
move 2 from 7 to 3
move 3 from 7 to 3
move 1 from 4 to 9
move 9 from 6 to 9
move 2 from 6 to 2
move 1 from 8 to 2
move 2 from 8 to 7
move 5 from 1 to 5
move 1 from 1 to 4
move 13 from 5 to 7
move 5 from 3 to 7
move 1 from 5 to 6
move 1 from 4 to 6
move 3 from 2 to 8
move 1 from 3 to 5
move 1 from 3 to 8
move 14 from 7 to 4
move 1 from 5 to 6
move 7 from 6 to 9
move 6 from 7 to 9
move 2 from 8 to 9
move 2 from 8 to 1
move 31 from 9 to 1
move 13 from 4 to 2
move 1 from 4 to 3
move 10 from 2 to 7
move 1 from 3 to 4
move 1 from 2 to 7
move 3 from 7 to 8
move 1 from 4 to 1
move 3 from 8 to 5
move 32 from 1 to 5
move 3 from 9 to 7
move 4 from 9 to 6
move 2 from 2 to 7
move 2 from 1 to 7
move 1 from 6 to 1
move 1 from 9 to 4
move 3 from 6 to 4
move 1 from 1 to 8
move 15 from 5 to 1
move 1 from 8 to 4
move 9 from 5 to 7
move 1 from 9 to 8
move 1 from 8 to 1
move 10 from 1 to 9
move 1 from 4 to 2
move 2 from 9 to 5
move 4 from 9 to 6
move 1 from 2 to 7
move 3 from 4 to 2
move 1 from 1 to 5
move 5 from 1 to 5
move 1 from 4 to 9
move 3 from 6 to 7
move 23 from 7 to 6
move 1 from 2 to 4
move 1 from 2 to 5
move 9 from 5 to 4
move 1 from 2 to 5
move 9 from 5 to 6
move 1 from 9 to 7
move 1 from 9 to 3
move 3 from 9 to 4
move 14 from 6 to 3
move 5 from 7 to 4
move 1 from 7 to 5
move 1 from 5 to 9
move 2 from 5 to 6
move 16 from 6 to 2
move 2 from 6 to 1
move 7 from 4 to 8
move 2 from 1 to 2
move 4 from 3 to 5
move 5 from 4 to 7
move 2 from 6 to 7
move 4 from 4 to 1
move 4 from 8 to 9
move 1 from 4 to 5
move 1 from 6 to 8
move 1 from 4 to 9
move 4 from 1 to 7
move 1 from 9 to 4
move 2 from 2 to 7
move 7 from 3 to 9
move 15 from 2 to 3
move 4 from 8 to 6
move 1 from 4 to 7
move 2 from 9 to 7
move 1 from 6 to 8
move 2 from 7 to 2
move 5 from 7 to 2
move 1 from 5 to 2
move 6 from 2 to 9
move 3 from 7 to 1
move 3 from 1 to 2
move 3 from 7 to 1
move 2 from 2 to 9
move 2 from 6 to 9
move 1 from 8 to 3
move 19 from 3 to 9
move 1 from 6 to 3
move 3 from 7 to 4
move 1 from 2 to 5
move 2 from 1 to 9
move 2 from 2 to 3
move 33 from 9 to 7
move 1 from 1 to 7
move 3 from 3 to 7
move 1 from 3 to 2
move 1 from 5 to 8
move 4 from 9 to 7
move 1 from 5 to 2
move 2 from 4 to 9
move 4 from 9 to 7
move 3 from 2 to 1
move 1 from 4 to 3
move 1 from 9 to 7
move 1 from 8 to 3
move 7 from 7 to 3
move 3 from 1 to 9
move 4 from 9 to 7
move 4 from 5 to 8
move 3 from 3 to 4
move 3 from 4 to 5
move 3 from 3 to 6
move 2 from 6 to 5
move 38 from 7 to 5
move 40 from 5 to 3
move 4 from 8 to 9
move 1 from 6 to 9
move 1 from 5 to 1
move 3 from 7 to 6
move 1 from 7 to 5
move 38 from 3 to 8
move 1 from 1 to 9
move 3 from 9 to 6
move 5 from 3 to 9
move 4 from 8 to 6
move 1 from 7 to 1
move 3 from 5 to 9
move 1 from 1 to 2
move 10 from 8 to 3
move 5 from 8 to 1
move 3 from 1 to 2
move 9 from 6 to 7
move 9 from 3 to 5
move 1 from 7 to 6
move 1 from 3 to 8
move 1 from 7 to 9
move 1 from 1 to 5
move 1 from 1 to 3
move 1 from 9 to 2
move 4 from 2 to 3
move 1 from 2 to 4
move 9 from 8 to 1
move 2 from 9 to 5
move 2 from 1 to 2
move 2 from 3 to 4
move 6 from 8 to 6
move 10 from 5 to 3
move 7 from 3 to 2
move 2 from 1 to 2
move 5 from 1 to 7
move 7 from 9 to 6
move 7 from 6 to 5
move 1 from 4 to 3
move 7 from 7 to 4
move 5 from 3 to 9
move 7 from 2 to 6
move 4 from 7 to 8
move 5 from 8 to 9
move 1 from 2 to 6
move 1 from 3 to 5
move 2 from 2 to 8
move 8 from 4 to 6
move 7 from 9 to 7
move 4 from 7 to 9
move 7 from 9 to 3
move 8 from 3 to 1
move 6 from 5 to 9
move 8 from 1 to 8
move 13 from 8 to 4
move 3 from 9 to 6
move 1 from 8 to 6
move 1 from 7 to 3
move 2 from 4 to 1
move 5 from 9 to 1
move 1 from 3 to 7
move 15 from 6 to 1
move 1 from 7 to 9
move 10 from 4 to 7
move 11 from 7 to 5
move 17 from 1 to 6
move 1 from 9 to 3
move 6 from 6 to 1
move 3 from 5 to 3
move 2 from 4 to 5
move 2 from 7 to 8
move 12 from 5 to 3
move 13 from 6 to 9
move 2 from 8 to 2
move 2 from 5 to 1
move 16 from 3 to 8
move 3 from 2 to 3
move 2 from 3 to 7
move 2 from 7 to 9
move 1 from 3 to 7
move 4 from 8 to 4
move 2 from 4 to 8
move 5 from 1 to 5
move 2 from 4 to 7
move 6 from 6 to 8
move 2 from 8 to 5
move 2 from 1 to 4
move 5 from 8 to 7
move 5 from 6 to 3
move 6 from 9 to 8
move 2 from 9 to 2
move 1 from 1 to 7
move 4 from 5 to 3
move 2 from 2 to 3
move 1 from 4 to 9
move 10 from 3 to 6
move 1 from 3 to 7
move 10 from 7 to 2
move 2 from 5 to 3
move 1 from 4 to 2
move 2 from 6 to 8
move 3 from 6 to 5
move 1 from 6 to 1
move 7 from 2 to 3
move 6 from 8 to 7
move 4 from 6 to 3
move 14 from 8 to 6
move 11 from 6 to 8
move 1 from 1 to 4
move 6 from 7 to 2
move 3 from 5 to 8
move 4 from 1 to 7
move 1 from 2 to 8
move 1 from 2 to 6
move 1 from 3 to 4
move 1 from 5 to 6
move 7 from 8 to 6
move 9 from 3 to 2
move 1 from 8 to 5`

func main() {
	input := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(input)

	stacks := initStacks()

	for scanner.Scan() {
		n, from, to := extractOp(scanner.Text())

		for i := 0; i < n; i++ {
			stacks[to].Stack(stacks[from].Unstack())
		}
	}

	var message []rune
	for _, s := range stacks {
		message = append(message, s.Unstack())
	}

	fmt.Println(string(message))
}

func extractOp(line string) (int, int, int) {
	re := regexp.MustCompile(`\d[\d,]*`)
	digits := re.FindAllString(line, -1)
	if len(digits) != 3 {
		panic("invalid line format")
	}

	n, err := strconv.ParseInt(digits[0], 10, 64)
	if err != nil {
		panic(err)
	}

	from, err := strconv.ParseInt(digits[1], 10, 64)
	if err != nil {
		panic(err)
	}

	to, err := strconv.ParseInt(digits[2], 10, 64)
	if err != nil {
		panic(err)
	}

	return int(n), int(from) - 1, int(to) - 1
}

func initStacks() []*DoublyLinkedStack {
	stacks := make([]*DoublyLinkedStack, 9)

	// 1
	stacks[0] = NewDoublyLinkedStack()
	stacks[0].Stack('Q')
	stacks[0].Stack('F')
	stacks[0].Stack('M')
	stacks[0].Stack('R')
	stacks[0].Stack('L')
	stacks[0].Stack('W')
	stacks[0].Stack('C')
	stacks[0].Stack('V')

	// 2
	stacks[1] = NewDoublyLinkedStack()
	stacks[1].Stack('D')
	stacks[1].Stack('Q')
	stacks[1].Stack('L')

	// 3
	stacks[2] = NewDoublyLinkedStack()
	stacks[2].Stack('P')
	stacks[2].Stack('S')
	stacks[2].Stack('R')
	stacks[2].Stack('G')
	stacks[2].Stack('W')
	stacks[2].Stack('C')
	stacks[2].Stack('N')
	stacks[2].Stack('B')

	// 4
	stacks[3] = NewDoublyLinkedStack()
	stacks[3].Stack('L')
	stacks[3].Stack('C')
	stacks[3].Stack('D')
	stacks[3].Stack('H')
	stacks[3].Stack('B')
	stacks[3].Stack('Q')
	stacks[3].Stack('G')

	// 5
	stacks[4] = NewDoublyLinkedStack()
	stacks[4].Stack('V')
	stacks[4].Stack('G')
	stacks[4].Stack('L')
	stacks[4].Stack('F')
	stacks[4].Stack('Z')
	stacks[4].Stack('S')

	// 6
	stacks[5] = NewDoublyLinkedStack()
	stacks[5].Stack('D')
	stacks[5].Stack('G')
	stacks[5].Stack('N')
	stacks[5].Stack('P')

	// 7
	stacks[6] = NewDoublyLinkedStack()
	stacks[6].Stack('D')
	stacks[6].Stack('Z')
	stacks[6].Stack('P')
	stacks[6].Stack('V')
	stacks[6].Stack('F')
	stacks[6].Stack('C')
	stacks[6].Stack('W')

	// 8
	stacks[7] = NewDoublyLinkedStack()
	stacks[7].Stack('C')
	stacks[7].Stack('P')
	stacks[7].Stack('D')
	stacks[7].Stack('M')
	stacks[7].Stack('S')

	// 9
	stacks[8] = NewDoublyLinkedStack()
	stacks[8].Stack('Z')
	stacks[8].Stack('N')
	stacks[8].Stack('W')
	stacks[8].Stack('T')
	stacks[8].Stack('V')
	stacks[8].Stack('M')
	stacks[8].Stack('P')
	stacks[8].Stack('C')

	return stacks
}

type DoublyLinkedStack struct {
	head   *DoublyLinkedNode
	tail   *DoublyLinkedNode
	length int
}

type DoublyLinkedNode struct {
	data rune
	next *DoublyLinkedNode
	prev *DoublyLinkedNode
}

func NewDoublyLinkedStack() *DoublyLinkedStack {
	return &DoublyLinkedStack{nil, nil, 0}
}

func (l *DoublyLinkedStack) Empty() bool {
	return l.length == 0
}

func (l *DoublyLinkedStack) Prepend(data rune) {
	newNode := &DoublyLinkedNode{data, nil, nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.length = 1
		return
	}

	l.head.prepend(newNode)
	if l.length == 1 {
		l.tail = l.head
	}

	l.head = newNode
	l.length++
}

func (l *DoublyLinkedStack) Stack(data rune) {
	l.Prepend(data)
}

func (l *DoublyLinkedStack) Shift() rune {
	var r rune

	if l.Empty() {
		return r
	}

	elem := l.head
	l.head = elem.Next()

	if l.length == 1 {
		l.tail = nil
	}

	elem.unlink()
	l.length--

	return elem.Data()
}

func (l *DoublyLinkedStack) Unstack() rune {
	return l.Shift()
}

func (n *DoublyLinkedNode) Data() rune {
	var r rune

	if n == nil {
		return r
	}

	return n.data
}

func (n *DoublyLinkedNode) Next() *DoublyLinkedNode {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *DoublyLinkedNode) Prev() *DoublyLinkedNode {
	if n == nil {
		return nil
	}
	return n.prev
}

func (n *DoublyLinkedNode) prepend(new *DoublyLinkedNode) {
	new.next = n
	newPrev := n.Prev()
	new.prev = newPrev
	if newPrev != nil {
		newPrev.next = new
	}
	n.prev = new
}

func (n *DoublyLinkedNode) unlink() {
	prev := n.Prev()
	next := n.Next()
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
}
