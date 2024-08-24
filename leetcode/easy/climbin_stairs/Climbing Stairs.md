[[easy]]
https://leetcode.com/problems/climbing-stairs/

Мда, задание классное тем, что алгоритм в лоб не пройдет по времени выполнения. Я имею ввиду так:
```go
func climb(n int) int {
	if n <= 1 {
		return 1
	}
	return climb(n - 1) + climb(n - 2)
}
```
Ничего не напоминает `return`? Мне вот тоже не напоминал. А если так:
```cpp
int fibbonacci(int n) {
   if(n == 0){
      return 0;
   } else if(n == 1) {
      return 1;
   } else {
      return (fibbonacci(n-1) + fibbonacci(n-2));
   }
}
```
М, так это числа Фибоначчи. Получается, если перепишем рекурсивный поиск числа Фибоначчи на линейный, то все пройдет по времени! Вот так она и решается...