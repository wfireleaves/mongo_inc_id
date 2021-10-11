package main

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

const MaxIDInc = 10

type IDFactory struct {
	id     int64
	maxID  int64
	idLock sync.Mutex
}

func NewIDFactory() *IDFactory {
	instance := &IDFactory{}
	instance.gen()
	fmt.Println("start id factory", instance.id, instance.maxID)
	return instance
}

func (f *IDFactory) gen() {
	collection, err := GetMongoCollection()
	if err != nil {
		fmt.Println("get collection failed,err:", err.Error())
		return
	}
	filter := bson.M{"_id": "userid"}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	result := collection.FindOneAndUpdate(context.Background(), filter, update)
	if result.Err() != nil {
		fmt.Println("findAndModify failed, err:", result.Err().Error())
		return
	}
	var res bson.M
	err = result.Decode(&res)
	if err != nil {
		fmt.Println("decode failed,err:", err.Error())
		return
	}

	fmt.Println("res", res["seq"], fmt.Sprintf("%T", res["seq"]))

	f.id = int64(res["seq"].(float64)) * MaxIDInc
	f.maxID = f.id + MaxIDInc
}

func (f *IDFactory) NextID() (int64, error) {
	f.idLock.Lock()
	defer f.idLock.Unlock()

	if f.id == f.maxID {
		f.gen()
	}
	if f.id == f.maxID {
		return 0, fmt.Errorf("factory after gen not working")
	}
	id := f.id
	f.id++
	return id, nil
}
