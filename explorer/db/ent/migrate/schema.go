// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlocksColumns holds the columns for the "blocks" table.
	BlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "source_chain_id", Type: field.TypeUint64},
		{Name: "block_height", Type: field.TypeUint64},
		{Name: "block_hash", Type: field.TypeBytes, Size: 32},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// BlocksTable holds the schema information for the "blocks" table.
	BlocksTable = &schema.Table{
		Name:       "blocks",
		Columns:    BlocksColumns,
		PrimaryKey: []*schema.Column{BlocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "block_source_chain_id_block_height",
				Unique:  false,
				Columns: []*schema.Column{BlocksColumns[1], BlocksColumns[2]},
			},
		},
	}
	// ChainsColumns holds the columns for the "chains" table.
	ChainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "chain_id", Type: field.TypeUint64},
		{Name: "name", Type: field.TypeString},
	}
	// ChainsTable holds the schema information for the "chains" table.
	ChainsTable = &schema.Table{
		Name:       "chains",
		Columns:    ChainsColumns,
		PrimaryKey: []*schema.Column{ChainsColumns[0]},
	}
	// MsgsColumns holds the columns for the "msgs" table.
	MsgsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "source_msg_sender", Type: field.TypeBytes, Size: 20},
		{Name: "dest_address", Type: field.TypeBytes, Size: 20},
		{Name: "data", Type: field.TypeBytes},
		{Name: "dest_gas_limit", Type: field.TypeUint64},
		{Name: "source_chain_id", Type: field.TypeUint64},
		{Name: "dest_chain_id", Type: field.TypeUint64},
		{Name: "stream_offset", Type: field.TypeUint64},
		{Name: "tx_hash", Type: field.TypeBytes, Size: 32},
		{Name: "block_hash", Type: field.TypeBytes, Size: 32},
		{Name: "block_height", Type: field.TypeUint64},
		{Name: "receipt_hash", Type: field.TypeBytes, Nullable: true, Size: 32},
		{Name: "status", Type: field.TypeString, Nullable: true, Default: "PENDING"},
		{Name: "block_time", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "block_id", Type: field.TypeInt, Nullable: true},
	}
	// MsgsTable holds the schema information for the "msgs" table.
	MsgsTable = &schema.Table{
		Name:       "msgs",
		Columns:    MsgsColumns,
		PrimaryKey: []*schema.Column{MsgsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msgs_blocks_Msgs",
				Columns:    []*schema.Column{MsgsColumns[16]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "msg_source_chain_id_dest_chain_id_stream_offset_block_id",
				Unique:  false,
				Columns: []*schema.Column{MsgsColumns[6], MsgsColumns[7], MsgsColumns[8], MsgsColumns[16]},
			},
		},
	}
	// ReceiptsColumns holds the columns for the "receipts" table.
	ReceiptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "gas_used", Type: field.TypeUint64},
		{Name: "success", Type: field.TypeBool},
		{Name: "relayer_address", Type: field.TypeBytes, Size: 20},
		{Name: "source_chain_id", Type: field.TypeUint64},
		{Name: "dest_chain_id", Type: field.TypeUint64},
		{Name: "stream_offset", Type: field.TypeUint64},
		{Name: "tx_hash", Type: field.TypeBytes, Size: 32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "block_id", Type: field.TypeInt, Nullable: true},
	}
	// ReceiptsTable holds the schema information for the "receipts" table.
	ReceiptsTable = &schema.Table{
		Name:       "receipts",
		Columns:    ReceiptsColumns,
		PrimaryKey: []*schema.Column{ReceiptsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "receipts_blocks_Receipts",
				Columns:    []*schema.Column{ReceiptsColumns[10]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "receipt_source_chain_id_dest_chain_id_stream_offset_block_id",
				Unique:  false,
				Columns: []*schema.Column{ReceiptsColumns[5], ReceiptsColumns[6], ReceiptsColumns[7], ReceiptsColumns[10]},
			},
		},
	}
	// XproviderCursorsColumns holds the columns for the "xprovider_cursors" table.
	XproviderCursorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "chain_id", Type: field.TypeUint64},
		{Name: "height", Type: field.TypeUint64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// XproviderCursorsTable holds the schema information for the "xprovider_cursors" table.
	XproviderCursorsTable = &schema.Table{
		Name:       "xprovider_cursors",
		Columns:    XproviderCursorsColumns,
		PrimaryKey: []*schema.Column{XproviderCursorsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "xprovidercursor_chain_id",
				Unique:  true,
				Columns: []*schema.Column{XproviderCursorsColumns[2]},
			},
		},
	}
	// MsgReceiptsColumns holds the columns for the "msg_Receipts" table.
	MsgReceiptsColumns = []*schema.Column{
		{Name: "msg_id", Type: field.TypeInt},
		{Name: "receipt_id", Type: field.TypeInt},
	}
	// MsgReceiptsTable holds the schema information for the "msg_Receipts" table.
	MsgReceiptsTable = &schema.Table{
		Name:       "msg_Receipts",
		Columns:    MsgReceiptsColumns,
		PrimaryKey: []*schema.Column{MsgReceiptsColumns[0], MsgReceiptsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_Receipts_msg_id",
				Columns:    []*schema.Column{MsgReceiptsColumns[0]},
				RefColumns: []*schema.Column{MsgsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "msg_Receipts_receipt_id",
				Columns:    []*schema.Column{MsgReceiptsColumns[1]},
				RefColumns: []*schema.Column{ReceiptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlocksTable,
		ChainsTable,
		MsgsTable,
		ReceiptsTable,
		XproviderCursorsTable,
		MsgReceiptsTable,
	}
)

func init() {
	MsgsTable.ForeignKeys[0].RefTable = BlocksTable
	ReceiptsTable.ForeignKeys[0].RefTable = BlocksTable
	MsgReceiptsTable.ForeignKeys[0].RefTable = MsgsTable
	MsgReceiptsTable.ForeignKeys[1].RefTable = ReceiptsTable
}
