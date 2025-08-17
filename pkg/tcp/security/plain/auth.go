package plain

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"strings"
)

var users = map[string]struct {
	Password string
}{
	"admin": {Password: "password"},
}

func Plain(conn net.Conn, r *bufio.Reader) (err error) {
	line, err := r.ReadString('\n')
	if err != nil {
		io.WriteString(conn, "ERROR failed to read line\n")
		return err
	}

	data, err := base64.StdEncoding.DecodeString(strings.TrimSpace(line))
	if err != nil {
		io.WriteString(conn, "ERROR failed to decode base64\n")
		return err
	}

	parts := strings.Split(string(data), "\x00")
	if len(parts) != 3 {
		io.WriteString(conn, "ERROR bad PLAIN format\n")
		return
	}

	username, password := parts[1], parts[2]
	user, ok := users[username]
	if !ok || user.Password != password {
		io.WriteString(conn, "ERROR invalid credentials\n")
		return fmt.Errorf("invalid credentials for user %s", username)
	}

	io.WriteString(conn, "OK\n")
	fmt.Printf("✅ [PLAIN] %s authenticated\n", username)

	return nil
}
