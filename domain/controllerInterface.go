package domain


import (
	"githun.com/yonatannn/ciyana-e-commerce-website/domain/models"
)


type ApiController interface {
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id string) (domain.Category, error)
	CreateCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(id string, category domain.Category) (domain.Category, error)
	DeleteCategory(id string) error
	GetAllItems() ([]domain.Item, error)
	GetItemByID(id string) (domain.Item, error)
	CreateItem(item domain.Item) (domain.Item, error)
	UpdateItem(id string, item domain.Item) (domain.Item, error)
	DeleteItem(id string) error
	GetAllCarts() ([]domain.Cart, error)
	GetCartByID(id string) (domain.Cart, error)
	CreateCart(cart domain.Cart) (domain.Cart, error)
	UpdateCart(id string, cart domain.Cart) (domain.Cart, error)
	DeleteCart(id string) error
	GetAllOrders() ([]domain.Order, error)
	GetOrderByID(id string) (domain.Order, error)
	CreateOrder(order domain.Order) (domain.Order, error)
	UpdateOrder(id string, order domain.Order) (domain.Order, error)
	DeleteOrder(id string) error
}


type UserController interface {
	GetAllUsers() ([]domain.User, error)
	GetUserByID(id string) (domain.User, error)
	RegisterUser(user domain.User) (domain.User, error)
	LoginUser(user domain.User) (domain.User, error)
	UpdateUser(id string, user domain.User) (domain.User, error)
	DeleteUser(id string) error
}


type ItemController interface {
	GetAllItems() ([]domain.Item, error)
	GetItemByID(id string) (domain.Item, error)
	CreateItem(item domain.Item) (domain.Item, error)
	UpdateItem(id string, item domain.Item) (domain.Item, error)
	DeleteItem(id string) error
}


type CartController interface {
	GetAllCarts() ([]domain.Cart, error)
	GetCartByID(id string) (domain.Cart, error)
	CreateCart(cart domain.Cart) (domain.Cart, error)
	UpdateCart(id string, cart domain.Cart) (domain.Cart, error)
	DeleteCart(id string) error
}


type OrderController interface {
	GetAllOrders() ([]domain.Order, error)
	GetOrderByID(id string) (domain.Order, error)
	CreateOrder(order domain.Order) (domain.Order, error)
	UpdateOrder(id string, order domain.Order) (domain.Order, error)
	DeleteOrder(id string) error
}

type CategoryController interface {
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id string) (domain.Category, error)
	CreateCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(id string, category domain.Category) (domain.Category, error)
	DeleteCategory(id string) error
}