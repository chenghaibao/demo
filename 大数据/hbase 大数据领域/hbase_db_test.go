package data_conversion

import (
	"fmt"
	"github.com/sdming/goh"
	"github.com/sdming/goh/Hbase"
	"testing"
)

func TestHbaseDb(t *testing.T) {
	address := "localhost:9098"

	client, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Open(); err != nil {
		fmt.Println(err)
		return
	}
	table := "test"
	defer client.Close()

	fmt.Println(client.IsTableEnabled(table))
	fmt.Println(client.DisableTable(table))
	fmt.Println(client.EnableTable(table))
	descriptors, err := client.GetColumnDescriptors(table)
	if err != nil {
		return
	}
	fmt.Println(descriptors)

	fmt.Println(client.Compact(table))

	rowBatches := goh.NewBatchMutation([]byte("2"), []*Hbase.Mutation{goh.NewMutation("order:1", []byte("a"))})
	err = client.MutateRows(table, []*Hbase.BatchMutation{rowBatches}, map[string]string{})
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkPutHbasePut(b *testing.B) {
	address := "localhost:9098"

	client, err := goh.NewTcpClient(address, goh.TBinaryProtocol, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Open(); err != nil {
		fmt.Println(err)
		return
	}
	table := "test"
	defer client.Close()
	//
	//fmt.Println(client.IsTableEnabled(table))
	//fmt.Println(client.DisableTable(table))
	//fmt.Println(client.EnableTable(table))
	//descriptors, err := client.GetColumnDescriptors(table)
	//if err != nil {
	//	return
	//}
	//fmt.Println(descriptors)
	//
	//fmt.Println(client.Compact(table))

	for i := 0; i < b.N; i++ {
		var batchMutation []*Hbase.Mutation
		for j := 0; j < 10; j++ {
			batchMutation = append(batchMutation, goh.NewMutation(fmt.Sprintf("order:%d%d", i, j), []byte(fmt.Sprintf("%d%da", i, j))))
		}
		rowBatches := goh.NewBatchMutation([]byte(fmt.Sprintf("%d", i)), batchMutation)
		err = client.MutateRows(table, []*Hbase.BatchMutation{rowBatches}, map[string]string{})
	}

}
