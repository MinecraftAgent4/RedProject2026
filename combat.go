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
	fmt.Println("            COMBAT")
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

		fmt.Println("\n--------------------------------")

		// ============================
		// MENU SELON LA CLASSE
		// ============================

		switch perso.classe {

		case "merc":
			fmt.Println("1 - Attaque au CAC")
			fmt.Println("2 - Tirer au pistolet")
			fmt.Println("3 - Se soigner")
			fmt.Println("4 - Fuir")

		case "netrunner":
			fmt.Println("1 - Hacks")
			fmt.Println("2 - Se soigner")
			fmt.Println("3 - Fuir")

		case "cyberpsycho":
			fmt.Println("1 - Tir gauche")
			fmt.Println("2 - Tir droit")
			fmt.Println("3 - Double tir")
			fmt.Println("4 - Se soigner")
			fmt.Println("5 - Fuir")
		}

		fmt.Print("Choix : ")

		var choix int
		fmt.Scanln(&choix)

		attaqueEffectuee := false
		actionValide := false

		// ============================
		// MERC
		// ============================

		if perso.classe == "merc" {

			switch choix {

			case 1:
				fmt.Printf("\n Vous attaquez avec %s !\n",
					perso.weapon.Nom())

				Attack(perso, &monstre, 1)

				attaqueEffectuee = true
				actionValide = true

			case 2:
				fmt.Printf("\n Vous tirez avec %s !\n",
					perso.weapon2.Nom())

				ancienneArme := perso.weapon
				perso.weapon = perso.weapon2

				Attack(perso, &monstre, 1)

				perso.weapon = ancienneArme

				attaqueEffectuee = true
				actionValide = true

			case 3:
				soin := 15
				perso.AddPV(soin)

				if perso.Hp() > perso.MaxHp() {
					perso.pv_actuelle = perso.MaxHp()
				}

				fmt.Printf("\n Vous récupérez %d PV.\n", soin)

				actionValide = true

			case 4:
				fmt.Println("\nVous fuyez le combat.")
				return false

			default:
				fmt.Println("Choix invalide.")
			}
		}

		// ============================
		// NETRUNNER
		// ============================

		if perso.classe == "netrunner" {

			switch choix {

			case 1:
				if netrunnerAttack(perso, &monstre) {
					attaqueEffectuee = true
					actionValide = true
				}

			case 2:
				soin := 15
				perso.AddPV(soin)

				if perso.Hp() > perso.MaxHp() {
					perso.pv_actuelle = perso.MaxHp()
				}

				fmt.Printf("\n Vous récupérez %d PV.\n", soin)

				actionValide = true

			case 3:
				fmt.Println("\nVous fuyez le combat.")
				return false

			default:
				fmt.Println("Choix invalide.")
			}
		}

		// ============================
		// CYBERPSYCHO
		// ============================

		if perso.classe == "cyberpsycho" {

			switch choix {

			case 1:
				fmt.Printf("\n Tir avec %s !\n",
					perso.weapon.Nom())

				Attack(perso, &monstre, 1)

				attaqueEffectuee = true
				actionValide = true

			case 2:
				fmt.Printf("\n Tir avec %s !\n",
					perso.weapon2.Nom())

				ancienneArme := perso.weapon
				perso.weapon = perso.weapon2

				Attack(perso, &monstre, 1)

				perso.weapon = ancienneArme

				attaqueEffectuee = true
				actionValide = true

			case 3:
				fmt.Println("\n DOUBLE TIR !")

				Attack(perso, &monstre, 1)

				if monstre.Hp() > 0 {
					ancienneArme := perso.weapon
					perso.weapon = perso.weapon2

					Attack(perso, &monstre, 1)

					perso.weapon = ancienneArme
				}

				attaqueEffectuee = true
				actionValide = true

			case 4:
				soin := 15
				perso.AddPV(soin)

				if perso.Hp() > perso.MaxHp() {
					perso.pv_actuelle = perso.MaxHp()
				}

				fmt.Printf("\n Vous récupérez %d PV.\n", soin)

				actionValide = true

			case 5:
				fmt.Println("\nVous fuyez le combat.")
				return false

			default:
				fmt.Println("Choix invalide.")
			}
		}

		// ============================
		// LE MONSTRE ATTAQUE
		// ============================

		if monstre.Hp() <= 0 {
			break
		}

		if actionValide && attaqueEffectuee {
			turn++

			fmt.Printf("\n %s attaque !\n",
				monstre.Name())

			monstre.goblinPattern(turn, perso)
		}
	}

	// ============================
	// DÉFAITE
	// ============================

	if perso.Hp() <= 0 {
		fmt.Println("\n================================")
		fmt.Println("            DÉFAITE")
		fmt.Println("================================")

		fmt.Println("Vous avez été vaincu.")

		perso.isDead()

		return false
	}

	// ============================
	// VICTOIRE
	// ============================

	fmt.Println("\n================================")
	fmt.Println("           VICTOIRE")
	fmt.Println("================================")

	fmt.Printf("Vous avez vaincu %s !\n",
		monstre.Name())

	perso.money += 20

	fmt.Println("Vous gagnez 20 po.")

	monstre.exp_reward(perso)

	return true
}


