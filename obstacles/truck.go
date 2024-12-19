package obstacles

import (
	"GONKI/participants"
	"fmt"
)

type Truck struct {
	name     string
	Distance int
}

func (t *Truck) GetName() string{
	 return t.name
}

func (t *Truck) Overcome(p participants.Participant) bool {
	result := p.Run(t.Distance)
    if result {
        fmt.Printf("Participant [%s] successfully overcame the treadmill obstacle [%d]\n", p.GetName(), t.Distance)
    } else {
        fmt.Printf("Participant [%s] unsuccessfully overcame the treadmill obstacle [%d]\n", p.GetName(), t.Distance)
    }
    return result
}
