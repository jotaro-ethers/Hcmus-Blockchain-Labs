require("@nomicfoundation/hardhat-toolbox");
const INFURA_API_KEY = "4395960d44db47568543b11c1b38aadd";
const SEPOLIA_PRIVATE_KEY = "9fe1a1c3deaeff95847281a1f89b7978cf8a8632f2a6d7d6f215cc92744e9fb4";

module.exports = {
  solidity: "0.8.19",
  networks: {
    sepolia: {
      url: `https://sepolia.infura.io/v3/${INFURA_API_KEY}`,
      accounts: [SEPOLIA_PRIVATE_KEY]
    }
  },
  etherscan: {
    apiKey: "ZZDZ4NMQ4KWRZY77B7C1HN2SIBYD2N4DBJ",
  },
};