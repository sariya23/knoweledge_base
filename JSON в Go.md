## Маршализация и демаршализация
### Демаршализация
В Go можно преобразовывать байты в структуру. Если они смогут спарситься.
```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Job  string `json:"job"`
}

func main() {
    data := []byte(`{
        "name": "Nikita",
        "age": 20,
        "job": "QA"
    }`)
    var user User
    err := json.Unmarshal(data, &user)
    if err != nil {
        panic(fmt.Sprintf("Прозошла ошибка: %v", err))
    }
    fmt.Println(user) // {Nikita 20 QA}
}
```
То, что написано после поля структуры в обратных кавычках - это то, как это будет называться в JSON. Поля структуры также должны быть названы с большой буквы, чтобы быть экспортируемыми. Если какое-то поле будет с маленькой буквы, его значение будет нулевым значением этого поля в результате демаршаллизации
 Функция `Unmarshal` принимает срез байтов и записывает их в структуру, если это возможно. Также можно записывать несколько json-ов в срез:
```go
func main() {
    data := []byte(`[
        {
            "name": "Nikita",
            "age": 20,
            "job": "QA"
        },
        {
            "name": "Andrew",
            "age": 22,
            "job": "Content-maker"
        }
    ]`)
    var user []User
    err := json.Unmarshal(data, &user)
    if err != nil {
        panic(fmt.Sprintf("Прозошла ошибка: %v", err))
    }
    fmt.Println(user)
}
```
### Маршализация
Доступна также и обратная операция - из структуры в байты
```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    Job  string `json:"job"`
}

func main() {
    user := User{
        Name: "Nikita",
        Age:  20,
        Job:  "QAA",
    }
    bytes, err := json.Marshal(user)
    if err != nil {
        panic(err)
    }
    fmt.Println(bytes) // [123 34 110 97 ...]
    
    var userFromBytes User
    err = json.Unmarshal(bytes, &userFromBytes)
    if err != nil {
        panic(err)
    }
    fmt.Println(userFromBytes) // {Nikita 20 QAA}
}
```
Функция `Marshal` принимает на вход структуру или их срез и возвращает байты с этими данными. В примере выше я перевел маршализировал структуру в набор байтов, а потом демаршализировал обратно в структуру

Важно заметить, что взаимодействие тут происходит именно в байтами - это сделано специально и удобно. Так как интерфейсы Writer и Reader тоже работают с байтами

## Запись и чтение JSON в файл из из файла
### Пребразование данных из файла JSON в структуру
```json
[
    {
        "name": "John Doe",
        "age": 20,
        "job": "QA"
    },
    {
        "name": "Brian Flemming",
        "age": 21,
        "job": "QAA"
    },
    {
        "name": "Lucy Black",
        "age": 10,
        "job": null
    },
    {
        "name": "John Doe",
        "age": 45,
        "job": "Bank"
    }
]
```
Для чтение данных из файла потребуются пакеты `os` для управления состоянием файла и `io` - для чтение данных. 
```go
func main() {
    file, err := os.Open("data.json")
    defer file.Close()
    if err != nil {
        fmt.Printf("Error: %v", err)
    }
    
    data, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error: %v", err)
    }
    
    var result []User
    jsonErr := json.Unmarshal(data, &result)
    if jsonErr != nil {
        fmt.Printf("Error: %v", jsonErr)
    }
    fmt.Println(result) // [{John Doe 20 QA} {Brian Flemming 21 QAA} {Lucy Black 10 } {John Doe 45 Bank}]
}
```
`io.ReadAll` принимает интерфейс `io.Reader`, который реализует файловый объект. Поля, которые в json имеют `null` становятся нулевым значением типа поля в структуре

### Запись структуры в файл
```go
func main() {
    data := []User{
        {Name: "Nikita", Age: 20, Job: "QA"},
        {Name: "Andrew", Age: 25, Job: "McJob"},
    }
    user := User{
        Name: "qwe",
        Age:  2,
        Job:  "",
    }
    bytesData, err := json.Marshal(data)
    if err != nil {
        fmt.Printf("Error: %v", err)
    }
    bytesUser, err := json.Marshal(user)
    if err != nil {
        fmt.Printf("Error: %v", err)
    }

    err = os.WriteFile("res.json", bytesData, 0644)
    err = os.WriteFile("user.json", bytesUser, 0644)
}
```
И вот прикол байтов - им не важно срез там или нет. Ничего при записи не изменится