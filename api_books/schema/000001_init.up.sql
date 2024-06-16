CREATE TABLE authors (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name VARCHAR(100) NOT NULL,
  biography TEXT,
  date_of_birth DATE
);

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author_id INT NOT NULL REFERENCES authors(id),
  publication_year INT,
  isbn VARCHAR(13) UNIQUE NOT NULL
);