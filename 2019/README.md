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

### Day 8: Space Image Format

It's one of the easiest ones (until now) and it basically consists on some loops and a bit of 
arrays / matrix manipulation. No tricky at all.

There were only a couple of things (from part two) that were tricky to understand (at first catch).

Firstly, I understood that the black "color" was the most "powerful" one, but not. 
You have to iterate over the layers from top to bottom until you find a non-transparent pixel. 

Secondly, when I was printing the obtained message from merging the given layers, I realised that 
white pixels were the ones with content, instead of what's the common sense on a RGB system
(where white means nothing). After that, I decided to  printed the white pixels ("X") and left the
others (black and transparent) with no content at all (" ") and I finally found it.

### Day 9: Sensor Boost

Another day with new instructions and modes for your processor (Intcode computer).

Both parts are technically the same, there's just a single difference: the input (the const, 
not the program) you have to configure. It seems (by the statement) that the execution with
the input of the part two should be a bit slow depending on the implementation, but it was 
not my case, despite of the ugly code I did 0:)

The statement samples and the feedback (output) from the executions are very intuitive (it
tells you what operations are not working properly), what makes it easier.

Finally, I'd suggest to use a map instead an array / slice to store the program instructions
due to the requirement that asks you to be able to access any (positive) memory position.

### Day 10: Monitoring Station  

Honestly, it was quite tricky, at least for me. I was a bit lost (I solved the part one
luckily) until I discovered I could solve it using the angles between the asteroids
(as suggested indirectly on the statement).

The idea is to use the reference asteroid as the (0,0) coordinates and use the values
from the other asteroids as vectors (in order to calculate their angles). Same angle
mean same sight line, so, taking that into account, the part one is easy.

On the part two, it becomes a bit trickier (specially if you didn't realise about the
angles approach before). For that part, uou can save the found asteroids for each angle,
sorted by distance. Additionally, you can sort the angles (from 0 to 360) and just loop
over them in order to find the one is evaporated the Xth.

### Day 11: Space Police

Another funny statement related with the processor (Intcode computer) built up during
the previous days, here you have to combine the processor with another component (I/O),
similarly to the day `7` challenge but with a robot instead of amplificators.

The part one was quite simple (IMHO), despite of my choice of storing each visited panel
on a map instead of drawing the entire matrix. Here, you just need to keep the robot position (x,y),
update it with the program and finally count the number of visited unique positions.

The part two was a bit more difficult, specially due to my initial decision of using
a map to store all the visited panels. Fortunately, despite of the map, I was storing
data enough to draw the visited panels an get the required password.

Additionally, I spent a funny development time with this challenge because I took the
change to rewrite entirely my Intcode computer in a better architectured application.

You can look at the solution from days `9` and `11` and compare it, if curious.

### Day 12: The N-Body Problem

Honestly, it was the first challenge (so far) that I needed external help to solve it.
However, the part one was very easy, as the total energy in the system is very simple 
to calculate and it can be solved with just few loops and some conditionals.

Regarding the part two, I supposed from the very beginning that the path to the solution 
would be somehow related with calculus of the LCM (least common divisor) of some values 
(no idea which exactly which ones). 

Unfortunately, I was not able (in a reasonable time) to determine exactly which calculus 
I had to do to get the result.

Finally, the strategy I followed (with external advisory) consisted on calculating the length
of the cycle (all moons again on the initial position) for each axis independently and then 
calculate the LCM of these values. 
That was all! Another lesson that makes me a better problem solver.

### Day 13: Care Package

Already as a tradition, it was another Intcode computer related challenge. Fortunately,
both parts can be solved by adapting a bit the solution from the day 11, replacing the robot
by an arcade cabinet which simulates a game.

The part one is almost equal to the part one from the day 11 but counting by type (blocks)
instead of counting the total number of tiles, so you can solve it with minor changes.

The part two is a bit trickier to understand (the game part), but also easy to solve. You
just need to determine the next movement of the joystick based on the paddle and the ball
locations. Then provide these movements as input to the Intcode computer program. That's all!

Finally, if you are inspired, you can refactor a bit the solution from the part two to let
the user know the status of the game simulated by the Intcode computer (print the tile map),
and then get the input joystick movements also from the user (standard input), so you'll have
a  playable game.

### Day 14: Space Stoichiometry

### Day 15: Oxygen System

### Day 16: Flawed Frequency Transmission

It was the classical problem that basically consists on doing basic calculus with some numbers.
Then those numbers becomes bigger on the part two, so you need to identify how you can improve
the performance of your first solution (unless you have done well at first).

In this case, the performance improvements can be done by (a) ignoring a slice of the input data
(as can be deduced from the message offset) and (b) simplifying the calculus done to obtain the output.

As expected (as I'm not very good at this kind of challenges), it took me too much time to discover
what kind of improvements I had to do to solve it. However, I finally could solve it. 