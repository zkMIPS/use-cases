## ERC20

### Mint Tokens

```
./caller mint 0x21f59Cfb0d41FA2c0eeF0Fe1593F46f704C1Db50
```

### Deploy a new ERC20 token

```
npx hardhat ignition deploy ./ignition/modules/erc20.js --network sepolia
```

Replace the `erc20TokenAddr` in `caller.go` and run `go build caller.go`
