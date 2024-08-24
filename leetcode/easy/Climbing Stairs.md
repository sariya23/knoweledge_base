[[easy]]
https://leetcode.com/problems/climbing-stairs/

```go
var cache = map[int]bool{}
func climbStairs(n int) int {
    if n <= 1 {
        return 1
    }
    if _, ok := cache[n]; ok {
        return n
    } else {
        newN := climbStairs(n - 1) + climbStairs(n - 2)
        cache[newN] = true
    }
    return climbStairs(n - 1) + climbStairs(n - 2)
}
```
