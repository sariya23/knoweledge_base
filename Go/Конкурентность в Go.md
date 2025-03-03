[[Go]]
## Горутины
Одной из ключевых концепций в модели конкурентности Go являются **горутины**. Горутины - это абстракция на потоками ОС. Их основное преимущество в том, что они легкие и управление ими происходит на уровне программы, а не на уровне ОС. 
Чтобы запустить горутину, перед функцией нужно приписать ключевое слово `go`
`main` - это тоже горутина, главная. Когда она завершается - завершается все

## Каналы
Каналы - это средство общения горутин. 
Создание канала происходит с помощью функции `make`
```go
ch1 := make(chan int) // небуферезированный канал
ch2 := make(chan int, 10) // буферезированный канал
```
Каналы - это ссылочный тип. Поэтому так же как и в случае со слайсами, при передаче канала в функцию, ей передает указатель на канал. Нулевым значением для канала является `nil`

### Чтение, запись и буферизация
Взаимодействие с каналами возможно через символ `<-`. В зависимости от того, где находится канал от этого символ происходит либо чтение, либо запись
```go
ch := make(chan int)
ch <- 10 // запись
v := <- ch // чтение
```
Чтобы показать, что происходит только чтение из канала: `ch <- chan type`, чтобы показать, что только запись в канал: `ch chan <- type`

### Цикл `for-range` для каналов
В цикле можно получить значения из канала:
```go
for v := range ch {
	///
}
```

### Закрытие канала
Чтобы явно закрыть канал, нужно использовать функцию `close(chan)`. Запись в закрытый канал или повторное его закрытия приведут к панике. Но чтение из закрытого канала возможно. Если канал закрыт и без значений, то чтение из любого типа канала вернет нулевое значение. 
Чтобы определить, нулевое это значение или настоящее нужно как и в случае в мапами использовать идиому ok-запятой:
```
v, ok := <- ch
```
Если `ok=false`, значит нам вернулось нулевое значение типа.
Best practice - закрывает канал писатель.
### Различие поведение каналов

|          | unbuff., открытый                      | unbuff., закрытый           | buff. открытый         | buff., закрытый                                                    | `nil`                 |
| -------- | -------------------------------------- | --------------------------- | ---------------------- | ------------------------------------------------------------------ | --------------------- |
| Чтение   | Lock, пока не будет произведена запись | Возвращает нулевое значение | Lock, если буфер пуст  | Возвращает значение из буфера. Если пуст, то нулевое значение типа | Бесконечное зависание |
| Запись   | Lock, пока не будет чтение             | **ПАНИКА**                  | Lock, если буфер полон | **ПАНИКА**                                                         | Бесконечное зависание |
| Закрытие | Работает                               | **ПАНИКА**                  | Работает               | **ПАНИКА**                                                         | **ПАНИКА**            |
## Оператор `select`
Оператор `select` похож на `switch`:
```go
select {
	case v := <- ch1:
	case v := <- ch2:
	case ch3 <- x:
	case <- ch4:
	default:
	
}
```
Каждая ветвь пытается прочитать или записать данные в канал. Если это возможно, то выполняет чтение или запись и тело `case`. Если же выполняет чтение или запись в канал в разных ветках, то выберется случайная из них
Преимущество `select` в том, что он позволяет избегать deadlock - это когда ни одна из горутин не может продолжить работу, так как ожидает действий от другой. 
При доступе к каналу в разных горутинах должен соблюдаться порядок. Пример дедлока:
```go
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        v1 := 1
        ch1 <- v1
        v2 := <-ch2
        fmt.Println(v1, v2)
    }()
    v1 := 2
    ch2 <- v1
    v2 := <-ch1
    fmt.Println(v1, v2)
}
```
Мы кладем в `ch1` значение. Канал небуферезированный, значит выполнение горутины остановиться, пока из этого канала не прочитают значение. Дальше в `ch2` тоже кладется значение и ситуация та же - пока из него не прочитают, будет лок. И все - обе горутины ждут действий друг от друга и не могут ничего сделать. Это можно поправить:
```go
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        v1 := 1
        ch1 <- v1
        v2 := <-ch2
        fmt.Println(v1, v2)
    }()
    v1 := <-ch1
    v2 := 2
    ch2 <- v2
    fmt.Println(v1, v2)
```
Но лучше использовать `select`. Если мы обернем обращение к каналам в него, то дедлока не будет какой бы порядок не был:
```go
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        v1 := 1
        ch1 <- v1
        v2 := <-ch2
        fmt.Println(v1, v2)
    }()
    v1 := 10
    var v2 int
    select {
    case ch2 <- v1:
    case v2 = <-ch1:
    }
    fmt.Println(v1, v2)
```
После записи в `ch1` мы перекинемся в `select`, где произойдет чтение из него и никакого дедлока не будет
Часто `select` работает с несколькими каналами, поэтому его помещают в бесконечный `for` (зач, я так и не понял)

