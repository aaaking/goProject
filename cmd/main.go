package main
import "../core"
func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to ZZH")
	bc.SendData("Send 1 EOS to Jack")
	bc.Print()
}