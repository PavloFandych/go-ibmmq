package main

import (
	"fmt"
	"github.com/ibm-messaging/mq-golang-jms20/mqjms"
	"go-ibmmq/core"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex = sync.Mutex{}

	cf, err := mqjms.CreateConnectionFactoryFromJSON("/home/total/go/src/go-ibmmq/config-samples/connection_info.json",
		"/home/total/go/src/go-ibmmq/config-samples/applicationApiKey.json")
	if err != nil {
		log.Fatal(err)
	}
	jmsContext, jmsException := cf.CreateContext()
	if jmsException != nil {
		log.Fatal(jmsException)
	}
	defer func() {
		jmsContext.Close()
	}()

	fmt.Println("-- Connection has been established --")

	queue := jmsContext.CreateQueue("DEV.QUEUE.1")

	for i := 0; i < 3; i++ {
		wg.Add(2)
		go core.Send(&jmsContext, &wg, &queue, &mutex, "**TESTING MESSAGE**")
		go core.Receive(&jmsContext, &wg, &queue, &mutex)
	}

	wg.Wait()
	fmt.Println("-- Done --")
}
