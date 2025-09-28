[[Go]]
## Минимальный сервер и логика
```go
const (
	deafultPort = ":9000"
)

func main() {
	http.HandleFunc("/ping", rootHandler)

	if err := http.ListenAndServe(deafultPort, nil); err != nil {
		log.Fatal()
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	w.Write([]byte("Root route"))
}
```
`HadleFunc` говорит серверу, что нужно выполнять по переданному паттерну. `ListenAndServe` запускает сервер.
## Роутинг
### Дерево роутинга
На базовый роут `/` обычно не вешают бизнес-логику. Так как большие бэкенды живут в k8s или подобных окружениях и от них зависят другие сервисы, то ми нужно как-то знать, что наше прилоежние живо. Поэтому на такой  базовый роут обычно вешают метрики, дебажную информацию или вот такие штуки
```go
	http.HandleFunc("/ping", rootHandler)
	http.HandleFunc("/alive", rootHandler)
	http.HandleFunc("/ready", rootHandler)
```
Считается плохой практикой вешать на такой базовый роут какую-то бизнес-логику (если не PWA и не фронт)
Роутинг можно представить в виде дерева, где есть родительские роуты и их потомки. Роуты бывают также *фиксированные* и *многоуровневые*. Фиксированный роут, это роут от которого дальше нет потомков, например вот так
```
/home
/articles/hello
```
В фиксированных роутах обычно отдаются какие-то статические данные, так как картинки, текст и тд. 
Многоуровневый роут это роут, у которого дальше уже могут быть потомки, например
```
articles/it/
```
Если роут заканчивается слэшом, то подразумевается, что у него есть потомки
### Редиректы
Редиректы можно также сделать в `net/http`
```go
	http.HandleFunc("/home", rootHandler)
	http.HandleFunc(
		"/home/articles/",
		http.RedirectHandler("http://localhost:9000/home", http.StatusMovedPermanently).ServeHTTP)
	if err := http.ListenAndServe(deafultPort, nil); err != nil {
		log.Fatal()
	}
```
ServeHTTP запускает мультиплексер (сервер) в отдельной горутине, это нужно, чтобы этот редирект нам что-то отдал. 
Либо можно сделать так
```go
http.HandleFunc("/home/movies/", redirectHandler)
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://localhost:9000/home", http.StatusMovedPermanently)
}
```
## Middleware
Middleware - это механизм, который позволяет выполнять какую-либо логику перед вызовом хендлера и после вызова хендлера. В middleware стоит помещать логику, которая требуется для каждого запроса, что-то общее. Например, логирование или проверка базовой авторизации (BaseAuth). 
Что-то тяжелое и сложное пихать в middleware не стоит, так как можно сильно нагрузить систему, что приведет к ее падению. 
Middleware сырым `net/http`
```go
func main() {
	http.HandleFunc("/home", middleware(
		func() http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("HELLO"))
			}
		}()),
	)
	if err := http.ListenAndServe(basePort, nil); err != nil {
		log.Fatal()
	}
}

func middleware(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// До вызова
		w.Write([]byte("before call"))

		// вызов
		handler.ServeHTTP(w, r)

		// После вызова
		w.Write([]byte("After call"))
	}
}
```
Можно выстраивать цепочку middleware
## Коды и ошибки
==ЗАПОМНИТЬ РАЗ==
Нельзя возвращать неверные коды ошибок, особенно везде выкидывать Internal, так как 5хх говорит о том, что бэкенд *сломался*. Если пользователь не найден, нужно выкидывать 4хх и по аналогии с другими ситуациями. Если же упала база (истек таймаут по ней), то вот тогда можно выкинуть 5хх
==ЗАПОМНИТЬ ДВА==
Нельзя выкидывать наружу внутренние ошибки. Например той же базы.