# Sorting algo

## Basic

### Bubble Sort

Swap if in wrong order
for i:0 -> n-1; for j:0 -> n-i-1

This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high

- best: Ω(n), sorted already
- average: θ(n^2)
- worst: O(n^2)
- worst(space): O(1)

### Insert Sort

Swap if in wrong order until it's at right order position
for i:1 -> n-1; for j:i-1 -> 0

Insertion sort is efficient for small data values

- best: Ω(n), sorted already
- average: θ(n^2)
- worst: O(n^2)
- worst(space): O(1)

### Selectiong Sort

find minimun number each iteration and put it in first place
for i:0 -> n-1; for j:i -> n-1

- best: Ω(n^2)
- average: θ(n^2)
- worst: O(n^2)
- worst(space): O(1)

## Advanced

### Merge Sort

Based on Divide and Conquer method. Sorting in smaller list and Merge back.

- best: Ω(n log(n))
- average: θ(n log(n))
- worst: O(n log(n))
- worst(space): O(n)

### Quick Sort

Like Merge Sort, but choose a pivot for partition.

How to choose pivot:

- Always pick the first element as a pivot.
- Always pick the last element as a pivot (implemented below)
- Pick a random element as a pivot.
- Pick median as the pivot.

- best: Ω(n log(n))
- average: θ(n log(n))
- worst: O(n^2), pivot on the left/right most position
- worst(space): O(n)

### Bucket Sort

Create a buckets for range of value. Put the elements into in the range of bucket. Do Insertion sort to each bucket. Concatenate all buckets

- best: Ω(n +k)
- average: θ(n +k)
- worst: O(n^2), All on the same bucket
- worst(space): O(n)

### Heap Sort

- best: Ω(n log(n))
- average: θ(n log(n))
- worst: O(n log(n))
- worst(space): O(1)
