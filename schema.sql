-- Схема базы данных сайта новостного агрегатора
DROP TABLE IF EXISTS posts, authors;

-- Авторы публикаций
CREATE TABLE IF NOT EXISTS authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Публикации
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT DEFAULT EXTRACT(epoch FROM clock_timestamp())
);

-- Очистка таблиц перед начальным заполнение БД
TRUNCATE TABLE posts, authors;

-- Начальное заполнение таблиц БД
INSERT INTO authors (name) VALUES
    ('Михаил Воскресенский'),
    ('Никита Абрамов'),
    ('Екатерина Карпова'),
    ('Максим Ветров'),
    ('Александр Иванов');

INSERT INTO posts (author_id, title, content) VALUES
    (1, 'Число жертв урагана "Ида" на юго-востоке США возросло до шести', 'Содержание публикации 1 ...'),
    (3, 'Apple отложит презентацию смарт-часов нового поколения', 'Содержание публикации 2 ...'),
    (5, 'В Латвии заявили, что не откажутся от перевода всех школ на латышский язык', 'Содержание публикации 3 ...'),
    (4, 'В Хельсинки ограничили скорость электросамокатов', 'Содержание публикации 4 ...'),
    (3, 'США и Россия готовят новые контакты по кибербезопасности, заявил Лавров', 'Содержание публикации 5 ...');