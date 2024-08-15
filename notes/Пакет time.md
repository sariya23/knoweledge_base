
Статья - https://proglib.io/p/samouchitel-po-go-dlya-nachinayushchih-chast-13-rabota-s-datoy-i-vremenem-paket-time-2024-06-26
## Пакет time. Общее
В Go время хранится как кол-во секунд прошедших с начала эпохи UNIX - 01.01.1970.
Для получения текущего локального времени:
```go
fmt.Println(time.Now()) // 2024-08-15 11:59:37.588920427 +0300 MSK m=+0.000032257
```
- дата + время `2024-08-15 11:59:37.588920427`
- смещение относительно UTC `+0300`
- Часовой пояс `MSK`
- Вспомогательная информация `m=+0.000032257`

Также можно выводить информацию в определенном формате:
```go
    fmt.Println(time.Now().Unix()) // 1723712585
    fmt.Println(time.Now().UTC()) // 2024-08-15 09:03:05.843144588 +0000 UTC
```
Каждый объект `Time` связан с определенным местоположением. Мы можем преобразовывать время к какой-то локации:
```go
    location, err := time.LoadLocation("Europe/Samara")
    if err != nil {
        panic(err)
    }
	fmt.Println(time.Now()) // 2024-08-15 12:07:49.460715959 +0300 MSK m=+0.000114935
    fmt.Println(time.Now().In(location).Location()) // Europe/Samara
    fmt.Println(time.Now().In(location)) // 2024-08-15 13:07:49.460822072 +0400 +04
```

## Компоненты времени
Доступны следующие компоненты времени:
```go
    t := time.Now()
    fmt.Println(t.Year())
    fmt.Println(t.Month())
    fmt.Println(t.Day())
    fmt.Println(t.Hour())
    fmt.Println(t.Minute())
    fmt.Println(t.Second())
    fmt.Println(t.Nanosecond())
    fmt.Println(t.Date()) // 2024 August 15
```
## Создание времени
```go
t := time.Date(2024, time.April, 1, 12, 12, 0, 0, time.UTC)
```
Очень удобно, что месяц представлен как именная константа. Также если например передать день 34, то ошибки не будет - он просто перенесется.
```go
    t := time.Date(2024, time.August, 50, 12, 12, 12, 12, time.UTC)
    fmt.Print(t.Format(time.DateOnly))  // 2024-09-19
```

## Форматирование времени. Объект -> строка
Форматирование происходит с помощью метода `time.Format`. Он переводит объект `Time` в строку в соответствии с переданным шаблоном
```
Year: "2006" "06"
Month: "Jan" "January" "01" "1"
Day of the week: "Mon" "Monday"
Day of the month: "2" "_2" "02"
Day of the year: "__2" "002"
Hour: "15" "3" "03" (PM or AM)
Minute: "4" "04"
Second: "5" "05"
AM/PM mark: "PM"
```
Вот это прикол. У них тут собственный форматер. То есть если мы хотим получить только год, то мы должны передать либо 2006, либо 06
```go
    t := time.Now()
    fmt.Println(t.Format("06")) // 2024
```
Смотрится непривычно:
```go
    t := time.Now()
    fmt.Println(t.Format("Сейчас 06 год. Jan. А время 15:4")) // Сейчас 24 год. Aug. А время 12:19
```
Есть также предопределенные шаблоны. `time.DateOnly`, `time.TimeOnly`, `time.DateTime`
```go
    t := time.Now()
    fmt.Println(t.Format(time.DateTime)) // 2024-08-15 12:21:48
    fmt.Println(t.Format(time.DateOnly)) // 2024-08-15
    fmt.Println(t.Format(time.TimeOnly)) // 12:21:48
```

## Парсинг времени. Строка -> объект
Парсинг времени происходит с помощью функции `time.Parse(patter, value string) (Time, error)`. 
```go
    parsedTime, err := time.Parse("2-Jan-2006-03:04AM", "11-Apr-2024-10:25AM")
    if err != nil {
        panic(err)
    }
    fmt.Println(parsedTime)
```
В шаблоне нужно указывать те модели, которые выше в таблице

## Длительность
Продолжительность между двумя промежутками времени представляется типом `type Duration int64`, то есть это просто алиас.
Длительность можно вычислить можно с помощью метода `time.Sub`
```go
    t1 := time.Date(2024, time.August, 2, 15, 16, 0, 0, time.UTC)
    t2 := time.Date(2024, time.August, 2, 15, 12, 0, 0, time.UTC)

    delta := t1.Sub(t2)
    fmt.Println(delta)           //4m0s
    fmt.Println(delta.Minutes()) // 4
```

Есть 2 вспомогательные функции на основе `time.Sub`
- `time.Since` - вычисляет длительность между текущим и прошлым моментом времени. Сокращение от `time.Now().Sub(t)`
- `time.Until` - вычисляет разницу между текущим и будущим временем. Сокращение от `t.Sub(time.Now())`

## Арифметика времени
Для арифметики времени есть 2 функции: `time.Add` и `time.AddDate`
```go
    t := time.Now()
    fmt.Println(t.Add(2 * time.Minute)) // добавится 2 минуты
    fmt.Println(t.AddDate(1, 0, 0)) // добавится 1 год
```
Для добавления времени - `time.Add`, для добавление года/месяца/дня - `time.AddDate(year, month, day int)`
Для сравнения используется 3 метода - `time.Equal(t time.Time)`, `time.After(t time.Time)`, `time.Before(t time.Time)`



