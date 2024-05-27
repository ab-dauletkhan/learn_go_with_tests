package main

// func main() {
// 	// arrForSlice := [5]int{1, 2, 3, 4, 5}
// 	// sliceFromArr := arrForSlice[1:4]
// 	// fmt.Println("Slice from array:", sliceFromArr)
// 	// arrForSlice[3] = 9
// 	// fmt.Println(arrForSlice)
// 	// fmt.Println(sliceFromArr)
// 	// sliceFromArr[2] = 1
// 	// fmt.Println(arrForSlice)
// 	// fmt.Println(sliceFromArr)

// 	// fmt.Printf("%d - %d - %d\n", 5, 4)
// }

// import "fmt"

// func main() {
// 	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
// 	slice := arr[2:5] // slice contains {3, 4, 5}, len(slice) == 3, cap(slice) == 5

// 	fmt.Println("Original array:", arr)  // Outputs: [1 2 3 4 5 6 7]
// 	fmt.Println("Initial slice:", slice) // Outputs: [3 4 5]

// 	slice = append(slice, 8)                     // Still within capacity
// 	fmt.Println("After appending 8:", slice)     // Outputs: [3 4 5 8]
// 	fmt.Println("Array after appending 8:", arr) // Outputs: [1 2 3 4 5 8 7]

// 	slice = append(slice, 9)                     // Still within capacity
// 	fmt.Println("After appending 9:", slice)     // Outputs: [3 4 5 8 9]
// 	fmt.Println("Array after appending 9:", arr) // Outputs: [1 2 3 4 5 8 9]

// 	slice = append(slice, 10)                     // Exceeds capacity, new array allocation
// 	fmt.Println("After appending 10:", slice)     // Outputs: [3 4 5 8 9 10]
// 	fmt.Println("Array after appending 10:", arr) // Outputs: [1 2 3 4 5 8 9]
// }

import "fmt"

func main() {
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	slice1 := arr[2:5]    // slice1 contains {3, 4, 5}, len(slice1) == 3, cap(slice1) == 5
	slice2 := slice1[1:3] // slice2 contains {4, 5}, len(slice2) == 2, cap(slice2) == 4

	fmt.Println("Original array:", arr)    // Outputs: [1 2 3 4 5 6 7]
	fmt.Println("Initial slice1:", slice1) // Outputs: [3 4 5]
	fmt.Println("Initial slice2:", slice2) // Outputs: [4 5]

	slice2[0] = 10 // Modify slice2, which modifies the underlying array
	fmt.Println("After modifying slice2[0]:")
	fmt.Println("Array:", arr)     // Outputs: [1 2 3 10 5 6 7]
	fmt.Println("slice1:", slice1) // Outputs: [3 10 5]
	fmt.Println("slice2:", slice2) // Outputs: [10 5]

	slice1 = append(slice1, 8, 9) // Append to slice1, within capacity
	fmt.Println("After appending to slice1:")
	fmt.Println("Array:", arr)     // Outputs: [1 2 3 10 5 8 9]
	fmt.Println("slice1:", slice1) // Outputs: [3 10 5 8 9]
	fmt.Println("slice2:", slice2) // Outputs: [10 5]

	slice1 = append(slice1, 11) // Exceeds capacity, new array allocation
	fmt.Println("After exceeding capacity in slice1:")
	fmt.Println("Array:", arr)     // Outputs: [1 2 3 10 5 8 9]
	fmt.Println("slice1:", slice1) // Outputs: [3 10 5 8 9 11]
	fmt.Println("slice2:", slice2) // Outputs: [10 5]

	slice2 = append(slice2, 12) // Exceeds capacity, new array allocation
	fmt.Println("After appending to slice2:")
	fmt.Println("Array:", arr)     // Outputs: [1 2 3 10 5 8 9]
	fmt.Println("slice1:", slice1) // Outputs: [3 10 5 8 9 11]
	fmt.Println("slice2:", slice2) // Outputs: [10 5]

}
