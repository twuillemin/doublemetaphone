package doublemetaphone

import "strings"

const maxLength = 32

// DoubleMetaphone computes the primary and secondary Metaphone of the given string. The orignal string
// is not modified by the function
func DoubleMetaphone(str string) (string, string) {

	// We need the real length and last prior to padding
	var length = uint(len(str))
	var last = length - 1

	// Pad original so we can index beyond end
	var original = strings.ToUpper(str)
	original = original + "     "

	var current = uint(0)
	var primary = ""
	var secondary = ""

	// Skip these when at start of word
	if matchStrings(original, 0, []string{"GN", "KN", "PN", "WR", "PS"}) {
		current += 1
	}

	// Initial 'X' is pronounced 'Z' e.g. 'Xavier'
	if getAt(original, 0) == 'X' {
		// 'Z' maps to 'S'
		primary += "S"
		secondary += "S"
		current += 1
	}

	// Main loop
	for (len(primary) < maxLength) || (len(secondary) < maxLength) {
		if current >= length {
			break
		}

		switch getAt(original, current) {
		case 'A', 'E', 'I', 'O', 'U', 'Y':
			if current == 0 {
				// All init vowels now map to 'A'
				primary += "A"
				secondary += "A"
			}
			current += 1
			break

		case 'B':
			// "-mb", e.g", "dumb", already skipped over...
			primary += "P"
			secondary += "P"

			if getAt(original, current+1) == 'B' {
				current += 2
			} else {
				current += 1
			}
			break

		case 'Ç':
			primary += "S"
			secondary += "S"
			current += 1
			break

		case 'C':
			// Various germanic
			if (current > 1) &&
				!isVowel(original, current-2) &&
				matchString(original, current-1, "ACH") &&
				((getAt(original, current+2) != 'I') &&
					((getAt(original, current+2) != 'E') ||
						matchStrings(original, current-2, []string{"BACHER", "MACHER"}))) {
				primary += "K"
				secondary += "K"
				current += 2
				break
			}

			// Special case 'caesar'
			if (current == 0) && matchString(original, current, "CAESAR") {
				primary += "S"
				secondary += "S"
				current += 2
				break
			}

			// Italian 'chianti'
			if matchString(original, current, "CHIA") {
				primary += "K"
				secondary += "K"
				current += 2
				break
			}

			if matchStrings(original, current, []string{"CH"}) {
				// Find 'michael'
				if (current > 0) && matchString(original, current, "CHAE") {
					primary += "K"
					secondary += "X"
					current += 2
					break
				}

				// Greek roots e.g. 'chemistry', 'chorus'
				if (current == 0) &&
					matchStrings(original, current+1, []string{"HARAC", "HARIS", "HOR", "HYM", "HIA", "HEM"}) &&
					!matchString(original, 0, "CHORE") {
					primary += "K"
					secondary += "K"
					current += 2
					break
				}

				// Germanic, greek, or otherwise 'ch' for 'kh' sound
				if matchStrings(original, 0, []string{"VAN ", "VON ", "SCH"}) ||
				// 'architect but not 'arch', 'orchestra', 'orchid'
					matchStrings(original, current-2, []string{"ORCHES", "ARCHIT", "ORCHID"}) ||
					matchStrings(original, current+2, []string{"T", "S"}) ||
					((matchStrings(original, current-1, []string{"A", "O", "U", "E"}) || (current == 0)) &&
					// e.g., 'wachtler', 'wechsler', but not 'tichner' 
						matchStrings(original, current+2, []string{"L", "R", "N", "M", "B", "H", "F", "V", "W", " "})) {
					primary += "K"
					secondary += "K"
				} else {
					if current > 0 {
						if matchString(original, 0, "MC") {
							// E.g., "McHugh"
							primary += "K"
							secondary += "K"
						} else {
							primary += "X"
							secondary += "K"
						}
					} else {
						primary += "X"
						secondary += "X"
					}
				}
				current += 2
				break
			}
			// E.g, 'czerny'
			if matchString(original, current, "CZ") && !matchString(original, current-2, "WICZ") {
				primary += "S"
				secondary += "X"
				current += 2
				break
			}

			// E.g., 'focaccia'
			if matchString(original, current+1, "CIA") {
				primary += "X"
				secondary += "X"
				current += 3
				break
			}

			// Double 'C', but not if e.g. 'McClellan'
			if matchString(original, current, "CC") && !((current == 1) && (getAt(original, 0) == 'M')) {
				// 'bellocchio' but not 'bacchus' 
				if matchStrings(original, current+2, []string{"I", "E", "H"}) && !matchString(original, current+2, "HU") {
					// 'accident', 'accede' 'succeed' 
					if ((current == 1) && (getAt(original, current-1) == 'A')) || matchStrings(original, current-1, []string{"UCCEE", "UCCES"}) {
						primary += "KS"
						secondary += "KS"
						// 'bacci', 'bertucci', other italian 
					} else {
						primary += "X"
						secondary += "X"
					}
					current += 3
					break
				} else { // Pierce's rule 
					primary += "K"
					secondary += "K"
					current += 2
					break
				}
			}

			if matchStrings(original, current, []string{"CK", "CG", "CQ"}) {
				primary += "K"
				secondary += "K"
				current += 2
				break
			}

			if matchStrings(original, current, []string{"CI", "CE", "CY"}) {
				// italian vs. english 
				if matchStrings(original, current, []string{"CIO", "CIE", "CIA"}) {
					primary += "S"
					secondary += "X"
				} else {
					primary += "S"
					secondary += "S"
				}
				current += 2
				break
			}

			// else 
			primary += "K"
			secondary += "K"

			// name sent in 'mac caffrey', 'mac gregor 
			if matchStrings(original, current+1, []string{" C", " Q", " G"}) {
				current += 3
			} else {
				if matchStrings(original, current+1, []string{"C", "K", "Q"}) &&
					!matchStrings(original, current+1, []string{"CE", "CI"}) {
					current += 2
				} else {
					current += 1
				}
			}
			break

		case 'D':
			if matchString(original, current, "DG") {
				if matchStrings(original, current+2, []string{"I", "E", "Y"}) {
					// e.g. 'edge' 
					primary += "J"
					secondary += "J"
					current += 3
					break
				} else {
					// e.g. 'edgar' 
					primary += "TK"
					secondary += "TK"
					current += 2
					break
				}
			}

			if matchStrings(original, current, []string{"DT", "DD"}) {
				primary += "T"
				secondary += "T"
				current += 2
				break
			}

			// else 
			primary += "T"
			secondary += "T"
			current += 1
			break

		case 'F':
			if getAt(original, current+1) == 'F' {
				current += 2
			} else {
				current += 1
				primary += "F"
				secondary += "F"
			}
			break

		case 'G':
			if getAt(original, current+1) == 'H' {
				if (current > 0) && !isVowel(original, current-1) {
					primary += "K"
					secondary += "K"
					current += 2
					break
				}

				if current < 3 {
					// 'ghislane', ghiradelli 
					if current == 0 {
						if getAt(original, current+2) == 'I' {
							primary += "J"
							secondary += "J"
						} else {
							primary += "K"
							secondary += "K"
						}
						current += 2
						break
					}
				}
				// Parker's rule (with some further refinements) - e.g., 'hugh' 
				if ((current > 1) && matchStrings(original, current-2, []string{"B", "H", "D"})) ||
				// E.g., 'bough'
					((current > 2) && matchStrings(original, current-3, []string{"B", "H", "D"})) ||
				// E.g., 'broughton'
					((current > 3) && matchStrings(original, current-4, []string{"B", "H"})) {
					current += 2
					break
				} else {
					// E.g., 'laugh', 'McLaughlin', 'cough', 'gough', 'rough', 'tough'
					if (current > 2) && (getAt(original, current-1) == 'U') && matchStrings(original, current-3, []string{"C", "G", "L", "R", "T"}) {
						primary += "F"
						secondary += "F"
					} else if (current > 0) && getAt(original, current-1) != 'I' {
						primary += "K"
						secondary += "K"
					}

					current += 2
					break
				}
			}

			if getAt(original, current+1) == 'N' {
				if (current == 1) && isVowel(original, 0) && !isSlavoGermanic(original) {
					primary += "KN"
					secondary += "N"
				} else
				// Not e.g. 'cagney'
				if !matchStrings(original, current+2, []string{"EY"}) && (getAt(original, current+1) != 'Y') && !isSlavoGermanic(original) {
					primary += "N"
					secondary += "KN"
				} else {
					primary += "KN"
					secondary += "KN"
				}
				current += 2
				break
			}

			// 'tagliaro' 
			if matchStrings(original, current+1, []string{"LI"}) && !isSlavoGermanic(original) {
				primary += "KL"
				secondary += "L"
				current += 2
				break
			}

			// -ges-,-gep-,-gel-, -gie- at beginning 
			if (current == 0) &&
				((getAt(original, current+1) == 'Y') ||
					matchStrings(original, current+1, []string{"ES", "EP", "EB", "EL", "EY", "IB", "IL", "IN", "IE", "EI", "ER"})) {
				primary += "K"
				secondary += "J"
				current += 2
				break
			}

			//  -ger-,  -gy- 
			if (matchString(original, current+1, "ER") || (getAt(original, current+1) == 'Y')) &&
				!matchStrings(original, 0, []string{"DANGER", "RANGER", "MANGER"}) &&
				!matchStrings(original, current-1, []string{"E", "I"}) &&
				!matchStrings(original, current-1, []string{"RGY", "OGY"}) {
				primary += "K"
				secondary += "J"
				current += 2
				break
			}

			//  Italian e.g, 'biaggi'
			if matchStrings(original, current+1, []string{"E", "I", "Y"}) || matchStrings(original, current-1, []string{"AGGI", "OGGI"}) {
				// Obvious germanic
				if (matchStrings(original, 0, []string{"VAN ", "VON "}) ||
					matchStrings(original, 0, []string{"SCH"})) ||
					matchStrings(original, current+1, []string{"ET"}) {
					primary += "K"
					secondary += "K"
				} else {
					// Always soft if french ending
					if (matchStrings(original, current+1, []string{"IER "})) {
						primary += "J"
						secondary += "J"
					} else {
						primary += "J"
						secondary += "K"
					}
				}
				current += 2
				break
			}

			if getAt(original, current+1) == 'G' {
				current += 2
			} else {
				current += 1
				primary += "K"
				secondary += "K"
			}
			break

		case 'H':
			// Only keep if first & before vowel or btw. 2 vowels
			if ((current == 0) || isVowel(original, current-1)) && isVowel(original, current+1) {
				primary += "H"
				secondary += "H"
				current += 2
			} else { // also takes care of 'HH' 
				current += 1
			}
			break

		case 'J':
			// Obvious spanish, 'jose', 'san jacinto'
			if matchString(original, current, "JOSE") || matchStrings(original, 0, []string{"SAN "}) {
				if ((current == 0) && (getAt(original, current+4) == ' ')) || matchStrings(original, 0, []string{"SAN "}) {
					primary += "H"
					secondary += "H"
				} else {
					primary += "J"
					secondary += "H"
				}
				current += 1
				break
			}

			// Yankelovich/Jankelowicz
			if (current == 0) && !matchString(original, current, "JOSE") {
				primary += "J"
				secondary += "A"
			} else {
				// Spanish pron. of e.g. 'bajador'
				if isVowel(original, current-1) && !isSlavoGermanic(original) &&
					((getAt(original, current+1) == 'A') || (getAt(original, current+1) == 'O')) {
					primary += "J"
					secondary += "H"
				} else {
					if current == last {
						primary += "J"
						secondary += ""
					} else {
						if !matchStrings(original, current+1, []string{"L", "T", "K", "S", "N", "M", "B", "Z"}) &&
							!matchStrings(original, current-1, []string{"S", "K", "L"}) {
							primary += "J"
							secondary += "J"
						}
					}
				}
			}

			if getAt(original, current+1) == 'J' { // it could happen! 
				current += 2
			} else {
				current += 1
			}
			break

		case 'K':
			if getAt(original, current+1) == 'K' {
				current += 2
			} else {
				current += 1
				primary += "K"
				secondary += "K"
			}
			break

		case 'L':
			if getAt(original, current+1) == 'L' {
				// Spanish e.g. 'cabrillo', 'gallegos'
				if ((current == (length - 3)) &&
					matchStrings(original, current-1, []string{"ILLO", "ILLA", "ALLE"})) ||
					((matchStrings(original, last-1, []string{"AS", "OS"}) ||
						matchStrings(original, last, []string{"A", "O"})) &&
						matchString(original, current-1, "ALLE")) {
					primary += "L"
					secondary += ""
					current += 2
					break
				}
				current += 2
			} else {
				current += 1
				primary += "L"
				secondary += "L"
			}
			break

		case 'M':
			if (matchStrings(original, current-1, []string{"UMB"}) &&
				(((current + 1) == last) ||
					matchStrings(original, current+2, []string{"ER"}))) ||
			// 'dumb','thumb' 
				(getAt(original, current+1) == 'M') {
				current += 2
			} else {
				current += 1
			}
			primary += "M"
			secondary += "M"
			break

		case 'N':
			if getAt(original, current+1) == 'N' {
				current += 2
			} else {
				current += 1
			}
			primary += "N"
			secondary += "N"
			break

		case 'Ñ':
			current += 1
			primary += "N"
			secondary += "N"
			break

		case 'P':
			if getAt(original, current+1) == 'H' {
				primary += "F"
				secondary += "F"
				current += 2
				break
			}

			// Also account for "campbell", "raspberry"
			if matchStrings(original, current+1, []string{"P", "B"}) {
				current += 2
			} else {
				current += 1
				primary += "P"
				secondary += "P"
			}
			break

		case 'Q':
			if getAt(original, current+1) == 'Q' {
				current += 2
			} else {
				current += 1
				primary += "K"
				secondary += "K"
			}
			break

		case 'R':
			// French e.g. 'rogier', but exclude 'hochmeier'
			if (current == last) &&
				!isSlavoGermanic(original) &&
				matchString(original, current-2, "IE") &&
				!matchStrings(original, current-4, []string{"ME", "MA"}) {
				primary += ""
				secondary += "R"
			} else {
				primary += "R"
				secondary += "R"
			}

			if getAt(original, current+1) == 'R' {
				current += 2
			} else {
				current += 1
			}
			break

		case 'S':
			// Special cases 'island', 'isle', 'carlisle', 'carlysle'
			if matchStrings(original, current-1, []string{"ISL", "YSL"}) {
				current += 1
				break
			}

			// Special case 'sugar-'
			if (current == 0) && matchString(original, current, "SUGAR") {
				primary += "X"
				secondary += "S"
				current += 1
				break
			}

			if matchString(original, current, "SH") {
				// Germanic
				if matchStrings(original, current+1, []string{"HEIM", "HOEK", "HOLM", "HOLZ"}) {
					primary += "S"
					secondary += "S"
				} else {
					primary += "X"
					secondary += "X"
				}
				current += 2
				break
			}

			// Italian & armenian
			if matchStrings(original, current, []string{"SIO", "SIA", "SIAN"}) {
				if !isSlavoGermanic(original) {
					primary += "S"
					secondary += "X"
				} else {
					primary += "S"
					secondary += "S"
				}
				current += 3
				break
			}

			// German & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
			// also, -sz- in slavic language altho in hungarian it is pronounced 's'
			if ((current == 0) && matchStrings(original, current+1, []string{"M", "N", "L", "W"})) ||
				matchString(original, current+1, "Z") {
				primary += "S"
				secondary += "X"
				if matchString(original, current+1, "Z") {
					current += 2
				} else {
					current += 1
				}
				break
			}

			if matchString(original, current, "SC") {
				// Schlesinger's rule 
				if getAt(original, current+2) == 'H' {
					// dutch origin, e.g. 'school', 'schooner' 
					if matchStrings(original, current+3, []string{"OO", "ER", "EN", "UY", "ED", "EM"}) {
						// 'schermerhorn', 'schenker' 
						if matchStrings(original, current+3, []string{"ER", "EN"}) {
							primary += "X"
							secondary += "SK"
						} else {
							primary += "SK"
							secondary += "SK"
						}
						current += 3
						break
					} else {
						if (current == 0) && !isVowel(original, 3) && (getAt(original, 3) != 'W') {
							primary += "X"
							secondary += "S"
						} else {
							primary += "X"
							secondary += "X"
						}
						current += 3
						break
					}
				}

				if matchStrings(original, current+2, []string{"I", "E", "Y"}) {
					primary += "S"
					secondary += "S"
					current += 3
					break
				}
				// else 
				primary += "SK"
				secondary += "SK"
				current += 3
				break
			}

			// french e.g. 'resnais', 'artois' 
			if (current == last) &&
				matchStrings(original, current-2, []string{"AI", "OI"}) {
				primary += ""
				secondary += "S"
			} else {
				primary += "S"
				secondary += "S"
			}

			if matchStrings(original, current+1, []string{"S", "Z"}) {
				current += 2
			} else {
				current += 1
			}
			break

		case 'T':
			if matchString(original, current, "TION") {
				primary += "X"
				secondary += "X"
				current += 3
				break
			}

			if matchStrings(original, current, []string{"TIA", "TCH"}) {
				primary += "X"
				secondary += "X"
				current += 3
				break
			}

			if matchStrings(original, current, []string{"TH", "TTH"}) {
				// special case 'thomas', 'thames' or germanic
				if matchStrings(original, current+2, []string{"OM", "AM"}) || matchStrings(original, 0, []string{"VAN ", "VON ", "SCH"}) {
					primary += "T"
					secondary += "T"
				} else {
					primary += "0" // yes, zero
					secondary += "T"
				}
				current += 2
				break
			}

			if matchStrings(original, current+1, []string{"T", "D"}) {
				current += 2
			} else {
				current += 1
			}
			primary += "T"
			secondary += "T"
			break

		case 'V':
			if getAt(original, current+1) == 'V' {
				current += 2
			} else {
				current += 1
			}
			primary += "F"
			secondary += "F"
			break

		case 'W':
			// can also be in middle of word
			if matchString(original, current, "WR") {
				primary += "R"
				secondary += "R"
				current += 2
				break
			}

			if (current == 0) &&
				(isVowel(original, current+1) || matchString(original, current, "WH")) {
				// Wasserman should match Vasserman
				if isVowel(original, current+1) {
					primary += "A"
					secondary += "F"
				} else {
					// need Uomo to match Womo
					primary += "A"
					secondary += "A"
				}
			}

			// Arnow should match Arnoff
			if ((current == last) && isVowel(original, current-1)) ||
				matchStrings(original, current-1, []string{"EWSKI", "EWSKY", "OWSKI", "OWSKY"}) ||
				matchStrings(original, 0, []string{"SCH"}) {
				primary += ""
				secondary += "F"
				current += 1
				break
			}

			// polish e.g. 'filipowicz'
			if matchStrings(original, current, []string{"WICZ", "WITZ"}) {
				primary += "TS"
				secondary += "FX"
				current += 4
				break
			}

			// else skip it
			current += 1
			break

		case 'X':
			// french e.g. breaux
			if !((current == last) &&
				(matchStrings(original, current-3, []string{"IAU", "EAU"}) || matchStrings(original, current-2, []string{"AU", "OU"}))) {
				primary += "KS"
				secondary += "KS"
			}

			if matchStrings(original, current+1, []string{"C", "X"}) {
				current += 2
			} else {
				current += 1
			}
			break

		case 'Z':
			// chinese pinyin e.g. 'zhao'
			if getAt(original, current+1) == 'H' {
				primary += "J"
				secondary += "J"
				current += 2
				break
			} else if matchStrings(original, current+1, []string{"ZO", "ZI", "ZA"}) ||
				(isSlavoGermanic(original) &&
					((current > 0) &&
						getAt(original, current-1) != 'T')) {
				primary += "S"
				secondary += "TS"
			} else {
				primary += "S"
				secondary += "S"
			}

			if getAt(original, current+1) == 'Z' {
				current += 2
			} else {
				current += 1
			}
			break

		default:
			current += 1
		}
		// printf("PRIMARY: %s\n", primary.str);
		//printf("SECONDARY: %s\n", secondary.str);
	}

	if len(primary) > maxLength {
		primary = primary[0:maxLength]
	}
	if len(secondary) > maxLength {
		secondary = secondary[0:maxLength]
	}

	return primary, secondary
}

