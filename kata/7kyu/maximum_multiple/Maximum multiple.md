[[kata]]
[[7kyu]]

Сама ката - https://www.codewars.com/kata/5aba780a6a176b029800041c/go

## Условие
**_Given_** a **_Divisor and a Bound_** , _Find the largest integer N_ , Such That ,

# Conditions :

- **_N_** is _divisible by divisor_
    
- **_N_** is _less than or equal to bound_
    
- **_N_** is _greater than 0_.
    
    ---
    

# Notes

- The **_parameters (divisor, bound)_** passed to the function are _only positive values_ .
- _It's guaranteed that_ a **divisor is Found** .
    
    ---
    

# Input >> Output Examples

```
divisor = 2, bound = 7 ==> return (6)
```

## Explanation:

**_(6)_** is divisible by **_(2)_** , **_(6)_** is less than or equal to bound **_(7)_** , and **_(6)_** is > 0 .

---

```
divisor = 10, bound = 50 ==> return (50)
```

## Explanation:

**_(50)_** _is divisible by_ **_(10)_** , **_(50)_** is less than or equal to bound **_(50)_** , and **_(50)_** is > 0 .*

---

```
divisor = 37, bound = 200 ==> return (185)
```

## Explanation:

**_(185)_** is divisible by **_(37)_** , **_(185)_** is less than or equal to bound **_(200)_** , and **_(185)_** is > 0 .