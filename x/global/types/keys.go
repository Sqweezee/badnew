package types

const (
	// ModuleName defines the module name
	ModuleName = "global"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_global"
)

var ParamsKey = []byte{0x00}
