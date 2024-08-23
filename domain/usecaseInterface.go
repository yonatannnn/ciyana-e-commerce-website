package domain


import (
	"githun.com/yonatannn/ciyana-e-commerce-website/domain/models"
)


type UserUsecase interface {
	GetUserByID(id string) (domain.User, error)
	RegisterUser(user domain.User) (domain.User, error)
	LoginUser(user domain.User) (domain.User, error)
	UpdateUser(id string, user domain.User) (domain.User, error)
	DeleteUser(id string) error
}

type ItemUsecase interface {
	GetItemByID(id string) (domain.Item, error)
	CreateItem(item domain.Item) (domain.Item, error)
	UpdateItem(id string, item domain.Item) (domain.Item, error)
	DeleteItem(id string) error
}

type OrderUsecase interface {
	GetOrderByID(id string) (domain.Order, error)
	CreateOrder(order domain.Order) (domain.Order, error)
	UpdateOrder(id string, order domain.Order) (domain.Order, error)
	DeleteOrder(id string) error
}

type CartUsecase interface {
	GetCartByID(id string) (domain.Cart, error)
	CreateCart(cart domain.Cart) (domain.Cart, error)
	UpdateCart(id string, cart domain.Cart) (domain.Cart, error)
	DeleteCart(id string) error
}

type CategoryUsecase interface {
	GetCategoryByID(id string) (domain.Category, error)
	CreateCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(id string, category domain.Category) (domain.Category, error)
	DeleteCategory(id string) error
}

type AddressUsecase interface {
	GetAddressByID(id string) (domain.Address, error)
	CreateAddress(address domain.Address) (domain.Address, error)
	UpdateAddress(id string, address domain.Address) (domain.Address, error)
	DeleteAddress(id string) error
}
