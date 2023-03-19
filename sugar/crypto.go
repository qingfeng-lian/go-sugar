package sugar

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	has := md5.New()       // 创建md5算法
	has.Write([]byte(str)) // 写入需要加密的数据
	b := has.Sum(nil)      // 获取hash值字符切片；Sum函数接受一个字符切片，这个切片的内容会原样的追加到abc123加密后的hash值的前面，这里我们不需要这么做，所以传入nil
	// fmt.Println(b)         // 打印一下 [233 154 24 196 40 203 56 213 242 96 133 54 120 146 46 3]
	// 上面可以看到加密后的数据为长度为16位的字符切片，一般我们会把它转为16进制，方便存储和传播，下一步转换16进制
	// fmt.Println(hex.EncodeToString(b)) // 通过hex包的EncodeToString函数，将数据转为16进制字符串； 打印 e99a18c428cb38d5f260853678922e03
	// 还有一种方法转换为16进制,通过fmt的格式化打印方法， %x表示转换为16进制
	return fmt.Sprintf("%x", b) // 打印 e99a18c428cb38d5f260853678922e03
}
