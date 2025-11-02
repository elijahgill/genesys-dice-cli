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

/* =================================================== */
/* ================== Custom Types =================== */
/* =================================================== */

type RollResult struct {
	success, failure, advantage, threat, triumph, despair int
}
// Tracks all axis of a Genesys roll. Can be used as a dice face, a starting point to manually modify a pool, or the result of a roll
// RollResult Methods
	// *Add(res RollResult) - Adds the values from res to the calling RollResult - modifies the calling RollResult
	// *Balance() - Cancels out all success, failures, threats and advantages - modifies the calling RollResult


type Die []RollResult
// Holds a RollResult for each face of the die
// Die Methods:
	// Roll() - Selects a random side of the die and returns it as a RollResult

// Dice pool type methods and usage
type DicePool struct {
	dice []Die
	result RollResult
}
// Contains an array of dice and a result. May be initialized with values in the RollResult for talents or items that add flat results like advantage or threat
// DicePool Methods
	// *AddDie() - Adds a dice to the pool by appending it to the existing slice - Modifies the calling DicePool
	// *Roll() - Rolls all of the dice, balances the result, and saves it in the RollResult of the DicePool - Modifies the calling DicePool
// DicePool Constructors
	// NewDicePool(pool string) - takes a string of dice runes (gybprk) and turns them into a pool

/* =================================================== */
/* ================== Global Vars ==================== */
/* =================================================== */

// Green - Ability - D8
var Ability Die=[]RollResult{
	{},			//1
	{success: 1},		//2
	{success: 1},		//3
	{success:2},		//4
	{advantage:1},		//5
	{advantage:1},		//6
	{success:1,advantage:1},//7
	{advantage:2},		//8
}

// Yellow - Proficiency - D12
var Proficiency Die =[]RollResult{
	{},			//1
	{success:1},		//2
	{success:1},		//3
	{success:2},		//4
	{success:2},		//5
	{advantage:1},		//6
	{success:1,advantage:1},//7
	{success:1,advantage:1},//8
	{success:1,advantage:1},//9
	{advantage:2},		//10
	{advantage:2},		//11
	{success:1,triumph:1},	//12
}

// Blue - Boost - D6
var Boost Die=[]RollResult {
	{},			//1
	{},			//2
	{success:1}, 		//3
	{success:2,advantage:1},//4
	{advantage:2},		//5
	{advantage:1},		//6
}

// Purple - Difficulty - D8
var Difficulty Die=[]RollResult {
	{},			//1
	{failure:1}, 		//2
	{failure:2},		//3
	{threat:1},		//4
	{threat:1},		//5
	{threat:1},		//6
	{threat:2},		//7
	{threat:1,failure:1},	//8
}

// Red - Challenge - D12
var Challenge Die=[]RollResult {
	{},			//1
	{failure:1},		//2
	{failure:1},		//3
	{failure:2},		//4
	{failure:2},		//5
	{threat:1},		//6
	{threat:1},		//7
	{threat:1,failure:1},	//8
	{threat:1,failure:1},	//9
	{threat:2},		//10
	{threat:2},		//11
	{despair:1,failure:1},	//12
}

// Black - Setback - D6
var Setback Die=[]RollResult{
	{},		//1
	{},		//2
	{failure:1},	//3
	{failure:1},	//4
	{threat:1},	//5
	{threat:1},	//6
}


// Set colors for dice for use within the module
var gColor = color.New(color.BgBlack,color.FgHiGreen)
var yColor = color.New(color.BgBlack,color.FgHiYellow)
var bColor = color.New(color.BgBlack,color.FgHiBlue)
var pColor = color.New(color.BgBlack,color.FgHiMagenta)
var rColor = color.New(color.BgBlack,color.FgHiRed)
var kColor = color.New()



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
	gColor.Println("g = Green ability die")
	yColor.Println("y = Yellow proficiency die")
	bColor.Println("b = Blue boost die")
	pColor.Println("p = Purple difficulty die")
	rColor.Println("r = Red challenge die")
	kColor.Println("k = Black setback die")
}

