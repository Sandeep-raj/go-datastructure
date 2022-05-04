package bst

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Left View Of A Binary Tree

1. Given a Binary Tree, print Right view of it.
2. Right view of a Binary Tree is set of nodes visible when tree is visited from Right side.


*/

func TestLRViewBT() {
	bst := utils.InitBST()
	bst.BSTInsert(*utils.InitBSTNode(20, ""))
	bst.BSTInsert(*utils.InitBSTNode(10, ""))
	bst.BSTInsert(*utils.InitBSTNode(30, ""))
	bst.BSTInsert(*utils.InitBSTNode(5, ""))
	bst.BSTInsert(*utils.InitBSTNode(15, ""))
	bst.BSTInsert(*utils.InitBSTNode(12, ""))
	bst.BSTInsert(*utils.InitBSTNode(18, ""))
	bst.BSTInsert(*utils.InitBSTNode(25, ""))
	bst.BSTInsert(*utils.InitBSTNode(35, ""))

	//bst.BSTTraversal("pre")

	lrView(bst)
}

func lrView(bst *utils.BST) {
	res1 := make([]int, 0)
	leftViewhelper(bst.Root, 0, &res1)

	res2 := make([]int, 0)
	rightViewhelper(bst.Root, 0, &res2)

	log.Printf("%+v", res1)
	log.Printf("%+v", res2)
}

func leftViewhelper(node *utils.BST_node, level int, resArr *[]int) {
	if node == nil {
		return
	}

	if len(*resArr)-1 < level {
		*resArr = append(*resArr, node.Key)
	}

	leftViewhelper(node.Left, level+1, resArr)
	leftViewhelper(node.Right, level+1, resArr)
}

func rightViewhelper(node *utils.BST_node, level int, resArr *[]int) {
	if node == nil {
		return
	}

	if len(*resArr)-1 < level {
		*resArr = append(*resArr, node.Key)
	}

	rightViewhelper(node.Right, level+1, resArr)
	rightViewhelper(node.Left, level+1, resArr)
}
