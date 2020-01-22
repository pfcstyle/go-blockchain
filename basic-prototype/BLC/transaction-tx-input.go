package BLC

type TXInput struct {
	// 1. 交易的Hash  之前transaction的hash,用于指向input的来源（output）
	TxHash []byte
	// 2. 存储TXOutput在Vout里面的索引 一个transaction可能有多个output（找零也需要记录，重新生成新的输出）
	Vout int
	// 3. 用户名 钱的来源者
	ScriptSig string
}
