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