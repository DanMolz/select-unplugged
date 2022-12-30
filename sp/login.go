package sp

import (
	"crypto/md5"
	"fmt"
)

func CalculateLoginHash(password string, readHash Message) Message {
	// pad the password with spaces to 32 characters
	passwordMessage := Message(fmt.Sprintf("%-32s", password))

	// combine password with hash
	buffer := append(readHash, passwordMessage...)
	result := md5.Sum([]byte(buffer))

	// adjust to little endian pairs
	resultLe := []byte{}
	for i := 0; i < len(result); i = i + 2 {
		resultLe = append(resultLe, result[i+1], result[i])
	}

	return Message(resultLe[:])
}
