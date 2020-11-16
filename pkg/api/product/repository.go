package product

import (
	"coding-challenge-go/pkg/api/seller"
	"database/sql"
)

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (r *repository) delete(product *product) error {
	rows, err := r.db.Query("DELETE FROM product WHERE uuid = ?", product.UUID)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) insert(product *product) error {
	rows, err := r.db.Query(
		"INSERT INTO product (id_product, name, brand, stock, fk_seller, uuid) VALUES(?,?,?,?,(SELECT id_seller FROM seller WHERE uuid = ?),?)",
		product.ProductID, product.Name, product.Brand, product.Stock, product.SellerUUID, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) update(product *product) error {
	rows, err := r.db.Query(
		"UPDATE product SET name = ?, brand = ?, stock = ? WHERE uuid = ?",
		product.Name, product.Brand, product.Stock, product.UUID,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *repository) list(offset int, limit int) ([]*product, error) {
	rows, err := r.db.Query(
		"SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p "+
			"INNER JOIN seller s ON(s.id_seller = p.fk_seller) LIMIT ? OFFSET ?",
		limit, offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*product

	for rows.Next() {
		product := &product{}

		err = rows.Scan(&product.ProductID, &product.Name, &product.Brand, &product.Stock, &product.SellerUUID, &product.UUID)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *repository) findByUUID(uuid string) (*product, error) {
	rows, err := r.db.Query(
		"SELECT p.id_product, p.name, p.brand, p.stock, s.uuid, p.uuid FROM product p "+
			"INNER JOIN seller s ON(s.id_seller = p.fk_seller) WHERE p.uuid = ?",
		uuid,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	product := &product{}

	err = rows.Scan(&product.ProductID, &product.Name, &product.Brand, &product.Stock, &product.SellerUUID, &product.UUID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) getTopSellers(resultsCount int) ([]*seller.Seller, error) {
	rows, err := r.db.Query(
		"SELECT s.id_seller, s.name, s.email, s.phone, s.uuid FROM seller s "+
			"INNER JOIN (SELECT fk_seller, COUNT(*) top FROM product.product GROUP BY fk_seller) p "+
			"ON (p.fk_seller = s.id_seller) ORDER BY p.top DESC LIMIT ?", resultsCount,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var topSellers []*seller.Seller
	for rows.Next() {
		seller := &seller.Seller{}
		err = rows.Scan(&seller.SellerID, &seller.Name, &seller.Email, &seller.Phone, &seller.UUID)
		if err != nil {
			return nil, err
		}

		topSellers = append(topSellers, seller)
	}

	return topSellers, nil
}
