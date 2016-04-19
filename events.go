package goose

/* Events */
type Event struct {
    name string
    id int
}

/* Commands */

func (Goose *Goose) SetHandler(e Event, handler func(line string)) {

}
