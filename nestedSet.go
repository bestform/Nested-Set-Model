package main

import "fmt"

type Node struct {
	name        string
	left, right int
	children    []*Node
}

func (n *Node) visit(c int) int {
	c += 1
	n.left = c
	for _, child := range n.children {
		c = child.visit(c)
	}
	c += 1
	n.right = c

	return c
}

func (n *Node) Init() {
	n.visit(0)
}

func (n *Node) append(nn *Node) {
	n.children = append(n.children, nn)
}

func (n *Node) appendNew(name string) *Node {
	nn := NewNode(name)
	n.append(&nn)

	return &nn
}

func (n Node) String() string {
	output := fmt.Sprintf("%v (%v, %v)\n", n.name, n.left, n.right)
	for _, child := range n.children {
		output += fmt.Sprint(child)
	}

	return output
}

func NewNode(name string) Node {
	var n Node
	n.name = name
	n.children = make([]*Node, 0, 0)

	return n
}

func main() {
	clothing := NewNode("Clothing")
	suits := clothing.appendNew("Men's").appendNew("Suits")
	suits.appendNew("Slacks")
	suits.appendNew("Jackets")
	womens := clothing.appendNew("Womens's")
	dresses := womens.appendNew("Dresses")
	dresses.appendNew("Evening Growns")
	dresses.appendNew("Sun Dresses")
	womens.appendNew("Skirts")
	womens.appendNew("Blouses")

	clothing.Init()
	fmt.Println(clothing)
}
