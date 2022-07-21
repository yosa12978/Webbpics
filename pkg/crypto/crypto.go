package crypto

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func NewMD5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func NewToken(buff_size int) string {
	buffer := make([]byte, buff_size)
	rand.Read(buffer)
	return fmt.Sprintf("%x", buffer)
}
