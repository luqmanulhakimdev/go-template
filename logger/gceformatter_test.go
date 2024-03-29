package logger

import (
	"errors"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	GceFormatter GCEFormatter = GCEFormatter{}
)

func TestFormat(t *testing.T) {
	res, err := GceFormatter.Format(&log.Entry{})
	assert.Equal(t, res, []uint8([]byte{0x7b, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x3a, 0x22, 0x22, 0x2c, 0x22, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x22, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x22, 0x2c, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x30, 0x30, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x7d, 0xa}))
	assert.NoError(t, err)

	res, err = GceFormatter.Format(&log.Entry{
		Data: map[string]interface{}{"field": make(chan int)},
	})
	assert.Equal(t, res, []byte(nil))
	assert.Error(t, err)

	res, err = GceFormatter.Format(&log.Entry{
		Data: map[string]interface{}{"error": errors.New("error_test")},
	})
	assert.Equal(t, res, []byte{0x7b, 0x22, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3a, 0x22, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x22, 0x6d, 0x73, 0x67, 0x22, 0x3a, 0x22, 0x22, 0x2c, 0x22, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x22, 0x3a, 0x22, 0x41, 0x4c, 0x45, 0x52, 0x54, 0x22, 0x2c, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x30, 0x30, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x54, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x5a, 0x22, 0x7d, 0xa})
	assert.NoError(t, err)
}
