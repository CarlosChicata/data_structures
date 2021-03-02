# Set data structure
Contain my implementation of several type of sets to learn golang.

# available implementations 
Those are types of sets will implement
- [X] Unique elements in Set (set)
- [ ] Non-unique elements in set (dst)
- [ ] Disjoint-union set (dus)

# Methods
All lists contain those methods in common.

- **add**: add element in set if not exists.
- **remove**: remove element in set if exists.
- **intersect**: all elements will intersect in two sets.
- **belong**: element belong this set.
- **difference**: all element will not intersect in two sets.
- **len**: count all elements of universe set.
- **display**: display all elements in set.

But disjoint-union version, i need to add some methods.

- **union**: merge two set
- **unionIn**: merge element into set
- **parentIn**: get parent of set
- **sizeIn**: get element inside set

# Optimization of Disjoint-union set
I can implement several methods to optimize this data structure:

- [ ] union by size: use the parent by set with more member.
- [ ] union by rank: use the parent by set with more depth.
- [ ] path compression: when insert element in set, directly point to parent.
- [ ] path halving: ??
- [ ] path splitting: ?? 


# Table of implemented methods
Contain counmmon methods implemented in those sets.

| Methods | set | dst | dus |
| ---- | ---- | ---- | ---- |
| add | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| remove | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| intersect | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| belong | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| difference | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| len| :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| insert | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |
| display | :heavy_check_mark: | :white_check_mark: | :white_check_mark: |

# Table of time complexity - worst case
Contain all time complexity of implementations

| Methods | set | dst | dus |
| ---- | ---- | ---- | ---- |
| add        | O(n)  |  |  |
| remove     | O(n^2) |  |  |
| intersect  | O(n^2) |  |  |
| belong     | O(n) |  |  |
| difference | O(n^2) |  |  |
| len        | O(1) |  |  |
| insert     | O(1) |  |  |
| display    | O(n) |  |  |

## version: 0.1.0
