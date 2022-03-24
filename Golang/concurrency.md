# Golang Concurrency

- ### Gorountines

1. Goroutines cực kỳ nhẹ so với các luồng hệ điều hành.
2. Kích thước ngăn xếp rất nhỏ 2KB so với 8MB kích thước ngăn xếp cho các luồng hệ điều hành.
3. Context switching is very cheap as it happens in the user space.
4. Goroutines có rất ít state được lưu trữ.
5. Hàng trăm nghìn goroutines có thể được tạo ra trên một máy machine.

```go
func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("Hello, world")

	// goroutine function call
	go fun("goroutine-1")

	// goroutine with anonymous function
	go func() {
		fun("goroutine-2")
	}()

	// goroutine with function value call
	fv := fun
	go fv("goroutine-3")

	// wait for goroutines to end
	fmt.Println("waiting for goroutines to complete..")
	time.Sleep(1 * time.Second)

	fmt.Println("done..")
}
```
