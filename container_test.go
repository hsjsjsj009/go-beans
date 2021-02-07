package beans

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()
	assert.Equal(t, container,&ProviderContainer{dependencyInitiator: map[reflect.Type]*bean{}})
}
