package main

import(
	"crypto/sha256"
	"fmt"
)
func main(){
	//计算哈希值
	hash := sha256.Sum256([]byte("welcome to blockchain"))
	fmt.Printf("%x\n",hash)
}

//在sha256的包中有Sum256函数，可以直接得到一个32字节的哈希值
func Sum256(data []btye) [32]byte{
	var d digest
	d.Reset()
	d.Write(data)
	return d.checkSum()
}