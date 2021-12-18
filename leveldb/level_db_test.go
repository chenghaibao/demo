package leveldb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"sync"
	"testing"
)

const path = "./test"

func TestPut(t *testing.T) {
	db := NewLevelDbInstance(path)
	defer db.Close()
	err := db.Put("key", "dd")
	if err != nil {
		t.Error(err)
	}
}

func TestGetOne(t *testing.T) {
	db := NewLevelDbInstance(path)
	defer db.Close()
	data, err := db.GetOne("key")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "dd" {
		t.Error("获取结果不一致")
	}
}

func TestGetAll(t *testing.T) {
	db := NewLevelDbInstance(path)
	defer db.Close()
	db.GetAll(func(key string, value string) {
		fmt.Println(key, value)
	})
}

func TestDelete(t *testing.T) {
	db := NewLevelDbInstance(path)
	defer db.Close()
	err := db.Delete("key")
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkPut(b *testing.B) {
	worker := 1
	dbMap := make(map[int]*LevelDb)
	for i := 0; i < worker; i++ {
		dbMap[i] = NewLevelDbInstance(fmt.Sprintf("./test%d", i))
	}
	defer func() {
		for _, db := range dbMap {
			db.Close()
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		for j := 0; j < worker; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dbMap[j].Put("key", `data`)
			}()
		}
		wg.Wait()
	}
}

func BenchmarkWrite(b *testing.B) {
	db := NewLevelDbInstance(path)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		batch := new(leveldb.Batch)
		for j := 0; j < 18; j++ {
			batch.Put([]byte(fmt.Sprintf("key%d", j)), []byte(`data`))
		}
		_ = db.Handler.Write(batch, nil)
		// batch.Reset()
	}
}

func BenchmarkWriteWithManyDB(b *testing.B) {

	dbMap := make(map[int]*LevelDb, 0)
	for j := 0; j < 18; j++ {
		dbMap[j] = NewLevelDbInstance(fmt.Sprintf("./test-%d", j))
	}
	defer func() {
		for j := 0; j < 18; j++ {
			dbMap[j].Close()
		}
	}()

	for i := 0; i < b.N; i++ {
		batch := new(leveldb.Batch)
		for j := 0; j < 18; j++ {
			batch.Put([]byte(fmt.Sprintf("key%d", j)), []byte(`data`))
		}
		_ = dbMap[i%18].Handler.Write(batch, nil)
		// batch.Reset()
	}
}

func BenchmarkBloom(b *testing.B) {
	//o := &opt.Options{
	//	Filter: filter.NewBloomFilter(500),
	//}
	db, _ := leveldb.OpenFile(path, nil)
	defer db.Close()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 18; j++ {
			_ = db.Put([]byte("key"), []byte(`data`), nil)
		}
	}
}
