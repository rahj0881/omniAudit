// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
)

// ReceiptCreate is the builder for creating a Receipt entity.
type ReceiptCreate struct {
	config
	mutation *ReceiptMutation
	hooks    []Hook
}

// SetUUID sets the "UUID" field.
func (rc *ReceiptCreate) SetUUID(u uuid.UUID) *ReceiptCreate {
	rc.mutation.SetUUID(u)
	return rc
}

// SetNillableUUID sets the "UUID" field if the given value is not nil.
func (rc *ReceiptCreate) SetNillableUUID(u *uuid.UUID) *ReceiptCreate {
	if u != nil {
		rc.SetUUID(*u)
	}
	return rc
}

// SetGasUsed sets the "GasUsed" field.
func (rc *ReceiptCreate) SetGasUsed(u uint64) *ReceiptCreate {
	rc.mutation.SetGasUsed(u)
	return rc
}

// SetSuccess sets the "Success" field.
func (rc *ReceiptCreate) SetSuccess(b bool) *ReceiptCreate {
	rc.mutation.SetSuccess(b)
	return rc
}

// SetRelayerAddress sets the "RelayerAddress" field.
func (rc *ReceiptCreate) SetRelayerAddress(b []byte) *ReceiptCreate {
	rc.mutation.SetRelayerAddress(b)
	return rc
}

// SetSourceChainID sets the "SourceChainID" field.
func (rc *ReceiptCreate) SetSourceChainID(u uint64) *ReceiptCreate {
	rc.mutation.SetSourceChainID(u)
	return rc
}

// SetDestChainID sets the "DestChainID" field.
func (rc *ReceiptCreate) SetDestChainID(u uint64) *ReceiptCreate {
	rc.mutation.SetDestChainID(u)
	return rc
}

// SetStreamOffset sets the "StreamOffset" field.
func (rc *ReceiptCreate) SetStreamOffset(u uint64) *ReceiptCreate {
	rc.mutation.SetStreamOffset(u)
	return rc
}

// SetTxHash sets the "TxHash" field.
func (rc *ReceiptCreate) SetTxHash(b []byte) *ReceiptCreate {
	rc.mutation.SetTxHash(b)
	return rc
}

// SetCreatedAt sets the "CreatedAt" field.
func (rc *ReceiptCreate) SetCreatedAt(t time.Time) *ReceiptCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (rc *ReceiptCreate) SetNillableCreatedAt(t *time.Time) *ReceiptCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetBlockID sets the "Block" edge to the Block entity by ID.
func (rc *ReceiptCreate) SetBlockID(id int) *ReceiptCreate {
	rc.mutation.SetBlockID(id)
	return rc
}

// SetNillableBlockID sets the "Block" edge to the Block entity by ID if the given value is not nil.
func (rc *ReceiptCreate) SetNillableBlockID(id *int) *ReceiptCreate {
	if id != nil {
		rc = rc.SetBlockID(*id)
	}
	return rc
}

// SetBlock sets the "Block" edge to the Block entity.
func (rc *ReceiptCreate) SetBlock(b *Block) *ReceiptCreate {
	return rc.SetBlockID(b.ID)
}

// Mutation returns the ReceiptMutation object of the builder.
func (rc *ReceiptCreate) Mutation() *ReceiptMutation {
	return rc.mutation
}

