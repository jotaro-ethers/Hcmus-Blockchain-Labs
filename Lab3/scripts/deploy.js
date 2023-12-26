// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const { expect } = require('chai');
const { ethers,run } = require('hardhat');
const verifyContract = async (contractAddress,args) => {
  console.log("Verifying contract...");
  try {
    await run("verify:verify", {
      address: contractAddress,
      constructorArguments: args,
      
    });

    console.log("Contract verified!");
  } catch (err) {
    console.log(err);
  }
};
async function main() {
  const supplychain = await ethers.getContractFactory("SupplyChain");
  const supplychainContract = await supplychain.deploy();
  console.log("wait for deployment")
  await supplychainContract.waitForDeployment(30);
  const address = await supplychainContract.getAddress();
  console.log("SupplyChain deployed to:", address);
  // await verifyContract(address,[]);
  
}
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
