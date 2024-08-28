[[projects]]

Не украл, а вдохновился pog - https://www.youtube.com/watch?v=rCJvW2xgnk0&t=2550s
## Цель
Написать REST API сервис для сокращения ссылок.

## Зависимости
- `cleanenv` используется для хранения и парсинга конфига. Минималистичная либа - https://github.com/ilyakaznacheev/cleanenv. Читаеть из yaml, .env. И все одной функцией.
- `log/slog` - либа для логирования
- `pgx` - драйвер для postgres
- `chi` - написание хендлеров. Полностью совместим с `net/http`

## Этапы
- Настройка конфига
- Инициализация логгера
- Инициализация хранилища
- Создание роутеров
- Запуск сервера

## Настраиваем конфиг
Логика настройки конфига будет находится в `internal/config`. Так как в этой папке мы храним не общедоступные функции.
Файл с конфигом в виде .yaml файла:
```yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
  address: "localhost:8082"
  timeout: 4s  # время на чтение и отправу запроса
  iddle_timeout: 60s   # время жизни соединения
```
Теперь нужно написать функцию, которая будет парсить эти данные в гошную структуру. В файле `internal/config/config.go`
```go
type Config struct {
    Env         string `yaml:"env" env-required:"true"`
    StoragePath string `yaml:"storage_path" env-required:"true"`
    HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
    Address      string        `yaml:"address" env-default:"localhost:8080"`
    Timeout      time.Duration `yaml:"timeout" env-default:"4s"`
    IddleTimeout time.Duration `yaml:"iddle_timeout" env-default:"60s"`
}
```
Здесь используются struct-теги для указания того, как это поля называется в .yaml или .env файле. Struct-тег `env-required` означает, что если это значения не задано, то будет ошибка. Так как поле `http_server` сложное, то для него нужно создавать отдельную структуру и встраивать ее в конфиг и также с struct-тегом.
Теперь напишем функцию, которая будет парсить этот конфиг:
```go
func MustLoad() *Config {
    configPath := os.Getenv("CONFIG_PATH")
    if configPath == "" {
        log.Fatal("CONFIG PATH is not set")
    }
    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        log.Fatalf("config file %s does not exist", configPath)
    }
    var config Config
    if err := cleanenv.ReadConfig(configPath, &config); err != nil {
        log.Fatalf("cannot read cofnig file: (%v)", err)
    }
    return &config
}
```
Функция начинается с Must, так как она выдает панику (log.Fatal). Так как это первый этап запуска, то ничего кроме паники не остается делать. Также используется стандартный логгер, так как `slog` еще не настроен. 
Путь до файла конфига будем брать из переменных окружения. 

## Логгер
Теперь настроим логгер. Сделаем так, чтобы при окружении local сообщения выводились с помощью простого текста, а при окружении dev в виде json.
В файле `cmd/url-shortener/main.go` напишем функцию, которая будет создавать логгер с нужными параметрами
```go
const (
    envLocal = "local"
    envDev   = "dev"
)

func setUpLogger(env string) *slog.Logger {
    var logger *slog.Logger
    handlerOptions := &slog.HandlerOptions{Level: slog.LevelDebug}
    switch env {
    case envLocal:
        logger = slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
    case envDev:
        logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
    }
    return logger
}
```
`NewTextHandler` означает, что сообщения будут выводится в виде текста, `newJSONHandler` соответственно в виде json. 
Добавим в функцию `main` первые сообщения:
```go
    log := setUpLogger(config.Env)

    log.Info("starting url-shortener", slog.String("env", config.Env))
    log.Debug("debuf messages are enabled")
```
`slog.String` можно добавить в логгер по умолчанию. Чтобы к каждому сообщению добавлялось окружение. Это можно сделать с помощью `With`
```go
    log := setUpLogger(config.Env)
    log = log.With(slog.String("env", config.Env))
    
    log.Info("starting url-shortener")
    log.Debug("debuf messages are enabled")
```
![[Pasted image 20240826181316.png]]
Это полезно, но не сейчас.

