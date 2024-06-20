require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.0",
  defaultNetwork: "localhost",
  networks: {
    hardhat: {},
    // example
    sepolia: {
      url: "https://eth-sepolia.g.alchemy.com/v2/RH793ZL_pQkZb7KttcWcTlOjPrN0BjOW",
      accounts: [
        "df4bc5647fdb9600ceb4943d4adff3749956a8512e5707716357b13d5ee687d9",
      ],
      chainId: 11155111,
      gasPrice: 1000000000,
    },
  },
};
