# Object-oriented programming

## Kế thừa (Inheritance)

  - Sử dụng lại hàm của lớp khác thông qua kế thừa lớp đó.

    ```java
    class Animal {
    void eat() {
        System.out.println("eating...");
        }
    }

    class Dog extends Animal {
        void bark() {
            System.out.println("barking...");
        }
    }

    public class TestInheritance1 {
        public static void main(String args[]) {
            Dog d = new Dog();
            d.bark();
            d.eat();
        }
    }
    ```

## Đa hình (Polymorphism)

  - Đa hình lúc runtime là quá trình gọi phương thức đã được ghi đè trong thời gian thực thi chương trình. Trong quá trình này, một phương thức được ghi đè được gọi thông qua biến tham chiếu của một lớp cha.

    ```java
    class Bike {
        void run() {
            System.out.println("running");
        }
    }

    public class Splender extends Bike {
        void run() {
            System.out.println("running safely with 60km");
        }

        public static void main(String args[]) {
            Bike b = new Splender();
            b.run();
        }
    }
    ```

## Trừu tượng (Abstraction)

  - Tính trừu tượng là một tiến trình ẩn các cài đặt chi tiết và chỉ hiển thị tính năng tới người dùng.
  - Có 2 cách dùng là interface và abstract.

    - Abstract

    ```java
    abstract class Bike{
        abstract void run();
    }

    class Honda4 extends Bike{
        void run() {
            System.out.println("running safely..");
        }

        public static void main(String args[]) {
            Bike obj = new Honda4();
            obj.run();
        }
    }
    ```

    - Interface
    
    ```java
    interface printable {
        void print();
    }

    class A6 implements printable {
        public void print() {
            System.out.println("Hello");
        }

        public static void main(String args[]){
            A6 obj = new A6();
            obj.print();
        }
    }
    ```

## Đóng gói (Encapsulation)

  - Bạn có thể tạo lớp read-only hoặc write-only bằng việc cài đặt phương thức setter hoặc getter.

  - Bạn có thể kiểm soát đối với dữ liệu. Giả sử bạn muốn đặt giá trị của id chỉ lớn hơn 100 bạn có thể viết logic bên trong lớp setter.

    ```java
    public class Student {
        private String name;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }
    }

    class Test {
        public static void main(String[] args) {
            Student s = new Student();
            s.setName("Hai");
            System.out.println(s.getName());
        }
    }
    ```