## Storage
В общем в видосе используется sqlite, но я решил заменить ее на postgres, так как изучаю ее сейчас. 
В `/internal` создадим каталог `storage`, а в нем пакет `postgres`.  А в нем файл `postgres.go`. В этом файле будет структура с соединением и также инициализация нужных сущностей. 
```go
type Storage struct {
    connection *pgx.Conn
}
```
Далее создадим функцию `New`, которая будет создавать таблицу, индекс и функцию закрытия соединения:
```go
func New(ctx context.Context, storagePath string) (*Storage, func(s Storage), error) {
    const operationPlace = "storage.postgres.New"
    cancel := func(s Storage) {
        err := s.connection.Close(ctx)
        if err != nil {
            panic(err)
        }
    }
    conn, err := pgx.Connect(ctx, storagePath)
    if err != nil {
        return &Storage{connection: conn}, cancel, fmt.Errorf("%s: %w", operationPlace, err)
    }
    _, err = conn.Exec(ctx, `
    create table if not exists url (
        url_id bigint generated always as identity primary key,
        alias text not null unique,
        url text not null
    );
    `)
    if err != nil {
        return &Storage{connection: conn}, cancel, fmt.Errorf("%s: %w", operationPlace, err)
    }
    _, err = conn.Exec(ctx, `create unique index if not exists url_idx on url(alias)`)
    if err != nil {
      return &Storage{connection: conn}, cancel, fmt.Errorf("%s: %w", operationPlace, err)
    }
    return &Storage{connection: conn}, cancel, nil
}
```
Эту функцию писал сам, не с видоса. Поэтому вероятнее всего она булщит. Контекст передают из вне, так как так и надо - использовать общий контекст для программы и модифицировать  его.
Функция отмены нужна, чтобы вызывать ее из вне с помощью `defer`. 
В данном случае не нужно создавать соединение на каждый запрос, так как это жрет ресурсы. 
Также важно - при создании таблиц, вставки и тд надо использовать `Exec`, так как `Query` заблокирует соединение до тех пор, пока он не будет прочитан из переменной. 

Создание таблицы и индекса тут все понятно - обычный postgres.
Использование в `main`
```go
    storage, cancel, err := postgres.New(ctx, os.Getenv("DATABASE_URL"))
    defer cancel(*storage)
    if err != nil {
        log.Error("failed to init storage", xslog.Err(err))
        os.Exit(1)
    }
    log.Info("Storage init success. Create table and index")
    _ = storage
```

Еще в каталоге `postgres` надо создать файл `storage.go`.
```go
package storage
import "errors"

var (
    ErrURLNotFound = errors.New("url now found")
    ErrURLExists   = errors.New("url exists")
)
```
Я пока не понял зачем, но раз надо, то ок
### Сохраняем url в БД
Напишем метод БД, который будет принимать url и его псевдоним, сохранять его и возвращать id вставленной записи
```go
func (s *Storage) SaveURL(ctx context.Context, urlToSave string, alias string) (int, error) {
    const operationPlace = "storage.postgres.SaveURL"
    var insertedId int
    var pgErr *pgconn.PgError
    
    query := "insert into url(url, alias) values ($1, $2) returning url_id"
    err := s.connection.QueryRow(ctx, query, urlToSave, alias).Scan(&insertedId)
    if ok := errors.As(err, &pgErr); ok && pgErr.Code == pgerrcode.UniqueViolation {
	    return -1, fmt.Errorf("%s: %w", operationPlace, storage.ErrURLExists)
    }
    if err != nil {
        return -1, fmt.Errorf("%s: %w", operationPlace, err)
    }
    return insertedId, nil
}
```
Первый блок - объявления переменных. Тут следует обратить внимание на строчку с `pgErr`. Нам понадобиться тип такой ошибки, чтобы получить код ошибки и сравнить с ожидаемым.
Вставка ничего интересного.
Обработка ошибки происходит с помощью `As`, чтобы проверить, что в цепочке ошибок есть нужна. Если в цепочке обернутых ошибок есть ошибка типа PgErr, то мы проверяем, что код ошибки равен коду нарушения уникальности в SQL (23505). 

***TODO: подумать, может это лучше будет реализовать через upsert.***
###  Получение алиаса по url
Эту колбасу я тоже сам писал, поэтому, вероятно, можно сделать лучше
```go
func (s *Storage) GetURLByAlias(ctx context.Context, alias string) (string, error) {
    const operationPlace = "storage.postgres.GetURLByAlias"
    var urlByAlias string
    query := `select url from url where alias=$1`
    err := s.connection.QueryRow(ctx, query, alias).Scan(&urlByAlias)

    if errors.Is(err, pgx.ErrNoRows) {
        return "", storage.ErrURLNotFound
    }
    if err != nil {
        return "", fmt.Errorf("%s: %w", operationPlace, err)
    }
    return urlByAlias, nil
}
```
Вот почему ошибку пустого возврата надо брать из `pgx`, а другие из `pgerrcode` и сравнивать с кодом из `pgconn`. Жесть какая-то.
Ну тут в принципе ничего интересного - обычный селект и обработка ошибок.

