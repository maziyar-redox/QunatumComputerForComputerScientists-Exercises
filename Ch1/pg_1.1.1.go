package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseComplex(osArgument []string) (isMatched bool, favError error) {
	for _, val := range osArgument[1:] {
		match, err := regexp.MatchString(`^[+-]?(\d+(\.\d*)?|\.\d+)[+-](\d+(\.\d*)?|\.\d+)i$`, val)
		if err != nil {
			return false, errors.New("There was an error")
		}
		if match == false {
			return false, errors.New("Format of complex number must be like a+bi, exp: 0+0i")
		}
	}
	return true, nil
}

func sumComplex(osArgument []string) (sum string, favError error) {
	realPartSlice := []int{}
	imPartSlice := []int{}
	for _, val := range osArgument[1:] {
		infixNotation := "+"
		indexOfSign := strings.LastIndex(val, infixNotation)
		if indexOfSign == -1 {
			infixNotation = "-"
			indexOfSign = strings.LastIndex(val, infixNotation)
		} else {
			infixNotation = "+"
			indexOfSign = strings.LastIndex(val, infixNotation)
		}
		splitRealNumber := val[:indexOfSign]
		convertRealToInt, err := strconv.Atoi(splitRealNumber)
		if err != nil {
			return "", errors.New("There was an error while converting real number to int")
		}
		realPartSlice = append(realPartSlice, convertRealToInt)
		splitImNumber := strings.Split(val[indexOfSign:], "i")
		convertImToInt, err := strconv.Atoi(splitImNumber[0])
		if err != nil {
			return "", errors.New("There was an error while converting im number to int")
		}
		imPartSlice = append(imPartSlice, convertImToInt)
	}
	var realSum, imSum int
	var isNegative string = ""
	for index := range len(osArgument[1:]) {
		realSum = realSum + realPartSlice[index]
		imSum = imSum + imPartSlice[index]
	}
	if imSum > 0 || imSum == 0 {
		isNegative = "+"
	}
	return strconv.Itoa(int(realSum)) + isNegative + strconv.Itoa(int(imSum)) + "i", nil
}

func productComplex(osArgument []string) (product string, favError error) {
	realPartSlice := []int{}
	imPartSlice := []int{}
	for _, val := range osArgument[1:] {
		infixNotation := "+"
		indexOfSign := strings.LastIndex(val, infixNotation)
		if indexOfSign == -1 {
			infixNotation = "-"
			indexOfSign = strings.LastIndex(val, infixNotation)
		} else {
			infixNotation = "+"
			indexOfSign = strings.LastIndex(val, infixNotation)
		}
		splitRealNumber := val[:indexOfSign]
		convertRealToInt, err := strconv.Atoi(splitRealNumber)
		if err != nil {
			return "", errors.New("There was an error while converting real number to int")
		}
		realPartSlice = append(realPartSlice, convertRealToInt)
		splitImNumber := strings.Split(val[indexOfSign:], "i")
		convertImToInt, err := strconv.Atoi(splitImNumber[0])
		if err != nil {
			return "", errors.New("There was an error while converting im number to int")
		}
		imPartSlice = append(imPartSlice, convertImToInt)
	}

	var isNegative string = ""

	var realProductFirst int
	var imProductFirst int
	for _, value := range realPartSlice[1:] {
		realProductFirst += realPartSlice[0] * value
	}
	for _, value := range imPartSlice[1:] {
		imProductFirst += realPartSlice[0] * value
	}
	var realProductSeccond int
	var imProductSeccond int
	for _, value := range realPartSlice[1:] {
		realProductSeccond += imPartSlice[0] * value
	}
	for _, value := range imPartSlice[1:] {
		imProductSeccond += (-1) * (imPartSlice[0] * value)
	}
	finalResultIm := imProductFirst + realProductSeccond
	if finalResultIm > 0 || finalResultIm == 0 {
		isNegative = "+"
	}
	finalResultReal := imProductSeccond + realProductFirst
	return strconv.Itoa(finalResultReal) + isNegative + strconv.Itoa(finalResultIm) + "i", nil
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide more arguments")
		return
	}
	_, favError := parseComplex(arguments)
	if favError != nil {
		fmt.Println("There was an error")
		return
	}
	sum, err := sumComplex(arguments)
	if err != nil {
		fmt.Println("There was an error while trying to sum")
		return
	}
	product, err := productComplex(arguments)
	if err != nil {
		fmt.Println("There was an error while trying to product")
		return
	}
	fmt.Println("Sum of your complex numbers is :", sum)
	fmt.Println("Product of your complex numbers is :", product)
}