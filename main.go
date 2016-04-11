package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang/protobuf/proto"

	ocpb "github.com/openconfig/reference/rpc/openconfig"
)

func main() {
	a := &ocpb.Notification{
		Timestamp: time.Now().UnixNano(),
		Prefix: &ocpb.Path{
			Element: []string{"foo", "bar"},
		},
	}
	b := &ocpb.Update{
		Path: &ocpb.Path{
			Element: []string{"value"},
		},
		Value: &ocpb.Value{
			Type:  ocpb.Type_JSON,
			Value: []byte("{\"a\":42}"),
		},
	}
	a.Update = append(a.Update, b)
	out, err := proto.Marshal(a)
	if err != nil {
		fmt.Printf("Failed to marshal %s: %v", a, err)
	}
	ioutil.WriteFile("/tmp/foo", out, 0644)
	c := &ocpb.SubscribeResponse{}
	c.Response = &ocpb.SubscribeResponse_Update{
		Update: &ocpb.Notification{},
	}
	switch c.GetResponse().(type) {
	default:
		fmt.Println("Unknown type")
	case nil:
		fmt.Println("Unset Response")
	case *ocpb.SubscribeResponse_Update:
		fmt.Println("Update")
	case *ocpb.SubscribeResponse_Heartbeat:
		fmt.Println("Heartbeat")
	case *ocpb.SubscribeResponse_SyncResponse:
		fmt.Println("Sync")
	}
}
