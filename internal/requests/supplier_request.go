package requests

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,max=150"`
	Email   string `json:"email" validate:"required,max=150,email"`
	Address string `json:"address" validate:"required,max=255"`
}

type UpdateSupplierRequest struct {
	Name    *string `json:"name" validate:"required,max=150"`
	Email   *string `json:"email" validate:"required,max=150,email"`
	Address *string `json:"address" validate:"required,max=255"`
}
