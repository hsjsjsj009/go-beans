package beans

import (
	"reflect"
)

type ProviderContainer struct {
	dependencyInitiator map[reflect.Type]*bean
}

func NewContainer() *ProviderContainer {
	return &ProviderContainer{dependencyInitiator: map[reflect.Type]*bean{}}
}