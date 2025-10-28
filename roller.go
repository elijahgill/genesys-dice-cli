package main

import (
	"fmt"
	"math/rand"
	"strings"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"errors"
)

// Set colors for dice for use within the module
var g = color.New(color.BgBlack,color.FgHiGreen)
var y = color.New(color.BgBlack,color.FgHiYellow)
var b = color.New(color.BgBlack,color.FgHiBlue)
var p = color.New(color.BgBlack,color.FgHiMagenta)
var r = color.New(color.BgBlack,color.FgHiRed)
var k = color.New()

// TYPES
type RollResult struct {
	success, failure, advantage, threat, triumph, despair int
}


/* =================================================== */
/* ==== Printing functions for use in the CLI app ==== */
/* =================================================== */

func ClearScreen(){
	if runtime.GOOS == "windows"{
		cmd := exec.Command("cmd","/c","cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// Print all of the valid dice options with colors
func PrintValidDice(){
	g.Println("g = Green ability die")
	y.Println("y = Yellow proficiency die")
	b.Println("b = Blue boost die")
	p.Println("p = Purple difficulty die")
	r.Println("r = Red challenge die")
	k.Println("k = Black setback die")
}

// Print a dice pool using the defined colors for each dice type
func PrintPrettyPool(pool string) {
	for _,dice:=range pool {
		switch dice {
		case 'g':
			g.Print("g")
		case 'y':
			y.Print("y")
		case 'b':
			b.Print("b")
		case 'p':
			p.Print("p")
		case 'r':
			r.Print("r")
		case 'k':
			k.Print("k")
		}
	}		
}

// Method to print the results of the roll, color coded, with a line for each type of result
// Does not print 0 values (e.g. Threat: 0 would be omitted)
func (res *RollResult) PrintResult() {

	if res.success > 0{
		g.Printf("Success: %d\n",res.success)
	}
	if res.failure > 0{
		p.Printf("Failure: %d\n",res.failure)
	}
	if res.advantage > 0 {
		b.Printf("Advantage: %d\n", res.advantage)
	}
	if res.threat > 0 {
		k.Printf("Threat: %d\n", res.threat)
	}
	if res.triumph > 0 {
		y.Printf("Triumph: %d\n", res.triumph)
	}
	if res.despair > 0 {
		r.Printf("Despair: %d\n", res.despair)
	}
}

/* =================================================== */
/* = Exported Methods and functions for rolling dice = */
/* =================================================== */

// Rolls a single dice - if dice does not match a valid dice, it will be ignored, returning an empty result (all 0s)
// Dice results for a corresponding d6, d8, and d12  are defined in the Genesys CRB, page 10
func RollDiceColor(dice rune) RollResult {
	var result RollResult
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

// Balance out the values of the results - Modifies the original RollResult
// Successes vs Failures, Threat vs Advantage
// Triumph and Despair do NOT cancel out
// The success and failure portions of these are added in the roll function
func (res *RollResult) Balance() {
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

}

// Add the values of a RollResult to the existing RollResult
// This can be used to combine results of multiple dice to get the result for the entire pool
func (res *RollResult) Add(res2 RollResult) {

	res.success += res2.success
	res.advantage += res2.advantage
	res.failure += res2.failure
	res.threat += res2.threat
	res.triumph += res2.triumph
	res.despair += res2.despair
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

// Evaluate a string and return the TOTAL (unbalanced) result
// Throws error if pool contains invalid characters
func RollPool(pool string) (RollResult, error) {

	var total RollResult
	valid := validatePool(pool)
	if valid==false{
		return total, errors.New("Invalid dice pool. Valid options are gybprk")

	}
	for _,char:=range pool {
		total.Add(RollDiceColor(char))
	}		
	return total,nil
}

/* =================================================== */
/* =========== CLI Dice Roller - main loop =========== */
/* =================================================== */
// This can funciton as an example of how to use the functions, or can be used on its own as a fully functional dice roller
func main() {
	ClearScreen()

	fmt.Println("Genesys Dice Roller")
	fmt.Println(`Enter a dice pool in a single string, e.g. "ggypp".`)
	fmt.Println("Valid dice are:")
	PrintValidDice()

	var pool string

	for pool!="exit" {
		fmt.Print("Enter dice pool or \"exit\" to quit:")
		_,err := fmt.Scanln(&pool)
		if err != nil {
			fmt.Println(err)
		}
		ClearScreen()

		if pool!="exit" {
			testRoll, err := RollPool(pool)
			if err != nil{
				fmt.Println("Invalid die pool. Valid values are:")
				PrintValidDice()
			}else{
				fmt.Print("Rolling ")
				PrintPrettyPool(pool)
				fmt.Print("...\n")
				testRoll.Balance()
				testRoll.PrintResult()
			}
		}
	}
}
