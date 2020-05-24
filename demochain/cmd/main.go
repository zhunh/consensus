package main

import "going/demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to jack")
	bc.SendData("Send 2 ETH to Tom")
	bc.Print()
}
