# Golang Quiz

## What is Golang?

Go is an opened programming language, which helps us easily in developing simple, reliable and efficient applications.
Applications are builed from packages that have attributes allows managing dependencies effectively.

## Workspace in Golang?

A workspace is a hierarchical folder system with three roots folders.

1. `src` which contains root GO files is organized to packges.
2. `pkg` contains package objects.
3. `bin` contains executed commands.

## How you can do testing in Golang?

There is a framework to do testing. it includes Go test and testing packages.

To write a test, you need to create a file that has name ended by \_testing. It contains functions named TESTXXX with function signature ( t \* tests. T) . Framework testing run one by one feature.

## What are the advantages of Golang?

1. Go compiles very fast.
2. Go supports concurrency in level language.
3. Functions are first objects in Go.
4. Go has garbage collection
5. String maps are integrated."

## How arrays in GO works differently then C?

1. Array are value, which are assigned
2. If array is transfered to function, It will received the clone of array, not a pointer.
3. The size of an array is a part of its type. Type [10] int and [20] are different.

## Khác nhau giữa Array và Slices?

Sự khác nhau giữa slices và array:

1. **Array** có kích thước cố định (fixed size) và phần tử phải cùng loại dữ liệu, còn **Slices** có kích thước động.
2. **Array** dạng tham trị (value types) khi gán biến mới sẽ tạo một array khác tốn bộ nhớ hơn, còn **Slices** dạng tham chiếu tới array và có thể tạo ra từ một **Array**
3. Do **Slice** chỉ là tham chiếu đến **Array**, do đó thay đổi giá trị của **Slice** sẽ làm thay đổi giá trị của **Array** mà nó tham chiếu đến. Nếu có nhiều **Slice** cùng tham chiếu đến một **Array** thì khi thay đổi giá trị một **Slice** có thể làm thay đổi giá trị các **Slice** khác

## Length và Capacity trong Golang?

Một Slice sẽ có 2 thuộc tính là length (len) và capacity (cap).

- Length là số phần tử chứa trong Slice
- Capacity là số phần tử chứa trong Array mà Slice tham chiếu đến (bắt đầu tính từ phần tử đầu tiên của Slice)
- Mặc định len là 0 và với cap là độ dài của mảng
