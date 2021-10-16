package core

import (
	"fmt"
	"github.com/ibm-messaging/mq-golang-jms20/jms20subset"
	"log"
	"sync"
	"time"
)

func Send(jmsContext *jms20subset.JMSContext, wg *sync.WaitGroup, queue *jms20subset.Queue, mutex *sync.Mutex, payload string) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	jmsProducer := (*jmsContext).CreateProducer()
	fmt.Println("Sending a message: " + payload + ". Time " + time.Now().String())
	exc := jmsProducer.SendString(*queue, payload)
	if exc != nil {
		log.Fatal(exc)
	}
}

func Receive(jmsContext *jms20subset.JMSContext, wg *sync.WaitGroup, queue *jms20subset.Queue, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	jmsConsumer, exc := (*jmsContext).CreateConsumer(*queue)
	if exc != nil {
		log.Fatal(exc)
	}
	defer jmsConsumer.Close()
	response, exc := jmsConsumer.ReceiveStringBody(30000)
	fmt.Println("Receiving a message: " + *response + ". Time " + time.Now().String())
	if exc != nil {
		log.Fatal(exc)
	}
	if response == nil {
		fmt.Println("No message received")
	}
}
