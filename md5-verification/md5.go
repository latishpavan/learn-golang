package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func parseSignatureFile(path string) (map[string]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	sigs := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for lnum := 1; scanner.Scan(); lnum++ {
		fields := strings.Fields(scanner.Text())

		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d\n", "bad line", path, lnum)
		}

		sigs[fields[0]] = fields[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

type result struct {
	path  string
	match bool
	err   error
}

func fileMd5(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func md5Worker(path string, sig string, out chan *result) {
	res := &result{path: path}
	s, err := fileMd5(path)

	res.err = err
	res.match = (sig == s)

	out <- res
}

func main() {
	path := filepath.Join("D:/", "nasa-logs", "md5sum.txt")
	sigs, err := parseSignatureFile(path)

	if err != nil {
		log.Fatalf("error can't read signature file: %s", err)
	}

	out := make(chan *result)

	for sig, fileName := range sigs {
		go md5Worker(filepath.Join(filepath.Dir(path), fileName), sig, out)
	}

	for range sigs {
		res := <-out
		if res.err != nil {
			fmt.Printf("Error while verifying the signature for file %s - %v\n", res.path, res.err)
		} else {
			if res.match {
				fmt.Printf("Hash signature matched for %s\n", res.path)
			} else {
				fmt.Printf("Hash signature did not match for %s\n", res.path)
			}
		}
	}
}
