// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"github.com/omni-network/omni/explorer/db/ent/schema\",\"Package\":\"github.com/omni-network/omni/explorer/db/ent\",\"Schemas\":[{\"name\":\"Block\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"msgs\",\"type\":\"Msg\"},{\"name\":\"receipts\",\"type\":\"Receipt\"}],\"fields\":[{\"name\":\"hash\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"height\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"offset\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"timestamp\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088245+03:00\",\"default_kind\":25,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088245+03:00\",\"default_kind\":25,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"unique\":true,\"fields\":[\"chain_id\",\"hash\"]},{\"unique\":true,\"fields\":[\"chain_id\",\"height\"]}]},{\"name\":\"Chain\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088663+03:00\",\"default_kind\":25,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"name\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}}]},{\"name\":\"Msg\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"block\",\"type\":\"Block\",\"ref_name\":\"msgs\",\"inverse\":true},{\"name\":\"receipts\",\"type\":\"Receipt\"}],\"fields\":[{\"name\":\"sender\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":20,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"to\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":20,\"validators\":1,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"data\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"gas_limit\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"source_chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"dest_chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"offset\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"tx_hash\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"receipt_hash\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":32,\"optional\":true,\"validators\":1,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"status\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"optional\":true,\"default\":true,\"default_value\":\"PENDING\",\"default_kind\":24,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088671+03:00\",\"default_kind\":25,\"position\":{\"Index\":10,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"sender\"]},{\"fields\":[\"to\"]},{\"fields\":[\"status\"]},{\"fields\":[\"tx_hash\"]},{\"unique\":true,\"fields\":[\"source_chain_id\",\"dest_chain_id\",\"offset\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}]},{\"name\":\"Receipt\",\"config\":{\"Table\":\"\"},\"edges\":[{\"name\":\"block\",\"type\":\"Block\",\"ref_name\":\"receipts\",\"inverse\":true},{\"name\":\"msgs\",\"type\":\"Msg\",\"ref_name\":\"receipts\",\"inverse\":true}],\"fields\":[{\"name\":\"block_hash\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"gas_used\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"success\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"relayer_address\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":20,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"source_chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"dest_chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"offset\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"tx_hash\",\"type\":{\"Type\":5,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":true,\"RType\":null},\"size\":32,\"validators\":1,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088704+03:00\",\"default_kind\":25,\"position\":{\"Index\":8,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"revert_reason\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"size\":255,\"optional\":true,\"validators\":1,\"position\":{\"Index\":9,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"fields\":[\"tx_hash\"]},{\"fields\":[\"source_chain_id\",\"dest_chain_id\",\"offset\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}]},{\"name\":\"XProviderCursor\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"chain_id\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"height\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"offset\",\"type\":{\"Type\":18,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088772+03:00\",\"default_kind\":25,\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":\"2024-06-03T11:37:16.088772+03:00\",\"default_kind\":25,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"unique\":true,\"fields\":[\"chain_id\"]}]}],\"Features\":[\"entql\",\"sql/versioned-migration\",\"privacy\",\"schema/snapshot\"]}"
