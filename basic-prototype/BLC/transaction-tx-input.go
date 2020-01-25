package BLC

import (
	"bytes"
)

type TXInput struct {
	// 1. 交易的Hash  之前transaction的hash,用于指向input的来源（output）
	TxHash []byte
	// 2. 存储TXOutput在Vout里面的索引 一个transaction可能有多个output（找零也需要记录，重新生成新的输出）
	Vout      int
	Signature []byte // 数字签名

	PublicKey []byte // 公钥，钱包里面  钱的来源者
}

// 判断当前的消费是谁的钱
func (txInput *TXInput) UnLockRipemd160Hash(ripemd160Hash []byte) bool {

	publicKey := Ripemd160Hash(txInput.PublicKey)

	return bytes.Compare(publicKey, ripemd160Hash) == 0
}
