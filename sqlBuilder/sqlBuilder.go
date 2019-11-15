package sqlBuilder

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
)

func GetSqlDb(username, password, hostname, port string) *sql.DB {
	query := url.Values{}
	query.Add("Authentication API", "AuthenticationAPI")

	u := url.URL{
		Scheme: "sqlserver",
		User: url.UserPassword(username, password),
		Host: fmt.Sprintf("%s:%d", hostname, port),
		RawQuery: query.Encode(),
	}

	if db, err:= sql.Open("sqlserver", u.String()); err == nil {
		return db
	} else {
		panic(err)
	}
}

