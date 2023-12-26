// SupplyChain.test.js
const { expect } = require('chai');
const { ethers } = require('hardhat');
const {
  loadFixture,
} = require("@nomicfoundation/hardhat-toolbox/network-helpers");
describe("SupplyChain Test Unit", function () {
  async function deployFixture() {
    let [owner,authorized,newOwner,random] = await ethers.getSigners();
    const supplyChain = await ethers.deployContract("SupplyChain");
    console.log("SupplyChain deployed to:", await supplyChain.getAddress());
    console.log("Owner address:", await owner.getAddress());
    console.log("Authorized address:", await authorized.getAddress());
    console.log("NewOwner address:", await newOwner.getAddress());
    return { supplyChain, owner, authorized, newOwner, random };
  };
  it('should create an item', async function () {
    const { supplyChain, owner, authorized, newOwner } = await loadFixture(deployFixture);
    expect(await supplyChain.itemCount()).to.equal(0);
    await supplyChain.connect(owner).createItem("item1", 100,10,"Location1");
    const item = await supplyChain.items(1);
    expect(item.name).to.equal("item1");
  });
  it("should authorized can do something with an item", async function () {
    const { supplyChain, owner, authorized, newOwner ,random} = await loadFixture(deployFixture);
    await supplyChain.connect(owner).createItem("item1", 100,10,"Location1");
    await supplyChain.connect(owner).authorizeParty(1, authorized.address);
    await expect(supplyChain.connect(random).shipItem(1, "Location2")).to.be.revertedWith("You are not authorized"); 
  });
  it("Only owner or third party can transfer ownership if meet status conditions ", async function () {
    const { supplyChain, owner, authorized, newOwner ,random} = await loadFixture(deployFixture);
    await supplyChain.connect(owner).createItem("item1", 100,10,"Location1");
    await supplyChain.connect(owner).authorizeParty(1, authorized.address);
    await supplyChain.connect(authorized).shipItem(1, "Location2");
    await expect(supplyChain.connect(random).transferOwnership(1, newOwner.address)).to.be.revertedWith("You don't own this item"); 
    const itemBeforeDelivery = await supplyChain.items(1);
    expect(itemBeforeDelivery.status).to.equal(1);
    await supplyChain.connect(authorized).deliverItem(1,newOwner.address);
    const item = await supplyChain.items(1);
    await expect(item.owner).to.equal(newOwner.address); 
  });



});
