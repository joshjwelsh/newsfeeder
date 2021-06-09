package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

// Holds each user post with their respective title
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// the users news feed
type Feed struct {
	Items []Item
}

// Create a new news feed
func New() *Feed {
	return &Feed{
		Items: []Item{},
	}
}

// Add a post to the news feed
func (f *Feed) Add(item Item) {
	f.Items = append(f.Items, item)
}

// Return a list of all posts for the news feed
func (f *Feed) GetAll() []Item {
	return f.Items
}
