# Word Frequency Counter

Time Limit: 40 minutes

## Instructions

Create a function in Golang that analyzes a given text (string) and returns a
map where keys are words and values are their frequency counts.

Normalize the input by converting all letters to lowercase.
Remove punctuation and handle special characters (e.g., symbols and line breaks).
Efficiently update word counts using a map.

Example Input and Output:
---
Input: "Hello, world! Hello."
Output: map["hello": 2, "world": 1]

Input: "Go is great! Go, go, go!"
Output: map["go": 4, "is": 1, "great": 1]
---

Notes:
1) Use the strings package to help with string manipulation tasks.
2) The function should handle large texts efficiently, with minimal performance overhead.
3) Treat words separated by whitespace or punctuation as distinct words.
