//go:build !wasm

package main

import (
	"github.com/mavryk-network/mavpay/cmd"
)

func main() {
	cmd.Execute()
}
