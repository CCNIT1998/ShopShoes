package controller

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/TechMaster/golang/15GoMySQL/model"
	"github.com/TechMaster/golang/15GoMySQL/repo"
	"github.com/gofiber/fiber/v2"
)

func GetProductsAll(c *fiber.Ctx) error {
	var products []model.Product

	result := repo.Db.Preload("Category").Find(&products)
	fmt.Println(result)
	if result.Error != nil {
		return c.JSON(result.Error)
	} else {
		return c.JSON(products)
	}
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product := repo.FindProductById(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	return c.JSON(product)
}

func GetSearchProducts(c *fiber.Ctx) error {
	var products []model.Product
	var products2 []model.Product
	// var products2 []model.Product

	limit, _ := strconv.Atoi(c.Query("limit"))
	newest, _ := strconv.Atoi(c.Query("newest"))
	order := c.Query("order")
	categorys, _ := strconv.Atoi(c.Query("categorys"))

	// category := c.Query("category")
	repo.Db.Preload("Category").Find(&products)
	if order == "desc" {
		sort.Slice(products[:], func(i, j int) bool {
			return products[i].PriceCurrent > products[j].PriceCurrent
		})
	} else if order == "asc" {
		sort.Slice(products[:], func(i, j int) bool {
			return products[i].PriceCurrent < products[j].PriceCurrent
		})
	}

	if categorys == 1 {

		for _, product := range products {
			if product.CategoryID == 1 {
				products2 = append(products2, product)
			}
		}
		products = products2
	}

	var results []model.Product
	if limit <= len(products) {
		if newest <= len(products) {
			if newest+limit > len(products)-1 {
				results = products[newest:]
			} else {
				results = products[newest : newest+limit]
			}
		} else {
			results = products
		}
	} else {
		results = []model.Product{}
	}

	data := make(map[string]interface{})
	data["items"] = results
	data["total_products"] = len(products)
	if len(products) >= limit {
		if len(products)%limit == 0 {
			data["total_pages"] = len(products) / limit
		} else {
			data["total_pages"] = len(products)/limit + 1
		}
	} else {
		data["total_pages"] = 1
	}
	return c.JSON(data)
}

//url/products?cate=...&page=...
func GetSearchSortProducts(c *fiber.Ctx) error {
	var products []model.Product

	// limit, _ := strconv.Atoi(c.Query("limit"))
	// newest, _ := strconv.Atoi(c.Query("newest"))
	// category := c.Query("category")
	repo.Db.Preload("Category").Find(&products)
	// var results []model.Product
	sort.Slice(products[:], func(i, j int) bool {
		return products[i].PriceCurrent > products[j].PriceCurrent
	})
	fmt.Println(products[1].PriceCurrent)

	return c.JSON(products)
}
