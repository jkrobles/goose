package goose

import (
    "net"
    "net/textproto"
)

/* Goose object */
type Goose struct{
        Server string
        Port string
        Nick string
        User string
        Channels []string
        Password string
        events chan string
        conn net.Conn
        Reader *textproto.Reader
}
