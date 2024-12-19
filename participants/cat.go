package participants

import "fmt"

type Cat struct {
	Name        string
	MaxDistance int
	MaxHeight   int
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Run(distance int) bool {
	if distance > c.MaxDistance {
		fmt.Printf("Cat [%s] did not ran a [%d] distance\n", c.Name, distance)
		return false
	}
	fmt.Printf("Cat [%s] ran [%d] distance\n", c.Name, distance)
	return true
}

func (c *Cat) Jump(height int) bool {
	if height > c.MaxHeight {
		fmt.Printf("Cat [%s] did not jumped a [%d] height\n", c.Name, height)
		return false
	}
	fmt.Printf("Cat [%s] jumped [%d] height\n", c.Name, height)
	return true
}