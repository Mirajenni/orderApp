package db

import "fmt"

func Create(product ProductDTO) (err error) {

	result := db.Create(&ProductDTO{Name: product.Name, Price: product.Price})

	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	return nil
}

func Get(name string) (product ProductDTO, err error) {
	var p ProductDTO

	result := db.Where("name = ?", name).First(&p)

	if result.Error != nil {
		fmt.Println(result.Error)
		return p, result.Error
	}

	return p, nil
}

func GetAll() (products []ProductDTO, err error) {
	var p []ProductDTO

	result := db.Find(&p)

	if result.Error != nil {
		fmt.Println(result.Error)
		return p, result.Error
	}

	return p, nil
}
