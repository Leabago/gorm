package main

import (
	"fmt"
	"go-orm/model"
)

func main() {
	fmt.Println("kek")

	a := App{}
	a.Initialize("test.db")
	a.initializeRoutes()

	a.DB.AutoMigrate(&model.Note{})

	// Create

	a.DB.Create(&model.Note{
		Title:       "Title1",
		Description: "Description1",
	})

	var note model.Note
	a.DB.First(&note, 1)

	a.DB.Model(&note).Update("Title", "Title1-Update")

	// // Migrate the schema
	// db.AutoMigrate(&model.Product{})

	// // Create
	// db.Create(&model.Product{Code: "D42", Price: 100})

	// // Read
	// var product model.Product
	// // db.First(&product, 2)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(model.Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)

	a.Run(":8080")
}
