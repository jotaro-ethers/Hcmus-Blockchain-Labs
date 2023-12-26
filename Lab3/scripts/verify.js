const { ethers,run } = require('hardhat');
const verifyContract = async (contractAddress) => {
  console.log("Verifying contract...");
  contractAddress = "0x88DfD61FBF4Ad47Ca9aA7769C8636833771deF9E"
  try {
    await run("verify:verify", {
      address: contractAddress,
    });

    console.log("Contract verified!");
  } catch (err) {
    console.log(err);
  }
};
verifyContract();