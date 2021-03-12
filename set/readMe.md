# Set data structure
Contain my implementation of several type of sets to learn golang.

# available implementations 
Those are types of sets will implement
- [X] Unique elements in Set (set)
- [X] Non-unique elements in set (dst)
- [X] Disjoint-union set (dus)

# Methods
All lists contain those methods in common.

- **add**: add element in set if not exists.
- **remove**: remove element in set if exists.
- **intersect**: all elements will intersect in two sets.
- **belong**: element belong this set.
- **difference**: all element will not intersect in two sets.
- **len**: count all elements of universe set.
- **display**: display all elements in set.

But disjoint-union data structure, i need to others methods.

- **add**: add element in set if not exists.
- **belong**: element belong this set.
- **len**: count all elements of universe set.
- **display**: display all elements in set.
- **union**: merge two set
- **parentIn**: get parent of set
- **sizeIn**: get element inside set

# Optimization of Disjoint-union set
I can implement several methods to optimize this data structure:

- [x] union by size: use the parent with more member.
- [x] union by rank: use the parent with more depth.
- [x] path compression: when insert element in set, directly point to parent.
- [ ] path halving: ??
- [ ] path splitting: ?? 



# Table of implemented methods
Contain counmmon methods implemented in those sets.

| Methods | set | dst | dus |
| ---- | ---- | ---- | ---- |
| add | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| remove | :heavy_check_mark: | :heavy_check_mark: | :x: |
| intersect | :heavy_check_mark: | :heavy_check_mark: | :x: |
| belong | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| difference | :heavy_check_mark: | :heavy_check_mark: | :x: |
| len| :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: |
| insert | :heavy_check_mark: | :heavy_check_mark: | :x: |
| display | :heavy_check_mark: | :heavy_check_mark: | :white_check_mark: |
| union | :x: | :x: | :heavy_check_mark: |
| sizeIn | :x: | :x: | :heavy_check_mark: |
| sizeSet | :x: | :x: | :heavy_check_mark: |
| parentIn | :x: | :x: | :heavy_check_mark: |

# Table of time complexity - worst case
Contain all time complexity of implementations

| Methods | set | dst | dus |
| ---- | ---- | ---- | ---- |
| add        | O(n)  | O(1) | O(1) |
| remove     | O(n) | O(n) | - |
| intersect  | O(n^2) | O(n^2) | - |
| belong     | O(n) | O(n) | O(1) |
| difference | O(n^2) | O(n^2) | - |
| len        | O(1) | O(1) | O(1) |
| display    | O(n) | O(n) |  |
| union      | - | - | O(1) |
| sizeIn     | - | - |  |
| sizeSet    | - | - | O(1) |
| parentIn   | - | - | O(n) |

## version: 0.1.0
