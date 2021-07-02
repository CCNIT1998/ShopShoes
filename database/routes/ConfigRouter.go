package routes

import (
	"github.com/TechMaster/golang/15GoMySQL/controller"
	"github.com/gofiber/fiber/v2"
)

func ConfigRouter(app *fiber.App) {
	productAPI := app.Group("/api/product")
	routeProduct(&productAPI)
}



func routeProduct(router *fiber.Router) {
	(*router).Get("/", controller.GetProductsAll) //Liệt kê
	(*router).Get("/:id", controller.GetProductById)// tim theo id
	// (*router).Get("/:idCategory", controller.GetProductsWithCategories)// tim theo id
	// (*router).Get("/:product", controller.GetProductsWithCategories)
	(*router).Get("/search/search_items", controller.GetSearchProducts)
	(*router).Get("/search/search_items/sort", controller.GetSearchSortProducts)
}


