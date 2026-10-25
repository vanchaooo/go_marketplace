package main
 
import (
	"fmt"
 
	"github.com/vanchaooo/go-marketplace/internal/repository"
	"github.com/vanchaooo/go-marketplace/internal/service"
)
 
func main() {
	users := repository.NewUserRepository()
	wallets := repository.NewWalletRepository()
	products := repository.NewProductRepository()
	carts := repository.NewCartRepository()
	orders := repository.NewOrderPerository()
	history := repository.NewHistoryRepository()
 
	userService := service.NewUserService(users, wallets, carts, orders)
	walletService := service.NewWalletService(users, wallets)
	productService := service.NewProductService(products)
	cartService := service.NewCartService(users, carts, products)
	orderService := service.NewOrderService(users, wallets, products, carts, orders, history)
 
	userID, ok := userService.AddUser("Иван")
	fmt.Println("AddUser:", userID, ok)
	secondID, ok := userService.AddUser("Мария")
	fmt.Println("AddUser (второй):", secondID, ok)
	fmt.Println("DeleteUser (второй, пустой):", userService.DeleteUser(secondID))
	thirdID, ok := userService.AddUser("Алексей")
	fmt.Println("AddUser (после удаления, номер не повторяется):", thirdID, ok)
	fmt.Println("AddUser (пустое имя):")
	_, ok = userService.AddUser("   ")
	fmt.Println(" ", ok)
 
	fmt.Println("AddProduct:", productService.AddNewProduct(1, "Книга 'Изучаем Go'", 2600, 10))
	fmt.Println("TopUp:", walletService.TopUpBalance(userID, 10000))
	fmt.Println("AddToCart:", cartService.AddToCart(userID, 1, 2))
 
	total, _ := cartService.CalculateCartTotal(userID)
	fmt.Println("Сумма корзины:", total)
 
	orderID, ok := orderService.MakeOrder(userID)
	fmt.Println("Checkout:", orderID, ok)
 
	balance, _ := walletService.GetBalance(userID)
	product, _ := productService.GetProduct(1)
	fmt.Println("Баланс после покупки:", balance, "| остаток товара:", product.Stock)
 
	fmt.Println("DeleteUser (есть оплаченный заказ):", userService.DeleteUser(userID))
 
	fmt.Println("CancelOrder:", orderService.CancelOrder(userID, orderID))
	fmt.Println("CancelOrder (повторно):", orderService.CancelOrder(userID, orderID))
 
	balance, _ = walletService.GetBalance(userID)
	product, _ = productService.GetProduct(1)
	fmt.Println("Баланс после отмены:", balance, "| остаток товара:", product.Stock)
 
	fmt.Println("Заказы пользователя:", orderService.GetUserOrders(userID))
	for _, line := range orderService.GetUserHistory(userID) {
		fmt.Println(" -", line)
	}
}