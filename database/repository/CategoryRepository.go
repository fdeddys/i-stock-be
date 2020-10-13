package repository

import (
	"com.ddabadi/estock/database"
	"com.ddabadi/estock/models/dbmodels"
)

func GetAllCategory() ([]dbmodels.ProductCategory, error) {
	db := database.GetDbCon()

	var productCategories []dbmodels.ProductCategory

	err := db.Find(&productCategories).Error

	return productCategories, err
}
