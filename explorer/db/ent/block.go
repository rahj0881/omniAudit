// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/omni-network/omni/explorer/db/ent/block"
)

// Block is the model entity for the Block schema.
type Block struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash []byte `json:"hash,omitempty"`
	// ChainID holds the value of the "chain_id" field.
	ChainID uint64 `json:"chain_id,omitempty"`
	// Height holds the value of the "height" field.
	Height uint64 `json:"height,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlockQuery when eager-loading is set.
	Edges        BlockEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BlockEdges holds the relations/edges for other nodes in the graph.
type BlockEdges struct {
	// Msgs holds the value of the Msgs edge.
	Msgs []*Msg `json:"Msgs,omitempty"`
	// Receipts holds the value of the Receipts edge.
	Receipts []*Receipt `json:"Receipts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MsgsOrErr returns the Msgs value or an error if the edge
// was not loaded in eager-loading.
func (e BlockEdges) MsgsOrErr() ([]*Msg, error) {
	if e.loadedTypes[0] {
		return e.Msgs, nil
	}
	return nil, &NotLoadedError{edge: "Msgs"}
}

// ReceiptsOrErr returns the Receipts value or an error if the edge
// was not loaded in eager-loading.
func (e BlockEdges) ReceiptsOrErr() ([]*Receipt, error) {
	if e.loadedTypes[1] {
		return e.Receipts, nil
	}
	return nil, &NotLoadedError{edge: "Receipts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Block) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case block.FieldHash:
			values[i] = new([]byte)
		case block.FieldID, block.FieldChainID, block.FieldHeight:
			values[i] = new(sql.NullInt64)
		case block.FieldTimestamp, block.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Block fields.
func (b *Block) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case block.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case block.FieldHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value != nil {
				b.Hash = *value
			}
		case block.FieldChainID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field chain_id", values[i])
			} else if value.Valid {
				b.ChainID = uint64(value.Int64)
			}
		case block.FieldHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field height", values[i])
			} else if value.Valid {
				b.Height = uint64(value.Int64)
			}
		case block.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				b.Timestamp = value.Time
			}
		case block.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Block.
// This includes values selected through modifiers, order, etc.
func (b *Block) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryMsgs queries the "Msgs" edge of the Block entity.
func (b *Block) QueryMsgs() *MsgQuery {
	return NewBlockClient(b.config).QueryMsgs(b)
}

// QueryReceipts queries the "Receipts" edge of the Block entity.
func (b *Block) QueryReceipts() *ReceiptQuery {
	return NewBlockClient(b.config).QueryReceipts(b)
}

// Update returns a builder for updating this Block.
// Note that you need to call Block.Unwrap() before calling this method if this Block
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Block) Update() *BlockUpdateOne {
	return NewBlockClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Block entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Block) Unwrap() *Block {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Block is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Block) String() string {
	var builder strings.Builder
	builder.WriteString("Block(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("hash=")
	builder.WriteString(fmt.Sprintf("%v", b.Hash))
	builder.WriteString(", ")
	builder.WriteString("chain_id=")
	builder.WriteString(fmt.Sprintf("%v", b.ChainID))
	builder.WriteString(", ")
	builder.WriteString("height=")
	builder.WriteString(fmt.Sprintf("%v", b.Height))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(b.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Blocks is a parsable slice of Block.
type Blocks []*Block
