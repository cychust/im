package main

import "im/libs/proto"

type Operator interface {
	Connect(*proto.ConnArg) (string, error)
	Disconnect(*proto.DisconnArg) error
}
type DefaultOperator struct {
}

func (operator *DefaultOperator) Connect(arg *proto.ConnArg) (uid string, err error) {
	uid, err = connect(arg)
	return
}

func (operator *DefaultOperator) Disconnect(arg *proto.DisconnArg) (err error) {
	if err = disconnect(arg); err != nil {
		return
	}
	return
}
