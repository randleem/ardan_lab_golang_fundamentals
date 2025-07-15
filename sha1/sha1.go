package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main(){
	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig("sha1.go"))
}

// SHA1SIG returns SHA1 signature of uncompressed file.
// Exercise: Decompress only if file name ends ith ".gz"
// gunzip -c http.log.gz | shasum
func SHA1Sig(fileName string)(string, error){
	file, err := os.Open(fileName)
	if err != nil{
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

    //unzip
	// BUG: Creates new "r" that is only in "if scope."
	//shadowing
	if filepath.Ext(fileName) == ".gz"{
	gz, err := gzip.NewReader(file)
	if err != nil{
		return "", fmt.Errorf("%q - gzip: %w", fileName, err)
	}
	defer gz.Close()
	r = gz
	}
		w := sha1.New()
	if _, err := io.Copy(w,r); err != nil{
		return "", fmt.Errorf("%q - copy: %w", fileName, err)

	}
	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

/*
Go - performance
Type Reader interface{
Read(p []byte)(n int, err error)}
*/