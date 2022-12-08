## 2022 Problems

### Day 1: Calorie Counting

You need to accumulate the sum for every (N=3) lines, until next empty line, and keep a global maximum, which value is
the response for the first part. Then, for the second part you just need to keep up to three global maximum values.

### Day 2: Rock Paper Scissors

You need to build a conversion _map_ between opponent moves and the expected result and translate them to corresponding
score by following the conversion rules also defined in the statement. **Tip:** Use numeric "enums" to make score
calculations sugary. Then, for the second part you just need to adjust the conversion _map_.

### Day 3: Rucksack Reorganization

To extract common items in different compartments (part one) or rucksacks (part two) you can just sort the elements on
each and loop over them looking for matching items (or types). So, treating the input as arrays of chars (runes). Then,
to calculate the priority of each item, you can play with the ASCII value of each char.

### Day 4: Camp Cleanup

The key aspect to solve this exercise is that ranges of ids are continuous, which makes calculations of overlaps way
simpler than it could be. Therefore, you can just identify them with by conditionally checking for smaller/larger than
and slightly modifying the condition for the part two.
