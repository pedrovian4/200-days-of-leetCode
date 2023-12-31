# Simple Sorting Solution

## Idea:
The simplest way to find the top `k` frequent elements is to count the occurrences of each element, sort them by frequency, and then return the `k` elements with the highest frequency.

## Steps:
1. **Count Frequencies**: Iterate through the array and use a map to count the occurrences of each element.
2. **Sort by Frequency**: Convert the frequency map into a slice, then sort this slice based on the frequency.
3. **Extract Top K**: After sorting, the first `k` elements of this slice are the top `k` frequent elements.


# Min-Heap Solution

## Idea 

For a more efficient solution, especially when k is small compared to the number of unique elements, a min-heap can be used to maintain the top k frequent elements as the array is processed.

## Steps
1) **Count Frequencies**: Similar to the simple solution, count the occurrences of each element.
2)  **Use Min-Heap**: Create a min-heap to keep track of the top `k` elements
      - The heap contains pairs of elements and their frequencies
      - if the heap size exceeds `k` remove the smallest element
3) **Build Result**: Once all elements are processed, extract elements from the min-heap to build the result array.

# How MinHeap Works:
The MinHeap is a binary tree where each parent node is less than or equal to its child nodes. In the context of this problem, the "less than" condition is based on the frequency of elements. The root of the MinHeap is the element with the smallest frequency. When elements are added or removed, the tree restructures itself to maintain this property.

- **Push/Insert**: When a new element is added, it's initially placed at the bottom of the heap. Then, the heap "bubbles up" this element, swapping it with its parent until the parent is smaller or the element reaches the root.

- **Pop/Remove**: When the root is removed (the smallest element), the heap replaces it with the last element and then "bubbles down" this new root, swapping it with the smaller of its children until the children are larger or the element reaches a leaf.
