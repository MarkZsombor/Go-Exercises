package main

// Create an algorithm that will calculate the largest prime number less than argument.

// Pseudocode for determining if a number is prime
// Create an array from 0..max
// Starting at 2, delete every multiple of 2 from the array.
// Then, go back to the beginning, and delete every multiple of 3.
// Repeat this starting from the next available number at the beginning of the array.
// Do this until the square of number you are checking is greater than your max number.
// Finally, compact the original array.
