package product_mng_repo

type AttributesFilter struct {
	AttributesIds []string
	Attributes    map[string]interface{}
}

type ProductAttributes struct {
	ProductId  string
	Attributes map[string]interface{}
}
