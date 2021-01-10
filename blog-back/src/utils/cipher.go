package utils

import (
	"github.com/go-jar/crypto"
	"github.com/go-jar/encoding"
)

func newCrypter() *crypto.AesCBCCrypter{
	key := crypto.Md5String([]byte("blog"))
	iv := crypto.Md5String([]byte("back"))[:crypto.AesBlockSize]

	cipherCrypter, _ := crypto.NewAesCBCCrypter([]byte(key), []byte(iv))
	return cipherCrypter
}

func EncryptString(str string) string {
	cipherCrypter := newCrypter()
	crypted := encoding.Base64Encode(cipherCrypter.Encrypt([]byte(str)))

	return string(crypted)
}

func DecryptString(crypted string) string {
	cipherCrypter := newCrypter()
	decrypted := cipherCrypter.Decrypt(encoding.Base64Decode([]byte(crypted)))

	return string(decrypted)
}
