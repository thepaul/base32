// base32 utility by paul cannon <p@thepaul.org>

package main

import (
	"bufio"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/pflag"
)

var (
	decoding = pflag.BoolP("-decode", "d", false,
		"Decode instead of encode")
	useHex = pflag.BoolP("-hex", "x", false,
		"Expect hexadecimal input (or, if decoding, produce hexadecimal output")

	base32Encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567").WithPadding(base32.NoPadding)
)

func decodeAllFromBase32(r io.Reader, w io.Writer, useHex bool) (err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		result, err := base32Encoding.DecodeString(word)
		if useHex {
			_, err = fmt.Fprintf(w, "%x\n", result)
		} else {
			_, err = w.Write(result)
		}
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func encodeAllToBase32(r io.Reader, w io.Writer) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	result := base32Encoding.EncodeToString(data)
	_, err = w.Write([]byte(result))
	return err
}

func encodeAllHexToBase32(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		data, err := hex.DecodeString(word)
		if err != nil {
			return err
		}
		result := base32Encoding.EncodeToString(data)
		_, err = w.Write([]byte(result))
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func main() {
	pflag.Parse()

	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	var err error
	if *decoding {
		err = decodeAllFromBase32(input, output, *useHex)
	} else {
		if *useHex {
			err = encodeAllHexToBase32(input, output)
		} else {
			err = encodeAllToBase32(input, output)
		}
	}
	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
