package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

type Updater interface {
	Update()
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	GildedRose(items)
}

func GildedRose(items []Item) {
	for i := 0; i < len(items); i++ {
		if items[i].name == "Aged Brie" {
			agedBrieItem := NewAgedBrieItem(&items[i])
			agedBrieItem.Update()
		} else if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			backstagePassesItem := NewBackstagePassesItem(&items[i])
			backstagePassesItem.Update()
		} else {
			// The quality of an item is never negative
			if items[i].quality > 0 && items[i].name != "Sulfuras, Hand of Ragnaros" {
				// At the end of each day our system lowers both values for every item
				items[i].updateQuality(-1)
			}

			if items[i].name != "Sulfuras, Hand of Ragnaros" {
				items[i].updateSellIn(-1)
			}

			if items[i].sellIn < 0 && items[i].quality > 0 && items[i].name != "Sulfuras, Hand of Ragnaros" {
				items[i].updateQuality(-1)
			}
		}
	}
}

func (i *Item) updateSellIn(value int) {
	i.sellIn += value
}

func (i *Item) updateQuality(value int) {
	i.quality += value
}

type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) *AgedBrieItem {
	return &AgedBrieItem{
		Item: item,
	}
}

func (item *AgedBrieItem) Update() {
	item.updateSellIn(-1)

	// The quality of an item is never more than 50
	if item.quality < 50 {
		item.updateQuality(1)
	}

	if item.sellIn < 0 && item.quality > 0 && item.quality < 50 {
		item.updateQuality(1)
	}
}

type BackstagePassesItem struct {
	*Item
}

func NewBackstagePassesItem(item *Item) *BackstagePassesItem {
	return &BackstagePassesItem{
		Item: item,
	}
}

func (item *BackstagePassesItem) Update() {
	item.updateSellIn(-1)

	// The quality of an item is never more than 50
	if item.quality < 50 {
		item.updateQuality(1)
	}

	if item.sellIn < 0 {
		// "Backstage passes" quality drops to 0 after the concert
		item.updateQuality(-item.quality)
	}
	if item.sellIn < 11 && item.quality > 0 && item.quality < 50 {
		item.updateQuality(1)
	}
	if item.sellIn < 6 && item.quality > 0 && item.quality < 50 {
		item.updateQuality(1)
	}
}
