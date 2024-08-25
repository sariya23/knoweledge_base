[[easy]]

https://leetcode.com/problems/find-target-indices-after-sorting-array/description/

```go
import "slices"

func targetIndices(nums []int, target int) []int {
    slices.Sort(nums)
    var indicies []int
    for i, v := range nums {
        if v == target {
            indicies = append(indicies, i)
        }
    }
    return indicies
}
```
Решение не очень хорошее. Не понял куда тут можно прикрутить бинарный поиск
Вот решение без сортировки, а значит `O(n)`
```go
func f(nums []int, target int) []int {
    targetCount := 0
    smallerCoount := 0
    for _, v := range nums {
        if v == target {
            targetCount++
        } else if v < target {
            smallerCoount++
        }
    }

    var res []int
    for i := 0; i < targetCount; i++ {
        res = append(res, smallerCoount)
        smallerCoount++
    }
    return res
}
```
Тут вообще гениальная идея. Чтобы найти первое вхождение целевого значения в отсортированном списке, нам нужно посчитать сколько элементов меньше него. Это реально круто.
Ну вот например: `4 2 1 3 3 5`. Отсортированная версия: `1 2 3 3 4 5`. Меньше пяти 5 элементов и у пятерки как раз 5 индекс. БУМ! Это первая часть функции. 
Вторая часть еще круче. Мы знаем индекс первого вхождения элемента. И мы знаем сколько целевых элементов. Так как список отсортирован следующее вхождение элемента будет просто следующим. То есть нам нужно прокрутится в цикле столько раз, сколько у нас элементов. А добавлять мы будем увеличенный индекс первого вхождения на один каждую итерацию. Гениально!