// Print a dice pool using the defined colors for each dice type
func PrintPrettyPool(pool string) {
	for _,dice:=range pool {
		switch dice {
		case 'g':
			gColor.Print("g")
		case 'y':
			yColor.Print("y")
		case 'b':
			bColor.Print("b")
		case 'p':
			pColor.Print("p")
		case 'r':
			rColor.Print("r")
		case 'k':
			kColor.Print("k")
		}
	}		
}

// Method to print the results of the roll, color coded, with a line for each type of result
// Does not print 0 values (e.g. Threat: 0 would be omitted)
func (res *RollResult) PrintResult() {

	if res.success > 0{
		gColor.Printf("Success: %d\n",res.success)
	}
	if res.failure > 0{
		pColor.Printf("Failure: %d\n",res.failure)
	}
	if res.advantage > 0 {
		bColor.Printf("Advantage: %d\n", res.advantage)
	}
	if res.threat > 0 {
		kColor.Printf("Threat: %d\n", res.threat)
	}
	if res.triumph > 0 {
		yColor.Printf("Triumph: %d\n", res.triumph)
	}
	if res.despair > 0 {
		rColor.Printf("Despair: %d\n", res.despair)
	}
}

/* =================================================== */
/* = Exported Methods and functions for custom types = */
/* =================================================== */

// ***** DicePool Methods *****

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

func NewDicePool(pool string) (DicePool,error) { 
	var newDicePool DicePool

	valid := validatePool(pool)
	if valid==false{
		return newDicePool, errors.New("Invalid dice pool. Valid options are gybprk")
	}
	for _,dice:=range pool {
		switch dice {
			case 'g': // Green - Ability
			newDicePool.AddDie(Ability)
			case 'y': // Yellow - Proficiency
			newDicePool.AddDie(Proficiency)
			case 'b': // Blue - Boost
			newDicePool.AddDie(Boost)
			case 'p': // Purple - Difficulty
			newDicePool.AddDie(Difficulty)
			case 'r': // Red - Challenge
			newDicePool.AddDie(Challenge)
			case 'k': //Black - Setback
			newDicePool.AddDie(Ability)
		}
	}
	return newDicePool,nil
}

func (pool *DicePool) AddDie(die Die) {
	pool.dice=append(pool.dice,die )	
}

func (pool *DicePool) Roll() {
	for _,die:=range pool.dice {
		pool.result.Add(die.Roll())	
	}
	pool.result.Balance()
}

// ***** Die Methods *****
func (die Die) Roll() RollResult {
	// Roll to get a number based on the size of the Die
	num := rand.Intn(len(die))
	var result = die[num]
	// Return that element from the Die
	return result
}

// ***** RollResult Methods *****
// Balance out the values of the results - Modifies the original RollResult
// Successes vs Failures, Threat vs Advantage
// Triumph and Despair do NOT cancel out
// The success and failure portions of these are added in the Die.Roll() function
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

// ***** String/Rune methods *****
// Rolls a single dice by color - if dice does not match a valid dice, it will be ignored, returning an empty result (all 0s)
// Dice results for a corresponding d6, d8, and d12  are defined in the Genesys CRB, page 10
func RollDiceColor(dice rune) RollResult {
	var result RollResult
	switch dice {
		case 'g': // Green - Ability
		result=Ability.Roll()
		case 'y': // Yellow - Profficiency
		result=Proficiency.Roll()
		case 'b': // Blue - Boost
		result=Boost.Roll()
		case 'p': // Purple - Difficulty
		result=Difficulty.Roll()
		case 'r': // Red - Challenge
		result=Challenge.Roll()
		case 'k': //Black - Setback
		result=Setback.Roll()
	}

	return result
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
			dicePool, err := NewDicePool(pool)
			if err != nil{
				fmt.Println("Invalid die pool. Valid values are:")
				PrintValidDice()
			}else{
				fmt.Print("Rolling ")
				PrintPrettyPool(pool)
				fmt.Print("...\n")
				dicePool.Roll()
				dicePool.result.PrintResult()
			}
		}
	}
}
