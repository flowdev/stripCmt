package main

import (
	"testing"
//	"fmt"
)

func TestCellIsAlive(t *testing.T) {
	cell := NewCell()
	if !cell.isAlive() {
		t.Fatal("Test Cell is Alive: RED") 
	}
}

func TestCellIsDead(t *testing.T){
	cell := Cell{false}
	if cell.isAlive() {
		t.Fatal("Fatal: the cell is alive.")
	}
}

func TestCellDies(t *testing.T) {
	cell := Cell{true}
	cell.evolve()
	if cell.isAlive() {
		t.Fatal("Fatal: Cell is supposed to die")
	}
}

