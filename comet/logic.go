package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"
	"im/libs/proto"
)

var (
	logicRpcClient client.XClient
)

func InitLogicRpcClient() (err error) {
	d := client.NewZookeeperDiscovery(Conf.ZookeeperInfo.BasePath,
		Conf.ZookeeperInfo.ServerPathLogic,
		[]string{Conf.ZookeeperInfo.Host},
		nil)
	logicRpcClient = client.NewXClient(Conf.ZookeeperInfo.ServerPathLogic, client.Failtry,
		client.RandomSelect, d, client.DefaultOption)
	return
}
func connect(arg *proto.ConnArg) (uid string, err error) {
	reply := &proto.ConnReply{}
	err = logicRpcClient.Call(context.Background(), "Connect", arg, reply)
	if err != nil {
		log.Fatal("failed to call: %v", err)
	}
	uid = reply.Uid
	log.Infof("comet logic uid: %s", reply.Uid)

	return
}

func disconnect(arg *proto.DisconnArg) (err error) {
	reply := &proto.DisconnReply{}
	if err = logicRpcClient.Call(context.Background(), "Disconnect", arg, reply); err != nil {
		log.Fatal("failed to call: %v", err)
	}
	return
}
