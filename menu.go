package main

import "fmt"

const (
	reset  = "\x1b[0m"
	bold   = "\x1b[1m"
	italic = "\x1b[3m"

	red   = "\x1b[38;5;160m"
	gold  = "\x1b[38;5;220m"
	gray  = "\x1b[38;5;244m"
	white = "\x1b[97m"
	green = "\x1b[38;5;114m"
)

var menu = gold + "╔════════════════════════════════════╗" + reset + "\n" +
	gold + "║" + reset + "       " + bold + gold + "▓▓▓  LE QUARTIER  ▓▓▓" + reset + "        " + gold + "║" + reset + "\n" +
	gold + "╠════════════════════════════════════╣" + reset + "\n" +
	gold + "║" + reset + "                                    " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "1." + reset + " " + white + "Perso" + reset + "                         " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "2." + reset + " " + white + "Marchand" + reset + "                      " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "3." + reset + " " + white + "Charcudoc" + reset + "                     " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "4." + reset + " " + white + "Guerre de gang" + reset + "                " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "5." + reset + " " + white + "Inventaire" + reset + "                    " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "6." + reset + " " + white + "EXIT" + reset + "                          " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "                                    " + gold + "║" + reset + "\n" +
	gold + "║" + reset + gray + italic + "  › choisis un numéro pour agir" + reset + "     " + gold + "║" + reset + "\n" +
	gold + "╚════════════════════════════════════╝" + reset + "\n"

func marketMenu(market Market, character Character) string {
	return gold + "╔═════════════════════════════════════════╗" + reset + "\n" +
		gold + "║" + reset + "           " + bold + gold + "▓▓▓  FORGERON  ▓▓▓" + reset + "          " + gold + "║" + reset + "\n" +
		gold + "╠═════════════════════════════════════════╣" + reset + "\n" +
		gold + "║" + reset + "                                         " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "1." + reset + " " + white + "Steampack basse qualité" + reset + "        " + green + bold + "3 Ferraille" + reset + " " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "2." + reset + " " + white + "Steampack" + reset + "                    " + green + bold + "5 Ferraille, 2 Composants" + reset + " " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "3." + reset + " " + white + "Grenade à fragmentation" + reset + "      " + green + bold + "2 Ferraille, 3 Poudre" + reset + " " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "4." + reset + " " + white + "EXIT" + reset + "                         " + green + bold + fmt.Sprintf("%d Ferraille", character.resourceQuantity("Ferraille")) + reset + " " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "                                         " + gold + "║" + reset + "\n" +
		gold + "║" + reset + gray + italic + "  › choisis un numéro pour acheter" + reset + "       " + gold + "║" + reset + "\n" +
		gold + "╚═════════════════════════════════════════╝" + reset + "\n"

}

func charcudocMenu() string {
	return gold + "╔═════════════════════════════════════════╗" + reset + "\n" +
		gold + "║" + reset + "          " + bold + gold + "▓▓▓  CHARCUDOC  ▓▓▓" + reset + "          " + gold + "║" + reset + "\n" +
		gold + "╠═════════════════════════════════════════╣" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "IMPLANTS" + reset + "                              " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   1. Interface neurale                  " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   2. Optiques cybernétiques              " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   3. Réflexes augmentés                  " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   4. Système immunitaire                 " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   5. Module de piratage                  " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   6. Blindage dermique                   " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "TÊTE" + reset + "                                  " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   7. Scanner rétinien                   " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   8. Processeur cérébral                " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "CORPS" + reset + "                                 " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   9. Coeur synthétique                  " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   10. Poumons artificiels               " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   11. Bras mécaniques                   " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   " + red + bold + "JAMBES" + reset + "                                " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   12. Jambes renforcées                 " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   13. Bottes hydrauliques               " + gold + "║" + reset + "\n" +
		gold + "║" + reset + "   14. Retour                            " + gold + "║" + reset + "\n" +
		gold + "╚═════════════════════════════════════════╝" + reset + "\n"
}
