package domain

import (
	"githun.com/yonatannn/ciyana-e-commerce-website/domain/models"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(userID string) error
	GetUserByID(userID string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}

type ItemRepository interface {
	CreateItem(item domain.Item) error
	UpdateItem(item domain.Item) error
	DeleteItem(itemID string) error
	GetItemByID(itemID string) (domain.Item, error)
	GetItemsByCategory(category string) ([]domain.Item, error)
}

type OrderRepository interface {
	CreateOrder(order domain.Order) error
	UpdateOrder(order domain.Order) error
	DeleteOrder(orderID string) error
	GetOrderByID(orderID string) (domain.Order, error)
	GetOrdersByUserID(userID string) ([]domain.Order, error)
}

type CategoryRepository interface {
	CreateCategory(category domain.Category) error
	UpdateCategory(category domain.Category) error
	DeleteCategory(categoryID string) error
	GetCategoryByID(categoryID string) (domain.Category, error)
	GetAllCategories() ([]domain.Category, error)
}

type CartRepository interface {
	CreateCart(cart domain.Cart) error
	UpdateCart(cart domain.Cart) error
	DeleteCart(cartID string) error
	GetCartByID(cartID string) (domain.Cart, error)
	GetCartByUserID(userID string) (domain.Cart, error)
}

type AddressRepository interface {
	CreateAddress(address domain.Address) error
	UpdateAddress(address domain.Address) error
	DeleteAddress(addressID string) error
	GetAddressByID(addressID string) (domain.Address, error)
	GetAddressesByUserID(userID string) ([]domain.Address, error)
}
