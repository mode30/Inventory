package product

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Product struct{
	productId int
	name string
	description string
	categoryId int
	costPrice float64
	sellingPrice float64
	isActive bool
	// weight float64
	// taxRate float64

}


func query(prompt string)(userInput string,err error){
	fmt.Println(prompt)
	buffer:=bufio.NewReader(os.Stdin)
	bufferString,err:=buffer.ReadString('\n')
	if err !=nil{
		fmt.Println(err)
		return
	}
	bufferTrimmed:=strings.TrimSpace(bufferString)
	return bufferTrimmed,nil
}





//function call to return float from user input
func parseFloat(prompt string)(parsedFloat64 float64,err error){
	parsedFloat,parsedFloatError:=query(prompt)
	if parsedFloatError != nil {
		return 0,errors.New("cannot take string to parse to float")
	}
	parsedValue,err:=strconv.ParseFloat(parsedFloat, 64)
	if err !=nil{
		return 0,errors.New("cannot parse to float")
	}
	return parsedValue,nil
}


//method to return a new constructor type Product
func (product *Product)newProduct()(products *Product,err error){

	product.name,err=query("enter product name:")
	if err !=nil{
		return &Product{},errors.New("cannot return product name")
	}
	product.description,err=query("enter product description")
	if err !=nil{
		return &Product{},errors.New("cannot return product description")

	}
	product.costPrice,err=parseFloat("enter price")
	if err != nil{
		return &Product{},errors.New("cannot parse float")
	}
	product.sellingPrice,err=parseFloat("enter selling price:")
	if err != nil{
		return &Product{},errors.New("cannot parse float")
	}
	product.isActive=true

	if product.name==""||product.description==""||product.categoryId==0||product.sellingPrice==0{
		return nil,errors.New("cannot return empty format")
	}else{
		return &Product{
			productId:product.productId+1,
			name:product.name,
			description:product.description,
			costPrice: product.costPrice,
			sellingPrice:product.sellingPrice,
			categoryId: product.categoryId+1,
		},nil
	}
}

func (product *Product)clearCache()(clearedProduct *Product,err error){
	return &Product{
		product.categoryId:0.0,
		product.productId:0
		product.name:""
		product.description:""
		product.costPrice:0.0
		product.sellingPrice:0.0
		product.isActive:false

	},nil
	 // return *Product,nil

}
