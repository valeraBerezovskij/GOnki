package participants

import "fmt"

type Robot struct {
	Name        string
	MaxDistance int
	MaxHeight   int
}

func (r *Robot) GetName() string {
	return r.Name
}

func (r *Robot) Run(distance int) bool {
	if distance > r.MaxDistance {
		fmt.Printf("Robot [%s]  did not ran a [%d] distance\n", r.Name, distance)
		return false
	}
	fmt.Printf("Robot [%s] ran [%d] distance\n", r.Name, distance)
	return true
}

func (r *Robot) Jump(height int) bool {
	if height > r.MaxHeight {
		fmt.Printf("Robot [%s]  did not jumped a [%d] height\n", r.Name, height)
		return false
	}
	fmt.Printf("Robot [%s] jumped [%d] height\n", r.Name, height)
	return true
}