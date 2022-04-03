package ecb_test

import (
	"crypto/aes"
	"fmt"

	"github.com/clfs/ecb"
)

func ExampleNewEncrypter() {
	key := []byte{
		0x65, 0x61, 0x73, 0x74, 0x65, 0x72, 0x20, 0x65,
		0x67, 0x67, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	}

	plaintext := []byte("exampleplaintext")
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := ecb.NewEncrypter(block)
	ciphertext := make([]byte, len(plaintext))
	mode.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("%x\n", ciphertext)

	// Output:
	// 188dc8b928ae0a13b0e9ab01127fd341
}

func ExampleNewDecrypter() {
	key := []byte{
		0x65, 0x61, 0x73, 0x74, 0x65, 0x72, 0x20, 0x65,
		0x67, 0x67, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	}
	ciphertext := []byte{
		0x18, 0x8d, 0xc8, 0xb9, 0x28, 0xae, 0x0a, 0x13,
		0xb0, 0xe9, 0xab, 0x01, 0x12, 0x7f, 0xd3, 0x41,
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	mode := ecb.NewDecrypter(block)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	fmt.Printf("%s\n", plaintext)

	// Output:
	// exampleplaintext
}
