// Code generated by ent, DO NOT EDIT.

package block

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "UUID" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldUUID, v))
}

// SourceChainID applies equality check predicate on the "SourceChainID" field. It's identical to SourceChainIDEQ.
func SourceChainID(v uint64) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldSourceChainID, v))
}

// BlockHeight applies equality check predicate on the "BlockHeight" field. It's identical to BlockHeightEQ.
func BlockHeight(v uint64) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldBlockHeight, v))
}

// BlockHash applies equality check predicate on the "BlockHash" field. It's identical to BlockHashEQ.
func BlockHash(v []byte) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldBlockHash, v))
}

// Timestamp applies equality check predicate on the "Timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldTimestamp, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldCreatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "UUID" field.
func UUIDEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "UUID" field.
func UUIDNEQ(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "UUID" field.
func UUIDIn(vs ...uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "UUID" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "UUID" field.
func UUIDGT(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "UUID" field.
func UUIDGTE(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "UUID" field.
func UUIDLT(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "UUID" field.
func UUIDLTE(v uuid.UUID) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldUUID, v))
}

// SourceChainIDEQ applies the EQ predicate on the "SourceChainID" field.
func SourceChainIDEQ(v uint64) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldSourceChainID, v))
}

// SourceChainIDNEQ applies the NEQ predicate on the "SourceChainID" field.
func SourceChainIDNEQ(v uint64) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldSourceChainID, v))
}

// SourceChainIDIn applies the In predicate on the "SourceChainID" field.
func SourceChainIDIn(vs ...uint64) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldSourceChainID, vs...))
}

// SourceChainIDNotIn applies the NotIn predicate on the "SourceChainID" field.
func SourceChainIDNotIn(vs ...uint64) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldSourceChainID, vs...))
}

// SourceChainIDGT applies the GT predicate on the "SourceChainID" field.
func SourceChainIDGT(v uint64) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldSourceChainID, v))
}

// SourceChainIDGTE applies the GTE predicate on the "SourceChainID" field.
func SourceChainIDGTE(v uint64) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldSourceChainID, v))
}

// SourceChainIDLT applies the LT predicate on the "SourceChainID" field.
func SourceChainIDLT(v uint64) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldSourceChainID, v))
}

// SourceChainIDLTE applies the LTE predicate on the "SourceChainID" field.
func SourceChainIDLTE(v uint64) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldSourceChainID, v))
}

// BlockHeightEQ applies the EQ predicate on the "BlockHeight" field.
func BlockHeightEQ(v uint64) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldBlockHeight, v))
}

// BlockHeightNEQ applies the NEQ predicate on the "BlockHeight" field.
func BlockHeightNEQ(v uint64) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldBlockHeight, v))
}

// BlockHeightIn applies the In predicate on the "BlockHeight" field.
func BlockHeightIn(vs ...uint64) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldBlockHeight, vs...))
}

// BlockHeightNotIn applies the NotIn predicate on the "BlockHeight" field.
func BlockHeightNotIn(vs ...uint64) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldBlockHeight, vs...))
}

// BlockHeightGT applies the GT predicate on the "BlockHeight" field.
func BlockHeightGT(v uint64) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldBlockHeight, v))
}

// BlockHeightGTE applies the GTE predicate on the "BlockHeight" field.
func BlockHeightGTE(v uint64) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldBlockHeight, v))
}

// BlockHeightLT applies the LT predicate on the "BlockHeight" field.
func BlockHeightLT(v uint64) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldBlockHeight, v))
}

// BlockHeightLTE applies the LTE predicate on the "BlockHeight" field.
func BlockHeightLTE(v uint64) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldBlockHeight, v))
}

// BlockHashEQ applies the EQ predicate on the "BlockHash" field.
func BlockHashEQ(v []byte) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldBlockHash, v))
}

// BlockHashNEQ applies the NEQ predicate on the "BlockHash" field.
func BlockHashNEQ(v []byte) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldBlockHash, v))
}

// BlockHashIn applies the In predicate on the "BlockHash" field.
func BlockHashIn(vs ...[]byte) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldBlockHash, vs...))
}

// BlockHashNotIn applies the NotIn predicate on the "BlockHash" field.
func BlockHashNotIn(vs ...[]byte) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldBlockHash, vs...))
}

// BlockHashGT applies the GT predicate on the "BlockHash" field.
func BlockHashGT(v []byte) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldBlockHash, v))
}

// BlockHashGTE applies the GTE predicate on the "BlockHash" field.
func BlockHashGTE(v []byte) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldBlockHash, v))
}

// BlockHashLT applies the LT predicate on the "BlockHash" field.
func BlockHashLT(v []byte) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldBlockHash, v))
}

// BlockHashLTE applies the LTE predicate on the "BlockHash" field.
func BlockHashLTE(v []byte) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldBlockHash, v))
}

// TimestampEQ applies the EQ predicate on the "Timestamp" field.
func TimestampEQ(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "Timestamp" field.
func TimestampNEQ(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "Timestamp" field.
func TimestampIn(vs ...time.Time) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "Timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "Timestamp" field.
func TimestampGT(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "Timestamp" field.
func TimestampGTE(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "Timestamp" field.
func TimestampLT(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "Timestamp" field.
func TimestampLTE(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldTimestamp, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Block {
	return predicate.Block(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Block {
	return predicate.Block(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.Block {
	return predicate.Block(sql.FieldLTE(FieldCreatedAt, v))
}

// HasMsgs applies the HasEdge predicate on the "Msgs" edge.
func HasMsgs() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MsgsTable, MsgsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMsgsWith applies the HasEdge predicate on the "Msgs" edge with a given conditions (other predicates).
func HasMsgsWith(preds ...predicate.Msg) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := newMsgsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReceipts applies the HasEdge predicate on the "Receipts" edge.
func HasReceipts() predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceiptsTable, ReceiptsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceiptsWith applies the HasEdge predicate on the "Receipts" edge with a given conditions (other predicates).
func HasReceiptsWith(preds ...predicate.Receipt) predicate.Block {
	return predicate.Block(func(s *sql.Selector) {
		step := newReceiptsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Block) predicate.Block {
	return predicate.Block(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Block) predicate.Block {
	return predicate.Block(sql.NotPredicates(p))
}
