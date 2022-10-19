package xmezum

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/bits"
	"time"
)

func encrypt(key []byte, data []byte) (enc []byte, err error) {
	counter := 0
	mod := 0
	keylen := len(key) - 1
	for i := 0; i < len(data); i++ {
		data[i] ^= (bits.RotateLeft8(key[counter], mod) | uint8(keylen))
		if counter >= keylen {
			counter = 0
			mod++
		} else {
			counter++
		}
	}
	return data, nil

}

func decrypt(key []byte, data []byte) (enc []byte, err error) {
	counter := 0
	mod := 0
	keylen := len(key) - 1
	for i := 0; i < len(data); i++ {
		data[i] ^= (bits.RotateLeft8(key[counter], mod) | uint8(keylen))
		if counter >= keylen {
			counter = 0
			mod++
		} else {
			counter++
		}
	}
	return data, nil

}

func main() {
	key := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		panic(err)
	}
	original := []byte("Hello, World!!! How are you today?")
	// Casting to string and back to bytes is required to prevent modification of original slice
	enc, _ := encrypt(key, []byte(string(original)))
	fmt.Print("Key: ")
	fmt.Println(key)
	fmt.Print("OG: ")
	fmt.Println(original)
	fmt.Print("ENC: ")
	fmt.Println(enc)

	fmt.Println("OG Text: " + string(original))
	fmt.Println("ENC Text: " + string(enc))

	start := time.Now()
	randdata := make([]byte, 1024*(1e+6))
	_, err = io.ReadFull(rand.Reader, randdata)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Print("SPEED of filling 1024MB: ")
	fmt.Println(duration.Milliseconds())
	start = time.Now()
	encrypt(key, randdata)
	duration = time.Since(start)
	fmt.Print("SPEED of 1024MB: ")
	fmt.Println(duration.Milliseconds())

	fmt.Print("\nDEC: ")
	dec, _ := decrypt(key, enc)
	fmt.Println(string(dec))
}
