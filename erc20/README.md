## ERC20

### Mint Tokens

Place the `snark_proof_with_public_inputs.json` file in the `verifier` folder

```
./caller mint 0x21f59Cfb0d41FA2c0eeF0Fe1593F46f704C1Db50
```

### Deploy a new ERC20 token

```
npx hardhat ignition deploy ./ignition/modules/erc20.js --network sepolia
```

Replace the `erc20TokenAddr` in `caller.go` and run `go build caller.go`

### Create a new Verifier Contract based on a new Verifier Key

Replace the `verifying.key` file in the `verifier` folder and generate the new `verifier.sol` contract using

```
./caller generate
```
