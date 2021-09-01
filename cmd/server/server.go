package main

import (
	"go_news/pkg/api"
	"go_news/pkg/storage"
	"go_news/pkg/storage/memdb"
	"go_news/pkg/storage/mongo"
	"go_news/pkg/storage/postgres"
	"log"
	"net/http"
	"os"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	pgPass := os.Getenv("pgPass")
	if pgPass == "" {
		os.Exit(1)
	}
	// conn - строка подключения к базе данных
	pgConn := "postgres://postgres:" + pgPass + "@localhost:5432/news"

	db2, err := postgres.New(pgConn)
	if err != nil {
		log.Fatal(err)
	}
	// закрываем ресурс
	defer db2.Close()

	// Документная БД MongoDB.
	mongoConn := "mongodb://localhost:27017/"
	db3, err := mongo.New(mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	// закрываем ресурс
	defer db3.Close()

	_, _, _ = db, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