## Принципы и паттерны конкурентного программирования

### Горутины и изменяющиеся переменные
Раньше (до 1.21 или до 1.22)в переменная цикла создавалась один раз и ссылалась на один адрес в памяти, поэтому такой код:
```go
    arr := []int{2, 4, 6, 8, 10}
    ch := make(chan int, len(arr))
    for _, v := range arr {
        go func() {
            ch <- v * 2
        }()
    }
    for i := 0; i < len(arr); i++ {
        fmt.Println(<-ch)
    }
```
Выдавал одинаковые значение 20, так как горутины не успевали начать выполняться, а цикл уже закончился и в памяти лежало последнее значение из слайса. Сейчас такого вроде нет, но вот пути решения этого: надо либо передавать значение как параметр, либо затенять переменную в горутине
### Учетка горутин (goroutine leak)
Утечка горутин - это когда горутина продолжает существовать, но не выполняет никаких действий
```go
func main() {
    ch := make(chan int)
    go func() {
        for {
            select {
            case val := <-ch:
                fmt.Println("Received:", val)
            }
        }
    }()
    ch <- 42
    close(ch)
}
```
После закрытия канала, горутина все равно будет ждать новых значений

### Паттерн на основе канала `done`
С помощью этого паттерна можно сказать горутине, что можно прекращать выполнение
```go
func searchData(s string, searchers []func(string) []string) []string {
    done := make(chan struct{})
    result := make(chan []string)
    for _, v := range searchers {
        go func(searcher func(string) []string) {
            select {
            case result <- searcher(s):
            case <-done:
            }
        }(v)
    }
    r := <-result
    close(done)
    return r
}
```
Канал `done` имеет тип `struct{}` так как нам не важны значения. В этой функции мы ждем результата самый быстрой функции `searcher`. Когда произойдет чтение из канала `result`, то потом произойдет закрытие канала, а чтение из закрытого канала вернет нулевое значение
```go
func main() {
    searchers := []func(string) []string{s1, s2, sFast}
    done := make(chan struct{})
    result := make(chan []string)
    for _, searcher := range searchers {
        go func(searcher func(string) []string) {
            select {
            case result <- searcher("s"):
            case _, ok := <-done:
                fmt.Println(ok)
            }
            fmt.Println("горутина закончила работу")
        }(searcher)
    }
    r := <-result
    close(done)
    time.Sleep(2 * time.Second)
    fmt.Println(r)
}
```
### Прекращение выполнения горутины с помощью функции отмены
Можно возвращать замыкание, чтобы завершать горутины:
```go
func countToWithCancel(max int) (<-chan int, func()) {
    ch := make(chan int)
    done := make(chan struct{})
    cancel := func() {
        close(done)
    }
    go func() {
        for i := 0; i < max; i++ {
            select {
            case <-done:
                return
            default:
                ch <- i
            }
        }
        close(ch)
    }()
    return ch, cancel
}

func main() { 
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 { break } 
		fmt.Println(i) 
	} 
	cancel() 
}
```

## Когда использовать буферизированные и не буферизированные каналы
Буферизированные каналы стоит использовать тогда, когда вы знаете кол-во запущенных горутин и хотите ограничить кол-во горутин, которые еще будут запущены, или хотите ограничить объем работы, стоящей в очереди на выполнение.
Например:
```go
func processChan(ch chan int) []int {
    const conc = 10
    res := make(chan int, conc)
    for i := 0; i < conc; i++ {
        go func() {
            v := <-ch
            res <- process(v)
        }()
    }
    var out []int
    for i := range res {
        out = append(out, i)
    }
    return out
}
```

