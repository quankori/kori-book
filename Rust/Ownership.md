# Rust Ownership

```rs
fn hold_my_vec<T>(_: &Vec<T>) {}

fn main() {
  let x = vec![1, 2, 3];
  hold_my_vec(&x); // <--- &x

  let z = x.get(0);
  println!("Got: {:?}", z);
}
```
