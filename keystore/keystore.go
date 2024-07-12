package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./tmp/UTC--2024-07-12T06-23-24.995753339Z--a8312953c527cd2a89b54c17066ae70990bcd1fb"

	jsonByte, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "secret"

	account, err := ks.Import(jsonByte, password, password)

	if err != nil {
		fmt.Println(err)
		accounts := ks.Accounts()
		for _, account := range accounts {
			fmt.Println(account.Address.Hex())
		}
	} else {
		fmt.Println(account.Address.Hex())
	}
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
