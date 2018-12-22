package main

type Node struct {
	left  *Node
	Data  interface{}
	Right *Node
}
type Initer interface {
	SetData(data interface{})
}
type Operater interface {
	PrintBT()
	Depth() int
	LeafCount() int
}
type Order interface {
	PreOrder()
	InOrder()
	PostOrder()
}

func (n *Node) SetData(data interface{}) {
	n.Data = data
}
func (n *Node) printfBT() {
	printBT(n)
}
func (n *Node) Depth() int {
	return Depth(n)
}
func (n *Node) Depth() int {
	return Depth(n)
}
func (n *Node) leafCount() int {
	return leafCount(n)
}
func (n *Node) PreOrder() {
	preOreder(n)
}
func (n *Node) InOrder() {
	InOrder(n)
}
func (n *Node) postOrser() {
	postOrder(n)
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}
