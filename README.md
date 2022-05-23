# Interaction with Ethereum Smart Contracts using Go Application (geth)

Developed a Todolist smart contract in solidity and the development environment used is hardhat, deployed to rinkeby network by using infura API and infrastructure. Using `solc` solidity compiler generated `abi` and `bin` file and using  `abigen` generated the `.go` from the Todo solidity contract.
Also, written a go application which uses go ethereum packages and go bindings and interacts with the deployed smart contract via ethclient.


### Task Accomplishments
1. Ethereum based Smart Contract written in solidity [Code](https://github.com/AnilRaj27/go-dapp/blob/main/contracts/Todo.sol)
2. Todo contract deployed on rinkeby testnet using infura [URL](https://rinkeby.etherscan.io/address/0x067D02EE461F1334F7aea27b52Ad687708D53D19)
3. Generated the go bindings for the Todo smart contract
    ```
    solc --bin --abi contracts/Todo.sol -o build
    abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=main --out=gen/Todo.go
    ```
4. Written go code for interaction with smart contract [Code](https://github.com/AnilRaj27/go-dapp/blob/main/main/main.go)
