## 2020 Problems

### Day 1: Report Repair

It can be solved just by nesting loops, so basically checking the equality condition for each pair (or triplet) of
numbers. Of course, multiple optimizations can be applied, for example, by sorting the numbers in order to reduce
the total amount of pairs or triplets to check. However, there is no optimization really required to get the answer for
the given input in a reasonable time.

### Day 2: Password Philosophy

Extremely easy. Just a few of string manipulations and index accesses are required to solve it.

### Day 3: Toboggan Trajectory

It might be a bit tricky due to the infinite (undetermined) pattern repeats, but you can easily deal with it with a
simple modulus operator (%). For the second part, you can simply do the traverse as many times as slopes you have.
However, it can ideally be solved with a single traverse. To do that you can pre-calculate at which position the
tree should be to be reached for each slope. Then you can directly look up for those positions.

### Day 4: Passport Processing

There is no other complexity than keeping the amount of valid fields per password. However, the second part adds multiple
requirements into the equation, that is the reason why my code looks a bit ugly (didn't spend enough time to refactor it).
Ideally, you might read the input char by char as a micro-optimization. I used a set (`map[string]struct{}`) to keep
track of the required fields and then use it to check if all of them are present (equal length).

### Day 5: Binary Boarding

The challenge title might shed some light on what's the solution (or one of among all the existing ones, at least):
binary search. In order to find the id of your seat, there are also multiple valid strategies. The one I followed
consists on keeping the lowest and highest ids, store all of them in a set and finally loop from the lowest to the
the highest finding for the missing one.

## Day 6: Custom Customs

Easy one that can be solved by looping over the chars on each line and properly using maps to keep track of unique 
questions answered with "yes". So, to check whether everyone has answered "yes" to a given question or not, you can
simply check for the equality between the number of lines in the group and the amount of "yes" answers to a given
question.