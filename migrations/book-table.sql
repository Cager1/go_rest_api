DROP TABLE IF EXISTS book;
CREATE TABLE book (
                       id         INT AUTO_INCREMENT NOT NULL,
                       title      VARCHAR(128) NOT NULL UNIQUE,
                       author     VARCHAR(255) NOT NULL,
                       quantity      DECIMAL(5,2) NOT NULL,
                       PRIMARY KEY (`id`)
);

INSERT INTO book
(title, author, quantity)
VALUES
    ('Blue Train', 'John Coltrane', 56),
    ('Giant Steps', 'John Coltrane', 63),
    ('Jeru', 'Gerry Mulligan', 17),
    ('Sarah Vaughan', 'Sarah Vaughan', 34);