package adapter

import (
	"chuandao-sails-core/apps/platform-gateway/pkg/constants"
	"chuandao-sails-core/apps/platform-gateway/pkg/platform/events"
)

// AdapterInterface 适配器接口
type AdapterInterface interface {
	TransformToStandardOrder(interface{}) (*events.StandardOrderCreateEvent, error)
}

// AdapterFactory 适配器工厂
type AdapterFactory struct{}

func (f *AdapterFactory) GetAdapter(platCode string) AdapterInterface {
	switch platCode {
	case constants.PLATFORM_UU:
		return &UUAdapter{}
	default:
		return nil
	}
}
