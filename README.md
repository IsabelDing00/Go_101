Go_101
====
### 区别print/printIn/printf
  - Println 打印的每一项之间都会有空行，Print 没有  
  - Println 会自动换行，Print 不会  
   fmt.Println("go","python","php","javascript") // go python php javascript  
   fmt.Print("go","python","php","javascript") // gopythonphpjavascript
   
  - Printf 是格式化输出，在很多场景下比 Println 更方便  
    ```
    func main() {
      a:=10
      b:=20
      c:=30
      fmt.Println("a=", a , ",b=" , b , ",c=" , c)
      fmt.Printf("a=%d,b=%d,c=%d" , a , b , c)
    }
    ```
    %d 是占位符，表示数字的十进制表示。Printf 中的占位符与后面的数字变量一一对应。更多的占位符参考：http://docscn.studygolang.com/pkg/fmt/


   
