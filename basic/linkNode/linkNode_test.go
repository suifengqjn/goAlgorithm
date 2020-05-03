package linkNode

import (
	"fmt"
	"testing"
)

var testNode = new(LinkNode)
func init() {

	AddHead(testNode, 1)
	AddHead(testNode, 2)
	AddHead(testNode, 3)

}

func TestAddHead(t *testing.T) {
	node := new(LinkNode)

	AddHead(node, 1)
	AddHead(node, 2)
	AddHead(node, 3)

	Show(node)
}

func TestAppend(t *testing.T) {
	node := new(LinkNode)

	Append(node, 1)
	Append(node, 2)
	Append(node, 3)

	Show(node)
}

func TestGetLength(t *testing.T) {
	node := new(LinkNode)

	Append(node, 1)
	Append(node, 2)
	Append(node, 3)

	fmt.Println(GetLength(node))
}

func TestSearch(t *testing.T) {
	fmt.Println(Search(testNode, 2))
	fmt.Println(Search(testNode, 4))
}