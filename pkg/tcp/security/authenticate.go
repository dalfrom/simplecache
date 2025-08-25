package security

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"strings"

	"github.com/dalfrom/simplecache/pkg/tcp/security/plain"
)

func Authenticate(conn net.Conn) error {
	// Read the first frame to determine the authentication type
	r := bufio.NewReader(conn)
	io.WriteString(conn, "SASL: PLAIN SCRAM-SHA-256 OAUTH2.0\n")

	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)

	if !strings.HasPrefix(line, "AUTH ") {
		io.WriteString(conn, "ERROR expected AUTH\n")
		return errors.New("expected AUTH command")
	}

	mech := strings.TrimPrefix(line, "AUTH ")
	switch mech {
	case "PLAIN":
		log.Println("Handling PLAIN authentication")
		plain.Plain(conn, r)
	case "SCRAM-SHA-256":
		log.Println("Handling SCRAM-SHA-256 authentication")
		// TODO: implement
	default:
		return errors.New("unsupported mechanism: " + mech)
	}

	return nil
}
