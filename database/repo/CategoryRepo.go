package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initCategory() {
	 categories := map[int]string{
		 1:"Giày thể thao nam",
		 2:"Giày thể thao nữ",
		 3:"Giày da nam",
		 4:"Giày da nữ",
	 }

	for _, category := range categories {
		createCategory := model.Category{Name: category}
		result := Db.Create(&createCategory)
		if result.Error != nil {
			panic(result.Error)
		}
	}

}