### Тайм-аут для кода
```go
func timeLimit() (int, error) {
    var result int
    done := make(chan struct{})
    go func() {
        result = fib(1000)
        close(done)
    }()
    select {
    case <-done:
        return result, nil
    case <-time.After(2 * time.Second):
        return 0, errors.New("so long")
    }
}
```
Для установки тайм аута можно использовать `select` в сочетании с паттерном канала `done`. 
Если канал закроется за определенное время, то из него можно будет прочитать значение (см таблицу выше) и горутина вернет результат. Если же нет, то из канала `time.After` чтение станет возможным и вернется ошибка
Лучше так не делать, так как это потенциально может привести к утечке горутин. Под капотом задается таймер и закроется он только тогда, когда дотикает (неактуально с 1.23).
## Использование типа `WaitGroup`
Когда мы ждем завершения одной горутины подойдет паттерн на основе канала `done`. Для ожидания завершения нескольких горутин уже надо использовать wait group. 
```go
func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    go func() {
        defer wg.Done()
        f1()
    }()
    go func() {
        defer wg.Done()
        f2()
    }()
    go func() {
        defer wg.Done()
        f3()
    }()
    wg.Wait()
}
```
Для использования группы нужно только объявить этот тип. Есть три метода:
- `Add` - инкрементирование кол-ва ожидаемых горутин
- `Done` - декрементирование кол-ва ожидаемых горутин
- `Wait` - приостанавливает выполнение своей горутины, пока счетчик ожидаемых горутин не станет равен 0

Мы не передает никуда группу, так как в этом случае передастся копия и декрементирование или инкрементирование будет срабатывать для копии, а не для оригинала

Более реальный пример:
```go
func proccesAndGather(arr []int, proccesor func(int) int) []int {
    var wg sync.WaitGroup
    out := make(chan int, len(arr))
    wg.Add(len(arr))
    for _, v := range arr {
        go func() {
            defer wg.Done()
            out <- proccesor(v)
        }()
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    res := make([]int, 0, len(arr))
    for v := range out {
        res = append(res, v)
    }
    return res
}
```
Здесь мы запускаем n-ое кол-во горутин для обработки значений, а дальше запускаем следящую горутину, в которой ожидаем завершение всех запущенных горутин.
Как бы это не было удобно не стоит считать этот метод основным способ синхронизации горутин. Если следует использовать в случае, когда за горутинами нужно "убрать" за всеми завершенными горутинами. Например, закрыть канал, в который они писали

### Однократное выполнение кода
Для того, чтобы выполнить какой-то код один раз, нужно использовать `sync.Once`. 
```go
func slow() {
    once.Do(func() {
        time.Sleep(5 * time.Second)
    })
    fmt.Println("end")
}

var once sync.Once

func main() {
    defer timer("main")() // 5
    slow()
    slow()
}
```
Единожды можно выполнить, например, какую-то долгую настройку. Once не стоит объявлять внутри функции, так как каждый раз он будет забывать, что уже выполнялся.

## Когда вместо каналов следует использовать мьютексы
Мьютекс накладывает ограничение на конкурентное выполнение кода или доступа к данным. Эта защищаемая часть программы называется *критической секцией*

Хотя каналы намного понятнее, чем мьютексы, за счет того, что сразу видно движение данных, так как доступ значение выполняет только в одной горутине, иногда стоит использовать их. Под иногда подходят случаи, когда горутины совместно считывают или записывают значение без обработки

В Go 2 реализации мьютексов в пакете `sync`. `Mutex` имеет два метода - `Lock` и `Unlock`. Метод `Lock` приостанавливает значение текущей горутины, если критическая секция занята другой горутиной. `Unlock` же снимает это ограничение. Мьютекс всегда нужно разлачивать, лучше в `defer`.
Вторая реализация - `RWMutex`, который лочит и чтение и запись
Схема выбора между каналами и мьютексами из книги "Конкурентность в Go":
- Если нужно координировать горутины или отслеживать значение по мере его преобразования с помощью нескольких горутин, то использовать лучше **каналы**
- Если нужно обеспечить совместный доступ к полю структуры, то использовать лучше **мьютексы**
- Если выявили критическую проблему с производительностью при использовании каналов и нет других способов решение, кроме как использовать **мьютекс** 

Также мьютекс нельзя залочить еще раз, если он уже залочен - это приведет к зависанию, так как горутина будет ждать пока блокировка снимется