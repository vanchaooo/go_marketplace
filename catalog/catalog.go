package catalog

import (
	"fmt"
	// "errors"
	"math/rand/v2"
)

type Product struct{
	Name string
	Price int
	Quantity int
}

var Catalog = map[string]map[string]*Product {
	"Телефоны": {
		"iPhone 15": &Product{
			Name: "iPhone 15",
			Price: 95000,
			Quantity: rand.IntN(1000),
		},
		"Samsung S24": &Product{
			Name: "Samsung S24",
			Price: 85000,
			Quantity: rand.IntN(1000),
		},
	},
}

// var (
// 	Catalog = map[string]*Product{} {
// 		"Телефоны": {
// 			Product.Name: "Iphone"
// 			},
// 			"Samsung S24": {
// 				85000: rand.IntN(1000),
// 			},
// 			"Xiaomi 14": {
// 				65000: rand.IntN(1000),
// 			},
// 		},
		// "Ноутбуки": {
		// 	"MacBook Pro 16": {
		// 		250000: rand.IntN(1000),
		// 	},
		// 	"Asus ROG Strix": {
		// 		180000: rand.IntN(1000),
		// 	},
		// 	"Lenovo ThinkPad X1": {
		// 		140000: rand.IntN(1000),
		// 	},
		// 	"Huawei MateBook X": {
		// 		120000: rand.IntN(1000),
		// 	},
		// },
		// "Телевизоры": {
		// 	"LG OLED C3 55": {
		// 		160000: rand.IntN(1000),
		// 	},
		// 	"Samsung QLED QN90": {
		// 		110000: rand.IntN(1000),
		// 	},
		// 	"Sony Bravia XR": {
		// 		90000: rand.IntN(1000),
		// 	},
		// },
		// "Аудио": {
		// 	"AirPods Pro 2": {
		// 		23000: rand.IntN(1000),
		// 	},
		// 	"Sony WH-1000XM5": {
		// 		32000: rand.IntN(1000),
		// 	},
		// 	"JBL Charge 5": {
		// 		12000: rand.IntN(1000),
		// 	},
		// },
		// "Умный дом": {
		// 	"Яндекс Станция Макс": {
		// 		28000: rand.IntN(1000),
		// 	},
		// 	"Лампа Xiaomi Smart Bulb": {
		// 		1500: rand.IntN(1000),
		// 	},
		// 	"Датчик Aqara Hub": {
		// 		4000: rand.IntN(1000),
		// 	},
		// },
// 	}
// )

func GetCatalog() {
	fmt.Println("---------------------------------------------")
	for section, positions := range Catalog {
		fmt.Println("Категория:", section)
		fmt.Println()
		for product, info := range positions {
			fmt.Println("Товар:", product)
			fmt.Println(info)
		}
		fmt.Println("---------------------------------------------")
	}
}

// func AddNewSection(name string) bool {
// 	if name == "" {
// 		fmt.Println(errors.New("Name can't be empty."))
// 		return false
// 	}

// 	Catalog[name] = make(map[string]map[int]int)
// 	return true
// }

// func AddNewPosition(section string, position string, price int) bool {
// 	if section == "" || position == "" {
// 		fmt.Println(errors.New("Name can't be empty."))
// 		return false
// 	}

// 	if Catalog[section] == nil {
// 		Catalog[section] = make(map[string]map[int]int)
// 	}

// 	if Catalog[section][position] == nil {
// 		Catalog[section][position] = make(map[int]int)
// 	}

// 	Catalog[section][position][price] = rand.IntN(1000)
// 	return true
// }