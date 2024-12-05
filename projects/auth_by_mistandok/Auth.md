[[projects]]

Ссылка на проект - https://github.com/mistandok/auth

## Интересности
#### Асинхронный запуск разных серверов
`internal/app/app.go`
```go
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	runActions := []struct {
		action func() error
		errMsg string
	}{
		{action: a.runGRPCServer, errMsg: "ошибка при запуске GRPC сервера"},
		{action: a.runHTTPServer, errMsg: "ошибка при запуске HTTP сервера"},
		{action: a.runSwaggerServer, errMsg: "ошибка при запуске Swagger сервера"},
		{action: a.runPrometheusServer, errMsg: "ошибка при запуске Prometheus сервера"},
	}

	wg := sync.WaitGroup{}
	wg.Add(len(runActions))

	for _, runAction := range runActions {
		currentRunAction := runAction
		go func() {
			defer wg.Done()

			err := currentRunAction.action()
			if err != nil {
				log.Fatalf(currentRunAction.errMsg)
			}
		}()
	}

	wg.Wait()
```
#### Параметры фильтрации хранятся в структуре
`internal/repository/user/get_by_filter.go`
```go
func (u *Repo) GetByFilter(ctx context.Context, filter *serviceModel.UserFilter) (*serviceModel.User, error)
```

#### platform_common
`internal/repository/user/get_by_filter.go` да и везде где запросы к БД идут.
```go
	q := db.Query{
		Name:     "user_repository.GetByFilter",
		QueryRaw: query,
	}

	rows, err := u.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return nil, err
	}
```
Используется оборетка над стандартной либой для работы с постгрес.  Потом имя запроса можно где-то логировать и использоваться для трейсинга.

Также в этом пакете есть интерфейс - `Client`, который имплиментриуется всеми структурами для работы с БД. Удобненько. Получается, что и создание `conn` тоже не прибито к конкретной реализации

#### Архитектура
`internal/api` - самый верхний слой. Слой API. Он только принимает запрос, отдает данные на слой ниже - слой бизнес-логики, принимает ответ от него, хендлит ошибки от бизнес-логики  и отдает ответ клиенту.
`internal/service` - слой бизнес логики. Слой, принимающий данные от слоя API. 
`internla/repositoy` - слой данных. Принимает только обработанные данные от бизнес логики и идет в БД.

Разберу на примере хендлера `Create` user
##### Слой API
```go
func (i *Implementation) Create(ctx context.Context, request *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	userID, err := i.userService.Create(ctx, convert.ToServiceUserForCreateFromCreateRequest(request))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrEmailIsTaken):
			return nil, status.Error(codes.AlreadyExists, repository.ErrEmailIsTaken.Error())
		case errors.Is(err, user.ErrPassToLong):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, api.ErrInternal
		}
	}

	return &user_v1.CreateResponse{Id: int64(userID)}, nil
}
```
Мы принимаем запрос от клиента и посылаем данные на слой бизнес логики. ВСЕ! Потом только обрабатываем ошибки от него и возвращаем ответ. 
##### Слой логики
```go
func (s *Service) Create(ctx context.Context, userForCreate *model.UserForCreate) (int64, error) {
	hashedPassword, err := s.passManager.HashPassword(userForCreate.Password)
	if err != nil {
		s.logger.Error().Err(err).Msg("не удалось хэшировать пароль")
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			return 0, ErrPassToLong
		}
		return 0, err
	}

	userForCreate.Password = hashedPassword

	userID, err := s.userRepo.Create(ctx, userForCreate)
	if err != nil {
		s.logger.Error().Err(err).Msg("не удалось создать пользователя")
		return 0, fmt.Errorf("ошибка при попытке создать пользователя: %w", err)
	}

	return userID, nil
}
```
На слое логики мы уже хешируем пароль. А потом отправляем данные для сохранения слою данных. ВСЕ! Ну и хендлим ошибки от слоя данных. Вот тут прикольно то, что ошибка просто врапиться, а потом передается слою API, где уже обрабатывается.
##### Слой данных
```go
func (u *Repo) Create(ctx context.Context, in *serviceModel.UserForCreate) (int64, error) {
	queryFormat := `
	INSERT INTO "%s" (%s, %s, %s, %s, %s, %s)
	VALUES (@%s, @%s, @%s, @%s, @%s, @%s)
	RETURNING id
	`
	query := fmt.Sprintf(
		queryFormat,
		userTable, nameColumn, emailColumn, roleColumn, passwordColumn, createdAtColumn, updatedAtColumn,
		nameColumn, emailColumn, roleColumn, passwordColumn, createdAtColumn, updatedAtColumn,
	)

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	currentTime := time.Now()

	args := pgx.NamedArgs{
		nameColumn:      in.Name,
		emailColumn:     in.Email,
		roleColumn:      in.Role,
		passwordColumn:  in.Password,
		createdAtColumn: currentTime,
		updatedAtColumn: currentTime,
	}

	rows, err := u.db.DB().QueryContext(ctx, q, args)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var userID int64

	userID, err = pgx.CollectOneRow(rows, pgx.RowTo[int64])
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.ConstraintName == userEmailKeyConstraint {
			return 0, repository.ErrEmailIsTaken
		}
		return 0, err
	}

	return userID, nil
}
```
На слое данных мы просто выполняем сохранения в БД и все.