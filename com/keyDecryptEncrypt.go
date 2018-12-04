package com

import (
    "fmt"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "crypto"
    "errors"    
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