// Save creates the Receipt in the database.
func (rc *ReceiptCreate) Save(ctx context.Context) (*Receipt, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReceiptCreate) SaveX(ctx context.Context) *Receipt {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReceiptCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReceiptCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReceiptCreate) defaults() {
	if _, ok := rc.mutation.UUID(); !ok {
		v := receipt.DefaultUUID()
		rc.mutation.SetUUID(v)
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := receipt.DefaultCreatedAt
		rc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReceiptCreate) check() error {
	if _, ok := rc.mutation.UUID(); !ok {
		return &ValidationError{Name: "UUID", err: errors.New(`ent: missing required field "Receipt.UUID"`)}
	}
	if _, ok := rc.mutation.GasUsed(); !ok {
		return &ValidationError{Name: "GasUsed", err: errors.New(`ent: missing required field "Receipt.GasUsed"`)}
	}
	if _, ok := rc.mutation.Success(); !ok {
		return &ValidationError{Name: "Success", err: errors.New(`ent: missing required field "Receipt.Success"`)}
	}
	if _, ok := rc.mutation.RelayerAddress(); !ok {
		return &ValidationError{Name: "RelayerAddress", err: errors.New(`ent: missing required field "Receipt.RelayerAddress"`)}
	}
	if v, ok := rc.mutation.RelayerAddress(); ok {
		if err := receipt.RelayerAddressValidator(v); err != nil {
			return &ValidationError{Name: "RelayerAddress", err: fmt.Errorf(`ent: validator failed for field "Receipt.RelayerAddress": %w`, err)}
		}
	}
	if _, ok := rc.mutation.SourceChainID(); !ok {
		return &ValidationError{Name: "SourceChainID", err: errors.New(`ent: missing required field "Receipt.SourceChainID"`)}
	}
	if _, ok := rc.mutation.DestChainID(); !ok {
		return &ValidationError{Name: "DestChainID", err: errors.New(`ent: missing required field "Receipt.DestChainID"`)}
	}
	if _, ok := rc.mutation.StreamOffset(); !ok {
		return &ValidationError{Name: "StreamOffset", err: errors.New(`ent: missing required field "Receipt.StreamOffset"`)}
	}
	if _, ok := rc.mutation.TxHash(); !ok {
		return &ValidationError{Name: "TxHash", err: errors.New(`ent: missing required field "Receipt.TxHash"`)}
	}
	if v, ok := rc.mutation.TxHash(); ok {
		if err := receipt.TxHashValidator(v); err != nil {
			return &ValidationError{Name: "TxHash", err: fmt.Errorf(`ent: validator failed for field "Receipt.TxHash": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "Receipt.CreatedAt"`)}
	}
	return nil
}

func (rc *ReceiptCreate) sqlSave(ctx context.Context) (*Receipt, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReceiptCreate) createSpec() (*Receipt, *sqlgraph.CreateSpec) {
	var (
		_node = &Receipt{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(receipt.Table, sqlgraph.NewFieldSpec(receipt.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.UUID(); ok {
		_spec.SetField(receipt.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := rc.mutation.GasUsed(); ok {
		_spec.SetField(receipt.FieldGasUsed, field.TypeUint64, value)
		_node.GasUsed = value
	}
	if value, ok := rc.mutation.Success(); ok {
		_spec.SetField(receipt.FieldSuccess, field.TypeBool, value)
		_node.Success = value
	}
	if value, ok := rc.mutation.RelayerAddress(); ok {
		_spec.SetField(receipt.FieldRelayerAddress, field.TypeBytes, value)
		_node.RelayerAddress = value
	}
	if value, ok := rc.mutation.SourceChainID(); ok {
		_spec.SetField(receipt.FieldSourceChainID, field.TypeUint64, value)
		_node.SourceChainID = value
	}
	if value, ok := rc.mutation.DestChainID(); ok {
		_spec.SetField(receipt.FieldDestChainID, field.TypeUint64, value)
		_node.DestChainID = value
	}
	if value, ok := rc.mutation.StreamOffset(); ok {
		_spec.SetField(receipt.FieldStreamOffset, field.TypeUint64, value)
		_node.StreamOffset = value
	}
	if value, ok := rc.mutation.TxHash(); ok {
		_spec.SetField(receipt.FieldTxHash, field.TypeBytes, value)
		_node.TxHash = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(receipt.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := rc.mutation.BlockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   receipt.BlockTable,
			Columns: []string{receipt.BlockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.block_receipts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReceiptCreateBulk is the builder for creating many Receipt entities in bulk.
type ReceiptCreateBulk struct {
	config
	err      error
	builders []*ReceiptCreate
}

// Save creates the Receipt entities in the database.
func (rcb *ReceiptCreateBulk) Save(ctx context.Context) ([]*Receipt, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Receipt, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReceiptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReceiptCreateBulk) SaveX(ctx context.Context) []*Receipt {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReceiptCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReceiptCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
