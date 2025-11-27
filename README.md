# Genesys Dice CLI

I made this as a way to learn Go, while also building something for my favorite TTRPG system.

## Exported types

### RollResult

```
type RollResult struct {
    success, failure, advantage, threat, triumph, despair int
}
```

Tracks all axis of a Genesys roll. Can be used as a dice face, a starting point to manually modify a pool, or the result of a roll

**Methods**

- *Add(res RollResult) - Combine RollResults. Adds the values from res to the calling RollResult - modifies the calling RollResult
- *Balance() - Cancels out all success, failures, threats and advantages - modifies the calling RollResult
- *PrintResult() Print the results of the roll, color coded, with a line for each type of result

### Die

```
type Die []RollResult
```

Holds a RollResult for each face of the die

**Methods**

- Roll() - Selects a random side of the die and returns it as a RollResult

**Example Usage**
Standard dice types are already defined. You can add a custom dice type if desired.

Dice types are defined based on the conversion tables for using standard numeric dice in the CRB:
```
// Green - Ability - D8
var Ability Die=[]RollResult{
	{},
	{success: 1},
	{success: 1},
	{success:2},
	{advantage:1},
	{advantage:1},
	{success:1,advantage:1},
	{advantage:2},
}

// Yellow - Proficiency - D12
var Proficiency Die =[]RollResult{
	{},
	{success:1},
	{success:1},
	{success:2},
	{success:2},
	{advantage:1},
	{success:1,advantage:1},
	{success:1,advantage:1},
	{success:1,advantage:1},
	{advantage:2},
	{advantage:2},
	{success:1,triumph:1},
}

// Blue - Boost - D6
var Boost Die=[]RollResult {
	{},
	{},
	{success:1},
	{success:2,advantage:1},
	{advantage:2},
	{advantage:1},
}

// Purple - Difficulty - D8
var Difficulty Die=[]RollResult {
	{},
	{failure:1},
	{failure:2},
	{threat:1},
	{threat:1},
	{threat:1},
	{threat:2},
	{threat:1,failure:1},
}

// Red - Challenge - D12
var Challenge Die=[]RollResult {
	{},
	{failure:1},
	{failure:1},
	{failure:2},
	{failure:2},
	{threat:1},
	{threat:1},
	{threat:1,failure:1},
	{threat:1,failure:1},
	{threat:2},
	{threat:2},
	{despair:1,failure:1},
}

// Black - Setback - D6
var Setback Die=[]RollResult{
	{},
	{},
	{failure:1},
	{failure:1},
	{threat:1},
	{threat:1},
}
```

#### DicePool

```
type DicePool struct {
    dice []Die
    result RollResult
}
```

Contains an array of dice and a result. May be initialized with values in the RollResult for talents or items that add flat results like advantage or threat

**Methods**
- *AddDie() - Adds a dice to the pool by appending it to the existing slice - Modifies the calling DicePool
- *Roll() - Rolls all of the dice, balances the result, and saves it in the RollResult of the DicePool - Modifies the calling DicePool
  
**Constructors***
- NewDicePool(pool string) - takes a string of dice runes (gybprk) and turns them into a pool. Returns DicePool,error

**Example Usage**
```
dicePool, err := NewDicePool("ggypp")
dicePool.result.advantage += 1 // Add an advantage to the pool manually (for item quality, talent, etc.)
dicePool.Roll() // Will roll each dice and add all of the roll results to the dicePool
dicePool.Balance() // Cancel out all threat/advantage and success/failure.
dicePool.result.PrintResult() // Formats and prints out the result
```

## Command Line App

Running the app in the CLI, you are presented with instructions on how to use the program.

```
Genesys Dice Roller
Enter a dice pool in a single string, e.g. "ggypp".
Valid dice are:
g = Green ability die
y = Yellow proficiency die
b = Blue boost die
p = Purple difficulty die
r = Red challenge die
k = Black setback die
Enter dice pool or "exit" to quit:
```

The program will roll the dice as simple d6, d8, or d12, and calculate the results based on the conversion table in the Genesys CRB (page 10).

The results will be balanced before printing it out and prompting you for another pool.

```
Rolling ggypp...
Success: 1
Advantage: 1
Enter dice pool or "exit" to quit:
```

If any invalid dice are entered in the pool, the system will repeat the instructions:

```
Invalid die pool. Valid values are:
g = Green ability die
y = Yellow proficiency die
b = Blue boost die
p = Purple difficulty die
r = Red challenge die
k = Black setback die
Enter dice pool or "exit" to quit:
```

You can enter "exit" instead of a dice pool to quit the program
