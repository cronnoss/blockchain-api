// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.3;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/IndexManager.sol";

contract TestIndexManager {

    function testInitialOwnerShouldBeSet() public {
        IndexManager indexManager = new IndexManager();
        Assert.equal(indexManager.owner(), address(this), "Owner should be the deployer");
    }

    function testGetGroupIdsShouldReturnEmptyInitially() public {
        IndexManager indexManager = new IndexManager();
        uint256[] memory groupIds = indexManager.getGroupIds();
        Assert.equal(groupIds.length, 0, "Group IDs should be empty initially");
    }

    function testAddGroupShouldStoreGroup() public {
        IndexManager indexManager = new IndexManager();
        indexManager.addGroup(1, "Group 1");
        (string memory name, uint256[] memory indexes) = indexManager.getGroup(1);
        Assert.equal(name, "Group 1", "Group name should be 'Group 1'");
        Assert.equal(indexes.length, 0, "Group indexes should be empty initially");
    }

    function testAddIndexShouldStoreIndex() public {
        IndexManager indexManager = new IndexManager();
        indexManager.addGroup(1, "Group 1");
        indexManager.addIndex(1, 1, "Index 1", 1000, 200, 300, 10);
        (string memory name, uint256 ethPriceInWei, uint256 usdPriceInCents, uint256 usdCapitalization, int256 percentageChange) = indexManager.getIndex(1);
        Assert.equal(name, "Index 1", "Index name should be 'Index 1'");
        Assert.equal(ethPriceInWei, 1000, "ETH price should be 1000 wei");
        Assert.equal(usdPriceInCents, 200, "USD price should be 200 cents");
        Assert.equal(usdCapitalization, 300, "USD capitalization should be 300");
        Assert.equal(percentageChange, 10, "Percentage change should be 10");
    }

    function testAddIndexShouldAddToGroup() public {
        IndexManager indexManager = new IndexManager();
        indexManager.addGroup(1, "Group 1");
        indexManager.addIndex(1, 1, "Index 1", 1000, 200, 300, 10);
        (, uint256[] memory indexes) = indexManager.getGroup(1);
        Assert.equal(indexes.length, 1, "Group should have one index");
        Assert.equal(indexes[0], 1, "Index ID should be 1");
    }
}