package BLC

type TXOutput struct {
	Value        int64  //给多少
	ScriptPubKey string //用户名  to  把钱给谁
}

// 解锁
func (txOutput *TXOutput) UnLockScriptPubKeyWithAddress(address string) bool {

	return txOutput.ScriptPubKey == address
}
