package scheduler

type Item struct {
	Url string
}

func NewItem(url *string) *Item {
	return &Item{Url: *url}
}
