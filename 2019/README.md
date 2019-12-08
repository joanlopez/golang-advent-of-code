## 2019 Problems

### Day 1: The Tyranny of the Rocket Equation

Just simple math operations and loops. To warm up.

It basically consists on calculating the needed fuel applying a given formula with the given mass.

On the part two, the fuel is also considered as extra mass, so you need to recalculate the
total needed fuel many times until the final amount can cover the given mass and all the fuel itself.   

### Day 2: 1202 Program Alarm

Kind of a processor implementation with just three available operations (`ADD`, `MUL`, and `HALT`).

You have to be able to run the given sequence of instructions and tell the final value at some "memory" position.

On the part two, you need to find the operators that makes the final value equals to the given one.
In order to make it easier, I studied how each operator modifies the final value in order to
reduce the total amount of inputs to a very small range. 

### Day 3: Crossed Wires

It was quite funny. Drawing maps (matrix) is always funny.

My first approach was to draw a matrix for each wire and another one for the superposition of both.
However, it took a lot of time to run when the matrix is very big, so I decided to just store the positions.
Then I looped over all the positions of each wires in order to find all the matching positions and
to determine which one of these is the searched one (based on the given condition). 

On the part two, the condition to determine which matching position is the best one is changed, but
it's still almost equal to the first one (nearer to vs shorter wire path).   

### Day 4: Secure Container

I expected to need to discover something tricky but, as you can see on the solution, the problem
can be solved just by defining a bunch of IF statements. The code is ugly, I know it, but it worked.

Nothing else, a bit disappointing.

### Day 5: Sunny with a Chance of Asteroids

It consists on extending the little processor implemented during the day 2.

Basically, you have to add support for input and output operations. Additionally,
the operands of the given operations can work in two modes: a) positional (like memory) and
b) immediate (directly the given value). However, the operand that determines where
the result of an operation (i.e. `ADD`) is stored, works always (evidently) as positional. 

On the part two, you have to add support for more operations (jump if true, jump if false,
less than, equal). The test is exactly the same but with support for those operations.

### Day 6: Universal Orbit Map

It's a common graphs problem (BFS + Dijkstra). As *always* you can solve it in many different ways.

On the part one, you have to find how many nodes are reachable between them. For that, I looped over
them and used a BFS to determine if each pair is reachable. 
Probably it might be improved but it worked (a bit slow).

On the part two, you have to find the shortest path between a node YOU are orbiting, and a node
SANta is orbiting, so you can just run the Dijkstra algorithm between this few pairs and take
the minimum result.  

### Day 7: Amplification Circuit

Again, another problem that consists on working with the processor developed during the days 2 and 5.

Now you can implement some amplifiers that contains a running program inside it.

The key of this problem is understand properly how each amplifier works and how they're communicated
through the input / output operations. It took me a lot of time to really understand how it works,
specially for the part two (where it's more relevant for the feedback loop mode).

Once I took it, then it was quite easy. Be careful with misunderstandings, IMHO it can be tricky.
If you have a look at the leaderboard [here](https://adventofcode.com/2019/leaderboard/day/7),
you can see how big is the difference between the one star results from the two stars ones. So,
it seems even the experts suffered something similar.