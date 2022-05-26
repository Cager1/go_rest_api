DROP TABLE IF EXISTS book;
CREATE TABLE book (
                       id         INT AUTO_INCREMENT NOT NULL,
                       title      VARCHAR(128) NOT NULL UNIQUE,
                       author     VARCHAR(255) NOT NULL,
                       PRIMARY KEY (`id`)
);

# INSERT INTO book
# (title, author, quantity)
# VALUES
#     ('Blue Train', 'John Coltrane'),
#     ('Giant Steps', 'John Coltrane', 63),
#     ('Jeru', 'Gerry Mulligan', 17),
#     ('Sarah Vaughan', 'Sarah Vaughan', 34);