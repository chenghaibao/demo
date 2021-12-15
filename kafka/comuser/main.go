package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"sync"
)

func main() {
	consumer()
}

func consumer() {
	wg := sync.WaitGroup{}
	consumer, err := sarama.NewConsumer(strings.Split("localhost:9092", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumer:%s", err)
	}

	partitionList, err := consumer.Partitions("hb_test")
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("hb_test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}

		wg.Add(1)
		defer pc.AsyncClose()
		go func(sarama.PartitionConsumer) {

			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
