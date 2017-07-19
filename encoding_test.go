package lutils

import (
	"testing"
	"log"
	"fmt"
)

func TestGBK2UTF8(t *testing.T) {
	src:=string([]byte{214,208,206,196})
	log.Println(GBK2UTF8(src))
}

func TestUTF8_2GBK(t *testing.T) {
	src:=string([]byte{228,184,173,230,150,135})
	fmt.Println([]byte(src))
	log.Println([]byte(UTF8_2GBK(src)))
}
