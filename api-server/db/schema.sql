CREATE TABLE IF NOT EXISTS products (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  starred BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY,
  title TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS product_categories (
  product_id INTEGER NOT NULL,
  category_id INTEGER NOT NULL,
  PRIMARY KEY (product_id, category_id),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);