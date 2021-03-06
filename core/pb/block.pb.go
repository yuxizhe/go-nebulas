// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Account
	ContractMeta
	Data
	Transaction
	BlockHeader
	Block
	NetBlocks
	NetBlock
	DownloadBlock
	Random
*/
package corepb

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	dagpb "github.com/nebulasio/go-nebulas/common/dag/pb"

	consensuspb "github.com/nebulasio/go-nebulas/consensus/pb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address      []byte        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance      []byte        `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce        uint64        `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	VarsHash     []byte        `protobuf:"bytes,4,opt,name=vars_hash,json=varsHash,proto3" json:"vars_hash,omitempty"`
	BirthPlace   []byte        `protobuf:"bytes,5,opt,name=birth_place,json=birthPlace,proto3" json:"birth_place,omitempty"`
	ContractMeta *ContractMeta `protobuf:"bytes,6,opt,name=contract_meta,json=contractMeta" json:"contract_meta,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Account) GetVarsHash() []byte {
	if m != nil {
		return m.VarsHash
	}
	return nil
}

func (m *Account) GetBirthPlace() []byte {
	if m != nil {
		return m.BirthPlace
	}
	return nil
}

func (m *Account) GetContractMeta() *ContractMeta {
	if m != nil {
		return m.ContractMeta
	}
	return nil
}

type ContractMeta struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *ContractMeta) Reset()                    { *m = ContractMeta{} }
func (m *ContractMeta) String() string            { return proto.CompactTextString(m) }
func (*ContractMeta) ProtoMessage()               {}
func (*ContractMeta) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *ContractMeta) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Data struct {
	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Data) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      *Data  `protobuf:"bytes,7,opt,name=data" json:"data,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	GasPrice  []byte `protobuf:"bytes,9,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit  []byte `protobuf:"bytes,10,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Alg       uint32 `protobuf:"varint,11,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,12,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *Transaction) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockHeader struct {
	Hash          []byte                     `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash    []byte                     `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Coinbase      []byte                     `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp     int64                      `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId       uint32                     `protobuf:"varint,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Alg           uint32                     `protobuf:"varint,7,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign          []byte                     `protobuf:"bytes,8,opt,name=sign,proto3" json:"sign,omitempty"`
	StateRoot     []byte                     `protobuf:"bytes,9,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TxsRoot       []byte                     `protobuf:"bytes,10,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
	EventsRoot    []byte                     `protobuf:"bytes,11,opt,name=events_root,json=eventsRoot,proto3" json:"events_root,omitempty"`
	ConsensusRoot *consensuspb.ConsensusRoot `protobuf:"bytes,12,opt,name=consensus_root,json=consensusRoot" json:"consensus_root,omitempty"`
	Random        *Random                    `protobuf:"bytes,13,opt,name=random" json:"random,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTxsRoot() []byte {
	if m != nil {
		return m.TxsRoot
	}
	return nil
}

func (m *BlockHeader) GetEventsRoot() []byte {
	if m != nil {
		return m.EventsRoot
	}
	return nil
}

func (m *BlockHeader) GetConsensusRoot() *consensuspb.ConsensusRoot {
	if m != nil {
		return m.ConsensusRoot
	}
	return nil
}

func (m *BlockHeader) GetRandom() *Random {
	if m != nil {
		return m.Random
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Dependency   *dagpb.Dag     `protobuf:"bytes,3,opt,name=dependency" json:"dependency,omitempty"`
	Height       uint64         `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{5} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetDependency() *dagpb.Dag {
	if m != nil {
		return m.Dependency
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type NetBlocks struct {
	From   string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Batch  uint64   `protobuf:"varint,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Blocks []*Block `protobuf:"bytes,3,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *NetBlocks) Reset()                    { *m = NetBlocks{} }
func (m *NetBlocks) String() string            { return proto.CompactTextString(m) }
func (*NetBlocks) ProtoMessage()               {}
func (*NetBlocks) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{6} }

func (m *NetBlocks) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlocks) GetBatch() uint64 {
	if m != nil {
		return m.Batch
	}
	return 0
}

func (m *NetBlocks) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type NetBlock struct {
	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Batch uint64 `protobuf:"varint,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Block *Block `protobuf:"bytes,3,opt,name=block" json:"block,omitempty"`
}

func (m *NetBlock) Reset()                    { *m = NetBlock{} }
func (m *NetBlock) String() string            { return proto.CompactTextString(m) }
func (*NetBlock) ProtoMessage()               {}
func (*NetBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{7} }

func (m *NetBlock) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlock) GetBatch() uint64 {
	if m != nil {
		return m.Batch
	}
	return 0
}

func (m *NetBlock) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type DownloadBlock struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Sign []byte `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *DownloadBlock) Reset()                    { *m = DownloadBlock{} }
func (m *DownloadBlock) String() string            { return proto.CompactTextString(m) }
func (*DownloadBlock) ProtoMessage()               {}
func (*DownloadBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{8} }

func (m *DownloadBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DownloadBlock) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type Random struct {
	VrfSeed  []byte `protobuf:"bytes,1,opt,name=vrf_seed,json=vrfSeed,proto3" json:"vrf_seed,omitempty"`
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
}

func (m *Random) Reset()                    { *m = Random{} }
func (m *Random) String() string            { return proto.CompactTextString(m) }
func (*Random) ProtoMessage()               {}
func (*Random) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{9} }

