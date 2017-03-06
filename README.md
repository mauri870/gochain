A simple Blockchain implementation in Go

> A blockchain – originally block chain – is a distributed database that maintains a continuously growing list of ordered records called blocks. Each block contains a timestamp and a link to a previous block. By design, blockchains are inherently resistant to modification of the data — once recorded, the data in a block cannot be altered retroactively. Blockchains are "an open, distributed ledger that can record transactions between two parties efficiently and in a verifiable and permanent way. The ledger itself can also be programmed to trigger transactions automatically." [Wikipedia](https://en.wikipedia.org/wiki/Blockchain_(database))

The data is not persisted across nodes, each service(peer or server) will keep an in-memory blockchain in sync. The project goals are:
- [x] A HTTP API for showing and inserting blocks and peers
- [ ] A websocket peer P2P implementation (each peer will keep a exact copy of the blockchain and receive updates about it, even if *n* peers are connected)
- [ ] Add support for users(with auth) and transactions(data exchange, not cryptocoins/bitcoins or something like this)

## Installation
```
go get -u github.mauri870/gochain
cd $GOPATH/src/github.com/mauri870/gochain
make install
make
```
The last command will generate a binary into the bin folder. To start the http server run:
```
./bin/server --address ":3000"
```

## Server endpoints

`GET api/blocks` Return an array of blocks in a blockchain. Each block is composed of:
- Index - a number in the blockchain. This number is incremented by 1 in every new block
- Timestamp - a unix time of the block creation
- Data - the value stored in a block, this field can be any type(interface{})
- PreviousBlock - a pointer to the previous block
- Hash - a sha256 unique hash for the record

`GET api/peers` Return an array of registered peers. Each peer is simply a string address of that peer

`POST api/mine-block` Create a new block
- Receive a json body: `{"data": "hi!"}`. Data can be any supported json type

`POST api/add-peer` Add a new peer server to the broadcast system and return nothing
- Receive a json body: `{"peer": "ws://localhost:9000"}`

Continue...