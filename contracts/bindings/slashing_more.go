package bindings

import (
	_ "embed"
)

const (
	SlashingDeployedBytecode = "0x6080604052348015600f57600080fd5b506004361060285760003560e01c8063f679d30514602d575b600080fd5b60336035565b005b60405133907fc3ef55ddda4bc9300706e15ab3aed03c762d8afd43a7d358a7b9503cb39f281b90600090a256fea2646970667358221220fe8c79947bd3f76512d36b0f60be955866e0acb251c4c801c263ddddb64b5ffd64736f6c63430008180033"
)

//go:embed slashing_storage_layout.json
var slashingStorageLayoutJSON []byte

var SlashingStorageLayout = mustGetStorageLayout(slashingStorageLayoutJSON)
