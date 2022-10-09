package sp

import (
	"encoding/hex"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Request Message

func (r Request) String() string {
	return fmt.Sprintf("Request(0x%s)", hex.EncodeToString(r))
}

func (r Request) ResponseLength() int {
	requestType := string(r[0])
	requestLength := len(r)
	dataLength := (int(r[1]) + 1) * 2
	crcLength := 2
	if requestType == "W" {
		return requestLength
	}
	if requestType != "Q" {
		logrus.Fatalf("Unexpected request type %s", requestType)
	}
	return requestLength + dataLength + crcLength
}
