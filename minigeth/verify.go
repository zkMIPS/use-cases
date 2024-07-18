package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"minigeth/common"

	ecommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"

	"minigeth/consensus/misc"
	"minigeth/core"
	"minigeth/core/state"
	"minigeth/core/types"
	"minigeth/core/vm"
	"minigeth/crypto"
	"minigeth/oracle"
	"minigeth/params"
	"minigeth/rlp"
	"minigeth/trie"
)

var Network *string

func main() {
	Network = flag.String("network", "https://eth-mainnet.g.alchemy.com/v2/u9yBeOHR2P2ag3iq-vyAdAwzM46nywr0", "network")
	blockNumber, bigIntConversion := new(big.Int).SetString(os.Args[1], 10)
	if !bigIntConversion {
        log.Fatal("Failed to convert input to bigint")
    }
	transactionHash := ecommon.HexToHash(os.Args[2])
	burnAddress := ecommon.HexToAddress("0xdeaDDeADDEaDdeaDdEAddEADDEAdDeadDEADDEaD")
	
	flag.Parse()

	validate(int(blockNumber.Int64()))

	var block, err = clientInstance().BlockByNumber(context.Background(), blockNumber)
	if err != nil {
        log.Fatal(err)
    }

	if (*block.Transaction(transactionHash).To() != burnAddress) {
		log.Fatal("Transaction is not directed to the burn address")
	}

	// Get the transaction within the block and display the output value in ETH transferred
	fmt.Printf("Transaction Value: %v", block.Transaction(transactionHash).Value())
}


func clientInstance() (*ethclient.Client) {
	client, err := ethclient.Dial(*Network)
	if err != nil {
		log.Fatalf("Failed to create eth client: %v", err)
	}

	return client;
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func validate(blockNumber int) {
	// non mips
	if len(os.Args) > 1 {
		newNodeUrl, setNewNodeUrl := os.LookupEnv("NODE")
		if setNewNodeUrl {
			fmt.Println("override node url", newNodeUrl)
			oracle.SetNodeUrl(newNodeUrl)
		}
		basedir := os.Getenv("BASEDIR")
		if len(basedir) == 0 {
			basedir = "/tmp/cannon"
		}

		pkw := oracle.PreimageKeyValueWriter{}
		pkwtrie := trie.NewStackTrie(pkw)

		// TODO: get the chainid
		oracle.SetRoot(fmt.Sprintf("%s/0_%d", basedir, blockNumber))
		oracle.PrefetchBlock(big.NewInt(int64(blockNumber)), true, nil)
		oracle.PrefetchBlock(big.NewInt(int64(blockNumber)+1), false, pkwtrie)
		hash, err := pkwtrie.Commit()
		check(err)
		fmt.Println("committed transactions", hash, err)
	}

	// init secp256k1BytePoints
	crypto.S256()

	// get inputs
	inputBytes := oracle.Preimage(oracle.InputHash())
	var inputs [6]common.Hash
	for i := 0; i < len(inputs); i++ {
		inputs[i] = common.BytesToHash(inputBytes[i*0x20 : i*0x20+0x20])
	}

	// read start block header
	var parent types.Header
	check(rlp.DecodeBytes(oracle.Preimage(inputs[0]), &parent))

	// read header
	var newheader types.Header
	// from parent
	newheader.ParentHash = parent.Hash()
	newheader.Number = big.NewInt(0).Add(parent.Number, big.NewInt(1))
	newheader.BaseFee = misc.CalcBaseFee(params.MainnetChainConfig, &parent)

	// from input oracle
	newheader.TxHash = inputs[1]
	newheader.Coinbase = common.BigToAddress(inputs[2].Big())
	newheader.UncleHash = inputs[3]
	newheader.GasLimit = inputs[4].Big().Uint64()
	newheader.Time = inputs[5].Big().Uint64()

	bc := core.NewBlockChain(&parent)
	database := state.NewDatabase(parent)
	statedb, _ := state.New(parent.Root, database, nil)
	vmconfig := vm.Config{}
	processor := core.NewStateProcessor(params.MainnetChainConfig, bc, bc.Engine())
	fmt.Println("processing state:", parent.Number, "->", newheader.Number)

	newheader.Difficulty = bc.Engine().CalcDifficulty(bc, newheader.Time, &parent)

	// read txs
	//traverseStackTrie(newheader.TxHash)

	//fmt.Println(fn)
	//fmt.Println(txTrieRoot)
	var txs []*types.Transaction

	triedb := trie.NewDatabase(parent)
	tt, _ := trie.New(newheader.TxHash, &triedb)
	tni := tt.NodeIterator([]byte{})
	for tni.Next(true) {
		//fmt.Println(tni.Hash(), tni.Leaf(), tni.Path(), tni.Error())
		if tni.Leaf() {
			tx := types.Transaction{}
			var rlpKey uint64
			check(rlp.DecodeBytes(tni.LeafKey(), &rlpKey))
			check(tx.UnmarshalBinary(tni.LeafBlob()))
			// TODO: resize an array in go?
			for uint64(len(txs)) <= rlpKey {
				txs = append(txs, nil)
			}
			txs[rlpKey] = &tx
		}
	}
	fmt.Println("read", len(txs), "transactions")
	// TODO: OMG the transaction ordering isn't fixed

	var uncles []*types.Header
	check(rlp.DecodeBytes(oracle.Preimage(newheader.UncleHash), &uncles))

	var receipts []*types.Receipt
	block := types.NewBlock(&newheader, txs, uncles, receipts, trie.NewStackTrie(nil))
	fmt.Println("made block, parent:", newheader.ParentHash)

	// if this is correct, the trie is working
	// TODO: it's the previous block now
	if newheader.TxHash != block.Header().TxHash {
		panic("wrong transactions for block")
	}
	if newheader.UncleHash != block.Header().UncleHash {
		panic("wrong uncles for block " + newheader.UncleHash.String() + " " + block.Header().UncleHash.String())
	}

	// validateState is more complete, gas used + bloom also
	receipts, _, _, err := processor.Process(block, statedb, vmconfig)
	receiptSha := types.DeriveSha(types.Receipts(receipts), trie.NewStackTrie(nil))
	if err != nil {
		log.Fatal(err)
	}
	newroot := statedb.IntermediateRoot(bc.Config().IsEIP158(newheader.Number))

	fmt.Println("receipt count", len(receipts), "hash", receiptSha)
	fmt.Println("process done with hash", parent.Root, "->", newroot)
	oracle.Output(newroot, receiptSha)
}