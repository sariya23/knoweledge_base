[[Код]]
Этот паттерн говорит о том, что функции или классы должны зависеть не от конкретных реализаций, а от абстракций (интерфейсов)
Вот это Dependency Inversion
```go
type Sender interface {
}

func NewHandler(sender Sender) Handler {
	return &Handler{s: sender}
}
```
А вот это нарушение этого принципа
```go
func NewHandler(sender SenderImplementation) Handler {
	return &Handler{s: sender}
}
```