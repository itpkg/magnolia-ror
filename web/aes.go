package web

import "crypto/rand"

//RandomBytes random bytes
func RandomBytes(n uint) ([]byte, error) {
	buf := make([]byte, n)
	_, err := rand.Read(buf)
	return buf, err
}
