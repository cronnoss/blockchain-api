// SPDX-License-Identifier: UNLICENSED
// Specifies the version of Solidity, using semantic versioning.
// Learn more: https://solidity.readthedocs.io/en/v0.5.10/layout-of-source-files.html#pragma
pragma solidity >=0.7.3;

contract IndexManager {

    address public owner;

    struct Group {
        string name;
        uint256[] indexes;
    }

    struct Index {
        uint256 id;
        string name;
        uint256 ethPriceInWei;
        uint256 usdPriceInCents;
        uint256 usdCapitalization;
        int256 percentageChange;
    }

    uint256[] private groupIds;
    mapping(uint256 => Group) private groups;
    mapping(uint256 => Index) private indexes;

    constructor() {
        owner = msg.sender;
    }

    // Function to donate
    function donate() public payable {
        require(msg.value > 0, "Donation needs to be a positive value");
    }

    // Withdraw ether
    function withdraw() public payable {
        require(msg.sender == owner, "Only the owner can withdraw!");

        uint balance = address(this).balance;
        require(balance > 0, "No ether left to withdraw");

        (bool success,) = (msg.sender).call{value: balance}("");
        require(success, "Transfer failed.");
    }

    // Get all group IDs
    function getGroupIds() public view returns (uint256[] memory) {
        return groupIds;
    }

    // Get a specific group by ID
    function getGroup(uint256 _groupId) public view returns (string memory name, uint256[] memory groupIndexes) {
        Group memory group = groups[_groupId];
        return (group.name, group.indexes);
    }

    // Get a specific index by ID
    function getIndex(uint256 _indexId) public view returns (
        string memory name,
        uint256 ethPriceInWei,
        uint256 usdPriceInCents,
        uint256 usdCapitalization,
        int256 percentageChange
    ) {
        Index memory index = indexes[_indexId];
        return (
            index.name,
            index.ethPriceInWei,
            index.usdPriceInCents,
            index.usdCapitalization,
            index.percentageChange
        );
    }

    // (Optional) Function to add groups and indexes (for completeness)
    function addGroup(uint256 _groupId, string memory _name) public {
        require(msg.sender == owner, "Only the owner can add groups");
        Group storage group = groups[_groupId];
        group.name = _name;
        groupIds.push(_groupId);
    }

    function addIndex(uint256 _groupId, uint256 _indexId, string memory _name, uint256 _ethPriceInWei, uint256 _usdPriceInCents, uint256 _usdCapitalization, int256 _percentageChange) public {
        require(msg.sender == owner, "Only the owner can add indexes");
        Index storage index = indexes[_indexId];
        index.name = _name;
        index.ethPriceInWei = _ethPriceInWei;
        index.usdPriceInCents = _usdPriceInCents;
        index.usdCapitalization = _usdCapitalization;
        index.percentageChange = _percentageChange;
        groups[_groupId].indexes.push(_indexId);
    }
}