[[Go]]
## Connect

```go
import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		ctx = context.Background()
		dsn = "user=nikita dbname=db password=1234 host=localhost sslmode=disable"
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}

func tx(ctx context.Context, db *sql.DB) (err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// query

	return
}
```
Функция `Open` не всегда может гарантировать, что соединение установлено. Для проверки надо сделать `Ping`. Лучше использовать метод с `Context`, так как без него методы оставили для обратной совместимости. Соединение также нужно закрывать, чтобы оно вернулось в пул соединений. 
Функция `tx` показывает как можно работать с транзакциями. В случае ошибки делаем ролбэк, а иначе коммит. Будет работать как надо, только есть возвращать именованный параметр.
И у `db`, `tx`, `conn` одинаковый интерфейс.
## Exec
```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		ctx = context.Background()
		dsn = "user=nikita dbname=db password=1234 host=localhost sslmode=disable"
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	var res sql.Result
	// Exec - no returning rows
	if res, err = db.Exec("delete from account where account_id=$1", 1); err != nil {
		log.Fatal(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if count == 0 {
		log.Fatal("no rows")
	}
	fmt.Printf("deleted %d rows\n", count)
	// res.LastInsertId()

	var id int
	// QueryRow - one return value
	row := db.QueryRow(`insert into account (name, age) values ($1, $2) returning account_id`, "Test", 90)
	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		log.Fatal("not found")
	}
	fmt.Printf("inserted row with id %d\n", id)
	var ids []int
	rows, err := db.Query("select account_id from account")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var ID int
		err := rows.Scan(&ID)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, ID)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(ids)
}

```
`db.Exec` следует использовать, когда запрос не возвращает значений - `delete`, `insert` и тд. У возвращаемого результата есть 2 метода - кол-во зааффекченных строк и id последней вставленной строки. Последний метод не поддерживается постгресом. 
`db.QueryRows` следует использовать когда запрос возвращает одну строку.
`db.Query` - когда множество строк. Так же нужно помнить про ошибку [[Ошибка 78. Типичные ошибки, связанные с SQL]] и проверять ошибки при сканировании строк
### Null
```go
import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	var (
		ctx = context.Background()
		dsn = "user=nikita dbname=db password=1234 host=localhost sslmode=disable"
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	var res sql.NullString

	if err = db.QueryRow("select name from account where age = $1", 48).Scan(&res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res) // { false}

	var res1 sql.NullString

	if err = db.QueryRow("select name from account where age = $1", 22).Scan(&res1); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res1) // {Andrew true}

	var res2 string

	if err = db.QueryRow("select name from account where age = $1", 48).Scan(&res2); err != nil {
		log.Fatal(err) // error
	}

}
```
В случае, когда какое-то значение может быть `null`, то лучше использовать тип `sql.Null<Type>`, чтобы предотвратить ошибку. Также можно класть в указатель на тип `*string`, но лучше использовать специализированные типы. В [[Ошибка 78. Типичные ошибки, связанные с SQL]] про это подробнее
Тип - это структура следующего содержания
```go
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}
```

### sqlx
```go
import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Foo struct {
	AccountId int            `db:"account_id"`
	Name      sql.NullString `db:"name"`
	Age       sql.NullInt64  `db:"age"`
}

func main() {
	var (
		ctx = context.Background()
		dsn = "user=nikita dbname=db password=1234 host=localhost sslmode=disable"
	)
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	var foos []Foo
	rows, err := db.QueryxContext(ctx, "select * from account")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var tmp Foo
		err = rows.StructScan(&tmp)
		if err != nil {
			log.Fatal(err)
		}
		foos = append(foos, tmp)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(foos)
}

```
`sqlx` расширяет стандартные возможности `sql`. Можно сканить строки прям в структуры.
Также есть именованные параметры
```go
tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
```