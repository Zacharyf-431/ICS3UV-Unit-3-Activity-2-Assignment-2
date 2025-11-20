/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-20
 * @fileoverview: This file reads in a number of cents 
 */

// INPUT
// Ask the user for the number of cents
const input = prompt("Input the cents please: ") ||"0";
const totalCents = parseInt(input);

// PROCESS
// Calculate dollars and cents
let centsLeft = totalCents % 100;
let totalDollar = (totalCents - centsLeft) / 100;

// OUTPUT 
console.log(`That is ${totalDollar} dollars and ${centsLeft} cents`);
console.log("\nDone.")