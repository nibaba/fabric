// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import smartbft "github.com/hyperledger/fabric/orderer/consensus/smartbft"

// ReceiverGetter is an autogenerated mock type for the ReceiverGetter type
type ReceiverGetter struct {
	mock.Mock
}

// ReceiverByChain provides a mock function with given fields: channelID
func (_m *ReceiverGetter) ReceiverByChain(channelID string) smartbft.MessageReceiver {
	ret := _m.Called(channelID)

	var r0 smartbft.MessageReceiver
	if rf, ok := ret.Get(0).(func(string) smartbft.MessageReceiver); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(smartbft.MessageReceiver)
		}
	}

	return r0
}