func (m *Random) GetVrfSeed() []byte {
	if m != nil {
		return m.VrfSeed
	}
	return nil
}

func (m *Random) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*ContractMeta)(nil), "corepb.ContractMeta")
	proto.RegisterType((*Data)(nil), "corepb.Data")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*NetBlocks)(nil), "corepb.NetBlocks")
	proto.RegisterType((*NetBlock)(nil), "corepb.NetBlock")
	proto.RegisterType((*DownloadBlock)(nil), "corepb.DownloadBlock")
	proto.RegisterType((*Random)(nil), "corepb.Random")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x8a, 0x2b, 0x45,
	0x10, 0x26, 0xc9, 0xe4, 0xaf, 0x26, 0x59, 0x0e, 0xed, 0x22, 0xe3, 0xaa, 0x6c, 0x18, 0x51, 0x82,
	0x62, 0x02, 0xab, 0xb0, 0x7a, 0xe7, 0xd1, 0x73, 0x71, 0x14, 0x95, 0xa5, 0xf5, 0x46, 0x10, 0x42,
	0x4d, 0x4f, 0xef, 0x64, 0x30, 0xd3, 0x3d, 0x74, 0x77, 0xe2, 0xd9, 0x47, 0xf0, 0x5d, 0xbc, 0xf1,
	0x3d, 0x7c, 0x28, 0xe9, 0xea, 0x9e, 0xec, 0xec, 0xf1, 0x80, 0x78, 0x35, 0xfd, 0xd5, 0x37, 0x5f,
	0x53, 0xf5, 0x55, 0x55, 0x43, 0x5a, 0x1c, 0xb4, 0xf8, 0x6d, 0xd3, 0x1a, 0xed, 0x34, 0x9b, 0x08,
	0x6d, 0x64, 0x5b, 0x5c, 0xdd, 0x56, 0xb5, 0xdb, 0x1f, 0x8b, 0x8d, 0xd0, 0xcd, 0x56, 0xc9, 0xe2,
	0x78, 0x40, 0x5b, 0xeb, 0x6d, 0xa5, 0x3f, 0x8d, 0x60, 0x2b, 0x74, 0xd3, 0x68, 0xb5, 0x2d, 0xb1,
	0xda, 0xb6, 0x85, 0xff, 0x84, 0x0b, 0xae, 0xbe, 0xf8, 0x6f, 0xa1, 0xb2, 0x52, 0xd9, 0xa3, 0xf5,
	0x3a, 0xeb, 0xd0, 0xc9, 0xa0, 0xcc, 0xff, 0x1e, 0xc0, 0xf4, 0xb9, 0x10, 0xfa, 0xa8, 0x1c, 0xcb,
	0x60, 0x8a, 0x65, 0x69, 0xa4, 0xb5, 0xd9, 0x60, 0x35, 0x58, 0x2f, 0x78, 0x07, 0x3d, 0x53, 0xe0,
	0x01, 0x95, 0x90, 0xd9, 0x30, 0x30, 0x11, 0xb2, 0x4b, 0x18, 0x2b, 0xed, 0xe3, 0xa3, 0xd5, 0x60,
	0x9d, 0xf0, 0x00, 0xd8, 0xbb, 0x30, 0x3f, 0xa1, 0xb1, 0xbb, 0x3d, 0xda, 0x7d, 0x96, 0x90, 0x62,
	0xe6, 0x03, 0x2f, 0xd1, 0xee, 0xd9, 0x35, 0xa4, 0x45, 0x6d, 0xdc, 0x7e, 0xd7, 0x1e, 0x50, 0xc8,
	0x6c, 0x4c, 0x34, 0x50, 0xe8, 0xce, 0x47, 0xd8, 0x97, 0xb0, 0x14, 0x5a, 0x39, 0x83, 0xc2, 0xed,
	0x1a, 0xe9, 0x30, 0x9b, 0xac, 0x06, 0xeb, 0xf4, 0xe6, 0x72, 0x13, 0x6c, 0xda, 0x7c, 0x13, 0xc9,
	0x1f, 0xa4, 0x43, 0xbe, 0x10, 0x3d, 0x94, 0xaf, 0x61, 0xd1, 0x67, 0x7d, 0xe2, 0x27, 0x69, 0x6c,
	0xad, 0x15, 0x95, 0x34, 0xe7, 0x1d, 0xcc, 0x3f, 0x87, 0xe4, 0x05, 0x3a, 0x64, 0x0c, 0x12, 0xf7,
	0xd0, 0xca, 0x48, 0xd3, 0xd9, 0xab, 0x5a, 0x7c, 0x38, 0x68, 0x2c, 0xbb, 0x72, 0x23, 0xcc, 0xff,
	0x1c, 0x42, 0xfa, 0xb3, 0x41, 0x65, 0x51, 0xb8, 0x5a, 0x2b, 0xaf, 0xa6, 0x1a, 0x83, 0x5f, 0x74,
	0xf6, 0xb1, 0x7b, 0xa3, 0x9b, 0x28, 0xa5, 0x33, 0xbb, 0x80, 0xa1, 0xd3, 0xe4, 0xd1, 0x82, 0x0f,
	0x9d, 0xf6, 0xb6, 0x9d, 0xf0, 0x70, 0x94, 0xd1, 0x9c, 0x00, 0x1e, 0xcd, 0x1c, 0xf7, 0xcd, 0x7c,
	0x0f, 0xe6, 0xae, 0x6e, 0xa4, 0x75, 0xd8, 0xb4, 0x64, 0xc5, 0x88, 0x3f, 0x06, 0xd8, 0x0a, 0x92,
	0x12, 0x1d, 0x66, 0x53, 0xf2, 0x68, 0xd1, 0x79, 0xe4, 0x6b, 0xe3, 0xc4, 0xb0, 0x77, 0x60, 0x26,
	0xf6, 0x58, 0xab, 0x5d, 0x5d, 0x66, 0xb3, 0xd5, 0x60, 0xbd, 0xe4, 0x53, 0xc2, 0xdf, 0x96, 0xbe,
	0x4f, 0x15, 0xda, 0x5d, 0x6b, 0x6a, 0x21, 0xb3, 0x79, 0xe8, 0x53, 0x85, 0xf6, 0xce, 0xe3, 0x8e,
	0x3c, 0xd4, 0x4d, 0xed, 0x32, 0x38, 0x93, 0xdf, 0x7b, 0xcc, 0x9e, 0xc1, 0x08, 0x0f, 0x55, 0x96,
	0xd2, 0x7d, 0xfe, 0xe8, 0xcb, 0xb6, 0x75, 0xa5, 0xb2, 0x45, 0x28, 0xdb, 0x9f, 0xf3, 0x3f, 0x46,
	0x90, 0x7e, 0xed, 0x07, 0xfd, 0xa5, 0xc4, 0x52, 0x9a, 0x37, 0xda, 0x75, 0x0d, 0x69, 0x8b, 0x46,
	0x2a, 0x17, 0xa6, 0x25, 0xb8, 0x06, 0x21, 0x44, 0xf3, 0x72, 0x05, 0x33, 0xa1, 0x6b, 0x55, 0xa0,
	0xed, 0xec, 0x3a, 0xe3, 0xa7, 0xde, 0x8c, 0x5f, 0xf7, 0xa6, 0x5f, 0xf9, 0xe4, 0x69, 0xe5, 0x31,
	0xff, 0xe9, 0xbf, 0xf3, 0x9f, 0x3d, 0xe6, 0xcf, 0xde, 0x07, 0xa0, 0x65, 0xd9, 0x19, 0xad, 0x5d,
	0x34, 0x68, 0x4e, 0x11, 0xae, 0xb5, 0xf3, 0xf7, 0xbb, 0x57, 0x36, 0x90, 0xc1, 0xa0, 0xa9, 0x7b,
	0x65, 0x89, 0xba, 0x86, 0x54, 0x9e, 0xa4, 0x72, 0x91, 0x4d, 0x43, 0x55, 0x21, 0x44, 0x3f, 0x3c,
	0x87, 0x8b, 0xf3, 0x52, 0x86, 0x7f, 0x16, 0xd4, 0xc1, 0xab, 0xcd, 0x39, 0x1c, 0x46, 0x3d, 0x9c,
	0xbd, 0x86, 0x2f, 0x45, 0x1f, 0xb2, 0x8f, 0x60, 0x62, 0x50, 0x95, 0xba, 0xc9, 0x96, 0x24, 0xbd,
	0xe8, 0x9a, 0xcf, 0x29, 0xca, 0x23, 0xfb, 0x5d, 0x32, 0x1b, 0x3d, 0x4b, 0xf2, 0xbf, 0x06, 0x30,
	0xa6, 0x5e, 0xb0, 0x4f, 0x60, 0xb2, 0xa7, 0x7e, 0x50, 0x1f, 0xd2, 0x9b, 0xb7, 0x3a, 0x5d, 0xaf,
	0x55, 0x3c, 0xfe, 0xc2, 0x6e, 0x61, 0xe1, 0x1e, 0x07, 0xde, 0x66, 0xc3, 0xd5, 0xa8, 0x2f, 0xe9,
	0x2d, 0x03, 0x7f, 0xf2, 0x23, 0xfb, 0x18, 0xa0, 0x94, 0xad, 0x54, 0xa5, 0x54, 0xe2, 0x81, 0x46,
	0x3f, 0xbd, 0x81, 0x4d, 0x89, 0x15, 0x4d, 0x67, 0xc5, 0x7b, 0x2c, 0x7b, 0xdb, 0x67, 0x54, 0x57,
	0x7b, 0x47, 0x0d, 0x4e, 0x78, 0x44, 0xf9, 0xaf, 0x30, 0xff, 0x51, 0x3a, 0x4a, 0xcb, 0x9e, 0xf7,
	0x2a, 0x6e, 0x2a, 0xed, 0xd5, 0x25, 0x8c, 0x0b, 0x74, 0x22, 0x8c, 0x4d, 0xc2, 0x03, 0x60, 0x1f,
	0xc2, 0x84, 0x9e, 0x57, 0x9b, 0x8d, 0x28, 0xdb, 0xe5, 0x93, 0x02, 0x79, 0x24, 0xf3, 0x5f, 0x60,
	0xd6, 0xdd, 0xfe, 0x3f, 0x2e, 0xff, 0x00, 0xc6, 0xa4, 0x8f, 0x25, 0xbd, 0x76, 0x77, 0xe0, 0xf2,
	0x5b, 0x58, 0xbe, 0xd0, 0xbf, 0x2b, 0xff, 0x66, 0x9c, 0xef, 0x7f, 0xd3, 0x43, 0x41, 0x13, 0x37,
	0xec, 0x6d, 0xcc, 0x57, 0x30, 0x09, 0xdd, 0xf3, 0xc3, 0x75, 0x32, 0xf7, 0x3b, 0x2b, 0x65, 0xd9,
	0x3d, 0xc7, 0x27, 0x73, 0xff, 0x93, 0x94, 0xb4, 0xb6, 0x9e, 0x6a, 0x8d, 0xd6, 0xf7, 0x51, 0xed,
	0xff, 0xbd, 0xf3, 0xb8, 0x98, 0xd0, 0xc3, 0xfe, 0xd9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xc7, 0x49, 0xf3, 0x62, 0x06, 0x00, 0x00,
}
