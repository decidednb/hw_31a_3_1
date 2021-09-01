package postgres

import (
	"context"
	"go_news/pkg/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

// New - конструктор, conn - строка подключения к БД
func New(conn string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		return nil, err
	}

	s := Storage{
		db: db,
	}

	return &s, nil
}

// ConnClose - закрывает соединение к базе данных
func (s *Storage) Close() {
	s.db.Close()
}

// Posts - получает список всех публикаций из БД
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			posts.id,
			posts.title,
			posts.content,
			posts.author_id,
			authors.name,
			posts.created_at			
		FROM posts
		JOIN authors ON posts.author_id = authors.id;
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

// AddPost - добавляет публикацию в БД
func (s *Storage) AddPost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (author_id, title, content) VALUES
			($1, $2, $3);
	`, post.AuthorID, post.Title, post.Content)

	if err != nil {
		return err
	}
	return nil
}

// UpdatePost - обновляет публикацию в БД
func (s *Storage) UpdatePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		UPDATE posts SET author_id = $1, title = $2, content = $3 WHERE id=$4;
	`, post.AuthorID, post.Title, post.Content, post.ID)

	if err != nil {
		return err
	}
	return nil
}

// DeletePost - удаляет публикацию из БД
func (s *Storage) DeletePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		DELETE FROM posts WHERE id = $1;
	`, post.ID)

	if err != nil {
		return err
	}
	return nil
}
