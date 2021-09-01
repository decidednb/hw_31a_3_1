package storage

// Post - публикация.
type Post struct {
	ID         int    `bson:"id" json:"id"`                   // идентификатор публикации
	Title      string `bson:"title" json:"title"`             // заголовок публикации
	Content    string `bson:"content" json:"content"`         // содержание публикации
	AuthorID   int    `bson:"author_id" json:"author_id"`     // идентификатор автора публикации
	AuthorName string `bson:"author_name" json:"author_name"` // имя автора публикации
	CreatedAt  int64  `bson:"created_at" json:"created_at"`   // время создания публикации Unix
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(Post) error  // удаление публикации по ID
	Close()                 // закрывает соединение
}
