package main

import (
	"fmt"
)

// type Category struct{
// 	categoryId int
// 	name string
// 	description string
// }
//
type Product struct{
	productId int
	name string
	description string
	categoryId int
	costPrice float64
	sellingPrice float64
	isActive bool
	weight float64
	taxRate float64

}

func main(){
	var product1 Product
	var product2 Product
	var product3 Product
	var product4 Product
	var product5 Product

	product1=Product{
		productId: 1,
		name:"rice",
		description:"grain",
		categoryId:1,
		costPrice: 12.99,
		sellingPrice: 20.00,
		isActive :true,
		weight:0.9,
		taxRate: 0.1,

	}

	product2=Product{
		productId: 1,
		name:"rice",
		description:"grain",
		categoryId:1,
		costPrice: 12.99,
		sellingPrice: 20.00,
		isActive :true,
		weight:0.9,
		taxRate: 0.1,

	}

	product3=Product{
		productId: 1,
		name:"rice",
		description:"grain",
		categoryId:1,
		costPrice: 12.99,
		sellingPrice: 20.00,
		isActive :true,
		weight:0.9,
		taxRate: 0.1,

	}

	product4=Product{
		productId: 1,
		name:"rice",
		description:"grain",
		categoryId:1,
		costPrice: 12.99,
		sellingPrice: 20.00,
		isActive :true,
		weight:0.9,
		taxRate: 0.1,

	}

	product5=Product{
		productId: 1,
		name:"rice",
		description:"grain",
		categoryId:1,
		costPrice: 12.99,
		sellingPrice: 20.00,
		isActive :true,
		weight:0.9,
		taxRate: 0.1,

	}
	products=[]Product{
		product1,
		product2,
		product3,
		product4,
		product5,

	}

}
