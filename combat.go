package main

import "fmt"

func Attack(a Entity, t Entity, mod int) {
	damage := (a.Atk() * mod) - t.Def()

	if damage < 1 {
		damage = 1
	}

	t.Dmg(damage)
}

func (m *Monster) goblinPattern(turn int, t *Character) {
	if turn%3 == 2 {
		Attack(m, t, 2)
	} else {
		Attack(m, t, 1)
	}
}

func (m *Monster) exp_reward(c *Character) {
	if m.pv_actuelle <= 0 {
		c.exp_joueur += 50

		fmt.Printf("Vous gagnez 50 XP !\n")
		fmt.Printf("XP : %d/%d\n", c.exp_joueur, c.exp_required)

		if c.exp_joueur >= c.exp_required {
			c.lvlUp()
			fmt.Printf("Vous passez niveau %d !\n", c.niveau)
		}
	}
}

func combat(perso *Character, monstre Monster) bool {
	fmt.Println("\n================================")
	fmt.Println("           COMBAT")
	fmt.Println("================================")

	turn := 0

	for perso.Hp() > 0 && monstre.Hp() > 0 {
		fmt.Println()
		fmt.Printf("%s : %d/%d PV\n",
			perso.Name(),
			perso.Hp(),
			perso.MaxHp())

		fmt.Printf("%s : %d/%d PV\n",
			monstre.Name(),
			monstre.Hp(),
			monstre.MaxHp())

		fmt.Println("\n1 - Attaquer")
		fmt.Println("2 - Se soigner")
		fmt.Println("3 - Fuir")
		fmt.Print("Choix : ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {

		case "1":
			fmt.Printf("\n%s attaque avec %s !\n",
				perso.Name(),
				perso.weapon.Nom())

			Attack(perso, &monstre, 1)

			fmt.Printf("%s perd des PV.\n", monstre.Name())

		case "2":
			soin := 15
			perso.AddPV(soin)

			fmt.Printf("\nVous récupérez %d PV.\n", soin)

		case "3":
			fmt.Println("\nVous fuyez le combat.")
			return false

		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if monstre.Hp() <= 0 {
			break
		}

		turn++

		fmt.Printf("\n%s attaque !\n", monstre.Name())
		monstre.goblinPattern(turn, perso)
	}

	if perso.Hp() <= 0 {
		fmt.Println("\n================================")
		fmt.Println("          DÉFAITE")
		fmt.Println("================================")

		fmt.Println("Vous avez été vaincu.")

		perso.isDead()

		return false
	}

	fmt.Println("\n================================")
	fmt.Println("          VICTOIRE")
	fmt.Println("================================")

	fmt.Printf("Vous avez vaincu %s !\n", monstre.Name())

	perso.money += 20
	fmt.Println("Vous gagnez 20 po.")

	monstre.exp_reward(perso)

	return true
}