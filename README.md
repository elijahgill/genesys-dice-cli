# Genesys Dice CLI
I made this as a way to learn Go, while also building something for my favorite TTRPG system.

## Usage
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
