package services

import (
	"log"

	"com.ddabadi/estock/constants"
	"com.ddabadi/estock/database/repository"
	"com.ddabadi/estock/models"
)

type ProductCategoryServiceInterface struct {
}

func InitProductCategoryServiceInterface() *ProductCategoryServiceInterface {
	return &ProductCategoryServiceInterface{}
}

func (service *ProductCategoryServiceInterface) GetAll() models.Response {
	var res models.Response

	productCategories, err := repository.GetAllCategory()

	if err != nil {
		log.Println("err get from database : ", err)

		res.Rc = constants.ERR_CODE_11
		res.Msg = constants.ERR_CODE_11_MSG
		return res
	}

	log.Println("get data : ", res)

	res.Rc = constants.ERR_CODE_00
	res.Msg = constants.ERR_CODE_00_MSG
	res.Data = productCategories

	return res
}
