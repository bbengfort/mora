// Package echo provides the protocol buffer communications functionality
// needed by Mora but compiled as seperat types to ensure there are no
// conflicts with the defautl Mora types. This package also provides helper
// functionality for augmenting and parsing messages.
package echo

import "time"

// GetSentTime parses the sent time on an Echo message to a time.Time
func (m *Echo) GetSentTime() time.Time {
	nsecs := m.GetSent()
	return time.Unix(0, nsecs)
}

// GetReceivedTime parses the received time on an EchoReply message to a time.Time
func (m *EchoReply) GetReceivedTime() time.Time {
	nsecs := m.GetReceived()
	return time.Unix(0, nsecs)
}
