const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

module.exports = buildModule("ERC20Token", (m) => {
  const erc20Token = m.contract("ERC20Token");
  return { erc20Token };
});
