package eth

import (
	"context"
	"fmt"
	"testing"
)

var ctx *Client

func TestConnect(t *testing.T) {
	client, err := Connect("http://localhost:8545")

	if err != nil {
		t.Errorf(err.Error())
	}
	ctx = client
	TestGetBlockNumber(t)
}

func TestGetBlockNumber(t *testing.T) {
	b, err := ctx.GetBlockNumber(context.TODO())

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(b.String())
}
