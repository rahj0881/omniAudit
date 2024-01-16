// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/omni-network/omni/lib/db/ent/predicate"
	"github.com/omni-network/omni/lib/db/ent/xblock"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeXBlock = "XBlock"
)

// XBlockMutation represents an operation that mutates the XBlock nodes in the graph.
type XBlockMutation struct {
	config
	op            Op
	typ           string
	id            *int
	uuid          *uuid.UUID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*XBlock, error)
	predicates    []predicate.XBlock
}

var _ ent.Mutation = (*XBlockMutation)(nil)

// xblockOption allows management of the mutation configuration using functional options.
type xblockOption func(*XBlockMutation)

// newXBlockMutation creates new mutation for the XBlock entity.
func newXBlockMutation(c config, op Op, opts ...xblockOption) *XBlockMutation {
	m := &XBlockMutation{
		config:        c,
		op:            op,
		typ:           TypeXBlock,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withXBlockID sets the ID field of the mutation.
func withXBlockID(id int) xblockOption {
	return func(m *XBlockMutation) {
		var (
			err   error
			once  sync.Once
			value *XBlock
		)
		m.oldValue = func(ctx context.Context) (*XBlock, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().XBlock.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withXBlock sets the old XBlock of the mutation.
func withXBlock(node *XBlock) xblockOption {
	return func(m *XBlockMutation) {
		m.oldValue = func(context.Context) (*XBlock, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m XBlockMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m XBlockMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *XBlockMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *XBlockMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().XBlock.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *XBlockMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *XBlockMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the XBlock entity.
// If the XBlock object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *XBlockMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *XBlockMutation) ResetUUID() {
	m.uuid = nil
}

// Where appends a list predicates to the XBlockMutation builder.
func (m *XBlockMutation) Where(ps ...predicate.XBlock) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the XBlockMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *XBlockMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.XBlock, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *XBlockMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *XBlockMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (XBlock).
func (m *XBlockMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *XBlockMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.uuid != nil {
		fields = append(fields, xblock.FieldUUID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *XBlockMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case xblock.FieldUUID:
		return m.UUID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *XBlockMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case xblock.FieldUUID:
		return m.OldUUID(ctx)
	}
	return nil, fmt.Errorf("unknown XBlock field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *XBlockMutation) SetField(name string, value ent.Value) error {
	switch name {
	case xblock.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	}
	return fmt.Errorf("unknown XBlock field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *XBlockMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *XBlockMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *XBlockMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown XBlock numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *XBlockMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *XBlockMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *XBlockMutation) ClearField(name string) error {
	return fmt.Errorf("unknown XBlock nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *XBlockMutation) ResetField(name string) error {
	switch name {
	case xblock.FieldUUID:
		m.ResetUUID()
		return nil
	}
	return fmt.Errorf("unknown XBlock field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *XBlockMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *XBlockMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *XBlockMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *XBlockMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *XBlockMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *XBlockMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *XBlockMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown XBlock unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *XBlockMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown XBlock edge %s", name)
}
