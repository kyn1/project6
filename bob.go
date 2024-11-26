package main

import (
	"fmt"
)

// Bob the evaluator

func evaluateGarbledCircuit(inputs [][]byte, gates []Gate) []byte {
	n := len(inputs) // number of inputs
	m := len(gates)  // number of gates

	// array of signals have a size of n+m
	signals := make([][]byte, m+n)

	// setup inputs signals
	for i := 0; i < n; i++ {
		signals[i] = inputs[i]
		fmt.Printf("input %d=%x\n", i, signals[i][:2])
	}

	// add code below to evaluate the gates
	/*
		for _, gate := range gates {
			a := signals[gate.in0]
			b := signals[gate.in1]
			...
			signals[gate.out] = ...
		}
	*/

	// the last signal is the output
	return signals[n+m-1]
}
