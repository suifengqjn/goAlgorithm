package linkNode

import "fmt"

//链表常见操作
const ERROR = -2134567
type Element int

//链表结构
// 头节点不存元素
type LinkNode struct {
	Data Element
	NextNode *LinkNode
}


type LinkNodeIn interface {
	Add(head* LinkNode,data Element, isHead bool)
	AddHead(head* LinkNode,data Element)
	Append(head* LinkNode,data Element)
	Show(head* LinkNode)
	GetLength(head* LinkNode)int
	Search(head* LinkNode,data Element)bool
	GetData(head* LinkNode,id int)Element
	Delete(head* LinkNode,id int)Element
	clear(head *LinkNode)
	Insert(head *LinkNode,index int ,data Element)

	Reverse(head* LinkNode)
	BubbleSort(head* LinkNode)
	SelectSort(head* LinkNode)

	InsertSort(head* LinkNode)
	MergeSort(head* LinkNode)
	QuickSort(head* LinkNode)
}

func Add(head* LinkNode,data Element, isHead bool)  {
	if isHead {
		AddHead(head, data)
	} else {
		Append(head, data)
	}
}

// 头插
func AddHead(head* LinkNode,data Element)  {

	newNode := &LinkNode{
		Data:     data,
		NextNode: nil,
	}
	firstNode := head.NextNode
	head.NextNode = newNode
	newNode.NextNode = firstNode
}

//尾插
func Append(head* LinkNode,data Element) {
	newNode := &LinkNode{
		Data:     data,
		NextNode: nil,
	}

	cursor := head
	for cursor.NextNode != nil {
		cursor = cursor.NextNode
	}

	cursor.NextNode = newNode

}

func Show(head* LinkNode)  {
	node := head.NextNode
	for node != nil {
		fmt.Println(node.Data)
		node = node.NextNode
	}
	fmt.Println("show over")
}

// 链表长度
func GetLength(head* LinkNode)int  {
	var length = 0
	cursor := head
	for cursor.NextNode != nil {
		length ++
		cursor = cursor.NextNode
	}

	return length
}

// 查找元素是否存在
func Search(head* LinkNode,data Element)bool  {
	var exist = false
	cursor := head
	for cursor.NextNode != nil {
		if data == cursor.Data {
			exist = true
			break
		}
		cursor = cursor.NextNode
	}
	return exist
}