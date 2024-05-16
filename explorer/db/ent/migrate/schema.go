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
		{Name: "hash", Type: field.TypeBytes, Size: 32},
		{Name: "chain_id", Type: field.TypeUint64},
		{Name: "height", Type: field.TypeUint64},
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
				Name:    "block_chain_id_hash",
				Unique:  true,
				Columns: []*schema.Column{BlocksColumns[2], BlocksColumns[1]},
			},
			{
				Name:    "block_chain_id_height",
				Unique:  true,
				Columns: []*schema.Column{BlocksColumns[2], BlocksColumns[3]},
			},
		},
	}
	// ChainsColumns holds the columns for the "chains" table.
	ChainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "chain_id", Type: field.TypeUint64, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
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
		{Name: "sender", Type: field.TypeBytes, Size: 20},
		{Name: "to", Type: field.TypeBytes, Size: 20},
		{Name: "data", Type: field.TypeBytes},
		{Name: "gas_limit", Type: field.TypeUint64},
		{Name: "source_chain_id", Type: field.TypeUint64},
		{Name: "dest_chain_id", Type: field.TypeUint64},
		{Name: "offset", Type: field.TypeUint64},
		{Name: "tx_hash", Type: field.TypeBytes, Size: 32},
		{Name: "receipt_hash", Type: field.TypeBytes, Nullable: true, Size: 32},
		{Name: "status", Type: field.TypeString, Nullable: true, Default: "PENDING"},
		{Name: "created_at", Type: field.TypeTime},
	}
	// MsgsTable holds the schema information for the "msgs" table.
	MsgsTable = &schema.Table{
		Name:       "msgs",
		Columns:    MsgsColumns,
		PrimaryKey: []*schema.Column{MsgsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "msg_sender",
				Unique:  false,
				Columns: []*schema.Column{MsgsColumns[1]},
			},
			{
				Name:    "msg_to",
				Unique:  false,
				Columns: []*schema.Column{MsgsColumns[2]},
			},
			{
				Name:    "msg_status",
				Unique:  false,
				Columns: []*schema.Column{MsgsColumns[10]},
			},
			{
				Name:    "msg_tx_hash",
				Unique:  false,
				Columns: []*schema.Column{MsgsColumns[8]},
			},
			{
				Name:    "msg_source_chain_id_dest_chain_id_offset",
				Unique:  true,
				Columns: []*schema.Column{MsgsColumns[5], MsgsColumns[6], MsgsColumns[7]},
			},
		},
	}
	// ReceiptsColumns holds the columns for the "receipts" table.
	ReceiptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "block_hash", Type: field.TypeBytes, Size: 32},
		{Name: "gas_used", Type: field.TypeUint64},
		{Name: "success", Type: field.TypeBool},
		{Name: "relayer_address", Type: field.TypeBytes, Size: 20},
		{Name: "source_chain_id", Type: field.TypeUint64},
		{Name: "dest_chain_id", Type: field.TypeUint64},
		{Name: "offset", Type: field.TypeUint64},
		{Name: "tx_hash", Type: field.TypeBytes, Size: 32},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ReceiptsTable holds the schema information for the "receipts" table.
	ReceiptsTable = &schema.Table{
		Name:       "receipts",
		Columns:    ReceiptsColumns,
		PrimaryKey: []*schema.Column{ReceiptsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "receipt_tx_hash",
				Unique:  false,
				Columns: []*schema.Column{ReceiptsColumns[8]},
			},
			{
				Name:    "receipt_source_chain_id_dest_chain_id_offset",
				Unique:  false,
				Columns: []*schema.Column{ReceiptsColumns[5], ReceiptsColumns[6], ReceiptsColumns[7]},
			},
		},
	}
	// XproviderCursorsColumns holds the columns for the "xprovider_cursors" table.
	XproviderCursorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "chain_id", Type: field.TypeUint64, Unique: true},
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
				Columns: []*schema.Column{XproviderCursorsColumns[1]},
			},
		},
	}
	// BlockMsgsColumns holds the columns for the "block_msgs" table.
	BlockMsgsColumns = []*schema.Column{
		{Name: "block_id", Type: field.TypeInt},
		{Name: "msg_id", Type: field.TypeInt},
	}
	// BlockMsgsTable holds the schema information for the "block_msgs" table.
	BlockMsgsTable = &schema.Table{
		Name:       "block_msgs",
		Columns:    BlockMsgsColumns,
		PrimaryKey: []*schema.Column{BlockMsgsColumns[0], BlockMsgsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "block_msgs_block_id",
				Columns:    []*schema.Column{BlockMsgsColumns[0]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "block_msgs_msg_id",
				Columns:    []*schema.Column{BlockMsgsColumns[1]},
				RefColumns: []*schema.Column{MsgsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BlockReceiptsColumns holds the columns for the "block_receipts" table.
	BlockReceiptsColumns = []*schema.Column{
		{Name: "block_id", Type: field.TypeInt},
		{Name: "receipt_id", Type: field.TypeInt},
	}
	// BlockReceiptsTable holds the schema information for the "block_receipts" table.
	BlockReceiptsTable = &schema.Table{
		Name:       "block_receipts",
		Columns:    BlockReceiptsColumns,
		PrimaryKey: []*schema.Column{BlockReceiptsColumns[0], BlockReceiptsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "block_receipts_block_id",
				Columns:    []*schema.Column{BlockReceiptsColumns[0]},
				RefColumns: []*schema.Column{BlocksColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "block_receipts_receipt_id",
				Columns:    []*schema.Column{BlockReceiptsColumns[1]},
				RefColumns: []*schema.Column{ReceiptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// MsgReceiptsColumns holds the columns for the "msg_receipts" table.
	MsgReceiptsColumns = []*schema.Column{
		{Name: "msg_id", Type: field.TypeInt},
		{Name: "receipt_id", Type: field.TypeInt},
	}
	// MsgReceiptsTable holds the schema information for the "msg_receipts" table.
	MsgReceiptsTable = &schema.Table{
		Name:       "msg_receipts",
		Columns:    MsgReceiptsColumns,
		PrimaryKey: []*schema.Column{MsgReceiptsColumns[0], MsgReceiptsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "msg_receipts_msg_id",
				Columns:    []*schema.Column{MsgReceiptsColumns[0]},
				RefColumns: []*schema.Column{MsgsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "msg_receipts_receipt_id",
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
		BlockMsgsTable,
		BlockReceiptsTable,
		MsgReceiptsTable,
	}
)

func init() {
	BlockMsgsTable.ForeignKeys[0].RefTable = BlocksTable
	BlockMsgsTable.ForeignKeys[1].RefTable = MsgsTable
	BlockReceiptsTable.ForeignKeys[0].RefTable = BlocksTable
	BlockReceiptsTable.ForeignKeys[1].RefTable = ReceiptsTable
	MsgReceiptsTable.ForeignKeys[0].RefTable = MsgsTable
	MsgReceiptsTable.ForeignKeys[1].RefTable = ReceiptsTable
}
