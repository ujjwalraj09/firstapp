package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func LoadMenu(filePath string) [][]string {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	menu, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return menu
}

// GetItemsForMeal returns the items available for a particular meal on a given day
func GetItemsForMeal(menu [][]string, day, meal string) []string {
	var items []string
	mealIndex := -1

	// Find the column index for the specified meal
	for i, header := range menu[0] {
		if strings.EqualFold(header, meal) {
			mealIndex = i
			break
		}
	}

	if mealIndex == -1 {
		fmt.Println("Meal not found")
		return nil
	}

	for _, row := range menu {
		if strings.EqualFold(row[0], day) {
			items = append(items, strings.TrimSpace(row[mealIndex]))
		}
	}

	return items
}

func CountItemsForMeal(menu [][]string, day, meal string) int {
	items := GetItemsForMeal(menu, day, meal)
	return len(items)
}

func CheckItemInMeal(menu [][]string, day, meal, item string) bool {
	items := GetItemsForMeal(menu, day, meal)
	for _, i := range items {
		if strings.EqualFold(i, item) {
			return true
		}
	}
	return false
}

func ConvertToJSON(menu [][]string, jsonFilePath string) {
	jsonData := make(map[string]interface{})
	for _, row := range menu[1:] {
		if len(row) != len(menu[0]) {
			fmt.Println("Error: Row length does not match the header length.")
			return
		}
		day := row[0]
		meals := make(map[string]interface{})
		for i, meal := range menu[0][1:] {
			meals[meal] = strings.TrimSpace(row[i+1])
		}
		jsonData[day] = meals
	}

	jsonFile, err := os.Create(jsonFilePath)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(jsonData)
	if err != nil {
		fmt.Println("Error encoding JSON data:", err)
		return
	}
	fmt.Println("JSON file created successfully:", jsonFilePath)
}

func main() {

	menuData := LoadMenu("Sample-Menu.xlsx")

	fmt.Println("1) Items available for Breakfast on Monday:", GetItemsForMeal(menuData, "Monday", "Breakfast"))
	fmt.Println("2) Number of items for Dinner on Tuesday:", CountItemsForMeal(menuData, "Tuesday", "Dinner"))
	fmt.Println("3) Is 'Pancakes' available in Breakfast on Wednesday?", CheckItemInMeal(menuData, "Wednesday", "Breakfast", "Pancakes"))

	ConvertToJSON(menuData, "menu_data.json")
}
