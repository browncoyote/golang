//Real growth of top 10 nations

package main

import (
    "fmt"
    "math"
)
	var nation string
	var years, pop, rate, sum float64
func main() {
	fmt.Println("Welcome to Real Growth")
Nation:	
	fmt.Println("Please enter a nation:")
	fmt.Scan(&nation)
	fmt.Println("How many years do you want to predict?")
	fmt.Scan(&years)
	china := []string{"China", "china"} 
	india := []string{"India", "india"}
	usa := []string{"America", "america", "USA", "usa", "US", "us", "United States"}
	indonesia := []string{"indonesia", "Indonesia", "Indo Asia"}
	brazil := []string{"Brazil", "brazil", "brasil", "Brasil"}
	pakistan := []string{"paki", "Pakistan", "pakistan"}
	nigeria := []string{"Nigeria", "nigeria"}
	bangladesh := []string{"Bangladesh", "bangladesh"}
	russia := []string{"Russia", "russia", "Russian Federation"}
	japan := []string{"Nippon", "nippon", "Japan", "japan"}
	if containsString(china, nation) {
		Growth(1385566537, 0.0061, years)
		goto Nation
	} else if containsString(india, nation) {
		Growth(1252139596, 0.0124, years)
		goto Nation
	} else if containsString(usa, nation) {
		Growth(320050716, 0.0081, years)
		goto Nation
	} else if containsString(indonesia, nation) {
		Growth(249865631, 0.0121, years)
		goto Nation
	} else if containsString(brazil, nation) {
		Growth(200361925, 0.0085, years)
		goto Nation
	} else if containsString(pakistan, nation) {
		Growth(182142594, 0.0166, years)
		goto Nation
	} else if containsString(nigeria, nation) {
		Growth(173615345, 0.0278, years)
		goto Nation
	} else if containsString(bangladesh, nation) {
		Growth(156594962, 0.0119, years)
		goto Nation
	} else if containsString(russia, nation) {
		Growth(142833689, -0.0021, years)
		goto Nation
	} else if containsString(japan, nation) {
		Growth(127143577, -0.0008, years)
		goto Nation
	} else {
		fmt.Println("Error! Pick a nation:")
		goto Nation
	}
}

func Growth(pop, rate, years float64) {
		sum = pop * pow((1.0 + rate), years, 2)
		fmt.Println("The current population of", nation, "is", pop)
		fmt.Println("In", years, "it will be", sum)
}

//two functions containsString and posString for study
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
