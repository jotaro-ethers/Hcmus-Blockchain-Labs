// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract SupplyChain {

    enum ItemStatus { Created, Shipped, Delivered}

    struct Item {
        string name;
        uint256 itemId;
        address owner;
        ItemStatus status;
        string location;
        uint256 manufacturingDate;
        uint256 deliveryDate;
        uint256 price;
        uint256 quantity;
        mapping(address => bool) authorizedParties;
    }

    mapping(uint256 => Item) public items;
    uint256 public itemCount;

    event ItemCreated(uint256 itemId, address owner);
    event ItemShipped(uint256 itemId, string newLocation);
    event ItemDelivered(uint256 itemId);
    event OwnershipTransferred(uint256 itemId, address previousOwner, address newOwner);

    modifier onlyOwnerOrAuthorized(uint256 _itemId) {
        require(items[_itemId].authorizedParties[msg.sender] || items[_itemId].owner == msg.sender, "You are not authorized");
        _;
    }

    modifier onlyOwner(uint256 _itemId) {
        require(items[_itemId].owner == msg.sender, "You don't own this item");
        _;
    }

    modifier onlyAuthorized(uint256 _itemId) {
        require(items[_itemId].authorizedParties[msg.sender], "You are not authorized");
        _;
    }

    function createItem(string memory _name,uint256 _price, uint256 _quantity, string memory _location) external {
        require(_price > 0, "Price must be greater than zero");
        require(_quantity > 0, "Quantity must be greater than zero");
        itemCount++;
        Item storage newItem = items[itemCount];
        newItem.name = _name;
        newItem.itemId = itemCount;
        newItem.owner = msg.sender;
        newItem.status = ItemStatus.Created;
        newItem.location = _location;
        newItem.manufacturingDate = block.timestamp;
        newItem.price = _price;
        newItem.quantity = _quantity;
        emit ItemCreated(itemCount, msg.sender);
    }

    function shipItem(uint256 _itemId, string memory _newLocation) external onlyAuthorized(_itemId) {
        require(items[_itemId].status == ItemStatus.Created, "Item is not in a shippable state");

        items[_itemId].status = ItemStatus.Shipped;
        items[_itemId].location = _newLocation;

        emit ItemShipped(_itemId, _newLocation);
    }

    function deliverItem(uint256 _itemId,address _newOwner) external onlyAuthorized(_itemId) {
        require(items[_itemId].status == ItemStatus.Shipped, "Item is not in a deliverable state");
        items[_itemId].deliveryDate = block.timestamp;
        items[_itemId].status = ItemStatus.Delivered;
        transferOwnershipByAuthorize(_itemId, _newOwner);
        emit ItemDelivered(_itemId);
    }
    function transferOwnershipByAuthorize(uint256 _itemId,address _newOwner) internal {
        require(_newOwner != address(0), "Invalid new owner address");
        require(_newOwner != msg.sender, "You are Authorized already");
        require(items[_itemId].status == ItemStatus.Delivered, "Item is not in a deliverable state");
        address oldOwner = items[_itemId].owner;
        items[_itemId].owner = _newOwner;
        emit OwnershipTransferred(_itemId,oldOwner, _newOwner);
    }
    function transferOwnership(uint256 _itemId, address _newOwner) external onlyOwner(_itemId) {
        require(_newOwner != address(0), "Invalid new owner address");
        require(_newOwner != msg.sender, "You are Owner already");
        require(items[_itemId].status == ItemStatus.Created, "Item is not in a Created state");
        emit OwnershipTransferred(_itemId,msg.sender, _newOwner);
    }

    function authorizeParty(uint256 _itemId, address _authorizedParty) external onlyOwner(_itemId) {
        items[_itemId].authorizedParties[_authorizedParty] = true;
    }

    function deauthorizeParty(uint256 _itemId, address _authorizedParty) external onlyOwner(_itemId) {
        items[_itemId].authorizedParties[_authorizedParty] = false;
    }
}