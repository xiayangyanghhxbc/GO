package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

/*
	定义merkletree的数据结构
*/

// MerkleTree 结构
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode 结构
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

/*
	编写构造节点函数
	通用，既要支持中间节点的，也要支持叶子节点的
*/

// 创建节点
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	//如果左右为空，则代表其对应的data是最原始的数据节点
	if left == nil && right == nil {
		//计算哈希
		hash := sha256.Sum256(data)
		//将32字节的[32]byte转换为[]byte
		mNode.Data = hash[:]
	} else {
		//将左右子树的数据集合在一起
		prevHashes := append(left.Data, right.Data...)
		//计算哈希
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}
	//左右子树赋值
	mNode.Left = left
	mNode.Right = right

	return &mNode
}

/*
	将节点组建为树
	其实利用NewMerkleNode方法已经可以自动组建成树
	这一步，是将这个过程自动化了

	并使用MerkleNode类型的切片来保存所有的数据节点
*/
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode
	//必须是2的整数倍节点
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	//	两层完成节点的树型构造
	for i := 0; i < len(data)/2; i++ {
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		for _, v := range nodes {
			fmt.Printf("i = %d, %p,%p,%x\n", i, v.Left, v.Right, v.Data)
		}
		fmt.Println()
		nodes = newLevel
		for _, v := range nodes {
			fmt.Printf("i = %d, %p,%p,%x\n", i, v.Left, v.Right, v.Data)
		}
		fmt.Println()
	}
	fmt.Println("--------------\n\n")
	//构造MerkleTree
	mTree := MerkleTree{&nodes[0]}

	return &mTree
}
/*
	遍历函数
*/

func showMerkleTree(root *MerkleNode) {

	if root == nil {
		return
	} else {
		//打印节点信息
		PrintNode(root)
	}
	showMerkleTree(root.Left)
	showMerkleTree(root.Right)
}

/*
	节点信息打印函数
*/
func PrintNode(node *MerkleNode) {
	fmt.Printf("%p\n", node)
	if node != nil {
		fmt.Printf("letf[%p],right[%p],data(%x)\n", node.Left, node.Right, node.Data)
		fmt.Printf("check:%t\n", check(node))
	}
}

/*
	增加检查的逻辑
	将左右两个子树的data值联合起来计算hash值
	并与父节点的data进行比较
*/
func check(node *MerkleNode) bool {

	if node.Left == nil {
		return true
	}
	prevHashes := append(node.Left.Data, node.Right.Data...)
	hash32 := sha256.Sum256(prevHashes)
	hash := hash32[:]
	return bytes.Compare(hash, node.Data) == 0
}

/*
	是[]byte类型的比较方法，当返回值是0时，代表两个切片相等。
*/
func ByteSliceEqualBCE(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main(){
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
	}

	tree := NewMerkleTree(data)
	showMerkleTree(tree.RootNode)
}
