package apiserver

import (
	//"github.com/jackc/pgx/v5"
	"net/http"
)

func Start(config *Config) error {
	//db, err := newDB(config.DatabaseURL)
	//if err != nil {
	//	return err
	//}
	//
	//defer db.Close()
	//store := sqlstore.New(db)
	//sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(config)

	return http.ListenAndServe(config.BindAddr, srv.router)
}
