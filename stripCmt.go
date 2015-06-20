package main

type Cell struct {
	state bool
}

func (c *Cell) isAlive() bool {
	return c.state 
}

func NewCell() *Cell {
	return &Cell{true}
}

func (c *Cell) evolve() {
	c.state = false
}

