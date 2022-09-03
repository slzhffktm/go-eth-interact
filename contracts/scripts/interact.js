const API_KEY = process.env.API_KEY;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const CONTRACT_ADDRESS = process.env.CONTRACT_ADDRESS;

const contract = require("../artifacts/contracts/HelloWorld.sol/HelloWorld.json");
const {ethers} = require("ethers");

console.log(JSON.stringify(contract.abi));

// Provider
const alchemyProvider = new ethers.providers.AlchemyProvider(network="goerli", API_KEY);

// Signer
const signer = new ethers.Wallet(PRIVATE_KEY, alchemyProvider);

// Contract
const helloWorldContract = new ethers.Contract(CONTRACT_ADDRESS, contract.abi, signer);

async function main() {
  const resp = await helloWorldContract.hello();
  console.log("The respond from hello() is: " + resp);

  console.log("Updating owner name...");
  const tx = await helloWorldContract.updateOwnerName("Owner");
  await tx.wait(); // Wait for the transaction to complete.
  
  const resp2 = await helloWorldContract.greetOwner();
  console.log("The respond from greetOwner() is: " + resp2);
}

main();
