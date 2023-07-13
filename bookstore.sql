CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         author_name TEXT
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       book_name TEXT,
                       edition TEXT,
                       publication_year TEXT
);

CREATE TABLE books_authors (
                              book_id INT,
                              author_id INT,
                              FOREIGN KEY (book_id) REFERENCES books (id),
                              FOREIGN KEY (author_id) REFERENCES authors (id)
);

-- Insert authors (First populating before CSV)
INSERT INTO authors (author_name)
VALUES
    ('J.D. Salinger'),
    ('Herman Melville'),
    ('J.R.R. Tolkien'),
    ('Harper Lee'),
    ('George Orwell'),
    ('Jane Austen'),
    ('J.K. Rowling'),
    ('F. Scott Fitzgerald'),
    ('Leo Tolstoy'),
    ('Homer'),
    ('Lewis Carroll'),
    ('J.R.R. Tolkien'),
    ('Aldous Huxley'),
    ('Miguel de Cervantes'),
    ('C.S. Lewis'),
    ('Dante Alighieri'),
    ('Margaret Mitchell'),
    ('Mark Twain'),
    ('Charlotte Brontë'),
    ('Joseph Heller'),
    ('Mary Shelley'),
    ('Oscar Wilde'),
    ('Paulo Coelho'),
    ('William Golding'),
    ('Fyodor Dostoevsky'),
    ('John Steinbeck'),
    ('Emily Brontë'),
    ('Alexandre Dumas'),
    ('Fyodor Dostoevsky'),
    ('Gabriel García Márquez'),
    ('Leo Tolstoy'),
    ('Nathaniel Hawthorne'),
    ('Frances Hodgson Burnett'),
    ('Charles Dickens'),
    ('Charles Dickens'),
    ('Ernest Hemingway'),
    ('Victor Hugo'),
    ('Antoine de Saint-Exupéry');

-- Insert books
INSERT INTO books (book_name, edition, publication_year)
VALUES
    ('The Catcher in the Rye', 'First Edition', '1951'),
    ('Moby-Dick', 'Original Edition', '1851'),
    ('The Lord of the Rings', 'Anniversary Edition', '1954'),
    ('To Kill a Mockingbird', 'First Edition', '1960'),
    ('1984', 'First Edition', '1949'),
    ('Pride and Prejudice', 'First Edition', '1813'),
    ('Harry Potter and the Sorcerer''s Stone', 'First Edition', '1997'),
    ('The Great Gatsby', 'First Edition', '1925'),
    ('War and Peace', 'First Edition', '1869'),
    ('The Odyssey', 'First Edition', '8th century BC'),
    ('Alice''s Adventures in Wonderland', 'First Edition', '1865'),
    ('The Hobbit', 'First Edition', '1937'),
    ('Brave New World', 'First Edition', '1932'),
    ('Don Quixote', 'First Edition', '1605'),
    ('The Chronicles of Narnia', 'First Edition', '1950'),
    ('The Divine Comedy', 'First Edition', '1320'),
    ('Gone with the Wind', 'First Edition', '1936'),
    ('The Adventures of Huckleberry Finn', 'First Edition', '1884'),
    ('Jane Eyre', 'First Edition', '1847'),
    ('Catch-22', 'First Edition', '1961'),
    ('Frankenstein', 'First Edition', '1818'),
    ('The Picture of Dorian Gray', 'First Edition', '1890'),
    ('The Alchemist', 'First Edition', '1988'),
    ('Lord of the Flies', 'First Edition', '1954'),
    ('Crime and Punishment', 'First Edition', '1866'),
    ('The Grapes of Wrath', 'First Edition', '1939'),
    ('Wuthering Heights', 'First Edition', '1847'),
    ('The Count of Monte Cristo', 'First Edition', '1844'),
    ('The Brothers Karamazov', 'First Edition', '1880'),
    ('One Hundred Years of Solitude', 'First Edition', '1967'),
    ('Anna Karenina', 'First Edition', '1878'),
    ('The Scarlet Letter', 'First Edition', '1850'),
    ('The Secret Garden', 'First Edition', '1911'),
    ('Great Expectations', 'First Edition', '1861'),
    ('A Tale of Two Cities', 'First Edition', '1859'),
    ('The Old Man and the Sea', 'First Edition', '1952'),
    ('Les Misérables', 'First Edition', '1862'),
    ('The Little Prince', 'First Edition', '1943'),
    ('The Catcher in the Rye', 'First Edition', '1951'),
    ('The Brothers Karamazov', 'First Edition', '1880');

-- Insert data into books_authors
INSERT INTO books_authors (book_id, author_id)
VALUES
    -- The Catcher in the Rye by J.D. Salinger
    (1, 1),
    -- Moby-Dick by Herman Melville
    (2, 2),
    -- The Lord of the Rings by J.R.R. Tolkien
    (3, 3),
    -- To Kill a Mockingbird by Harper Lee
    (4, 4),
    -- 1984 by George Orwell
    (5, 5),
    -- Pride and Prejudice by Jane Austen
    (6, 6),
    -- Harry Potter and the Sorcerer's Stone by J.K. Rowling
    (7, 7),
    -- The Great Gatsby by F. Scott Fitzgerald
    (8, 8),
    -- War and Peace by Leo Tolstoy
    (9, 9),
    -- The Odyssey by Homer
    (10, 10),
    -- Alice's Adventures in Wonderland by Lewis Carroll
    (11, 11),
    -- The Hobbit by J.R.R. Tolkien
    (12, 12),
    -- Brave New World by Aldous Huxley
    (13, 13),
    -- Don Quixote by Miguel de Cervantes
    (14, 14),
    -- The Chronicles of Narnia by C.S. Lewis
    (15, 15),
    -- The Divine Comedy by Dante Alighieri
    (16, 16),
    -- Gone with the Wind by Margaret Mitchell
    (17, 17),
    -- The Adventures of Huckleberry Finn by Mark Twain
    (18, 18),
    -- Jane Eyre by Charlotte Brontë
    (19, 19),
    -- Catch-22 by Joseph Heller
    (20, 20),
    -- Frankenstein by Mary Shelley
    (21, 21),
    -- The Picture of Dorian Gray by Oscar Wilde
    (22, 22),
    -- The Alchemist by Paulo Coelho
    (23, 23),
    -- Lord of the Flies by William Golding
    (24, 24),
    -- Crime and Punishment by Fyodor Dostoevsky
    (25, 25),
    -- The Grapes of Wrath by John Steinbeck
    (26, 26),
    -- Wuthering Heights by Emily Brontë
    (27, 27),
    -- The Count of Monte Cristo by Alexandre Dumas
    (28, 28),
    -- The Brothers Karamazov by Fyodor Dostoevsky
    (29, 29),
    -- One Hundred Years of Solitude by Gabriel García Márquez
    (30, 30),
    -- Anna Karenina by Leo Tolstoy
    (31, 9),
    -- The Scarlet Letter by Nathaniel Hawthorne
    (32, 31),
    -- The Secret Garden by Frances Hodgson Burnett
    (33, 32),
    -- Great Expectations by Charles Dickens
    (34, 33),
    -- A Tale of Two Cities by Charles Dickens
    (35, 33),
    -- The Old Man and the Sea by Ernest Hemingway
    (36, 34),
    -- Les Misérables by Victor Hugo
    (37, 35),
    -- The Little Prince by Antoine de Saint-Exupéry
    (38, 36),
    -- Moby-Dick by Herman Melville (additional author)
    (2, 37),
    -- The Adventures of Huckleberry Finn by Mark Twain (additional author)
    (18, 38);