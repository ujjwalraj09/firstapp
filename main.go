package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

type DayOfWeek string
type Meal string

var Menu = map[DayOfWeek]map[string][]string{}

func parseExcel(filePath string) error {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}

	currentDay := ""
	currentMeal := ""
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			cellValue := strings.TrimSpace(row.Cells[0].String())
			if cellValue != "" {
				currentDay = cellValue
				Menu[DayOfWeek(currentDay)] = make(map[string][]string)
			} else {
				cellValue = strings.TrimSpace(row.Cells[1].String())
				if cellValue == "Date" {
					date := strings.TrimSpace(row.Cells[2].String())
					Menu[DayOfWeek(currentDay)]["Date"] = []string{date}
				} else {
					cellValue = strings.TrimSpace(row.Cells[2].String())
					if cellValue != "" {
						currentMeal = cellValue
						Menu[DayOfWeek(currentDay)][string(currentMeal)] = []string{}
					} else {
						item := strings.TrimSpace(row.Cells[3].String())
						if item != "" {
							Menu[DayOfWeek(currentDay)][string(currentMeal)] = append(Menu[DayOfWeek(currentDay)][string(currentMeal)], item)
						}
					}
				}
			}
		}
	}

	return nil
}

func GetMenuItems(day DayOfWeek, meal Meal) ([]string, bool) {
	items, ok := Menu[day][string(meal)]
	return items, ok
}

func CountMenuItems(day DayOfWeek, meal Meal) int {
	return len(Menu[day][string(meal)])
}

func IsItemInMeal(day DayOfWeek, meal Meal, item string) bool {
	items := Menu[day][string(meal)]
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

	jsonData, err := json.MarshalIndent(Menu, "", "    ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Println("Menu successfully saved as", filename)
	return nil
}

type MealInstance struct {
	Day   DayOfWeek
	Date  string
	Meal  Meal
	Items []string
}

func (m *MealInstance) PrintDetails() {
	fmt.Printf("Day: %s\n", m.Day)
	fmt.Printf("Date: %s\n", m.Date)
	fmt.Printf("Meal: %s\n", m.Meal)
	fmt.Println("Items:")
	for _, item := range m.Items {
		fmt.Printf("- %s\n", item)
	}
	fmt.Println()
}

func main() {
	err := parseExcel("Sample-Menu.xlsx")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

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
		items, ok := GetMenuItems(DayOfWeek(dayInput), Meal(mealInput))
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
		numItems := CountMenuItems(DayOfWeek(dayInput), Meal(mealInput))
		fmt.Printf("Number of items in %s %s: %d\n", dayInput, mealInput, numItems)
	case 3:
		var dayInput, mealInput, itemInput string

		fmt.Println("Enter the day:")
		fmt.Scanln(&dayInput)
		fmt.Println("Enter the meal:")
		fmt.Scanln(&mealInput)
		fmt.Println("Enter the item:")
		fmt.Scanln(&itemInput)

		isInMeal := IsItemInMeal(DayOfWeek(dayInput), Meal(mealInput), itemInput)
		fmt.Printf("Is '%s' in %s %s? %t\n", itemInput, dayInput, mealInput, isInMeal)
	case 4:
		err := SaveMenuAsJSON("menu.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	case 5:
		for day, meals := range Menu {
			for meal, items := range meals {
				date := Menu[day]["Date"][0]
				instance := MealInstance{
					Day:   day,
					Date:  date,
					Meal:  Meal(meal),
					Items: items,
				}
				instance.PrintDetails()
			}
		}

	default:
		fmt.Println("Invalid choice")
	}
}
