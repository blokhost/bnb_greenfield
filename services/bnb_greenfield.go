package services

import (
	"context"
	"github.com/bnb-chain/greenfield-go-sdk/client"
	"github.com/bnb-chain/greenfield-go-sdk/types"
	"github.com/bnb-chain/greenfield/sdk/keys"
	"github.com/cloakd/common/services"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/tendermint/tendermint/p2p"
	"log"
	"os"
)

type BnbGreenfieldService struct {
	services.DefaultService

	cli client.Client
}

const BNB_GREENFIELD_SVC = "bnb_greenfield_svc"

func (svc BnbGreenfieldService) Id() string {
	return BNB_GREENFIELD_SVC
}

func (svc *BnbGreenfieldService) Start() error {
	keyring, err := keys.NewMnemonicKeyManager(os.Getenv("PRIVATE_KEY"))

	account, err := types.NewAccountFromPrivateKey("test", keyring.GetPrivKey().String())
	if err != nil {
		log.Fatalf("New account from private key error, %v", err)
	}

	svc.cli, err = client.New("5601", "https://gnfd-bsc-testnet-dataseed1.bnbchain.org", client.Option{DefaultAccount: account})
	if err != nil {
		log.Fatalf("unable to new greenfield client, %v", err)
	}

	return nil
}

func (svc *BnbGreenfieldService) NodeInfo() (*p2p.DefaultNodeInfo, *tmservice.VersionInfo, error) {
	ctx := context.Background()
	nodeInfo, versionInfo, err := svc.cli.GetNodeInfo(ctx)
	if err != nil {
		log.Fatalf("unable to get node info, %v", err)
	}

	return nodeInfo, versionInfo, nil
}

func (svc *BnbGreenfieldService) LatesBlock() (*tmservice.Block, error) {
	ctx := context.Background()
	latestBlock, err := svc.cli.GetLatestBlock(ctx)
	if err != nil {
		log.Fatalf("unable to get latest block, %v", err)
	}

	return latestBlock, nil
}

func (svc *BnbGreenfieldService) CreateBucket(owner string) (string, error) {
	ctx := context.Background()

	cid, err := svc.cli.CreateBucket(ctx, "BlokHost Test", owner, types.CreateBucketOptions{
		Visibility: 1,
		//TxOpts:         &types2.TxOption{
		//	Mode:      nil,
		//	GasLimit:  0,
		//	Memo:      "",
		//	FeeAmount: nil,
		//	FeePayer:  nil,
		//	Nonce:     0,
		//},
		PaymentAddress: []byte(owner),
		ChargedQuota:   0,
	})
	if err != nil {
		return "", err
	}

	return cid, err
}
