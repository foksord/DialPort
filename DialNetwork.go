package dialnetwork

import (
	"net"
	"strconv"
	"strings"
)

// TryDialSocket
func TryDialSocket(port int) (bool, bool) {

	connectionTypes := [2]string{"UDP", "TCP"}

	var address string = "localhost:"

	var using bool
	var access bool

	for _, s := range connectionTypes {

		var stringBuilder strings.Builder
		stringBuilder.WriteString(address)
		stringBuilder.WriteString(strconv.FormatInt(int64(port), 10))

		u, a, dialConnection := DialSocket(s, stringBuilder.String())
		if dialConnection != nil {

			using = u
			access = a
			dialConnection.Close()

		}

		stringBuilder.Reset()

	}

	return using, access

}

// TryDialAllSocket
func TryDialAllSocket() map[int]bool {

	connectionTypes := [2]string{"UDP", "TCP"}

	var address string = "localhost:"
	var checkPortBeginNumber int = 1
	var checkPortEndNumber int = 65535

	resultList := make(map[int]bool)

	for _, s := range connectionTypes {

		for i := checkPortBeginNumber; i <= checkPortEndNumber; i++ {

			var stringBuilder strings.Builder
			stringBuilder.WriteString(address)
			stringBuilder.WriteString(strconv.FormatInt(int64(i), 10))

			_, access, dialConnection := DialSocket(s, stringBuilder.String())
			if dialConnection == nil {

				resultList[i] = access

			}

			stringBuilder.Reset()

		}

	}

	return resultList

}

// DialSocket
func DialSocket(ConnectionType string, ConnectionAddress string) (using bool, access bool, dialConnection net.Conn) {

	if ConnectionAddress != "" && ConnectionType != "" {

		var portUsing bool
		var portAccess bool

		defer func() {
			if r := recover(); r != nil {
				portUsing = true
			}
		}()

		newSocket, err := net.Dial(ConnectionType, ConnectionAddress)
		if err == nil {

			portAccess = true
			newSocket.Close()

		}

		portAccess = false

		return portUsing, portAccess, newSocket

	}

	return false, false, nil
}
