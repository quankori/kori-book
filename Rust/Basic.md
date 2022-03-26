# Rust Basic

## Cargo

- Package manager của Rust

```bash
cargo new name_project
```

```bash
cargo init
```

- Run source

```bash
cargo run --release || cargo run
```

## Variable

- Nối chuỗi

```rs
let missiles = 8;
let ready = 2;
println!("Firing {} of my {} missiles...", ready, missiles);
```

- Mutable and Immutable

Mặc định giá trị khai báo của Rust là immutable, nên muốn gán biến khi khai báo ta phải chuyển qua mutable

```rs
fn main() {
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);
}
```

- Constants

```rs
const MAX_POINTS: u32 = 100_000;
```

- Shadowing

Khác với mutable ở chỗ có thể sử dụng lại tên và có thể thay đổi kiểu của biến

```rs
fn main() {
    let x = 5;

    let x = x + 1;

    let x = x * 2;

    println!("The value of x is: {}", x);
}
```

## Data types

Có 2 kiểu là vô hướng (scalar) và phức hợp (compound)

- Có 4 loại scalar chính là integer, float, boolean và ký tự

```rs
// Integer
let guess: u32 = 10; // i8 ... i128 ... isize || u8 ... u128 ... usize

// Float
let x = 2.0; //f64
let y: f32 = 3.0 //f32

// Boolean
let t = true;
let f: bool = false;

// Char
let c = 'z';
```

- Có 2 kiểu compund là array và tuple

```rs
// Tuple
fn tuple() {
    let x: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = x.0;

    let six_point_four = x.1;

    let one = x.2;
}

// Array
fn main() {
    let a = [1, 2, 3, 4, 5];

    let first = a[0];
    let second = a[1];
}
```

## Function

```rs
fn main() {
    another_function(5, 6);
}

fn another_function(x: i32, y: i32) -> i32 {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
    x
}
```

## Module

```rs
mod back_of_house {
    pub struct Breakfast {
        pub toast: String,
        seasonal_fruit: String,
    }

    impl Breakfast {
        pub fn summer(toast: &str) -> Breakfast {
            Breakfast {
                toast: String::from(toast),
                seasonal_fruit: String::from("peaches"),
            }
        }
    }
}

pub fn eat_at_restaurant() {
    // Order a breakfast in the summer with Rye toast
    let mut meal = back_of_house::Breakfast::summer("Rye");
    // Change our mind about what bread we'd like
    meal.toast = String::from("Wheat");
    println!("I'd like {} toast please", meal.toast);

    // The next line won't compile if we uncomment it; we're not allowed
    // to see or modify the seasonal fruit that comes with the meal
    // meal.seasonal_fruit = String::from("blueberries");
}

mod back_of_house {
    pub enum Appetizer {
        Soup,
        Salad,
    }
}

pub fn eat_at_restaurant() {
    let order1 = back_of_house::Appetizer::Soup;
    let order2 = back_of_house::Appetizer::Salad;
}
```

## Control flow

```rs
// If else 
fn main() {
    let number = 3;

    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }
}

// If else 2
let number = if condition { 5 } else { 6 };

// For
fn main() {
    for number in (1..4).rev() {
        println!("{}!", number);
    }
    println!("LIFTOFF!!!");
}

// While
while index < 5 {
    println!("the value is: {}", a[index]);

    index += 1;
}

// Loop
let mut counter = 0;

let result = loop {
    counter += 1;

    if counter == 10 {
        break counter * 2;
    }
};

println!("The result is {}", result);
```

## Structs

## Traits

## Collections

## Enums

## Closures

## Threads

