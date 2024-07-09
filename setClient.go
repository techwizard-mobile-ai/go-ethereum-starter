package main

import (
	"fmt"

	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// client, err := ethclient.Dial("https://cloudflare-eth.com")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("We have a connection")
	// _ = client //we'll use this in the upcoming sections
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
}
