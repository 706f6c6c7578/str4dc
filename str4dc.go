package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"encoding/hex"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: echo -n <input_string> | str4dc")
		os.Exit(1)
	}

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text()+"\n") // Add "\n" to include newline
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Join the lines into a single string
	inputStr := strings.Join(inputLines, "")

	// Convert input string to hexadecimal representation
	encodedHex := hexEncode(inputStr)

	// Create a big.Int from the hexadecimal string
	x := new(big.Int)
	_, success := x.SetString(encodedHex, 16)
	if !success {
		fmt.Println("Error converting hexadecimal to big.Int")
		os.Exit(1)
	}

	y := ""
	f := true
	ten := big.NewInt(10)
	five := big.NewInt(5)

	zero := big.NewInt(0)
	one := big.NewInt(1)
	tenMinusOne := big.NewInt(10)

	// Lookup table for hexadecimal characters
	hexChars := "0123456789ABCDEF"

	for x.Cmp(zero) != 0 {
		var d big.Int
		d.Mod(x, ten)
		x.Div(x, ten)
		if x.Cmp(zero) != 0 && d.Cmp(five) <= 0 {
			if f {
				x.Sub(x, one)
				d.Add(&d, tenMinusOne)
			}
		} else {
			f = !f
		}
		y = string(hexChars[d.Int64()]) + y // Corrected line using hexChars
	}

	fmt.Println(y)
}

// hexEncode converts a string to its hexadecimal representation
func hexEncode(inputStr string) string {
	encodedBytes := hex.EncodeToString([]byte(inputStr))
	return strings.ToUpper(encodedBytes)
}
