package block

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"time"
)

type BaseBlock struct {

	// Header
	hash      []byte		// 64
	pHash     []byte		// 64
	rand      uint32		// 32
	timestamp int64			// 64

	// MetaData
	bodySize uint32			// 32

	// Body
	iBody IBaseData			// ...
}

func NewBlock() BaseBlock {
	b := BaseBlock{
		hash:  make([]byte, 32),
		pHash: make([]byte, 32),
	}
	return b
}

// BaseBlock Method

func (block *BaseBlock) SetHeader(p_hash []byte) {
	if p_hash != nil {
		copy(block.pHash, p_hash)
	}
	block.timestamp = time.Now().Unix()
	block.rand = rand.Uint32()
}

func (block *BaseBlock) SetData(data IBaseData) {
	block.iBody = data
	block.bodySize = data.GetDataSize()
}

func (block *BaseBlock) HashBlcok() {

	hasher := sha256.New()
	hasher.Write(block.pHash)
	t := []byte{byte(block.timestamp / 8), byte(block.timestamp % 8)}
	hasher.Write(t)

	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(block.timestamp))
	hasher.Write(b)

	binary.LittleEndian.PutUint32(b, block.rand)
	hasher.Write(b)

	hasher.Write(block.iBody.GetData())

	copy(block.hash, hasher.Sum(nil))
}

func (block *BaseBlock) GetHash() []byte {
	return block.hash
}

func (block *BaseBlock) BuildNextBlock(data IBaseData) BaseBlock {
	b := NewBlock()
	b.SetHeader(block.hash)
	b.SetData(data)
	return b
}

type IBaseData interface {
	GetDataSize() 		uint32
	GetData() 			[]byte
	WriteData([]byte)
}

