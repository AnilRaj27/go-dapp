const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy
  const todoFactory = await hre.ethers.getContractFactory("Todo");
  const todoContract = await todoFactory.deploy();
  const x = await todoContract.addTask("def");
  const y = await todoContract.getTask(1);
  const z = await todoContract.getTaskList();
  const a = await todoContract.updateTaskData(1, "updateditem");
  const c = await todoContract.getTaskList();
  const f = await todoContract.toggleCompleted(1);
  const taskList = await todoContract.getTaskList();

  console.log("taskList : ", taskList);

  await todoContract.deployed();
  console.log("Todo Contract deployed to:", todoContract.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
