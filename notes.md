

# Commands
`go mod init example.com/name` - creates a go.mod file to track all the code's dependencies.
`go run .` Runs the app stored in that folder
`go build` Compiles the created app into an .exe file for windows users to run it without needing go dependencies



# Notes

## Structs
- `structs` that start with **Uppercase** can be exported outside of their packages, while lowercases can only be used locally 
	- SAME APPLIES TO THEIR **FIELDS**!!
## Pointers
- `pointers` are good so you avoid creating duplicated variables, by passing pointers into functions, you are **literally pointing** to that position in memory to use those variables/structures instead of creating a new one.
- `Pointers to structs` are an exception where you *don't need to dereference a pointer* inside a function:
	- `(*structPtr).WantedProp`  - correct approach
	- `structPtr.WantedProp` - Allowed by GO

## Methods
- **Functions** attached to a **struct** is called a **method**.
	- AKA = A **function** that **only works/depends** on data that comes from a specific **struct**.
	- A good thing about methods is that you don't need to pass values into them, as they already expect to populate themselves when created.
	- When you **define methods** that should define/modify a struct, you **should not accept the struct input as a value**, because it **creates a copy** of that struct!
		- Solution: use `pointers` by adding * !
- Scan only works well for **single worded user input**.
	- use `bufio.NewReader()` function instead for longer inputs (with spaces)

## Functions: deep dive
- You can use functions has arguments in other functions, for exemple, creating a generic function to multiply numbers, create a function to double() the numbers and another to triple() them.
- Then just have a generic function that takes in the numbers array to transform and the function that will run the multiplication.
- Bonus:
	- Create a custom function type to avoid lengthy verbose functions by declaring them at the start of the code (before main)
	- `type transformFn func(int) int`
	- usage: `transform transformFn`
## Interfaces
- **Interfaces**: guarantees that a certain *value* has a certain *method* implemented.
	- Interfaces **don't define the logic** of the method, only that it exists.
	- If you define a method on your **interface**, it has to exist on the **structs** that will use your interface. 
- Good practices: IF an interface only has 1 method defined in it, usually we name the interface the same as the method in it, but add the suffix 'er' at the end
	- EX: `type saver interface { Save() error}`
## Arrays/lists
- When working with **Slices** in **arrays/lists** its important to know that **you don't copy the original array**, you create a **reference to that array in memory**, any data modified in the slice will be shown on the original array.
- When working with **cap()** function on **arrays**, you can find the capacity of the sliced array if you still have more elements to the **right side of the array**, but never to the **left side**.
	- `array[5]{1,2,3,4,5]`
	- `slicedArray := array[:1]`
	- `len(slicedArray) // 1`
	- `cap(slicedArray) //4`
- By default Go arrays are **not dynamic**, you need to use **slices** for that:
	- `arr := []float64` - example of dynamic array
	- `arr = append(arr,10)` - adding a new element to the array
	- `Result: [10]`
- Use **make** function whenever you know the final size of the array, since Go recreates the array/slice every time an item is added into it.

## Maps
- Its a data structure to group data together.
- Map is dividited into two parts = key:value
- The **main difference between Maps and structs** is that maps allow you to use any data struct as a key!
	- With structs you've got pre-defined data structures when you define it, you also **can't delete a key** pair 
	- Maps is better when you don't want to work with indexes but instead use custom labels!