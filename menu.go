package main

const (
	reset  = "\x1b[0m"
	bold   = "\x1b[1m"
	italic = "\x1b[3m"

	red   = "\x1b[38;5;160m"
	gold  = "\x1b[38;5;220m"
	gray  = "\x1b[38;5;244m"
	white = "\x1b[97m"
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
