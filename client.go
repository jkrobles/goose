package goose

import ("net"
        "log"
        "bufio"
        "fmt"
        "net/textproto"
        "strings"
      )


/* Commands */

func (Goose *Goose) Join(c string) (ret bool) {
    // TODO: join channel if not joined
    Goose.Write("JOIN "+c)
    return true
}

func (Goose *Goose) Part(c string) (ret bool) {
    // TODO: part channel
    Goose.Write("PART "+c)
    // TODO: remove from channels list
    return true
}

func (Goose *Goose) Names(c string) {
}

func (Goose *Goose) Connect() (conn net.Conn, err error){
    conn, err = net.Dial("tcp",Goose.Server + ":" + Goose.Port)
if err != nil{
    log.Fatal("unable to connect to IRC server ", err)
  }
  Goose.conn = conn
  log.Printf("Connected to IRC server %s (%s)\n", Goose.Server, Goose.conn.RemoteAddr())

  Goose.Write("NICK " + Goose.Nick)
  Goose.Write("USER "+ Goose.Nick+" 8 * :" + Goose.Nick)

 // defer Goose.conn.Close()

  reader := bufio.NewReader(conn)
    var tp *textproto.Reader = textproto.NewReader( reader )
    Goose.Reader = tp


  fmt.Println(Goose.Channels[0])
  Goose.Join("#goose")

  return nil, nil
}

func (Goose *Goose) Write(s string) {
    Goose.conn.Write([]byte(s+"\r\n"))
}

func (Goose *Goose) Reading() (ret bool) {
    return true
}

func (Goose *Goose) Pings(line string) {
    l := strings.Split(line,":")
        if l[0] == "PING " {
            fmt.Println("PONG sent\r\n")
            Goose.Write("PONG :" + l[1])
        }
}
