package main

import (
	"flag"

	"github.com/labstack/gommon/log"
	"github.com/mauri870/gochain"
)

var address string

func init() {
	flag.StringVar(&address, "address", ":3000", "The address to listen on")
	flag.Parse()
}

func main() {
	// Create our first block a.k.a the genesis block
	genesisBlock := gochain.NewBlock()

	// create our handler
	h := newBlockChainHandler(genesisBlock)

	h.Logger.SetLevel(log.INFO)

	// define routes
	api := h.Group("/api")
	api.GET("/blocks", h.handleGetBlocks)
	api.GET("/peers", h.handleGetPeers)
	api.POST("/mine-block", h.handlePostMineBlock)
	api.POST("/add-peer", h.handlePostAddPeer)

	h.Logger.Fatal(h.Start(address))
}
