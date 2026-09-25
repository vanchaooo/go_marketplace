package repository

import (
	"fmt"
	"errors"
	"math/rand/v2"
)

type Product struct {
	ID int
	Name  string
	Price int
	Stock int
}

var Catalog = map[int]*Product{
	// --- 📱 Электроника и Гаджеты (1-8) ---
	1: {ID: 1, Name: "Смартфон Флагман 256GB", Price: 89990, Stock: rand.IntN(1000)},
	2: {ID: 2, Name: "Беспроводные наушники ANC", Price: 14990, Stock: rand.IntN(1000)},
	3: {ID: 3, Name: "Ноутбук для работы 15.6", Price: 64990, Stock: rand.IntN(1000)},
	4: {ID: 4, Name: "Умные часы Amoled", Price: 17490, Stock: rand.IntN(1000)},
	5: {ID: 5, Name: "Портативная колонка 20W", Price: 5990, Stock: rand.IntN(1000)},
	6: {ID: 6, Name: "Планшет 10.5' Wi-Fi", Price: 29990, Stock: rand.IntN(1000)},
	7: {ID: 7, Name: "Внешний аккумулятор 20k мАч", Price: 2490, Stock: rand.IntN(1000)},
	8: {ID: 8, Name: "Игровая мышь беспроводная", Price: 4990, Stock: rand.IntN(1000)},

	// --- 🏡 Бытовая техника (9-15) ---
	9:  {ID: 9, Name: "Робот-пылесос с базой", Price: 27990, Stock: rand.IntN(1000)},
	10: {ID: 10, Name: "Кофемашина автоматическая", Price: 34990, Stock: rand.IntN(1000)},
	11: {ID: 11, Name: "Электрический чайник 1.7L", Price: 3490, Stock: rand.IntN(1000)},
	12: {ID: 12, Name: "Микроволновая печь 20L", Price: 7990, Stock: rand.IntN(1000)},
	13: {ID: 13, Name: "Вертикальный отпариватель", Price: 5490, Stock: rand.IntN(1000)},
	14: {ID: 14, Name: "Тостер на два слота", Price: 2190, Stock: rand.IntN(1000)},
	15: {ID: 15, Name: "Увлажнитель воздуха ультразвуковой", Price: 3890, Stock: rand.IntN(1000)},

	// --- 👕 Одежда и Обувь (16-22) ---
	16: {ID: 16, Name: "Футболка хлопковая Basic", Price: 1200, Stock: rand.IntN(1000)},
	17: {ID: 17, Name: "Худи оверсайз с начесом", Price: 4200, Stock: rand.IntN(1000)},
	18: {ID: 18, Name: "Джинсы прямые синие", Price: 3500, Stock: rand.IntN(1000)},
	19: {ID: 19, Name: "Кроссовки городские", Price: 7990, Stock: rand.IntN(1000)},
	20: {ID: 20, Name: "Куртка-бомбер демисезонная", Price: 6800, Stock: rand.IntN(1000)},
	21: {ID: 21, Name: "Носки спортивные (3 пары)", Price: 450, Stock: rand.IntN(1000)},
	22: {ID: 22, Name: "Кепка хлопковая черная", Price: 990, Stock: rand.IntN(1000)},

	// --- 🎒 Аксессуары и Сумки (23-28) ---
	23: {ID: 23, Name: "Рюкзак городской с отделом для ноутбука", Price: 4800, Stock: rand.IntN(1000)},
	24: {ID: 24, Name: "Кожаный кошелек", Price: 2900, Stock: rand.IntN(1000)},
	25: {ID: 25, Name: "Ремень кожаный черный", Price: 1800, Stock: rand.IntN(1000)},
	26: {ID: 26, Name: "Сумка на пояс (бананка)", Price: 1500, Stock: rand.IntN(1000)},
	27: {ID: 27, Name: "Зонт-автомат прочный", Price: 1990, Stock: rand.IntN(1000)},
	28: {ID: 28, Name: "Чехол для ноутбука 14'", Price: 1350, Stock: rand.IntN(1000)},

	// --- ⚽ Спорт и Активный отдых (29-35) ---
	29: {ID: 29, Name: "Коврик для йоги нескользящий", Price: 1600, Stock: rand.IntN(1000)},
	30: {ID: 30, Name: "Спортивная бутылка 0.7L", Price: 790, Stock: rand.IntN(1000)},
	31: {ID: 31, Name: "Набор фитнес-резинок (5 шт.)", Price: 850, Stock: rand.IntN(1000)},
	32: {ID: 32, Name: "Гантели неопреновые 2x2кг", Price: 1900, Stock: rand.IntN(1000)},
	33: {ID: 33, Name: "Эспандер кистевой", Price: 400, Stock: rand.IntN(1000)},
	34: {ID: 34, Name: "Спортивная сумка 30L", Price: 2600, Stock: rand.IntN(1000)},
	35: {ID: 35, Name: "Скакалка с подшипниками", Price: 690, Stock: rand.IntN(1000)},

	// --- 🛋️ Дом, Интерьер и Уют (36-42) ---
	36: {ID: 36, Name: "Настольная лампа LED", Price: 2490, Stock: rand.IntN(1000)},
	37: {ID: 37, Name: "Термокружка вакуумная 0.4L", Price: 1250, Stock: rand.IntN(1000)},
	38: {ID: 38, Name: "Ароматическая свеча в стекле", Price: 750, Stock: rand.IntN(1000)},
	39: {ID: 39, Name: "Плед флисовый 150x200", Price: 1850, Stock: rand.IntN(1000)},
	40: {ID: 40, Name: "Подушка ортопедическая", Price: 2990, Stock: rand.IntN(1000)},
	41: {ID: 41, Name: "Набор кухонных ножей (5 шт.)", Price: 4300, Stock: rand.IntN(1000)},
	42: {ID: 42, Name: "Полотенце махровое большое", Price: 950, Stock: rand.IntN(1000)},

	// --- 📚 Книги и Канцелярия (43-47) ---
	43: {ID: 43, Name: "Книга 'Изучаем Go' 2-е изд.", Price: 2600, Stock: rand.IntN(1000)},
	44: {ID: 44, Name: "Книга 'Чистый код'", Price: 1950, Stock: rand.IntN(1000)},
	45: {ID: 45, Name: "Блокнот в кожаной обложке A5", Price: 850, Stock: rand.IntN(1000)},
	46: {ID: 46, Name: "Ручка перьевая подарочная", Price: 1500, Stock: rand.IntN(1000)},
	47: {ID: 47, Name: "Набор маркеров для скетчинга", Price: 1100, Stock: rand.IntN(1000)},

	// --- 🧼 Уход и Красота (48-50) ---
	48: {ID: 48, Name: "Электрическая зубная щетка", Price: 3990, Stock: rand.IntN(1000)},
	49: {ID: 49, Name: "Фен для волос с ионизацией", Price: 4990, Stock: rand.IntN(1000)},
	50: {ID: 50, Name: "Триммер для бороды и усов", Price: 2800, Stock: rand.IntN(1000)},
}

