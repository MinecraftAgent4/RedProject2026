package main

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

var market = gold + "╔═════════════════════════════════════════╗" + reset + "\n" +
	gold + "║" + reset + "     " + bold + gold + "▓▓▓  MARCHÉ  ▓▓▓" + reset + "     " + gold + "║" + reset + "\n" +
	gold + "╠═════════════════════════════════════════╣" + reset + "\n" +
	gold + "║" + reset + "                                         " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "1." + reset + " " + white + "steampack de basse qualité " + reset + "      " + green + bold + "12 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "2." + reset + " " + white + "steampack" + reset + "     " + green + bold + "28 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "3." + reset + " " + white + "Livre de Sort : Boule de Feu " + reset + "      " + green + bold + "55 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "4." + reset + " " + white + "Grenade à fragmentation" + reset + "      " + green + bold + "40 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "5." + reset + " " + white + "Grenade fumigène" + reset + "             " + green + bold + "22 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "6." + reset + " " + white + "Grenade incendiaire" + reset + "          " + green + bold + "48 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "7." + reset + " " + white + "Grenade paralysante" + reset + "          " + green + bold + "35 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "   " + red + bold + "8." + reset + " " + white + "EXIT" + reset + "          " + green + bold + "35 po" + reset + " " + gold + "║" + reset + "\n" +
	gold + "║" + reset + "                                         " + gold + "║" + reset + "\n" +
	gold + "║" + reset + gray + italic + "  › choisis un numéro pour acheter" + reset + "       " + gold + "║" + reset + "\n" +
	gold + "╚═════════════════════════════════════════╝" + reset + "\n"
