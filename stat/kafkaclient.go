package main

import (
	"fmt"
	"math/rand"
	"github.com/Shopify/sarama"
	"log"
)

var(
	ProuceMsg chan string = make(chan string)
	EndChan    chan bool   = make(chan bool)
)      
      
func InitKafuka(kahost string) {
	go ProduceKafka(kahost, &ProuceMsg, &EndChan)
	//go ConsumeKafka(kahost, &EndChan)
}

func ProduceKafka(kahost string, produceMsg *chan string, endchan *chan bool) {
	fmt.Println("ProdceKafka") 
	config := sarama.NewConfig()
	// WaitForAll waits for alln-sync replicas to commit before responding.
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true	
	//config.Producer.Partitioner = sarama.NewManualPartitioner
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewSyncProducer([]string{kahost}, config)
	if err != nil {
		return
	}

	defer producer.Close()

		
	// 定义一个生产消息，包括Topic、消息内容、
	for{
		msg:= &sarama.ProducerMessage{}
		msg.Topic = "revolution"
		msg.Key = sarama.StringEncoder("iles")
		msg.Value = sarama.StringEncoder("hello world...")
		msg.Partition = int32(rand.Int()% 5)
		// 发送消息
		pid, offset, err := producer.SendMessage(msg)
		fmt.Println(pid, offset, err)

		msg2 := &sarama.ProducerMessage{}
		msg2.Topic = "revolution"
		msg2.Key = sarama.StringEncoder("onroe")
		msg2.Value = sarama.StringEncoder("hello world2...")
		msg2.Partition = 2 //int32(rand.It()%5)
		pid2, offset2, err := producer.SendMessage(msg2)
		fmt.Println(pid2,  offset2, err)
	} 
}


func ConsumeKafka(kahost string, endchan *chan bool) {
	fmt.Println("ConsumeKafka")
	//config := sarama.NewConfig() 
	
	//sarama.NewAscCosumer([]strig{kahost}, config) 
	consumer, err := sarama.NewConsumer([]string{kahost}, nil)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}
		
	defer consumer.Close()
		
	for {
		partitions, err := consumer.Partitions("revolution")
		if  err != nil {
			fmt.Println("geet partitions failed, err:", err)
			return
		}

		fmt.Println("partitions,", partitions)
/*
		for _, p := range partitions {
			partitionConsumer, err := consumer.ConsumePartition("revolution", p, sarama.OffsetOldest)
			if err != nil {
				fmt.Println("partitionConsumer err:", err)
				continue
			}
		
			for m := range partitionConsumer.Messages() {
				fmt.Print("p:%d key: %s, text: %s, offset: %d\n", p, string(m.Key), string(m.Value), m.Offset)
			}
		}
*/
	}
}
	

