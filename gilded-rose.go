package main

type Updatable interface {
	Update()
}

type Item struct {
	name            string
	sellIn, quality int
}

func (item *Item) changeSellIn(change int) {
	item.sellIn += change
}

func (item *Item) changeQuality(change int) {
	item.quality += change
}

func (item *Item) maximizeQuality(max int) {
	if item.quality > max {
		item.quality = max
	}
}

func (item *Item) minimizeQuality(min int) {
	if item.quality < min {
		item.quality = min
	}
}

//************ Regular ************
type RegularItem struct {
	*Item
}

func NewRegularItem(item *Item) *RegularItem {
	return &RegularItem{
		Item: item,
	}
}

func (item *RegularItem) Update() {
	item.changeSellIn(-1)
	if item.sellIn < 0 {
		item.changeQuality(-2)
	} else {
		item.changeQuality(-1)
	}
	item.minimizeQuality(0)
}

//************ AgedBrie ************
type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) *AgedBrieItem {
	return &AgedBrieItem{
		Item: item,
	}
}

func (item *AgedBrieItem) Update() {
	item.changeSellIn(-1)
	if item.sellIn < 0 {
		item.changeQuality(+2)
	} else {
		item.changeQuality(+1)
	}
	item.maximizeQuality(50)
}

//************ Sulfuras ************
type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) *SulfurasItem {
	return &SulfurasItem{
		Item: item,
	}
}

func (item *SulfurasItem) Update() {
}

//************ BackstagePasses ************
type BackstagePassesItem struct {
	*Item
}

func NewBackstagePassesItem(item *Item) *BackstagePassesItem {
	return &BackstagePassesItem{
		Item: item,
	}
}

func (item *BackstagePassesItem) Update() {
	item.changeSellIn(-1)
	sellIn := item.sellIn
	switch {
	case sellIn >= 11:
		item.changeQuality(+1)
	case sellIn <= 10 && sellIn > 5:
		item.changeQuality(+2)
	case sellIn <= 5 && sellIn >= 0:
		item.changeQuality(+3)
	case sellIn < 0:
		item.maximizeQuality(0)
	}
	item.maximizeQuality(50)
}

//************ Conjured ************
type ConjuredItem struct {
	*Item
}

func NewConjuredItem(item *Item) *ConjuredItem {
	return &ConjuredItem{
		Item: item,
	}
}

func (item *ConjuredItem) Update() {
	item.changeSellIn(-1)
	if item.sellIn < 0 {
		item.changeQuality(-4)
	} else {
		item.changeQuality(-2)
	}
	item.minimizeQuality(0)
}

type UpdatableItemFunc func(item *Item) Updatable

// UpdatableItemCreator creates an Updatable Item for a specific Item name in the map or a Regular Item if not found
func UpdatableItemCreator(itemCreatorMap map[string]UpdatableItemFunc, item *Item) Updatable {
	create, exists := itemCreatorMap[item.name]
	if exists {
		return create(item)
	} else {
		return NewRegularItem(item)
	}
}

// UpdateQuality updates the quality for all received items
func UpdateQuality(items []*Item) {
	creationMap := map[string]UpdatableItemFunc{
		"Aged Brie": func(item *Item) Updatable {
			return NewAgedBrieItem(item)
		},
		"Sulfuras, Hand of Ragnaros": func(item *Item) Updatable {
			return NewSulfurasItem(item)
		},
		"Backstage passes to a TAFKAL80ETC concert": func(item *Item) Updatable {
			return NewBackstagePassesItem(item)
		},
		"Conjured Mana Cake": func(item *Item) Updatable {
			return NewConjuredItem(item)
		},
	}

	for i := 0; i < len(items); i++ {
		updatableItem := UpdatableItemCreator(creationMap, items[i])
		updatableItem.Update()
	}
}
