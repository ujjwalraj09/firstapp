package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

var organizedData map[string]map[string][]string

func main() {
	// Load XLSX file
	f, err := excelize.OpenFile("Sample-Menu.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// Define a map to store column data
	columnData := make(map[string][]string)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through each row
	for _, row := range rows {
		// Iterate through each cell in the row
		for colIndex, cell := range row {
			// Convert the cell coordinate to column name
			colName, err := excelize.ColumnNumberToName(colIndex + 1)
			if err != nil {
				log.Fatal(err)
			}
			// Append cell value to the corresponding column array in the map
			columnData[colName] = append(columnData[colName], cell)
		}
	}

	// Create a map to store organized data
	organizedData = make(map[string]map[string][]string)

	// Iterate over each map with key starting with "A"
	for _, arr := range columnData {
		// Extract the day and date from the data
		day := arr[0]
		date := arr[1]

		// Initialize the inner map for this day
		organizedData[day] = make(map[string][]string)
		organizedData[day]["date"] = []string{date}

		// Find the indices for breakfast, lunch, and dinner
		breakfastStart := -1
		lunchStart := -1
		dinnerStart := -1
		for i, value := range arr {
			switch strings.ToLower(value) {
			case "breakfast":
				breakfastStart = i + 1
			case "lunch":
				lunchStart = i + 1
			case "dinner":
				dinnerStart = i + 1
			}
		}

		// Populate the breakfast, lunch, and dinner arrays in the inner map
		if breakfastStart != -1 && lunchStart != -1 && dinnerStart != -1 {
			organizedData[day]["breakfast"] = arr[breakfastStart:lunchStart]
			organizedData[day]["lunch"] = arr[lunchStart:dinnerStart]
			organizedData[day]["dinner"] = arr[dinnerStart:]
		}
	}

	// Menu instance
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. GetMenu")
	fmt.Println("2. CountMenuItems")
	fmt.Println("3. IsItemInMeal")
	fmt.Println("4. SaveMenuAsJSON")
	fmt.Println("5. PrintDetails")
	fmt.Println("The inputs are case sensitive")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var dayInput, mealInput string
		fmt.Println("Enter the day:")
		fmt.Scanln(&dayInput)
		fmt.Println("Enter the meal:")
		fmt.Scanln(&mealInput)
		items, ok := GetMenuItems(dayInput, mealInput)
		if ok {
			fmt.Printf("%s %s menu: %v\n", dayInput, mealInput, items)
		} else {
			fmt.Printf("Menu not found for %s %s\n", dayInput, mealInput)
		}
	case 2:
		var dayInput, mealInput string
		fmt.Println("Enter the day:")
		fmt.Scanln(&dayInput)
		fmt.Println("Enter the meal:")
		fmt.Scanln(&mealInput)
		numItems := CountMenuItems(dayInput, mealInput)
		fmt.Printf("Number of items in %s %s: %d\n", dayInput, mealInput, numItems)
	case 3:
		var dayInput, mealInput, itemInput string

		fmt.Println("Enter the day:")
		fmt.Scanln(&dayInput)
		fmt.Println("Enter the meal:")
		fmt.Scanln(&mealInput)
		fmt.Println("Enter the item:")
		fmt.Scanln(&itemInput)

		isInMeal := IsItemInMeal(dayInput, mealInput, itemInput)
		fmt.Printf("Is '%s' in %s %s? %t\n", itemInput, dayInput, mealInput, isInMeal)
	case 4:
		err := SaveMenuAsJSON("menu.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Menu saved as JSON successfully!")
	case 5:
		for day, meals := range organizedData {
			for meal, items := range meals {
				date := organizedData[day]["date"][0]
				instance := MealInstance{
					Day:   day,
					Date:  date,
					Meal:  meal,
					Items: items,
				}
				instance.PrintDetails()
			}
		}

	default:
		fmt.Println("Invalid choice")
	}
}

func GetMenuItems(day, meal string) ([]string, bool) {
	dayMenu, found := organizedData[day]
	if !found {
		return nil, false
	}
	items, found := dayMenu[meal]
	if !found {
		return nil, false
	}
	return items, true
}

func CountMenuItems(day, meal string) int {
	items, found := GetMenuItems(day, meal)
	if !found {
		return 0
	}
	return len(items)
}

func IsItemInMeal(day, meal, item string) bool {
	items, found := GetMenuItems(day, meal)
	if !found {
		return false
	}
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func SaveMenuAsJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(organizedData); err != nil {
		return err
	}

	return nil
}

type MealInstance struct {
	Day   string
	Date  string
	Meal  string
	Items []string
}

func (m MealInstance) PrintDetails() {
	fmt.Printf("%s %s menu:\n", m.Day, m.Date)
	fmt.Printf("  Meal: %s\n", m.Meal)
	fmt.Println("  Items:")
	for _, item := range m.Items {
		fmt.Printf("    %s\n", item)
	}
}
