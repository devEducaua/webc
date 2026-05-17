package internal

import (
	"encoding/json"
	"fmt"
	"net"
)

type Response struct {
	Ok bool
	Error string
	Data json.RawMessage
}

func connect() (net.Conn, error) {
	addr := "localhost:5353";
	conn, err := net.Dial("tcp", addr);
	if err != nil {
		return conn, err;
	}

	return conn, nil;
}

func Request(req string) (Response, error) {
	var r Response;
	conn, err := connect();
	if err != nil {
		return r, err;
	}
	defer conn.Close();

	n, err := conn.Write([]byte(req + "\n"));
	if err != nil {
		return r, err;
	}
	if n != len(req) +1 {
		return r, fmt.Errorf("failed to write the request.");
	}

	dec := json.NewDecoder(conn);
	if err := dec.Decode(&r); err != nil {
		return r, err;
	}
	return r, nil;
}

