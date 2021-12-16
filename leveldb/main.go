package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"strings"
)

func NewDbInstance(DbPath string) *leveldb.DB {
	db, err := leveldb.OpenFile(DbPath, nil)
	if err != nil {
		log.Panic(err)
	}
	return db
}

func NewLevelDbInstance(DbPath string) *LevelDb {
	levelDb := LevelDb{DbPath: DbPath}
	levelDb.Instance()
	return &levelDb
}

type LevelDb struct {
	DbPath  string
	Handler *leveldb.DB
}

func (t *LevelDb) Instance() {
	t.Handler = NewDbInstance(t.DbPath) // 创建LevelDB数据库实例
}

// Put Put
func (t *LevelDb) Put(key string, value string) error {
	return t.Handler.Put([]byte(key), []byte(value), nil)
}

// GetOne 获取单条数据
func (t *LevelDb) GetOne(key string) ([]byte, error) {
	return t.Handler.Get([]byte(key), nil)
}

// GetAll 获取全部数据
func (t *LevelDb) GetAll(callFunc func(key string, value string)) {
	iter := t.Handler.NewIterator(nil, nil)
	for iter.Next() {
		callFunc(string(iter.Key()), string(iter.Value()))
	}
}

func (t *LevelDb) FilterManyMapByPrefix(prefix string) ([]interface{}, error) {
	iter := t.Handler.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	ret := make([]interface{}, 0)
	for iter.Next() {
		key := string(iter.Key())
		orgKey := strings.TrimPrefix(key, prefix+"-")
		ret = append(ret, orgKey)
	}
	iter.Release()
	err := iter.Error()
	return ret, err
}

// Delete 删除指定的key
func (t *LevelDb) Delete(key string) error {
	return t.Handler.Delete([]byte(key), nil)
}

// Close 关闭连接
func (t *LevelDb) Close() {
	err := t.Handler.Close()
	if err != nil {
		panic(err)
	}
}
