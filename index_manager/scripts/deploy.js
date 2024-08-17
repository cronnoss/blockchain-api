async function main() {
    const HelloWorld = await ethers.getContractFactory("IndexManager");

    // Start deployment, returning a promise that resolves to a contract object
    const hello_world = await HelloWorld.deploy();
    console.log("Contract deployed to address:", hello_world.address);

    // Example: Adding some groups and indexes after deployment (optional)
    await hello_world.addGroup(1, "Group 1");
    await hello_world.addIndex(1, 101, "Index 101", ethers.utils.parseUnits("1.5", "ether"), 15000, 1000000, 5);
    await hello_world.addIndex(1, 102, "Index 102", ethers.utils.parseUnits("2.0", "ether"), 20000, 2000000, -3);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });