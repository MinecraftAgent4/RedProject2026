package main

import fmt

type Character struct {
nom         string
classe      string
niveau      int
pv_total    int
pv_actuelle int
inventaire  []string
}

func (c *Character) displayInfo() {
fmt.Println(
nom
