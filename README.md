![Github CI/CD](https://github.com/cronnoss/blockchain-api/actions/workflows/go.yml/badge.svg)
[![Go Report](https://goreportcard.com/badge/github.com/cronnoss/blockchain-api)](https://goreportcard.com/report/github.com/cronnoss/blockchain-api)
![Repository Top Language](https://img.shields.io/github/languages/top/cronnoss/blockchain-api.svg)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/cronnoss/blockchain-api/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/cronnoss/blockchain-api/?branch=main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cronnoss/blockchain-api.svg)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bf381a805bd34d95b9506f01a0113b66)](https://app.codacy.com/gh/cronnoss/blockchain-api?utm_source=github.com&utm_medium=referral&utm_content=cronnoss/blockchain-api&utm_campaign=Badge_Grade)
[![Coverage Status](https://coveralls.io/repos/github/cronnoss/blockchain-api/badge.svg)](https://coveralls.io/github/cronnoss/blockchain-api)
![Repository Size](https://img.shields.io/github/repo-size/cronnoss/blockchain-api?style=flat-square)
![GitHub Open Issues](https://img.shields.io/github/issues/cronnoss/blockchain-api.svg)
![GitHub last commit](https://img.shields.io/github/last-commit/cronnoss/blockchain-api.svg)
![GitHub contributors](https://img.shields.io/github/contributors/cronnoss/blockchain-api.svg)
![License](https://img.shields.io/badge/license-Apache%202-blue)


<img align="right" width="50%" src="./images/big-gopher.png">

# Back-end developer tech assignment

In this tech assignment your task is to create an API-server for blockchain indexes.

Base item is Group (Solidity smart contract code):
```solidity
struct Group { 
    string name; 
    uint256[] indexes;
}
```

Group contains a list of ids. Each id represents index: 
```solidity
struct Index {
    uint256 id;
    string  name;
    uint256 ethPriceInWei; 
    uint256 usdPriceInCents; 
    uint256 usdCapitalization; 
    uint256 percentageChange;
}
```

You should deploy a smart contract on **Sepolia** (ETH test network) and create an API-server that will allow you to interact with the smart contract.

Contract ABI’s serve as an interface to the smart contract: [https://github.com/HvrlK/abi-contract](https://github.com/HvrlK/abi-contract)

You contract address should be like: `0xA655555555555555555555555555555555555555`

You should implement next methods:
- GET groups/ (**getGroupIds** method in smart contract)
- GET groups/{groupId} (**getGroup(groupId)** method in smart contract)
- GET indexes/{indexId} (**getIndex(indexId)** method in smart contract)
- GET blocks/{BlockNumber | BlockHash | “latest“} (get block info from **Sepolia**)

Use [fiber](https://github.com/gofiber/fiber) or any other golang framework. Create [Swagger](https://github.com/go-swagger/go-swagger) documentation.

# 1. Deploy Smart Contract on Sepolia Testnet using Alchemy

See [README.md](index_manager/README.md) in `index_manager` directory.


# 2. Blockchain API

API-server for blockchain indexes

## Solution notes

- :trident: clean architecture (handler->service->repository)
- :book: standard Go project layout (well, more or less :blush:)
- :cd: github CI/CD + docker compose + Makefile included
- :white_check_mark: handler tests with mocks included
- :heart: Swagger auto-generated on `make generate` included

## HOWTO

- :running_man: run app in docker-compose with `make dc`
- :test_tube: run tests with `make test`
- :coin: generate contract from `abi` file with `make abigen`
- :sunflower: run linter with `make lint`

## A picture is worth a thousand words

<img src="./images/make-run.png">
