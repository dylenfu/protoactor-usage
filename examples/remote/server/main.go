package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/dylenfu/miner-protoactor/messages"
	"runtime"
)

const serverAddress = "127.0.0.1:9090"

func main() {
	//flag.Parse()
	//handle := &Handle{}
	//reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})

	runtime.GOMAXPROCS(runtime.NumCPU())
	remote.Register("hello", actor.FromProducer(newHelloActor))
	remote.Start(serverAddress)
	console.ReadLine()
}

type helloActor struct{}

func (*helloActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *messages.HelloRequest:
		ctx.Respond(&messages.HelloResponse{
			Message: "Hello from " + msg.Who,
		})
	}
}

func newHelloActor() actor.Actor {
	return &helloActor{}
}
