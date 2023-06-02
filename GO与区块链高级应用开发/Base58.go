package main
//将求出的结果从后到前表示出来即可
import(
	"fmt"
	"math/big"
) 

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
//将字符串逆序
func ReverseBytes(data []byte){
	for i,j := 0,len(data)-1; i < j ;i,j = i +1,j-1{
		data[i],data[j] = data[j],data[i]
	} 
}
//计算Bsae58编码
func Base58Encode(input int64) []byte {
	var result []byte
	//x为商值
	x := big.NewInt(input)
	//计算除数
	base := big.NewInt(int64(len(b58Alphabet)))
	//获取big.Int类型的0
	zero := big.NewInt(0)
	//用于存储余数
	mod := &big.Int{}
	//只要被除数不为0，就继续计算
	for x.Cmp(zero) != 0{
		//求余运算，x = 商值， mod = 余数
		//x/base=mod
		x.DivMod(x,base,mod)
		//取出编码，存储到result中
		result = append(result,b58Alphabet[mod.Int64()])
	}
	//结果逆序
	ReverseBytes(result)
	return result
}

func main(){
	result := Base58Encode(258)
	fmt.Println(string(result))
}
//DivMod方法中x被除数，y除数,m代表余数,z的值在调用后将等于除法的商
//func (z *int)DivMod(x,y,m *Int )(*Int,*Int)