### Удаление записи
Тут ничего интересного:
```go
func (s *Storage) DeleteURLByAlias(ctx context.Context, alias string) (string, error) {
    const operationPlace = "storage.postgres.GetURLByAlias"
    query := `delete from url where alias=$1`
    t, err := s.connection.Exec(ctx, query, alias)
    if err != nil {
        return "", fmt.Errorf("%s: %w", operationPlace, err)
    }
    deletedRows := string(strings.Split(t.String(), " ")[1])
    return deletedRows, nil
}
```
Единственное - это возврат кол-ва удаленных строк. Думаю, это можно будет в будущем убрать. Возвращаю это для логирования. Если не будет помогать - удалю.

## Создание объекта роутера
В `main()` нужно создать объект роутера из пакета `chi`
```go
router := chi.NewRouter()
```
## Добавление middleware
Теперь добавить middleware
```go
    // Добавляет request id к каждому запросу
    router.Use(middleware.RequestID)
    // Добавляет ip пользователя
    router.Use(middleware.RealIP)
    // Логирует входящие запросы
    router.Use(mwLogger.New(log))
    // При панике, чтобы не падало все приложение из-за одного запроса
    router.Use(middleware.Recoverer)
    // Фишка chi. Позволяет писать такие роуты: /articles/{id} и потом
    // получать этот id в хендлере
    router.Use(middleware.URLFormat)
```
mwLogger - это наш пакет. В `internal/http-server/middleware/logger` нужно создать файл `logger.go` и поместить туда этот код:
```go
func New(log *slog.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        log := log.With(
            slog.String("component", "middleware/logger"),
        )
        log.Info("logger middleware enabled")
        fn := func(w http.ResponseWriter, r *http.Request) {
            entry := log.With(
                slog.String("method", r.Method),
                slog.String("path", r.URL.Path),
                slog.String("remote_addr", r.RemoteAddr),
                slog.String("user_agent", r.UserAgent()),
                slog.String("request_id", middleware.GetReqID(r.Context())),
            )
            ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
            t1 := time.Now()
            defer func() {
                entry.Info("request completed",
                    slog.Int("status", ww.Status()),
                    slog.Int("bytes", ww.BytesWritten()),
                    slog.String("duration", time.Since(t1).String()),
                )
            }()
            next.ServeHTTP(ww, r)
        }
        return http.HandlerFunc(fn)
    }
}
```
Можно было бы использовать стандартный логгер, но он использует свой какой-то логгер. Лучше использовать тот, который уже настроен, чтобы было проще собирать информацию

## Делаем логгер красивым
![[Pasted image 20240827120250.png]]
Чтобы сделать такую красоту:
в `internal/lib/logger/handlers/slogpretty` нужно создать файл с кодом как тут - https://github.com/GolangLessons/url-shortener/blob/main/internal/lib/logger/handlers/slogpretty/slogpretty.go. Далее в `main.go` нужно создать функцию, которая засетапит такой логгер:
```go
func setupPrettySlog() *slog.Logger {
    opts := slogpretty.PrettyHandlerOptions{
        SlogOpts: &slog.HandlerOptions{
            Level: slog.LevelDebug,
        },
    }
    handler := opts.NewPrettyHandler(os.Stdout)
    return slog.New(handler)
}
```
И изменить при localEnv логгер:
```go
    case envLocal:
        logger = setupPrettySlog()
```


