package web_test

import (
	"crypto/aes"
	"testing"

	"github.com/magnolia/magnolia/web"
)

func TestTextEncryptor(t *testing.T) {
	cip, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		t.Fatal(err)
	}

	enc := web.TextEncryptor{Cipher: cip}
	hello := "Hello, magnolia!"
	code, err := enc.Encode([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	plain, err := enc.Decode(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("wang %s get %s", hello, string(plain))
	}
}

func TestPasswordEncryptor(t *testing.T) {
	enc := web.PasswordEncryptor{}
	plain := "123456"
	code, err := enc.Sum([]byte(plain), 8)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("doveadm pw -t {SSHA512}%s -p %s", code, plain)
	rst, err := enc.Equal([]byte(plain), code)
	if err != nil {
		t.Fatal(err)
	}
	if !rst {
		t.Fatalf("check password failed")
	}
}
