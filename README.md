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

- *Add(res RollResult) - Adds the values from res to the calling RollResult - modifies the calling RollResult
- *Balance() - Cancels out all success, failures, threats and advantages - modifies the calling RollResult

### Die

```
type Die []RollResult
```

Holds a RollResult for each face of the die

**Methods**

- Roll() - Selects a random side of the die and returns it as a RollResult

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
- NewDicePool(pool string) - takes a string of dice runes (gybprk) and turns them into a pool

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
