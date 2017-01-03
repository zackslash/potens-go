package breadcrumb

import "encoding/json"

type BreadcrumbItem struct {
	Url   string
	Title string
}

type Breadcrumb struct {
	Items []BreadcrumbItem
}

func (b *Breadcrumb) AddItem(item BreadcrumbItem) {
	b.Items = append(b.Items, item)
}

func (b *Breadcrumb) Json() string {
	bytes, _ := json.Marshal(b.Items)
	return string(bytes)
}

func (b *Breadcrumb) FromJson(rawJson string) error {
	items := make([]BreadcrumbItem, 0)
	err := json.Unmarshal([]byte(rawJson), &items)
	if err == nil {
		b.Items = items
	}
	return err
}
