package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io/ioutil"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		if fileExists(arg) {
			fileBytes, err := ioutil.ReadFile(arg)
			if err != nil {
				panic(err)
			}

			md5chan := make(chan string)
			go doHash(md5.New(), fileBytes, md5chan)

			sha1chan := make(chan string)
			go doHash(sha1.New(), fileBytes, sha1chan)

			sha256chan := make(chan string)
			go doHash(sha256.New(), fileBytes, sha256chan)

			sha512chan := make(chan string)
			go doHash(sha512.New(), fileBytes, sha512chan)

			if len(os.Args) > 2 {
				fmt.Printf("         [%s]\n", arg)
			}
			fmt.Printf("    MD5: %s\n", <-md5chan)
			fmt.Printf("  SHA-1: %s\n", <-sha1chan)
			fmt.Printf("SHA-256: %s\n", <-sha256chan)
			fmt.Printf("SHA-512: %s\n\n", <-sha512chan)
		} else {
			fmt.Printf("\"%s\" is not a file!\n\n", arg)
		}
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func doHash(h hash.Hash, b []byte, res chan string) {
	h.Write(b)
	res <- fmt.Sprintf("%x", h.Sum(nil))
}