## Хендлер для сохранения урла
Хендлеры будут лежать в `internal/http-server/handlers`. Создадим хендлер для сохранения урла и алиаса. В папке handlers создадим каталог url, а в нем пакет `save`
```go
type Request struct {
    URL   string `json:"url" validate:"required,url"`
    Alias string `json:"alias,omitempty"`
}

type Response struct {
    response.Response
    Alias string `json:"alias,omitempty"`
}  

type URLSaver interface {
    SaveURL(ctx context.Context, urlToSave string, alias string) (int, error)
}

func New(ctx context.Context, log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        const operationPlace = "handlers.url.save.New"
        log = log.With(
            slog.String("op", operationPlace),
            slog.String("requies_id", middleware.GetReqID(r.Context())),
        )
        var request Request
        err := render.DecodeJSON(r.Body, &request)
        if err != nil {
            log.Error("failed to decode request body", xslog.Err(err))
            render.JSON(w, r, response.Error("failed to decode request"))
            return
        }
        log.Info("request body decoded", slog.Any("request", request))
        err = validator.New(validator.WithRequiredStructEnabled()).Struct(request)
        if err != nil {
            validationErrors := err.(validator.ValidationErrors)
            log.Error("invalid request data", xslog.Err(err))
            render.JSON(w, r, response.ValidationError(validationErrors))
            return
        }
        alias := request.Alias
        if alias == "" {
            alias = random.NewRandomString(random.DefaultStringLen)
        }
        id, err := urlSaver.SaveURL(ctx, request.URL, alias)
        if errors.Is(err, storage.ErrURLExists) {
            log.Info("url already exists", "url", request.URL)
            render.JSON(w, r, response.Error("url already exists"))
            return
        }
        if err != nil {
            log.Error("failed to add url", xslog.Err(err))
            render.JSON(w, r, response.Error("failed to add url"))
            return
        }
        log.Info("url added", slog.Int("id", id))
        render.JSON(w, r, Response{
            Response: response.OK(),
            Alias:    alias,
        })
    }
}
```
Структура `Request`. В ней содержаться поля, которые прилетят нам во время запроса в виде json. Поэтому мы используем соответствующие struct теги. `omitempty` значит, что если поле пустое, то в json этого поля вообще не будет. `validate` - это тег из стороннего пакета, который проводит... валидацию.
Структура `Response` аналогично. Здесь мы используем встраивание. Остальная часть лежит в `internal/lib/api/response/response.go`
```go
package response

)

type Response struct {
    Status string `json:"status"`
    Error  string `json:"error,omitempty"`
}

const (
    StatusOK    = "OK"
    StatusError = "Error"
)

func OK() Response {
    return Response{
        Status: StatusOK,
    }
}

func Error(msg string) Response {
    return Response{
        Status: StatusError,
        Error:  msg,
    }
}

func ValidationError(errs validator.ValidationErrors) Response {
    var errMessages []string
    for _, err := range errs {
        switch err.ActualTag() {
        case "required":
            errMessages = append(errMessages, fmt.Sprintf("field %s is required", err.Field()))
        case "url":
            errMessages = append(errMessages, fmt.Sprintf("field %s is invalid URL", err.Field()))
        default:
            errMessages = append(errMessages, fmt.Sprintf("field %s is invalid", err.Field()))
        }
    }
    return Error(strings.Join(errMessages, ", "))
}
```
Интерфейс `URLSaver` нужен, чтобы не привязываться к конкретной реализации storage. Он определен здесь, так как интерфейсы должны располагаться на стороне Потребителя ([[Ошибка 6. Интерфейсы на стороне производителя]])
Создаем новую функцию хендлер и сначала декодируем запрос с помощью `render`. Это штука из `chi`.  Потом валидируем через сторонний пакет `validator`. Мы написали `ValidationError`, так как ошибки обычного валидатора не очень понятные. Так же мы используем type cast, чтобы проверить, что произошла именно ошибка валидации. Также если не указан alias мы генерируем его сами. В `internal/lib/random` нужно создать файл `random.go`
```go
package random

import "math/rand/v2"

const DefaultStringLen = 6

func NewRandomString(strLen int) string {
    min := 97
    max := 123
    var res string
    if strLen == 0 {
        strLen = DefaultStringLen
    }
    for i := 0; i < strLen; i++ {
        res += string(rune(rand.IntN(max-min) + min))
    }
    return res
}
```
И если нигде не было ошибок, то отправляем ответ с OK. В респонзе есть поле alias для тех случаев, когда он сгенерировался автоматически.

