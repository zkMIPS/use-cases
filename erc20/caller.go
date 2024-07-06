package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"erc20/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProofPublicData struct {
	Proof struct {
		Ar struct {
			X string
			Y string
		}
		Krs struct {
			X string
			Y string
		}
		Bs struct {
			X struct {
				A0 string
				A1 string
			}
			Y struct {
				A0 string
				A1 string
			}
		}
		Commitments []struct {
			X string
			Y string
		}
	}
	PublicWitness []string
}

type Receipt struct {
	proof token.Proof
	input [65]*big.Int
	proofCommitment [2]*big.Int 
}

var ChainId *int64
var Network *string
var HexPrivateKey *string // ("df4bc5647fdb9600ceb4943d4adff3749956a8512e5707716357b13d5ee687d9") // 0x21f59Cfb0d41FA2c0eeF0Fe1593F46f704C1Db50

func main() {
	ChainId = flag.Int64("chainId", 11155111, "chainId")
	Network = flag.String("network", "https://eth-sepolia.g.alchemy.com/v2/RH793ZL_pQkZb7KttcWcTlOjPrN0BjOW", "network")
	HexPrivateKey = flag.String("privateKey", "df4bc5647fdb9600ceb4943d4adff3749956a8512e5707716357b13d5ee687d9", "privateKey")
	erc20TokenAddr := flag.String("erc20TokenAddr", "0xA234F9049720EDaDF3Ae697eE8bF762f2A03949A", "erc20TokenAddr")
	proofPath := flag.String("proofPath", "./hardhat/test/snark_proof_with_public_inputs.json", "proofPath")
	if len(os.Args) < 2 {
		log.Printf("expected subcommands")
		os.Exit(1)
	}
	flag.CommandLine.Parse(os.Args[2:])
	switch os.Args[1] {
	case "mint":
		mintTokenFromProof(*erc20TokenAddr, *proofPath)
	}
}

func authInstance() (*bind.TransactOpts) {
	unlockedKey, err := crypto.HexToECDSA(*HexPrivateKey)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(unlockedKey, big.NewInt(*ChainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.GasLimit = 1000000

	return auth;
}

func clientInstance() (*ethclient.Client) {
	client, err := ethclient.Dial(*Network)
	if err != nil {
		log.Fatalf("Failed to create eth client: %v", err)
	}

	return client;
}

func generateReceipt(proofPath string) (Receipt) {
	jsonFile, err := os.Open(proofPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	proofPublicData := ProofPublicData{}
	err = json.Unmarshal(byteValue, &proofPublicData)
	if err != nil {
		log.Fatal(err)
	}

	var input [65]*big.Int
	for i := 0; i < len(proofPublicData.PublicWitness); i++ {
		input[i], _ = new(big.Int).SetString(proofPublicData.PublicWitness[i], 0)
	}

	var proof = token.Proof{}
	proof.A.X, _ = new(big.Int).SetString(proofPublicData.Proof.Ar.X, 0)
	proof.A.Y, _ = new(big.Int).SetString(proofPublicData.Proof.Ar.Y, 0)

	proof.B.X[0], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.X.A0, 0)
	proof.B.X[1], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.X.A1, 0)
	proof.B.Y[0], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.Y.A0, 0)
	proof.B.Y[1], _ = new(big.Int).SetString(proofPublicData.Proof.Bs.Y.A1, 0)

	proof.C.X, _ = new(big.Int).SetString(proofPublicData.Proof.Krs.X, 0)
	proof.C.Y, _ = new(big.Int).SetString(proofPublicData.Proof.Krs.Y, 0)

	var proofCommitment [2]*big.Int
	proofCommitment[0], _ = new(big.Int).SetString(proofPublicData.Proof.Commitments[0].X, 0)
	proofCommitment[1], _ = new(big.Int).SetString(proofPublicData.Proof.Commitments[0].Y, 0)

	return Receipt{proof, input, proofCommitment}
}

func mintTokenFromProof(addr string, proofPath string) {
	flag.Parse()

	receipt := generateReceipt(proofPath)

	erc20, err := token.NewTestToken(common.HexToAddress(addr), clientInstance())
	if err != nil {
		fmt.Printf("Failed to instantiate a Token contract: %v", err)
		return
	}

	name, err := erc20.Name(nil)
	if err != nil {
		fmt.Printf("Failed to get name: %v", err)
		return
	}
	fmt.Printf("name: %v\n", name)

	totalSupply, err := erc20.TotalSupply(nil)
	if err != nil {
		fmt.Printf("Failed to get name: %v", err)
		return
	}
	fmt.Printf("Total Supply: %v\n", totalSupply)

	// TODO: Find out way to get the value of ETH from the proof output
	tx, err := erc20.MintWithProof(authInstance(), common.HexToAddress(os.Args[2]), big.NewInt(1000000000000000000), receipt.proof, receipt.input, receipt.proofCommitment)
	if err != nil {
		log.Fatalf("Failed to VerifyProof,err:[%+v]", err)
	}
	log.Printf("mint with proof txHash: %+v\n", tx.Hash())
}
