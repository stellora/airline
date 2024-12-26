-- name: GetCategory :one
SELECT * FROM categories
WHERE id=? LIMIT 1;

-- name: ListCategorys :many
SELECT * FROM categories
ORDER BY id ASC;

-- name: CreateCategory :one
INSERT INTO categories (
  title
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateCategory :exec
UPDATE categories SET
title=?
WHERE id=?;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id=?;

-------------------------------------------------------------------------------

-- name: GetProduct :one
SELECT * FROM products
WHERE id=? LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id ASC;

-- name: CreateProduct :one
INSERT INTO products (
  title, starred
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE products SET
title=?,
starred=?
WHERE id=?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id=?;

-------------------------------------------------------------------------------

-- name: AddProductToCategory :exec
INSERT INTO product_categories (
  product_id,
  category_id
) VALUES (
  ?, ?
);

-- name: RemoveProductFromCategory :exec
DELETE FROM product_categories
WHERE product_id = ? AND category_id = ?;
