package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

type FileCrypter interface {
	EncryptFile(key []byte, inputFile, outputFile string) error
	DecryptFile(key []byte, inputFile, outputFile string) error
}

type AESFileCrypter struct{}

func (a AESFileCrypter) EncryptFile(key []byte, inputFile, outputFile string) error {
	// Open the input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Generate the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create the AES cipher stream
	stream := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])

	// Create a buffer to hold the encrypted data
	buffer := make([]byte, 4096)

	// Encrypt and write the data to the output file
	for {
		n, err := inFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = outFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

func (a AESFileCrypter) DecryptFile(key []byte, inputFile, outputFile string) error {
	// Open the input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// Create the output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Generate the AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Create the AES cipher stream
	stream := cipher.NewCFBDecrypter(block, key[:block.BlockSize()])

	// Create a buffer to hold the decrypted data
	buffer := make([]byte, 4096)

	// Decrypt and write the data to the output file
	for {
		n, err := inFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = outFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}
