package catalog

import (
	"fmt"
	"errors"
)

var (
	catalog = map[string]map[string]int {
		"Смартфоны": {
			"iPhone 15": 49900,
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

	bucket = make(map[string]int)
)

func GetCatalog() {
	for _, v := range catalog {
		for sk, sv := range v {
			fmt.Printf("Товар: %s | Цена: %d\n", sk, sv)
		}
	}
}

func AddToBucket(category string, device string) bool {
		if category == "" || device == "" {
			fmt.Println(errors.New("Name can't be empty."))
			return false
		}

		bucket[device] = catalog[category][device]
		return true
}

func DeleteFromBucket(position string) bool {
	if position == "" {
		fmt.Println(errors.New("Name can't be empty."))
		return false
	}

	delete(bucket, position)
	return true
}