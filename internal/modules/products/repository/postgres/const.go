package repository_postgres

const (
	queryCreateProduct = `
		INSERT INTO products (id, name, price, created_at) VALUES ($1, $2, $3, $4)
	`
	queryGetProductByID = `
		SELECT id, name, price FROM products WHERE id = $1
	`
	queryGetAllProduct = `
		SELECT id, name, price FROM products
	`
	queryUpdateProductByID = `
		UPDATE products set name = $2, price = $3 WHERE id = $1 RETURNING id, name, price
	`
	queryDeleteProductByID = `
		DELETE FROM products where id = $1
	`
)
