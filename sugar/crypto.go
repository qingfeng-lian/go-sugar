package sugar

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/base64"
	"fmt"
)

// md5 加密
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

// rc4 加密
// 当然解密同样可用这个方法
// 为了方便区分加密解密，建议只用此方法进行加密， 解密可以用：RC4Decrypt
func RC4Encrypt(key, str string) []byte {
	strByte := []byte(str)
	cipher, err := rc4.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	dst := make([]byte, len(strByte))
	cipher.XORKeyStream(dst, strByte)
	return dst
}

// rc4 解密
// 当然加密同样可用这个方法
// 为了方便区分加密解密、 建议只用此方法进行解密、建议加密用： RC4Encrypt
func RC4Decrypt(key, str string) []byte {
	return RC4Encrypt(key, str)
}

// rc4 加密
// 并且返回base64 encode 后的字符串
func RC4EncryptBaseEncode(key, str string) string {
	strByte := RC4Encrypt(key, str)
	return base64.StdEncoding.EncodeToString(strByte)
}

// rc4 解密
// str 是base64 encode 的字符串
func RC4EncryptBaseDecode(key, str string) string {
	src, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	strByte := RC4Decrypt(key, string(src))
	return string(strByte)
}