// matchStrings checks if the string str, starting at the position start is one of the given options
func matchStrings(str string, start uint, options []string) bool {

	// If starting out of range
	if start >= uint(len(str)) {
		return false
	}

	// Shift the original string
	var shifted = str[start:]

	// For all options
	for _, option := range options {
		if strings.HasPrefix(shifted, option) {
			return true
		}
	}

	// Nothing found
	return false
}

// matchStrings checks if the string str, starting at the position start is one of the given options
func matchString(str string, start uint, option string) bool {

	// If starting out of range
	if start >= uint(len(str)) {
		return false
	}

	// Shift the original string
	var shifted = str[start:]

	return strings.HasPrefix(shifted, option)
}

// getAt return a single char from the given string. If the given position is out of range
// it returns 0.
func getAt(str string, pos uint) uint8 {

	// If getting out of range
	if pos >= uint(len(str)) {
		return 0
	}

	return str[pos]
}

func isVowel(str string, pos uint) bool {

	// If getting out of range
	if pos >= uint(len(str)) {
		return false
	}

	var c = str[pos]
	if (c == 'A') || (c == 'E') || (c == 'I') || (c == 'O') || (c == 'U') || (c == 'Y') {
		return true
	}

	return false
}

func isSlavoGermanic(str string) bool {

	// TODO check original source because the "W" already handle "WITZ"
	if strings.Contains(str, "W") ||
		strings.Contains(str, "K") ||
		strings.Contains(str, "CZ") ||
		strings.Contains(str, "WITZ") {
		return true
	}
	return false
}
