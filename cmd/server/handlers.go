package main

import (
	"net/http"

	"sync"

	"github.com/labstack/echo"
	"github.com/mauri870/gochain"
)

type peers struct {
	Peers      []string `json:"peers"`
	sync.Mutex `json:"-"`
}

type blockChainHandler struct {
	*echo.Echo
	*sync.Mutex
	BlockChain *gochain.BlockChain
	Peers      peers
}

func newBlockChainHandler(genesisBlock *gochain.Block) *blockChainHandler {
	h := &blockChainHandler{
		Echo:       echo.New(),
		BlockChain: gochain.NewBlockChain(genesisBlock),
	}

	return h
}

func (h *blockChainHandler) handleGetBlocks(c echo.Context) error {
	return c.JSON(http.StatusOK, h.BlockChain)
}

func (h *blockChainHandler) handleGetPeers(c echo.Context) error {
	return c.JSON(http.StatusOK, h.Peers)
}

func (h *blockChainHandler) handlePostMineBlock(c echo.Context) error {
	r := struct {
		Data interface{} `json:"data"`
	}{}

	if err := c.Bind(&r); err != nil {
		return err
	}

	newBlock, err := h.BlockChain.NextBlock(r.Data)
	if err != nil {
		return err
	}

	h.Logger.Infof("Created block %s", newBlock.Hash)
	return c.JSON(http.StatusOK, newBlock)
}

func (h *blockChainHandler) handlePostAddPeer(c echo.Context) error {
	r := struct {
		Peer string `json:"peer"`
	}{}

	if err := c.Bind(&r); err != nil {
		return err
	}

	if r.Peer == "" {
		return c.JSON(http.StatusBadRequest, "Empty peer")
	}

	h.Peers.Lock()
	defer h.Peers.Unlock()

	h.Peers.Peers = append(h.Peers.Peers, r.Peer)

	h.Logger.Infof("Added peer: %s", r.Peer)

	return c.JSON(http.StatusOK, "")
}
