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

