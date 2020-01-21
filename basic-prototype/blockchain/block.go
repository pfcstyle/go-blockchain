package blockchain

import (
	"bytes"
	"strconv"
	"time"
)

//Block prototype
type Block struct {
	//1.区块高度
	height int64
	//2.上一个区块hash
	prevBlockHash []byte
	//3.交易数据
	data []byte
	//4.时间戳
	timestamp int64
	//5.Hash
	hash []byte
}

func (block *Block) SetHash() {
	//1. Height []byte
	heightBytes := IntToHex(block.height)
	//2. 将时间戳转[]byte
	timeString := strconv.FormatInt(block.timestamp, 2)
	timeBytes := []byte(timeString)
	//3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, timeBytes}, []byte(""))
	//4.
}

// NewBlock 创建新的区块
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	return &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
}
