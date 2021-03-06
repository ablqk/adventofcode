package company

import (
	"testing"

	"github.com/ablqk/adventofcode/2020/dec13/bus"
	"github.com/stretchr/testify/assert"
)

func TestNewCompany(t *testing.T) {
	c := NewCompany("7,13,x,x,59,x,31,19")
	assert.Equal(t, []bus.Bus{7, 13, 59, 31, 19}, c.(*busCompany).buses)
}
