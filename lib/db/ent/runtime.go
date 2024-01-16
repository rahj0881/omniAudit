// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/omni-network/omni/lib/db/ent/schema"
	"github.com/omni-network/omni/lib/db/ent/xblock"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	xblockFields := schema.XBlock{}.Fields()
	_ = xblockFields
	// xblockDescUUID is the schema descriptor for uuid field.
	xblockDescUUID := xblockFields[0].Descriptor()
	// xblock.DefaultUUID holds the default value on creation for the uuid field.
	xblock.DefaultUUID = xblockDescUUID.Default.(func() uuid.UUID)
}
