package sp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestResponseLength(t *testing.T) {
	i := func(i int) *int { return &i }
	for _, tt := range []struct {
		request Request
		length  *int
		err     error
	}{
		{NewRequestQuery(Area{0xa000, 0}), i(12), nil},
		{NewRequestQuery(Area{0xa000, 0xff}), i(522), nil},
		{Request("Z"), nil, errors.New("Unknown request type 'Z'")},
		{
			NewRequestWrite(Memory{
				address: 0x01f,
				data:    []byte("0000000000000000"),
			}),
			i(26),
			nil,
		},
	} {
		length, err := tt.request.ResponseLength()
		assert.Equal(t, tt.length, length)
		assert.Equal(t, tt.err, err)
	}
}

func TestRequestType(t *testing.T) {
	query := Query
	write := Write
	for _, tt := range []struct {
		request     Request
		requestType *RequestType
		err         error
	}{
		{Request("Q"), &query, nil},
		{Request("W"), &write, nil},
		{Request("Z"), nil, errors.New("Unknown request type 'Z'")},
	} {
		requestType, err := tt.request.Type()
		assert.Equal(t, tt.requestType, requestType)
		assert.Equal(t, tt.err, err)
	}
}

func TestNewRequestQuery(t *testing.T) {
	for _, tt := range []struct {
		area    Area
		request Request
	}{
		{Area{0x0000a000, 0}, Request("Q\x00\x00\xa0\x00\x00\x9d\x4b")},
		{Area{0x0000a093, 3}, Request("Q\x03\x93\xa0\x00\x00\x53\x9d")},
	} {
		request := NewRequestQuery(tt.area)
		assert.Equal(t, tt.request, request)
	}
}

func TestNewRequestWrite(t *testing.T) {
	expected := Request("W\x07\x00\x00\x1f\x00\x35\x7a\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35w\xaa")
	actual := NewRequestWrite(Memory{
		address: 0x001f0000,
		data:    Data("\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35"),
	})
	assert.Equal(t, expected, actual)
}