func mercAttack(perso *Character, monstre *Monster, choix int) {
	switch choix {
	case 1:
		fmt.Printf("%s attaque avec %s !\n",
			perso.Name(),
			perso.weapon.Nom())

		Attack(perso, monstre, 1)

	case 2:
		fmt.Printf("%s tire avec %s !\n",
			perso.Name(),
			perso.weapon2.Nom())

		ancienne := perso.weapon
		perso.weapon = perso.weapon2

		Attack(perso, monstre, 1)

		perso.weapon = ancienne
	}
}

func cyberpsychoAttack(perso *Character, monstre *Monster, choix int) {
	switch choix {

	case 1:
		fmt.Println("Tir du gros calibre gauche !")
		ancienne := perso.weapon
		perso.weapon = perso.weapon
		Attack(perso, monstre, 1)
		perso.weapon = ancienne

	case 2:
		fmt.Println("Tir du gros calibre droit !")

		ancienne := perso.weapon
		perso.weapon = perso.weapon2

		Attack(perso, monstre, 1)

		perso.weapon = ancienne

	case 3:
		fmt.Println("DOUBLE TIR !")

		Attack(perso, monstre, 1)

		if monstre.Hp() > 0 {
			ancienne := perso.weapon
			perso.weapon = perso.weapon2

			Attack(perso, monstre, 1)

			perso.weapon = ancienne
		}
	}
}

func netrunnerAttack(perso *Character, monstre *Monster) {
	fmt.Println("\n===== HACKS =====")
	fmt.Println("1 - Surcharge")
	fmt.Println("2 - Virus")
	fmt.Println("3 - Court-circuit")
	fmt.Println("4 - Retour")
	fmt.Print("Choix : ")

	var choix int
	fmt.Scanln(&choix)

	switch choix {

	case 1:
		// Surcharge : gros dégâts directs
		fmt.Println("\n💻 Surcharge du système !")

		damage := 20 - monstre.Def()
		if damage < 1 {
			damage = 1
		}

		monstre.Dmg(damage)

		fmt.Printf("Le hack inflige %d dégâts !\n", damage)

	case 2:
		// Virus : dégâts plus faibles mais poison
		fmt.Println("\n Injection d'un virus !")

		damage := 10 - monstre.Def()
		if damage < 1 {
			damage = 1
		}

		monstre.Dmg(damage)

		fmt.Printf("Le virus inflige %d dégâts !\n", damage)
		monstre.poison()

		

	case 3:
		// Court-circuit : dégâts + deuxième petit dégât
		fmt.Println("\n⚡ Court-circuit !")

		damage := 12 - monstre.Def()
		if damage < 1 {
			damage = 1
		}

		monstre.Dmg(damage)

		fmt.Printf("Le court-circuit inflige %d dégâts !\n", damage)

	case 4:
		fmt.Println("Retour.")
		return

	default:
		fmt.Println("Choix invalide.")
	}
}