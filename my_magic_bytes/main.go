package main

import (
	"io/ioutil"
	"os"
)

var (
	key = []byte{0x46, 0xcc, 0xf9, 0xa5, 0x71, 0xf0, 0xff, 0xb1, 0x7e, 0x41, 0xcb, 0x84}
)

func encrypt(input, key []byte) []byte {
	l := len(key)
	out := make([]byte, len(input))
	for i := range input {
		out[i] = input[i] ^ key[i%l]
	}
	return out
}

func main() {
	f, err := os.Open("./my_magic_bytes.jpg.enc")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := os.Create("./res.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	decrypted := encrypt(data, key)
	out.Write(decrypted)
}
