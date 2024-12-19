package obstacles

import "GONKI/participants"

type Obstacle interface{
    Overcome(p participants.Participant) bool
}