func AddProduct(product *Product) bool {
	product.ID = len(Catalog)+1
	if product.Name == "" {
		fmt.Println(errors.New("Имя товара не может быть пустым!"))
		return false
	}
	if product.Price == 0 {
		fmt.Println(errors.New("Цена не должна быть равна 0."))
		return false
	}
	if product.Stock == 0 {
		fmt.Println(errors.New("При добавлении товара, он должен быть в наличии."))
		return false
	}

	Catalog[product.ID] = product
	return true
}

func GetCatalog() {
	for _, v := range Catalog {
		fmt.Printf("ID: %d, Товар: %s, Цена: %d, В наличии: %d шт.\n", v.ID, v.Name, v.Price, v.Stock)
	}
}

func GetProduct(id int) *Product {
	if product, ok := Catalog[id]; ok {
		fmt.Printf("ID: %d || Товар: %s, Цена: %d, Кол-во на складе: %d шт.\n", product.ID, product.Name, product.Price, product.Stock)
		return product
	} else {
		fmt.Println(errors.New("С таким ID, товара не найдено."))
		return nil
	}
}

func DeleteProduct(id int) {
	if _, ok := Catalog[id]; ok {
		delete(Catalog, id)
	} else {
		fmt.Println(errors.New("С таким ID, товара не найдено."))
	}
}