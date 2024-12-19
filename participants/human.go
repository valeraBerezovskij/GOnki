package participants

import "fmt"

type Human struct {
	Name        string
	MaxDistance int
	MaxHeight   int
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) Run(distance int) bool {
	if distance > h.MaxDistance {
		fmt.Printf("Participant [%s] did not run a [%d] distance\n", h.Name, distance)
		return false
	}
	fmt.Printf("Participant [%s] ran [%d] distance\n", h.Name, distance)
	return true
}

func (h *Human) Jump(height int) bool {
	if height > h.MaxHeight {
		fmt.Printf("Participant [%s] did not jump a [%d] height\n", h.Name, height)
		return false
	}
	fmt.Printf("Participant [%s] jumped [%d] height\n", h.Name, height)
	return true
}