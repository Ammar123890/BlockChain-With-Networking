# Simple Blockchain in Go

This repository contains a basic implementation of a blockchain in Go. It's designed to help newcomers understand the fundamental concepts of blockchains and how they can be implemented in a networked environment.

## Features

- **Immutable Ledger**: Once a block has been added to the chain, it cannot be changed, ensuring data integrity.
- **Proof-of-Work Algorithm**: Implemented to adjust difficulty and ensure security against spam and Sybil attacks.
- **Chain Validation**: Ensures the integrity of the entire blockchain is maintained at every step.
- **Network Communication**: Nodes communicate to broadcast and verify new blocks, simulating a decentralized network.
- **Dynamic Difficulty Adjustment**: Difficulty of the Proof-of-Work algorithm adjusts depending on the rate of block creation.
- **Merkle Tree Implementation**: Enhances data verification and integrity within blocks.

## Setup

1. **Initialize Go Module**:
    ```bash
    go mod init <module-name>
    ```

2. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

## Usage

To start a node on the default port (8001), run:
```bash
go run main.go
```

To start a node on a different port, run:
```bash
go run main.go -port=<desired-port>
```

## Understanding the Code

- **Block Structure**: Each block contains data, a timestamp, the previous block's hash, its own hash, a nonce, and a difficulty target.
- **Proof-of-Work**: This mechanism ensures that creating a block requires computational effort. This process, also known as "mining," requires finding a hash that meets the dynamic difficulty criteria.
- **Chain Validation**: Every time an action is performed on the blockchain, the entire chain is validated to ensure its integrity.
- **Network Communication**: Nodes use TCP/IP to communicate new blocks and transactions. Each node independently verifies new blocks before adding them to their local blockchain.
- **Dynamic Difficulty**: The mining difficulty adjusts after a set number of blocks to ensure a steady rate of block creation.

## Future Improvements

- Extend the network communication to handle more complex scenarios and potential conflicts.
- Implement a peer-to-peer network to further distribute ledger capabilities.
- Integrate a more persistent form of storage (e.g., databases).
- Add transaction pooling and more sophisticated transaction selection for block mining.
- Include smart contract capabilities to extend the blockchain's functionality.
