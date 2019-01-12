package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	Assert := assert.New(t)
	m := MicroTimer{}

	nano := m.Get()
	Assert.True("1546112714106137771" < nano)
	Assert.True(nano < m.Get())
}
