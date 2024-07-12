package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}
	//7a18728395805110bff8f78fefdeca2dd1c2d18b41ae9a14aa0386e6c68bd4c0
	//8bae922F13d30731B29b5A4c98bB0054b2Df1809
	//8ddbb9dd0268fe9e24ebb5edcf21c5b577070c8eac444f03318e8f6308ac13f26868c59ea4ead74868151e47cb8d4fd1fd64ce400f20f6b2be460222cdf56500
	privateKeyByte := crypto.FromECDSA(privateKey)

	fmt.Println("This is for the privatekey")
	fmt.Println(hexutil.Encode(privateKeyByte)[2:])

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	hash := sha3.NewLegacyKeccak256()

	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
