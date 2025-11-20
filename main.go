/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-20
	* This file reads in a number of cents
	*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// INPUT
	// Ask the user for the number of cents
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input the cents please: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert to integer
	totalCents, _ := strconv.Atoi(input)

	// PROCESS
	// Calculate dollars and cents
	centsLeft := totalCents % 100
	totalDollar := (totalCents - centsLeft) / 100

	// OUTPUT
	fmt.Printf("That is %d dollars and %d cents\n", totalDollar, centsLeft)

	fmt.Println("\nDone.")
}