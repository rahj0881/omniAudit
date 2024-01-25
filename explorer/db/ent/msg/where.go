// Code generated by ent, DO NOT EDIT.

package msg

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "UUID" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldUUID, v))
}

// SourceMsgSender applies equality check predicate on the "SourceMsgSender" field. It's identical to SourceMsgSenderEQ.
func SourceMsgSender(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldSourceMsgSender, v))
}

// DestAddress applies equality check predicate on the "DestAddress" field. It's identical to DestAddressEQ.
func DestAddress(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestAddress, v))
}

// Data applies equality check predicate on the "Data" field. It's identical to DataEQ.
func Data(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldData, v))
}

// DestGasLimit applies equality check predicate on the "DestGasLimit" field. It's identical to DestGasLimitEQ.
func DestGasLimit(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestGasLimit, v))
}

// SourceChainID applies equality check predicate on the "SourceChainID" field. It's identical to SourceChainIDEQ.
func SourceChainID(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldSourceChainID, v))
}

// DestChainID applies equality check predicate on the "DestChainID" field. It's identical to DestChainIDEQ.
func DestChainID(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestChainID, v))
}

// StreamOffset applies equality check predicate on the "StreamOffset" field. It's identical to StreamOffsetEQ.
func StreamOffset(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldStreamOffset, v))
}

// TxHash applies equality check predicate on the "TxHash" field. It's identical to TxHashEQ.
func TxHash(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldTxHash, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldCreatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "UUID" field.
func UUIDEQ(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "UUID" field.
func UUIDNEQ(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "UUID" field.
func UUIDIn(vs ...uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "UUID" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "UUID" field.
func UUIDGT(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "UUID" field.
func UUIDGTE(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "UUID" field.
func UUIDLT(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "UUID" field.
func UUIDLTE(v uuid.UUID) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldUUID, v))
}

// SourceMsgSenderEQ applies the EQ predicate on the "SourceMsgSender" field.
func SourceMsgSenderEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldSourceMsgSender, v))
}

// SourceMsgSenderNEQ applies the NEQ predicate on the "SourceMsgSender" field.
func SourceMsgSenderNEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldSourceMsgSender, v))
}

// SourceMsgSenderIn applies the In predicate on the "SourceMsgSender" field.
func SourceMsgSenderIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldSourceMsgSender, vs...))
}

// SourceMsgSenderNotIn applies the NotIn predicate on the "SourceMsgSender" field.
func SourceMsgSenderNotIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldSourceMsgSender, vs...))
}

// SourceMsgSenderGT applies the GT predicate on the "SourceMsgSender" field.
func SourceMsgSenderGT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldSourceMsgSender, v))
}

// SourceMsgSenderGTE applies the GTE predicate on the "SourceMsgSender" field.
func SourceMsgSenderGTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldSourceMsgSender, v))
}

// SourceMsgSenderLT applies the LT predicate on the "SourceMsgSender" field.
func SourceMsgSenderLT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldSourceMsgSender, v))
}

// SourceMsgSenderLTE applies the LTE predicate on the "SourceMsgSender" field.
func SourceMsgSenderLTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldSourceMsgSender, v))
}

// DestAddressEQ applies the EQ predicate on the "DestAddress" field.
func DestAddressEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestAddress, v))
}

// DestAddressNEQ applies the NEQ predicate on the "DestAddress" field.
func DestAddressNEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldDestAddress, v))
}

// DestAddressIn applies the In predicate on the "DestAddress" field.
func DestAddressIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldDestAddress, vs...))
}

// DestAddressNotIn applies the NotIn predicate on the "DestAddress" field.
func DestAddressNotIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldDestAddress, vs...))
}

// DestAddressGT applies the GT predicate on the "DestAddress" field.
func DestAddressGT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldDestAddress, v))
}

// DestAddressGTE applies the GTE predicate on the "DestAddress" field.
func DestAddressGTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldDestAddress, v))
}

// DestAddressLT applies the LT predicate on the "DestAddress" field.
func DestAddressLT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldDestAddress, v))
}

// DestAddressLTE applies the LTE predicate on the "DestAddress" field.
func DestAddressLTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldDestAddress, v))
}

// DataEQ applies the EQ predicate on the "Data" field.
func DataEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "Data" field.
func DataNEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "Data" field.
func DataIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "Data" field.
func DataNotIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "Data" field.
func DataGT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "Data" field.
func DataGTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "Data" field.
func DataLT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "Data" field.
func DataLTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldData, v))
}

// DestGasLimitEQ applies the EQ predicate on the "DestGasLimit" field.
func DestGasLimitEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestGasLimit, v))
}

// DestGasLimitNEQ applies the NEQ predicate on the "DestGasLimit" field.
func DestGasLimitNEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldDestGasLimit, v))
}

// DestGasLimitIn applies the In predicate on the "DestGasLimit" field.
func DestGasLimitIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldDestGasLimit, vs...))
}

// DestGasLimitNotIn applies the NotIn predicate on the "DestGasLimit" field.
func DestGasLimitNotIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldDestGasLimit, vs...))
}

// DestGasLimitGT applies the GT predicate on the "DestGasLimit" field.
func DestGasLimitGT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldDestGasLimit, v))
}

// DestGasLimitGTE applies the GTE predicate on the "DestGasLimit" field.
func DestGasLimitGTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldDestGasLimit, v))
}

