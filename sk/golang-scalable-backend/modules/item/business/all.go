package business

type all struct {
	listItemBusiness
	getItemBusiness
	createItemBusiness
	updateItemBusiness
	deleteItemBusiness
}

func NewAllBusiness() *all {
	return &all{}
}
