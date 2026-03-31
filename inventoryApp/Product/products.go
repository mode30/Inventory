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
	if product.name==""||product.description==""||product.categoryId==0||product.sellingPrice==0{
		return nil,errors.New("cannot return empty format")
	}else{
		return &Product{
			product.productId:product.productId+1,
			product.name:product.name,
		},nil
	}
}
