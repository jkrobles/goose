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
    Goose.Write("PART "+c)
    var Chan int
    for i := 0; i <= len(Goose.Channels); i++ {
        if c == Goose.Channels[i] {
            Chan = i
        }
    }

    if (Chan+1) == len(Goose.Channels) {
        Goose.Channels = Goose.Channels[:Chan]
    } else {
        Goose.Channels = append(Goose.Channels[:Chan],Goose.Channels[Chan+1:]...)

    }
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
            fmt.Println("PONG sent\r")
            Goose.Write("PONG :" + l[1])
        }
}

func (Goose *Goose) loop_handler(line string) {
    // TODO: Check l[0] and if handler is set for that event
    // TODO: Call handler
}

func (Goose *Goose) Split(line string, c string) (seg []string) {
    var l []string = strings.Split(line,c)
    return l
}