// DestGasLimitLT applies the LT predicate on the "DestGasLimit" field.
func DestGasLimitLT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldDestGasLimit, v))
}

// DestGasLimitLTE applies the LTE predicate on the "DestGasLimit" field.
func DestGasLimitLTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldDestGasLimit, v))
}

// SourceChainIDEQ applies the EQ predicate on the "SourceChainID" field.
func SourceChainIDEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldSourceChainID, v))
}

// SourceChainIDNEQ applies the NEQ predicate on the "SourceChainID" field.
func SourceChainIDNEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldSourceChainID, v))
}

// SourceChainIDIn applies the In predicate on the "SourceChainID" field.
func SourceChainIDIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldSourceChainID, vs...))
}

// SourceChainIDNotIn applies the NotIn predicate on the "SourceChainID" field.
func SourceChainIDNotIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldSourceChainID, vs...))
}

// SourceChainIDGT applies the GT predicate on the "SourceChainID" field.
func SourceChainIDGT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldSourceChainID, v))
}

// SourceChainIDGTE applies the GTE predicate on the "SourceChainID" field.
func SourceChainIDGTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldSourceChainID, v))
}

// SourceChainIDLT applies the LT predicate on the "SourceChainID" field.
func SourceChainIDLT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldSourceChainID, v))
}

// SourceChainIDLTE applies the LTE predicate on the "SourceChainID" field.
func SourceChainIDLTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldSourceChainID, v))
}

// DestChainIDEQ applies the EQ predicate on the "DestChainID" field.
func DestChainIDEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldDestChainID, v))
}

// DestChainIDNEQ applies the NEQ predicate on the "DestChainID" field.
func DestChainIDNEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldDestChainID, v))
}

// DestChainIDIn applies the In predicate on the "DestChainID" field.
func DestChainIDIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldDestChainID, vs...))
}

// DestChainIDNotIn applies the NotIn predicate on the "DestChainID" field.
func DestChainIDNotIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldDestChainID, vs...))
}

// DestChainIDGT applies the GT predicate on the "DestChainID" field.
func DestChainIDGT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldDestChainID, v))
}

// DestChainIDGTE applies the GTE predicate on the "DestChainID" field.
func DestChainIDGTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldDestChainID, v))
}

// DestChainIDLT applies the LT predicate on the "DestChainID" field.
func DestChainIDLT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldDestChainID, v))
}

// DestChainIDLTE applies the LTE predicate on the "DestChainID" field.
func DestChainIDLTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldDestChainID, v))
}

// StreamOffsetEQ applies the EQ predicate on the "StreamOffset" field.
func StreamOffsetEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldStreamOffset, v))
}

// StreamOffsetNEQ applies the NEQ predicate on the "StreamOffset" field.
func StreamOffsetNEQ(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldStreamOffset, v))
}

// StreamOffsetIn applies the In predicate on the "StreamOffset" field.
func StreamOffsetIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldStreamOffset, vs...))
}

// StreamOffsetNotIn applies the NotIn predicate on the "StreamOffset" field.
func StreamOffsetNotIn(vs ...uint64) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldStreamOffset, vs...))
}

// StreamOffsetGT applies the GT predicate on the "StreamOffset" field.
func StreamOffsetGT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldStreamOffset, v))
}

// StreamOffsetGTE applies the GTE predicate on the "StreamOffset" field.
func StreamOffsetGTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldStreamOffset, v))
}

// StreamOffsetLT applies the LT predicate on the "StreamOffset" field.
func StreamOffsetLT(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldStreamOffset, v))
}

// StreamOffsetLTE applies the LTE predicate on the "StreamOffset" field.
func StreamOffsetLTE(v uint64) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldStreamOffset, v))
}

// TxHashEQ applies the EQ predicate on the "TxHash" field.
func TxHashEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldTxHash, v))
}

// TxHashNEQ applies the NEQ predicate on the "TxHash" field.
func TxHashNEQ(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldTxHash, v))
}

// TxHashIn applies the In predicate on the "TxHash" field.
func TxHashIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldTxHash, vs...))
}

// TxHashNotIn applies the NotIn predicate on the "TxHash" field.
func TxHashNotIn(vs ...[]byte) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldTxHash, vs...))
}

// TxHashGT applies the GT predicate on the "TxHash" field.
func TxHashGT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldTxHash, v))
}

// TxHashGTE applies the GTE predicate on the "TxHash" field.
func TxHashGTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldTxHash, v))
}

// TxHashLT applies the LT predicate on the "TxHash" field.
func TxHashLT(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldTxHash, v))
}

// TxHashLTE applies the LTE predicate on the "TxHash" field.
func TxHashLTE(v []byte) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldTxHash, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Msg {
	return predicate.Msg(sql.FieldLTE(FieldCreatedAt, v))
}

// HasBlock applies the HasEdge predicate on the "Block" edge.
func HasBlock() predicate.Msg {
	return predicate.Msg(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlockTable, BlockColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlockWith applies the HasEdge predicate on the "Block" edge with a given conditions (other predicates).
func HasBlockWith(preds ...predicate.Block) predicate.Msg {
	return predicate.Msg(func(s *sql.Selector) {
		step := newBlockStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Msg) predicate.Msg {
	return predicate.Msg(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Msg) predicate.Msg {
	return predicate.Msg(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Msg) predicate.Msg {
	return predicate.Msg(sql.NotPredicates(p))
}
