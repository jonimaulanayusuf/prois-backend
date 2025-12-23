package requests

type CreateItemRequest struct {
	Name  string  `json:"name" validate:"required,max=150"`
	Stock int     `json:"stock" validate:"gte=0"`
	Price float64 `json:"price" validate:"gte=0"`
}

type UpdateItemRequest struct {
	Name  *string  `json:"name" validate:"omitempty,max=150"`
	Stock *int     `json:"stock" validate:"omitempty,gte=0"`
	Price *float64 `json:"price" validate:"omitempty,gte=0"`
}
