package main

import (
    "GONKI/obstacles"
    "GONKI/participants"
	"fmt"
)

func main(){
	participantsList := []participants.Participant{
        &participants.Human{Name: "Valerii", MaxDistance: 500, MaxHeight: 10},
        &participants.Cat{Name: "Masha", MaxDistance: 50, MaxHeight: 1},
        &participants.Robot{Name: "Joe", MaxDistance: 200, MaxHeight: 5},
    }

    obstaclesList := []obstacles.Obstacle{
        &obstacles.Truck{Distance: 60},
        &obstacles.Wall{Height: 3},
        &obstacles.Truck{Distance: 150},
    }

	for _, p := range participantsList {
        fmt.Printf("\nParticipant [%s] started\n", p.GetName())
        passedAll := true

        for _, o := range obstaclesList {
            if !o.Overcome(p) {
                fmt.Printf("Participant [%s] was kicked!\n", p.GetName())
                passedAll = false
                break
            }
        }

        if passedAll {
            fmt.Printf("Participant [%s] finished!\n", p.GetName())
        }
    }
}