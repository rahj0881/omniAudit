// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/chain"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/db/ent/schema"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	blockFields := schema.Block{}.Fields()
	_ = blockFields
	// blockDescHash is the schema descriptor for hash field.
	blockDescHash := blockFields[0].Descriptor()
	// block.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	block.HashValidator = blockDescHash.Validators[0].(func([]byte) error)
	// blockDescTimestamp is the schema descriptor for timestamp field.
	blockDescTimestamp := blockFields[3].Descriptor()
	// block.DefaultTimestamp holds the default value on creation for the timestamp field.
	block.DefaultTimestamp = blockDescTimestamp.Default.(time.Time)
	// blockDescCreatedAt is the schema descriptor for created_at field.
	blockDescCreatedAt := blockFields[4].Descriptor()
	// block.DefaultCreatedAt holds the default value on creation for the created_at field.
	block.DefaultCreatedAt = blockDescCreatedAt.Default.(time.Time)
	chainFields := schema.Chain{}.Fields()
	_ = chainFields
	// chainDescCreatedAt is the schema descriptor for created_at field.
	chainDescCreatedAt := chainFields[1].Descriptor()
	// chain.DefaultCreatedAt holds the default value on creation for the created_at field.
	chain.DefaultCreatedAt = chainDescCreatedAt.Default.(time.Time)
	msgHooks := schema.Msg{}.Hooks()
	msg.Hooks[0] = msgHooks[0]
	msgFields := schema.Msg{}.Fields()
	_ = msgFields
	// msgDescBlockHash is the schema descriptor for block_hash field.
	msgDescBlockHash := msgFields[0].Descriptor()
	// msg.BlockHashValidator is a validator for the "block_hash" field. It is called by the builders before save.
	msg.BlockHashValidator = msgDescBlockHash.Validators[0].(func([]byte) error)
	// msgDescSender is the schema descriptor for sender field.
	msgDescSender := msgFields[3].Descriptor()
	// msg.SenderValidator is a validator for the "sender" field. It is called by the builders before save.
	msg.SenderValidator = msgDescSender.Validators[0].(func([]byte) error)
	// msgDescTo is the schema descriptor for to field.
	msgDescTo := msgFields[4].Descriptor()
	// msg.ToValidator is a validator for the "to" field. It is called by the builders before save.
	msg.ToValidator = msgDescTo.Validators[0].(func([]byte) error)
	// msgDescTxHash is the schema descriptor for tx_hash field.
	msgDescTxHash := msgFields[10].Descriptor()
	// msg.TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	msg.TxHashValidator = msgDescTxHash.Validators[0].(func([]byte) error)
	// msgDescReceiptHash is the schema descriptor for receipt_hash field.
	msgDescReceiptHash := msgFields[11].Descriptor()
	// msg.ReceiptHashValidator is a validator for the "receipt_hash" field. It is called by the builders before save.
	msg.ReceiptHashValidator = msgDescReceiptHash.Validators[0].(func([]byte) error)
	// msgDescStatus is the schema descriptor for status field.
	msgDescStatus := msgFields[12].Descriptor()
	// msg.DefaultStatus holds the default value on creation for the status field.
	msg.DefaultStatus = msgDescStatus.Default.(string)
	// msgDescCreatedAt is the schema descriptor for created_at field.
	msgDescCreatedAt := msgFields[13].Descriptor()
	// msg.DefaultCreatedAt holds the default value on creation for the created_at field.
	msg.DefaultCreatedAt = msgDescCreatedAt.Default.(time.Time)
	receiptHooks := schema.Receipt{}.Hooks()
	receipt.Hooks[0] = receiptHooks[0]
	receiptFields := schema.Receipt{}.Fields()
	_ = receiptFields
	// receiptDescBlockHash is the schema descriptor for block_hash field.
	receiptDescBlockHash := receiptFields[0].Descriptor()
	// receipt.BlockHashValidator is a validator for the "block_hash" field. It is called by the builders before save.
	receipt.BlockHashValidator = receiptDescBlockHash.Validators[0].(func([]byte) error)
	// receiptDescRelayerAddress is the schema descriptor for relayer_address field.
	receiptDescRelayerAddress := receiptFields[3].Descriptor()
	// receipt.RelayerAddressValidator is a validator for the "relayer_address" field. It is called by the builders before save.
	receipt.RelayerAddressValidator = receiptDescRelayerAddress.Validators[0].(func([]byte) error)
	// receiptDescTxHash is the schema descriptor for tx_hash field.
	receiptDescTxHash := receiptFields[7].Descriptor()
	// receipt.TxHashValidator is a validator for the "tx_hash" field. It is called by the builders before save.
	receipt.TxHashValidator = receiptDescTxHash.Validators[0].(func([]byte) error)
	// receiptDescCreatedAt is the schema descriptor for created_at field.
	receiptDescCreatedAt := receiptFields[8].Descriptor()
	// receipt.DefaultCreatedAt holds the default value on creation for the created_at field.
	receipt.DefaultCreatedAt = receiptDescCreatedAt.Default.(time.Time)
	xprovidercursorFields := schema.XProviderCursor{}.Fields()
	_ = xprovidercursorFields
	// xprovidercursorDescCreatedAt is the schema descriptor for created_at field.
	xprovidercursorDescCreatedAt := xprovidercursorFields[3].Descriptor()
	// xprovidercursor.DefaultCreatedAt holds the default value on creation for the created_at field.
	xprovidercursor.DefaultCreatedAt = xprovidercursorDescCreatedAt.Default.(time.Time)
	// xprovidercursorDescUpdatedAt is the schema descriptor for updated_at field.
	xprovidercursorDescUpdatedAt := xprovidercursorFields[4].Descriptor()
	// xprovidercursor.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	xprovidercursor.DefaultUpdatedAt = xprovidercursorDescUpdatedAt.Default.(time.Time)
	// xprovidercursorDescID is the schema descriptor for id field.
	xprovidercursorDescID := xprovidercursorFields[0].Descriptor()
	// xprovidercursor.DefaultID holds the default value on creation for the id field.
	xprovidercursor.DefaultID = xprovidercursorDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
