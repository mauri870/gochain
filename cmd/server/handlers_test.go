package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"github.com/labstack/echo"
	"github.com/mauri870/gochain"
)

func TestGetBlocks(t *testing.T) {
	h := newBlockChainHandler(gochain.NewBlock())

	req, err := http.NewRequest(echo.GET, "/api/blocks", nil)
	if err != nil {
		t.Error(err)
	}

	rec := httptest.NewRecorder()
	c := h.NewContext(req, rec)

	if err := h.handleGetBlocks(c); err != nil {
		t.Error(err)
	}

	var blockchain gochain.BlockChain

	err = json.NewDecoder(rec.Body).Decode(&blockchain)
	if err != nil {
		t.Error(err)
	}
}

func TestPostMineBlock(t *testing.T) {
	h := newBlockChainHandler(gochain.NewBlock())

	var jsonBody = []byte(`{"data": "Hello World!"}`)
	req, err := http.NewRequest(echo.POST, "/api/mine-block", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := h.NewContext(req, rec)

	if err := h.handlePostMineBlock(c); err != nil {
		t.Error(err)
	}

	var block gochain.Block

	err = json.NewDecoder(rec.Body).Decode(&block)
	if err != nil {
		t.Error(err)
	}

	if h.BlockChain.GetLatestBlock().Data != block.Data {
		t.Error("The block data must be the same as the latest block inserted")
	}
}

func TestPostAddPeer(t *testing.T) {
	h := newBlockChainHandler(gochain.NewBlock())

	var jsonBody = []byte(`{"peer": "ws://localhost:1000"}`)
	req, err := http.NewRequest(echo.POST, "/api/add-peer", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	c := h.NewContext(req, rec)

	if err := h.handlePostAddPeer(c); err != nil {
		t.Error(err)
	}

	if len(h.Peers.Peers) != 1 {
		t.Errorf("The number of peers don't match, expect 1 got %d", len(h.Peers.Peers))
	}
}
