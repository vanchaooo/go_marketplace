package catalog

import (
	"fmt"
	"errors"
)

var (
	Catalog = map[string]map[string]int {
		"Телефоны": {
			"iPhone 15": 49990,
			"Samsung S24": 75000,
			"Xiaomi 14":  40000,
		}, 
		"Ноутбуки": {
			"MacBook Pro 16": 150000,
			"Asus ROG Strix": 100000,
			"Lenovo ThinkPad X1": 70000,
			"Huawei MateBook X": 100000,
		},
		"Телевизоры": {
			"LG OLED C3 55": 130000,	
			"Samsung QLED QN90": 95000,
			"Sony Bravia XR": 69990,
		},
		"Аудио": {
			"AirPods Pro 2": 15000,
			"Sony WH-1000XM5": 14000,
			"JBL Charge 5": 8990,
		},
		"Умный дом": {
			"Яндекс Станция Макс": 20000,
			"Лампа Xiaomi Smart Bulb": 5590,
			"Датчик Aqara Hub": 35100,
		},
	}
)

func GetCatalog() {
	fmt.Println("---------------------------------------------")
	for k, v := range Catalog {
		fmt.Println("Категория:", k)
		fmt.Println()
		for sk, sv := range v {
			fmt.Printf("Товар: %s | Цена: %d\n", sk, sv)
		}
		fmt.Println("---------------------------------------------")
	}
}

func AddNewSection(name string, position map[string]int) bool {
	if name == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	Catalog[name] = position
	return true
}

func AddNewPosition(section string, position string, price int) bool {
	if section == "" || position == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	Catalog[section][position] = price
	return true
}