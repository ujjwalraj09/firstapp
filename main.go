package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DayOfWeek string
type Meal string

// bringing the excel data into map of map form
var Menu = map[DayOfWeek]map[Meal][]string{
	"Monday": {
		"Breakfast": {"CHOICE OF EGG", "CORNFLAKES", "BREAD +JAM", "POHA", "GRAPES", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"DISCO PAPAD", "KADHI PAKODA", "MADRASI ALOO", "VEG KHICHDI", "CHAPATI", "PLAIN RICE"},
		"Dinner":    {"MIRCH KE TAPORE/ LAHSUN CHUTNEY", "ALOO RASSA", "DAL FRY", "BATI/ CHAPATI", "PLAIN RICE", "RASAM"},
	},
	"Tuesday": {
		"Breakfast": {"NO EGG", "CORNFLAKES", "BREAD +JAM", "ALOO PARATHA", "CURD", "PICKEL", "KINU", "TEA+ COFFEE", "Milk", "TEA+ COFFEE"},
		"Lunch":     {"VEG KADHAI", "DAL MAHARANI", "PLAIN RICE", "SAMBAR", "SWEET LASSI"},
		"Dinner":    {"SWEET LASSI", "HOT CHOCO MILK", "CHOLE MASALA"},
	},
	"Wednesday": {
		"Breakfast": {"CHOICE OF EGG", "CORNFLAKES", "BREAD +JAM", "BREAD PAKODA", "TOMATO KETCHUP", "SWEET DALIYA", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"CHAPATI", "TAWA VEG (NO KARELA)", "PLAIN RICE"},
		"Dinner":    {"TAWA VEG (NO KARELA)", "DAL PALAK"},
	},
	"Thursday": {
		"Breakfast": {"CHOICE OF EGG", "CORNFLAKES", "BREAD +JAM", "MUTTER KULCHA", "KULCHA", "CUT ONION LEMON", "PAPAYA", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"GATTA CURRY", "CHANNA DAL TADKA"},
		"Dinner":    {"SAMBAR", "BUTTER DAL TADKA", "CHAPATI", "GULAB JAMUN"},
	},
	"Friday": {
		"Breakfast": {"CHOICE OF EGG", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"RAJMA MASALA", "SHIKANJI"},
		"Dinner":    {"VEG MANCHURIAN", "EGG FRIED RICE / VEG FRIED RICE", "CHAPATI", "BALIUSHAHI"},
	},
	"Saturday": {
		"Breakfast": {"CHOICE OF EGG", "CORNFLAKES", "BREAD +JAM", "SUJI UPMA", "ADRAK CHUTNEY", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"MIX VEG PARATHA", "RAGDA MUTTER", "FRENCH FRIES"},
		"Dinner":    {"GREEN CHUTNEY", "PASTA SALAD"},
	},
	"Sunday": {
		"Breakfast": {"CHOICE OF EGG", "CORNFLAKES", "BREAD +JAM", "MASALA DOSA", "SAMBHAR", "GRAPES", "TEA+ COFFEE", "Milk"},
		"Lunch":     {"PANEER LABABDAR", "DHABA CHICKEN", "NAAN / CHAPATI"},
		"Dinner":    {"METHI MALAI MUTTER", "DAL LASOONI", "PLAIN RICE", "HOT & SOUR SOUP"},
	},
}

// 1st part function
func GetMenu(day DayOfWeek, meal Meal) ([]string, bool) {

	menuForDay, ok := Menu[day]
	if !ok {
		return nil, false
	}

	menuForMeal, ok := menuForDay[meal]
	if !ok {
		return nil, false
	}

	return menuForMeal, true
}

// 2nd part function
func CountMenuItems(day DayOfWeek, meal Meal) int {
	menuItems, found := Menu[day][meal]
	if !found {
		return 0
	}
	return len(menuItems)
}

// 3rd part function
func IsItemInMeal(day DayOfWeek, meal Meal, item string) bool {

	menuItems, found := Menu[day][meal]
	if !found {
		return false
	}

	for _, menuItem := range menuItems {
		if menuItem == item {
			return true
		}
	}
	return false
}

// 4th converting to json
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

// creating structure
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
func getDateForMeal(day DayOfWeek) string {

	dateMap := map[DayOfWeek]string{
		"Monday":    "05-feb-24",
		"Tuesday":   "06-feb-24",
		"Wednesday": "07-feb-24",
		"Thursday":  "08-feb-24",
		"Friday":    "09-feb-24",
		"Saturday":  "10-feb-24",
		"Sunday":    "11s-feb-24",
	}

	date, found := dateMap[day]
	if !found {

		return ""
	}
	return date
}

func main() {
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. GetMenu")
	fmt.Println("2. CountMenuItems")
	fmt.Println("3. IsItemInMeal")
	fmt.Println("4. SaveMenuAsJSON")
	fmt.Println("5. PrintDetails")
	fmt.Println("The nputs are case sensitive")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var dayInput, mealInput string
		fmt.Println("Enter the day:")
		fmt.Scanln(&dayInput)
		fmt.Println("Enter the meal:")
		fmt.Scanln(&mealInput)
		if menu, ok := GetMenu(DayOfWeek(dayInput), Meal(mealInput)); ok {
			fmt.Printf("%s %s menu: %v\n", dayInput, mealInput, menu)
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
				date := getDateForMeal(day) // Get the date for the meal instance
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
