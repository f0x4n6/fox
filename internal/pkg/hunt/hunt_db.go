package hunt

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"

	"github.com/cuhsat/fox/v4/internal/pkg/types/event"
)

const schema = `
CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY,
    time INTEGER NOT NULL,
    host_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
	message TEXT NOT NULL,
	severity INTEGER NOT NULL,
	UNIQUE(time, host_id, user_id, message) ON CONFLICT ROLLBACK,
    FOREIGN KEY (host_id) REFERENCES hosts (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS hosts (
    id INTEGER PRIMARY KEY,
    value TEXT NOT NULL,
	UNIQUE(value) ON CONFLICT ROLLBACK
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    value TEXT NOT NULL,
	UNIQUE(value) ON CONFLICT ROLLBACK
);
`

var db *sql.DB

func UseDB(path string) {
	var err error

	db, err = sql.Open("sqlite", fmt.Sprintf("file:%s", path))

	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(path)

	if errors.Is(err, os.ErrNotExist) {
		_, err = db.Exec(schema)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("hunt: created %s\n", path)
	} else if err != nil {
		log.Fatal(err)
	}
}

func Save(evt *event.Event) error {
	var err error

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	hostId, err := insert(`hosts (value)`, evt.Host)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	userId, err := insert(`users (value)`, evt.User)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	id, err := insert(`events (time, message, severity, host_id, user_id)`,
		evt.Time.UTC(),
		evt.Message,
		evt.Severity,
		hostId,
		userId,
	)

	log.Println(id, hostId, userId)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func insert(table string, v ...any) (int64, error) {
	query := "INSERT OR IGNORE INTO %s VALUES (%s);"

	return execute(fmt.Sprintf(query, table, values(len(v))), v...)
}

func execute(query string, v ...any) (int64, error) {
	res, err := db.Exec(query, v...)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func values(n int) string {
	var sb strings.Builder

	sb.WriteRune('?')

	for range n - 1 {
		sb.WriteString(", ?")
	}

	return sb.String()
}
