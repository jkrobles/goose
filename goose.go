package goose

import (
    "net"
    "net/textproto"
)

type Goose struct{
        Server string
        Port string
        Nick string
        User string
        Channels []string
        Password string
        pread, pwrite chan string
        conn net.Conn
        Reader *textproto.Reader
}
