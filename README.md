# Benchmark

A straightforward benchmark function for measuring the execution time of synchronous tasks.

Example:

```go
func main() {
	_, timeTaken := Benchmark(func() {
		ls := &LinkedList{}
		ls.append(1)
		ls.append(2)
		ls.append(4)
		ls.append(5)
		ls.append(6)
		ls.append(7)

		current := ls.Head
		for current != nil {
			fmt.Println(current.Data)
			current = current.Next
		}
	})
	fmt.Println("Time taken to append 6 elements: ", timeTaken)
}
```
