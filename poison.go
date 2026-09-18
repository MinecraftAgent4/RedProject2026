package main

import "time"

func (c *Character) poison() {
	for i := 0; i <= 2; i++ {
		c.AddPV(-10)
		time.Sleep(1 * time.Second)
	}
}
