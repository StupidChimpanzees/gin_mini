package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"strings"
)

// Base64Encrypt Base64加密
func Base64Encrypt(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decrypt Base64解密
func Base64Decrypt(data []byte) ([]byte, error) {
	baseEncode := base64.StdEncoding
	buf := make([]byte, baseEncode.EncodedLen(len(data)))
	_, err := baseEncode.Decode(buf, data)
	return buf, err
}

func Md5Encode(data []byte) string {
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func Md5Check(content, encrypted string) bool {
	return strings.EqualFold(Md5Encode([]byte(content)), encrypted)
}

func Sha1Encode(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha1Check(content, encrypted string) bool {
	return strings.EqualFold(Sha1Encode([]byte(content)), encrypted)
}

func Sha256Encode(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha256Check(content, encrypted string) bool {
	return strings.EqualFold(Sha256Encode([]byte(content)), encrypted)
}

func Sha512Encode(data []byte) string {
	hash := sha512.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha512Check(content, encrypted string) bool {
	return strings.EqualFold(Sha512Encode([]byte(content)), encrypted)
}

// DesEncrypt DES加密
func DesEncrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	encryptBytes := pkcs7Padding(data, block.BlockSize())
	dst := make([]byte, len(encryptBytes))
	block.Encrypt(dst, encryptBytes)
	return dst, nil
}

// DesDecrypt DES解密
func DesDecrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, len(data))
	block.Decrypt(dst, data)
	dst, _ = pkcs7UnPadding(dst)
	return dst, nil
}

// AesEncrypt AES加密
func AesEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	encrypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return encrypted, nil
}

// AesDecrypt AES解密
func AesDecrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)
	crypto, err := pkcs7UnPadding(decrypted)
	if err != nil {
		return nil, err
	}
	return crypto, nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("string error: data is empty")
	}
	unPadding := int(data[length-1])
	return data[:length-unPadding], nil
}

// RsaEncrypt RSA加密
func RsaEncrypt(originalData, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	public := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, public, originalData)
}

// RsaDecrypt RSA解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, private, ciphertext)
}
