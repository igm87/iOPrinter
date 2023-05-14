package iOPrinter

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"errors"
	"fmt"
)

const fZero float64 = 0.0 

func GetFloat()(float64,error){
	myErr := errors.New("That ain't no float")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter float value ")
	value, err := reader.ReadString('\n')
	if err != nil{
		return fZero,myErr
	}
	fvalue, err := strconv.ParseFloat(strings.TrimSpace(value),64)
	if err != nil{
		return fZero,err
	}
	return fvalue,nil
}

