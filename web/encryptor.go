package web

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

//TextEncryptor text encryptor
type TextEncryptor struct {
	Cipher cipher.Block `inject:""`
}

//Encode encode buffer
func (p *TextEncryptor) Encode(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cipher, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

//Decode decode buffer
func (p *TextEncryptor) Decode(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cipher, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

//-----------------------------------------------------------------------------

//PasswordEncryptor password encryptor
type PasswordEncryptor struct {
}

//Sum sha512 with salt
func (p *PasswordEncryptor) Sum(plain []byte, num uint) (string, error) {
	salt, err := RandomBytes(num)
	if err != nil {
		return "", err
	}
	return p.sum(plain, salt)
}

//Equal compare
func (p *PasswordEncryptor) Equal(plain []byte, code string) (bool, error) {
	buf, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return false, err
	}

	if len(buf) <= sha512.Size {
		return false, err
	}
	salt := buf[sha512.Size:]
	rst, err := p.sum(plain, salt)
	if err != nil {
		return false, err
	}
	return rst == code, nil
}

func (p *PasswordEncryptor) sum(plain, salt []byte) (string, error) {
	buf := append([]byte(plain), salt...)
	code := sha512.Sum512(buf)
	return base64.StdEncoding.EncodeToString(append(code[:], salt...)), nil
	// return base64.StdEncoding.EncodeToString(append(p.Hash.Sum(append(plain, salt...)), salt...)), nil
}