## Запускаем сервер
```go
    router.Post("/url", save.New(ctx, log, storage))
    log.Info("starting server", "address", config.Address)
    server := &http.Server{
        Addr:         config.Address,
        Handler:      router,
        ReadTimeout:  config.HTTPServer.Timeout,
        WriteTimeout: config.HTTPServer.Timeout,
        IdleTimeout:  config.HTTPServer.IddleTimeout,
    }
    if err := server.ListenAndServe(); err != nil {
        log.Error("failed to start server", xslog.Err(err))
    }
    log.Error("server stopped")
```
Добавить в роутер хендлер - это будет пост запрос.
Далее запускаем сервер. `log.Error("server stopped")` этот лог нужен, чтобы дать понять, что сервер уже не але, так как выше значит произошла ошибка. 
Запустим сервер и отправим туда какой-то запрос:
![[Pasted image 20240828102026.png]]
Алиас и урл добавились. Сходим в БД и проверим это
![[Pasted image 20240828102500.png]]
Тут тоже все появилось.

## Тест на хэндлер save
Чтобы не засорять и не ходить в нашу БД в тестах, замокаем ее. Для этого используется утилита mockery - https://pkg.go.dev/github.com/vektra/mockery#section-readme. Не забыть кинуть его в go/bin. Для мока нужно указать интерфейс, который нужно сымитировать и в пакете с ним создаться новый пакет `mocks`, в котором будет лежать этот mock.
То етсь в нашем случае надо выполнить такую штуку:
```shell
mockery --name=URLSaver
```
Сам тест:
```go
func TestSaveHandler(t *testing.T) {
    cases := []struct {
        caseName    string
        urlToSave   string
        aliasForURL string
        responseErr string
        mockErr     error
    }{
        {
            caseName:    "Success save",
            urlToSave:   "http://test.ru",
            aliasForURL: "suc",
        },
        {
            caseName:    "Long url",
            urlToSave:   "http://chart.apis.google.com/chart?chs=500x500&chma=0,0,100,100&cht=p&chco=FF0000%2CFFFF00%7CFF8000%2C00FF00%7C00FF00%2C0000FF&chd=t%3A122%2C42%2C17%2C10%2C8%2C7%2C7%2C7%2C7%2C6%2C6%2C6%2C6%2C5%2C5&chl=122%7C42%7C17%7C10%7C8%7C7%7C7%7C7%7C7%7C6%7C6%7C6%7C6%7C5%7C5&chdl=android%7Cjava%7Cstack-trace%7Cbroadcastreceiver%7Candroid-ndk%7Cuser-agent%7Candroid-webview%7Cwebview%7Cbackground%7Cmultithreading%7Candroid-source%7Csms%7Cadb%7Csollections%7Cactivity",
            aliasForURL: "sucv2",
        },
        {
            caseName:  "No allias",
            urlToSave: "http://test.ru",
        },
        {
            caseName:    "empty url",
            urlToSave:   "",
            aliasForURL: "empty url",
            responseErr: fmt.Sprintf("%s %s", response.ErrMSgMissingRequiredField, "URL"),
        },
        {
            caseName:    "SaveURL error",
            urlToSave:   "http://qwe.ru",
            responseErr: save.ErrMsgFailedAddUrl,
            mockErr:     errors.New("unexpected error"),
        },
    }
    for _, testCase := range cases {
        t.Run(testCase.caseName, func(t *testing.T) {
            ctx := context.Background()
            t.Parallel()
            urlSaverMock := mocks.NewURLSaver(t)
            if testCase.responseErr == "" || testCase.mockErr != nil {
                urlSaverMock.On("SaveURL", ctx, testCase.urlToSave, mock.AnythingOfType("string")).
                    Return(1, testCase.mockErr).
                    Once()
            }
            handler := save.New(ctx, slogdiscard.NewDiscardLogger(), urlSaverMock)
            dataToRequest := fmt.Sprintf(`{"url":"%s", "alias":"%s"}`, testCase.urlToSave, testCase.aliasForURL)
            request, err := http.NewRequest(http.MethodPost, "/url", bytes.NewReader([]byte(dataToRequest)))
            require.NoError(t, err)
            rr := httptest.NewRecorder()
            handler.ServeHTTP(rr, request)
            require.Equal(t, rr.Code, http.StatusOK)
            body := rr.Body.String()
            var response save.Response
            require.NoError(t, json.Unmarshal([]byte(body), &response))
            require.Equal(t, testCase.responseErr, response.Error)
        })
    }
}
```
Я вот тоже в шоке - какая-то непонятная каша. Пока что не погружался в вопрос тестирования с помощью моков. Это разберу отдельно.