package utils

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
)
func padPwd(srcByte []byte,blockSize int)  []byte {
	// 16 13       13-3 = 10
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}

// 去掉填充的部分
func unPadPwd(dst []byte) ([]byte,error) {
	if len(dst) <= 0 {
		return dst, errors.New("长度有误")
	}
	// 去掉的长度
	unpadNum := int(dst[len(dst)-1])
	return dst[:(len(dst) - unpadNum)], nil
}

// 只支持8字节的长度
var desKey = []byte("20210604")
// 加密
func DesEncoding(src string) (string,error) {
	srcByte := []byte(src)
	block, err := des.NewCipher(desKey)
	if err != nil {
		return src, err
	}
	// 密码填充
	newSrcByte := padPwd(srcByte, block.BlockSize())
	dst := make([]byte, len(newSrcByte))
	block.Encrypt(dst, newSrcByte)
	// base64编码
	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd, nil
}

// 解密
func DesDecoding(pwd string) (string,error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd, err
	}
	block, errBlock := des.NewCipher(desKey)
	if errBlock != nil {
		return pwd, errBlock
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte)
	// 填充的要去掉
	dst, _ = unPadPwd(dst)
	return string(dst), nil
}