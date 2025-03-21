package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xmtp/xmtpd/contracts/pkg/groupmessages"
	"github.com/xmtp/xmtpd/contracts/pkg/identityupdates"
)

// Construct a raw blockchain listener that can be used to listen for events across many contract event types
type LogStreamBuilder interface {
	ListenForContractEvent(
		fromBlock uint64,
		contractAddress common.Address,
		topic common.Hash,
	) <-chan types.Log
	Build() (LogStreamer, error)
}

type LogStreamer interface {
	Start(ctx context.Context) error
}

type ChainClient interface {
	ethereum.BlockNumberReader
	ethereum.LogFilterer
	ethereum.ChainIDReader
	ethereum.ChainReader
}

type TransactionSigner interface {
	FromAddress() common.Address
	SignerFunc() bind.SignerFn
}

type IBlockchainPublisher interface {
	PublishIdentityUpdate(
		ctx context.Context,
		inboxId [32]byte,
		identityUpdate []byte,
	) (*identityupdates.IdentityUpdatesIdentityUpdateCreated, error)
	PublishGroupMessage(
		ctx context.Context,
		groupdId [32]byte,
		message []byte,
	) (*groupmessages.GroupMessagesMessageSent, error)
}
