/*
*  Define types and structures relate header
 */
package types

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ChainID       uint64   `json:"chainId"    gencodec:"required"`     // chainid
	PrevBlockHash Hash     `json:"prevHash"    gencodec:"required"`    // preblock hash
	StateRoot     Hash     `json:"stateRoot"    gencodec:"required"`   // statedb root
	TxRoot        Hash     `json:"txRoot"    gencodec:"required"`      // transactions root
	ReceiptsRoot  Hash     `json:"receipsRoot"    gencodec:"required"` // receipt root
	Height        uint64   `json:"height"    gencodec:"required"`      // block height
	Timestamp     uint64   `json:"timestamp"    gencodec:"required"`   // timestamp
	MixDigest     Hash     `json:"mixDigest"    gencodec:"required"`   // digest
	SigData       [][]byte `json:"sigData"    gencodec:"required"`     // signatures
	BlockHash     Hash     `json:"blockHash"    gencodec:"required"`   // block hash
}
