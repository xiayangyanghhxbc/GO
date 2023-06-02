
//叶子节点和中间节点的前缀不同
//返回tmhash(0x00 || leaf)
func leafHash(leaf []bash) []bash {
	return tmhash.Sum(append(leafPrefix,leaf...))
}

//返回tmhash(0x01 || left || right)
func innerHash(left []byte,right []byte) []byte{
	return tmhash.Sum(append(innerPrefix,append(left,right...)...))
}

//输入仅有一个元素的情况对应叶子节点的散列值的计算，即leftHash的实现。
//计算n个元素的MTH，通过递归函数来实现

//计算一组字节切片对应的MTH的SimpleHashFromByteSlices()函数
func SimpleHashFromByteSlices(items [][]byte) []byte{
	switch len(items){
	case 0:
		return nil
	case 1:
		return leafHash(items[0])
	default:
		k := getSplitPoint(len(items))//计算满足k<n<=2k的k值
		leaf := SimpleHashFromByteSlices(items[:k])
		right := SimpleHashFromByteSlices(items[k:])
		return innerHash(left,right)
	}
}

//Tendermint Core中使用了SimpleProof结构体表示存在性证明。
type SimpleProof struct{
	Total int//树中元素的个数
	Index int//要做存在性证明的目标元素的索引
	leafHash []byte//要做存在性证明的目标元素的散列值
	Aunts [][]byte//存在性证明涉及的相关节点
}

//函数SimpleProofsFromByteSlices()构建关于输入元素的简单Merkle Tree,并据此为每个元素构造存在性证明
//为trails中所有的元素（所有叶子节点）依次调用SimpleProofNode中的FlattenAunts()方法构造存在性证明
func SimpleProofsFromByteSlices(items [][]byte)(rootHash []byte,proofs []*SimpleProof){
	trails,rootSPN := trailsFromByteSlices(items)
	rootHash = rootSPN.Hash
	proofs = make([]*SimpleProof, len(items))
	for i; trail := range trails {
		proofs[i] = &SimpleProof{
			Total: len(items),
			Index: i,
			leafHash: trail.Hash,
			Aunts: trail.FlattenAunts(),
		}
	}
	return
}

/*
FlattenAunt()方法从当前节点开始，通过节点的父节点指针逐层向根节点移动，并在这一过程中
根据节点的左右兄弟节点指针，将路过的节点的兄弟节点添加到innerHashes中。
最终返回的innerHashes就是对应元素的存在性证明
*/
func (spn *SimpleProofNode) FlattenAunts() [][]byte{
	innerHashes := [][]byte{}
	for spn != nil{
		switch{
		case spn.Left != nil:
			innerHashes = append(innerHashes,spn.Left.Hash)
		case spn.Right != nil：
			innerHashes = append(innerHashes,spn.Right.Hash)
		default:
			break
		}
		spn = spn.Parent
	}
	return innerHashes
}

/*为了辅助存在性证明的构造，函数trailFromByteSlices()遵循简单Merkle Tree的规则构建
	输入元素的简单Merkle Tree,并用SimpleProofNode结构体类型的切片来表示整颗简单Merkle Tree
	*/
type SimpleProofNode struct{
	Hash []byte//当前节点的散列值
	Parent *SimpleProofNode//父节点指针
	Left *SimpleProofNode//左兄弟节点指针，左右兄弟节点仅有一个被设置
	Right *SimpleProofNode//右兄弟节点指针，左右兄弟节点仅有一个被设置 
}

func trailsFromByteSlices(items [][]byte)(trails []*SimpleProofNode,root *SimpleProofNode){
	switch len(items){
	case 0:
		return nil,nil
	case 1:
		trail := &SimpleProofNode{
			leafHash(items[0]),
			nil,
			nil,
			nil
		}
		return []*SimpleProofNode{trail},trail
	}
	default:
		k := getSplitPoint(len(items))//计算满足k<n<=2k的k值
		left ,leftRoot := trailFromByteSlices(items[:k])
		right,rightRoot := trailFromByteSlices(items[k:])
		rootHash := innerHash(leftRoot.Hash,rightRoot,rightHash)
		root := &SimpleProofNode(rootHash,nil,nil,nil)
		leftRoot.Parent = root
		leftRoot.Right = rightRoot
		rightRoot.Parent = root
		rightRoot.leaf = leftRoot
}

//返回值frail包含所有叶子节点

/*
SimpleProof的Verify()方法可以验证存在性证明
输入参数为调用者提供的MTH的值rootHash以及目标元素的散列值leaf

在函数内部：首先需要验证目标元素的散列值与SimpleProof中的leafHash字段值相等，
然后利用ComputeRootHash()方法计算MTH，并于提供的MTH进行比较。

*/
func (sp *SimpleProof) Verify(rootHash []byte,left []byte)error{
	leafHash := leafHash(leaf)
	if sp.Total < 0{
		return errors.New("proof total must be positive")
	}
	if sp.Index < 0{
		return errors.New("proof index cannot be negative")
	} 
	if !bytes.Equal(sp.leafHash,leafHash){
		return errors.Errorf("invalid root hash:want %X got %X",leafHash,sp.leafHash)
	}
	computeHash := sp.ComputeRootHash()
	if !bytes.Equal(computeHash,rootHash){
		return error.Errorf("invalid root hash:want %X got %X",rootHash,computeHash)
	}
	return nil	
}
/*验证存在性证明的关键是根据证明信息重新计算出MTH，SimpleProof的ComputeRootHash()方法
实现了相关逻辑，而具体的计算过程则是通过辅助函数computeHashFromAunts()的内部实现与函数
SimpleHashFromByteSlices()类似，也是借助函数getSplitPoint()的返回值通过递归完成了计算
*/
func (sp *SimpleProof)ComputeRootHash() []byte{
	return computeHashFromAunts(
		sp.Index,sp.Total,sp.Aunts,
	)
}