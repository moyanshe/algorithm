package classics

import (
	"container/list"
	"fmt"
)

//Tree 树接口
type Tree interface {
	Do()
}

//Node 树节点
type Node struct {
	*list.List
	Name string
}

//Do 定义树方法，同时也自动实现了树的接口
func (n *Node) Do() {
	fmt.Println(n.Name)
	for e := n.Front(); e != nil; e = e.Next() {
		e.Value.(Tree).Do()
	}
}

//AddSub 添加子节点
func (n *Node) AddSub(sub Tree) {
	n.PushBack(sub)
}

//Leaf 叶子节点
type Leaf struct {
	Name string
}

//Do 叶子节点方法，同时自动实现了树的接口
func (l *Leaf) Do() {
	fmt.Println(l.Name)
}
