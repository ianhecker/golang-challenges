# Custom Sorting with Interfaces

Time Limit: 45 minutes

## Instructions

Implement custom sorting for a slice of structs in Golang by defining the
sort.Interface. For example, sort a slice of Person structs by one of the
fields, such as Age.

    Define a struct, e.g., Person, with at least two fields (e.g., Name and Age).
    Implement the sort.Interface by providing the methods:
        Len() int
        Less(i, j int) bool
        Swap(i, j int)
    Use the sort.Sort function from the standard library to sort the slice by a specific field.
    (Optional) Extend the challenge by sorting based on multiple criteria (e.g., by Age and then by Name).

Example Input and Output:
---
type Person struct {
    Name string
    Age  int
}

Input: []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Charlie", 35},
}
Output (sorted by Age): []Person{
    {"Bob", 25},
    {"Alice", 30},
    {"Charlie", 35},
}
---
Requirements:

    Ensure the implementation is efficient and handles various slice sizes.
    Use the sort.Interface to customize the sorting logic.
    Allow flexibility to switch sorting criteria (e.g., Name vs. Age) by modifying the Less() method.

Notes:

    Utilize the sort.Slice() function if you want a more concise implementation for simple cases.
    Experiment with sorting nested or more complex structs if time allows.
