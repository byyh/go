package com

import (
    "fmt"    
    "crypto"
    "errors"  
    "strings"
    "crypto/md5"
    "crypto/sha256"
    "crypto/hmac"
    "crypto/sha1" 
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "encoding/hex"
)

// 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte) {
    //解密pem格式的公钥
    block, _ := pem.Decode(publicKey)
    if block == nil {
        fmt.Println("public key error")
        panic(errors.New("public key error"))
    }
    // 解析公钥
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        fmt.Println("解析公钥错误")
        panic(err)
    }
    // 类型断言
    pub := pubInterface.(*rsa.PublicKey)

    //加密
    ret,err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
    if(nil != err) {
        fmt.Println("加密错误")
        panic(err)
    }

    return ret
}

// 解密
func RsaDecrypt(ciphertext []byte, privateKey string) ([]byte, error) {
    //解密
    block, _ := pem.Decode([]byte(privateKey))
    if block == nil {
        fmt.Println("解析私钥错误")
        panic(errors.New("private key error"))
    }
    //解析PKCS1格式的私钥
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        fmt.Println("解析私钥错误")
        panic(err)
    }
    // 解密
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// hash = crypto.SHA256 or crypto.SHA1
// 返回的是base64编码
func SignRsa(src []byte, privateKey string, hash crypto.Hash) (string) {
    var h = hash.New()
    h.Write(src)
    var hashed = h.Sum(nil)

    var err error
    var block *pem.Block
    block, _ = pem.Decode([]byte(privateKey))
    if block == nil {
        fmt.Println("解析私钥错误")
        panic(errors.New("private key error"))
    }

    var prvte *rsa.PrivateKey
    prvte, err = x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        fmt.Println("解析私钥错误2")
        panic(err)
    }

    retData, err := rsa.SignPKCS1v15(rand.Reader, prvte, hash, hashed)
    if(nil != err) {
        fmt.Println("签名错误")
        panic(err)
    }

    return Base64(retData)
}

// 返回的是16进制的字符串，不用base64编码
func SignRsaToHex(src []byte, privateKey string, hash crypto.Hash) (string) {
    var h = hash.New()
    h.Write(src)
    var hashed = h.Sum(nil)

    var err error
    var block *pem.Block
    block, _ = pem.Decode([]byte(privateKey))
    if block == nil {
        fmt.Println("解析私钥错误")
        panic(errors.New("private key error"))
    }

    var prvte *rsa.PrivateKey
    prvte, err = x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        fmt.Println("解析私钥错误2")
        panic(err)
    }

    retData, err := rsa.SignPKCS1v15(rand.Reader, prvte, hash, hashed)
    if(nil != err) {
        fmt.Println("签名错误")
        panic(err)
    }

    return hex.EncodeToString(retData)
}

func CheckSignRsa(src []byte, sign string, publicKey []byte, hash crypto.Hash) bool {
    var h = hash.New()
    h.Write(src)
    var hashed = h.Sum(nil)

    var block *pem.Block
    block, _ = pem.Decode(publicKey)
    if block == nil {
        panic(errors.New("public key error"))
    }

    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        fmt.Println("解析公钥错误", err)

        panic(err)
    }
    var pub = pubInterface.(*rsa.PublicKey)

    err = rsa.VerifyPKCS1v15(pub, hash, hashed, []byte(sign))
    if err != nil {
        return false
    }

    return true
}

func Base64(data []byte) string {
    return base64.StdEncoding.EncodeToString(data)
}


func Md5(str string) string {
    h := md5.New()
    h.Write([]byte(str)) // 需要加密的字符串为 str
    byteMd5 := h.Sum(nil) 
    strMd5 := hex.EncodeToString(byteMd5)

    return strings.ToUpper(strMd5)
}

func Sha256Code(str string, key string) string {
    h := sha256.New()
    h.Write([]byte(str))
    byteRet := h.Sum([]byte(key))
    strReg := hex.EncodeToString(byteRet)

    return strings.ToUpper(strReg)
}

func HmacMd5(data string, key string) string {
    hmac := hmac.New(md5.New, []byte(key))
    hmac.Write([]byte(data))

    return Base64(hmac.Sum(nil))
}

func HmacSha1(data string, key string) string {
    hmac := hmac.New(sha1.New, []byte(key))
    hmac.Write([]byte(data))

    return Base64(hmac.Sum(nil))
}
 
func Sha1(data string, key string) string {
    sha1 := sha1.New()
    sha1.Write([]byte(data))

    //return hex.EncodeToString()
    return Base64(sha1.Sum([]byte(key)))
}