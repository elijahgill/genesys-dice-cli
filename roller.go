package main

import (
	"fmt"
	"math/rand"
	"strings"
	"github.com/fatih/color"
)

// Set colors for dice for use within the module
var g = color.New(color.BgBlack,color.FgHiGreen)
var y = color.New(color.BgBlack,color.FgHiYellow)
var b = color.New(color.BgBlack,color.FgHiBlue)
var p = color.New(color.BgBlack,color.FgHiMagenta)
var r = color.New(color.BgBlack,color.FgHiRed)
var k = color.New(color.BgWhite,color.FgBlack)

type rollResult struct {
	success, failure, advantage, threat, triumph, despair int
}


// TODO - new method prettyPool - add slices for each letter to make a pretty string with colors to print out: Rolling <pretty pool>
// TODO - On user input, can we change the color of the runes as they type? This would also be feedback on valid input

func rollDiceColor(dice rune) rollResult {
	var result rollResult
	switch dice {
		case 'g': // Green - Ability
		num := rand.Intn(8)
		if num==1||num==2 {
			result.success=1
		} else if num==3 {
			result.success=2
		} else if num==4||num==5 {
			result.advantage=1
		} else if num==6 {
			result.success=1
			result.advantage=1
		} else if num==7 {
			result.advantage=2
		}

		case 'y': // Yellow - Profficiency
		num := rand.Intn(12)
		if num==1||num==2 {
			result.success=1
		} else if num==3||num==4 {
			result.success=2
		} else if num==5 {
			result.advantage=1
		} else if num==6||num==7||num==8 {
			result.success=1
			result.advantage=1
		} else if num==9||num==10 {
			result.advantage=2
		} else if num==11 {
			result.triumph=1
			result.success=1
		}

		case 'b': // Blue - Boost
		num := rand.Intn(6)
		if num==2 {
			result.success=1
		} else if num==3 {
			result.success=2
			result.advantage=1
		} else if num==4 {
			result.advantage=2
		} else if num==5 {
			result.advantage=1
		}

		case 'p': // Purple - Difficulty
		num := rand.Intn(8)
		if num==1 {
			result.failure=1
		} else if num==2 {
			result.failure=2
		} else if num==3||num==4||num==5 {
			result.threat=1
		} else if num==6 {
			result.threat=2
		} else if num==7 {
			result.failure=1
			result.threat=1
		}

		case 'r': // Red - challenge
		num := rand.Intn(12)
		if num==1||num==2 {
			result.failure=1
		} else if num==3||num==4 {
			result.failure=2
		} else if num==5||num==6 {
			result.threat=1
		} else if num==7||num==8 {
			result.failure=1
			result.threat=1
		} else if num==9||num==10 {
			result.threat=2
		} else if num==11 {
			result.despair=1
			result.failure=1
		}

		case 'k': //Black - setback
		num := rand.Intn(6)
		if num==2||num==3 {
			result.failure=1
		} else if num==4||num==5 {
			result.threat=1
		}
	}

	return result
}

func (res *rollResult) Balance() {
	// Success and failure
	if res.success <= res.failure {
		res.failure -= res.success
		res.success = 0
	} else if res.success > res.failure {
		res.success -= res.failure
		res.failure = 0
	}

	// Threat and advantage
	if res.advantage <= res.threat {
		res.threat -= res.advantage
		res.advantage = 0
	} else if res.advantage > res.threat {
		res.advantage -= res.threat
		res.threat = 0
	}

	// Triumph and Despair
	if res.triumph <= res.despair {
		res.despair -= res.triumph
		res.triumph = 0
	} else if res.triumph > res.despair {
		res.triumph -= res.despair
		res.despair = 0
	}
}

func (res *rollResult) Add(res2 rollResult) {

	res.success += res2.success
	res.advantage += res2.advantage
	res.failure += res2.failure
	res.threat += res2.threat
	res.triumph += res2.triumph
	res.despair += res2.despair
}

func (res *rollResult) Format() string {

	return fmt.Sprintf("Success: %d \nFailure: %d \nAdvantage: %d \nThreat: %d \nTriumph: %d \nDespair: %d",res.success,res.failure,res.advantage,res.threat,res.triumph,res.despair)
}

func RollPool(pool string) rollResult {

	var total rollResult

	for _,char:=range pool {
		total.Add(rollDiceColor(char))
	}		

	return total
}

// Returns false if any invalid dice runes are included in the string
func validatePool(pool string) bool {

	validDice := "gybprk"
	for _,char:=range pool {
		//validate here
		if !strings.ContainsRune(validDice,char)  {
			return false
		}
	}

	return true
}

func main() {
	var pool string

	for pool!="exit" {
		fmt.Print("Enter dice pool or \"exit\" to quit:")
		_,err := fmt.Scanln(&pool)
		if err != nil {
			fmt.Println(err)
		}

		if pool!="exit" {
			if !validatePool(pool) {
				fmt.Println("Invalid die pool. Valid values are:")
				g.Println("g = Green ability die")
				y.Println("y = Yellow proficiency die")
				b.Println("b = Blue boost die")
				p.Println("p = Purple difficulty die")
				r.Println("r = Red challenge die")
				k.Println("k = Black setback die")
			} else {
				testRoll := RollPool(pool)
				testRoll.Balance()
				fmt.Println(testRoll)
				fmt.Println(testRoll.Format())
			}
		}
	}

}
