## 2024 Problems

### Day 1: Historian Hysteria

Both the total distance and the similarity score between both lists is as simple to calculate as iterating over
both lists and applying the rules described in the problem statement. For the first one, you need to sort both lists,
for the second one you need to count occurrences (in a map).

### Day 2: Red-Nosed Reports

The first part is as simple as defining a `isSafe` function that applies the rules described in the problem statement.
The second one, might be a more tricky if you want to avoid brute force. Although brute force is acceptable because
the input is pretty small (most of the Reddit solutions are based on brute force), you can optimize the heuristics by
just re-checking whether the report is "safe" when it makes sense, so not trying all the possible combinations.