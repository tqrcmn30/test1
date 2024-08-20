// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: cart.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING cart_id, cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id
`

type CreateCartParams struct {
	CartUserID   *int32           `json:"cart_user_id"`
	CartCoursID  *int32           `json:"cart_cours_id"`
	CartQty      *int32           `json:"cart_qty"`
	CartPrice    pgtype.Numeric   `json:"cart_price"`
	CartModified pgtype.Timestamp `json:"cart_modified"`
	CartStatus   *string          `json:"cart_status"`
	CartCartID   *int32           `json:"cart_cart_id"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (*Cart, error) {
	row := q.db.QueryRow(ctx, createCart,
		arg.CartUserID,
		arg.CartCoursID,
		arg.CartQty,
		arg.CartPrice,
		arg.CartModified,
		arg.CartStatus,
		arg.CartCartID,
	)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CartUserID,
		&i.CartCoursID,
		&i.CartQty,
		&i.CartPrice,
		&i.CartModified,
		&i.CartStatus,
		&i.CartCartID,
	)
	return &i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM carts WHERE cart_id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, cartID int32) error {
	_, err := q.db.Exec(ctx, deleteCart, cartID)
	return err
}

const getAllCarts = `-- name: GetAllCarts :many
SELECT cart_id, cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id FROM carts
`

func (q *Queries) GetAllCarts(ctx context.Context) ([]*Cart, error) {
	rows, err := q.db.Query(ctx, getAllCarts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Cart
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.CartID,
			&i.CartUserID,
			&i.CartCoursID,
			&i.CartQty,
			&i.CartPrice,
			&i.CartModified,
			&i.CartStatus,
			&i.CartCartID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCartByID = `-- name: GetCartByID :one
SELECT cart_id, cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id FROM carts WHERE cart_id = $1
`

func (q *Queries) GetCartByID(ctx context.Context, cartID int32) (*Cart, error) {
	row := q.db.QueryRow(ctx, getCartByID, cartID)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CartUserID,
		&i.CartCoursID,
		&i.CartQty,
		&i.CartPrice,
		&i.CartModified,
		&i.CartStatus,
		&i.CartCartID,
	)
	return &i, err
}

const updateCart = `-- name: UpdateCart :one
UPDATE carts
SET cart_user_id = $2, cart_cours_id = $3, cart_qty = $4, cart_price = $5, cart_modified = $6, cart_status = $7, cart_cart_id = $8
WHERE cart_id = $1
RETURNING cart_id, cart_user_id, cart_cours_id, cart_qty, cart_price, cart_modified, cart_status, cart_cart_id
`

type UpdateCartParams struct {
	CartID       int32            `json:"cart_id"`
	CartUserID   *int32           `json:"cart_user_id"`
	CartCoursID  *int32           `json:"cart_cours_id"`
	CartQty      *int32           `json:"cart_qty"`
	CartPrice    pgtype.Numeric   `json:"cart_price"`
	CartModified pgtype.Timestamp `json:"cart_modified"`
	CartStatus   *string          `json:"cart_status"`
	CartCartID   *int32           `json:"cart_cart_id"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) (*Cart, error) {
	row := q.db.QueryRow(ctx, updateCart,
		arg.CartID,
		arg.CartUserID,
		arg.CartCoursID,
		arg.CartQty,
		arg.CartPrice,
		arg.CartModified,
		arg.CartStatus,
		arg.CartCartID,
	)
	var i Cart
	err := row.Scan(
		&i.CartID,
		&i.CartUserID,
		&i.CartCoursID,
		&i.CartQty,
		&i.CartPrice,
		&i.CartModified,
		&i.CartStatus,
		&i.CartCartID,
	)
	return &i, err
}
