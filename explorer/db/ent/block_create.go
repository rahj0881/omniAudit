// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
)

// BlockCreate is the builder for creating a Block entity.
type BlockCreate struct {
	config
	mutation *BlockMutation
	hooks    []Hook
}

// SetHash sets the "hash" field.
func (bc *BlockCreate) SetHash(b []byte) *BlockCreate {
	bc.mutation.SetHash(b)
	return bc
}

// SetChainID sets the "chain_id" field.
func (bc *BlockCreate) SetChainID(u uint64) *BlockCreate {
	bc.mutation.SetChainID(u)
	return bc
}

// SetHeight sets the "height" field.
func (bc *BlockCreate) SetHeight(u uint64) *BlockCreate {
	bc.mutation.SetHeight(u)
	return bc
}

// SetOffset sets the "offset" field.
func (bc *BlockCreate) SetOffset(u uint64) *BlockCreate {
	bc.mutation.SetOffset(u)
	return bc
}

// SetTimestamp sets the "timestamp" field.
func (bc *BlockCreate) SetTimestamp(t time.Time) *BlockCreate {
	bc.mutation.SetTimestamp(t)
	return bc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (bc *BlockCreate) SetNillableTimestamp(t *time.Time) *BlockCreate {
	if t != nil {
		bc.SetTimestamp(*t)
	}
	return bc
}

// SetCreatedAt sets the "created_at" field.
func (bc *BlockCreate) SetCreatedAt(t time.Time) *BlockCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BlockCreate) SetNillableCreatedAt(t *time.Time) *BlockCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// AddMsgIDs adds the "msgs" edge to the Msg entity by IDs.
func (bc *BlockCreate) AddMsgIDs(ids ...int) *BlockCreate {
	bc.mutation.AddMsgIDs(ids...)
	return bc
}

// AddMsgs adds the "msgs" edges to the Msg entity.
func (bc *BlockCreate) AddMsgs(m ...*Msg) *BlockCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return bc.AddMsgIDs(ids...)
}

// AddReceiptIDs adds the "receipts" edge to the Receipt entity by IDs.
func (bc *BlockCreate) AddReceiptIDs(ids ...int) *BlockCreate {
	bc.mutation.AddReceiptIDs(ids...)
	return bc
}

// AddReceipts adds the "receipts" edges to the Receipt entity.
func (bc *BlockCreate) AddReceipts(r ...*Receipt) *BlockCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return bc.AddReceiptIDs(ids...)
}

// Mutation returns the BlockMutation object of the builder.
func (bc *BlockCreate) Mutation() *BlockMutation {
	return bc.mutation
}

// Save creates the Block in the database.
func (bc *BlockCreate) Save(ctx context.Context) (*Block, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockCreate) SaveX(ctx context.Context) *Block {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlockCreate) defaults() {
	if _, ok := bc.mutation.Timestamp(); !ok {
		v := block.DefaultTimestamp
		bc.mutation.SetTimestamp(v)
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := block.DefaultCreatedAt
		bc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockCreate) check() error {
	if _, ok := bc.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Block.hash"`)}
	}
	if v, ok := bc.mutation.Hash(); ok {
		if err := block.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Block.hash": %w`, err)}
		}
	}
	if _, ok := bc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Block.chain_id"`)}
	}
	if _, ok := bc.mutation.Height(); !ok {
		return &ValidationError{Name: "height", err: errors.New(`ent: missing required field "Block.height"`)}
	}
	if _, ok := bc.mutation.Offset(); !ok {
		return &ValidationError{Name: "offset", err: errors.New(`ent: missing required field "Block.offset"`)}
	}
	if _, ok := bc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Block.timestamp"`)}
	}
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Block.created_at"`)}
	}
	return nil
}

func (bc *BlockCreate) sqlSave(ctx context.Context) (*Block, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlockCreate) createSpec() (*Block, *sqlgraph.CreateSpec) {
	var (
		_node = &Block{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(block.Table, sqlgraph.NewFieldSpec(block.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.Hash(); ok {
		_spec.SetField(block.FieldHash, field.TypeBytes, value)
		_node.Hash = value
	}
	if value, ok := bc.mutation.ChainID(); ok {
		_spec.SetField(block.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := bc.mutation.Height(); ok {
		_spec.SetField(block.FieldHeight, field.TypeUint64, value)
		_node.Height = value
	}
	if value, ok := bc.mutation.Offset(); ok {
		_spec.SetField(block.FieldOffset, field.TypeUint64, value)
		_node.Offset = value
	}
	if value, ok := bc.mutation.Timestamp(); ok {
		_spec.SetField(block.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(block.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := bc.mutation.MsgsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   block.MsgsTable,
			Columns: block.MsgsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(msg.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ReceiptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   block.ReceiptsTable,
			Columns: block.ReceiptsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(receipt.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlockCreateBulk is the builder for creating many Block entities in bulk.
type BlockCreateBulk struct {
	config
	err      error
	builders []*BlockCreate
}

// Save creates the Block entities in the database.
func (bcb *BlockCreateBulk) Save(ctx context.Context) ([]*Block, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Block, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockCreateBulk) SaveX(ctx context.Context) []*Block {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
