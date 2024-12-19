package obstacles

import (
	"GONKI/participants"
	"fmt"
)

type Wall struct {
	name     string
	Height int
}

func (w *Wall) GetName() string{
	 return w.name
}

func (w *Wall) Overcome(p participants.Participant) bool {
	result := p.Jump(w.Height)
    if result {
        fmt.Printf("Participant [%s] successfully overcame the wall obstacle [%d]\n", p.GetName(), w.Height)
    } else {
        fmt.Printf("Participant [%s] unsuccessfully overcame the wall obstacle [%d]\n", p.GetName(), w.Height)
    }
    return result
}
