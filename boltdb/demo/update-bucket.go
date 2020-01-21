package demo

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func UpdateBucket() {
	db, err := OpenDB()
	if err != nil {
		CloseDB(db)
		return
	}
	// 创建表
	err = db.Update(func(tx *bolt.Tx) error {

		// 创建BlockBucket表
		b := tx.Bucket([]byte("BlockBucket"))
		if b == nil {
			b, err = tx.CreateBucket([]byte("BlockBucket"))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}

		// 往表里面存储数据
		if b != nil {
			err := b.Put([]byte("l"), []byte("Send 100 BTC To 强哥......"))
			if err != nil {
				log.Panic("数据存储失败......")
			}
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
	CloseDB(db)
}
