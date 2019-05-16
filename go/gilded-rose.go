package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []*Item{
	&Item{"+5 Dexterity Vest", 10, 20},
	&Item{"Aged Brie", 2, 0},
	&Item{"Elixir of the Mongoose", 5, 7},
	&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	&Item{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	GildedRose(items)
}

func GildedRose(items []*Item) {
	for _, item := range items {
		if item.name == "Aged Brie" || item.name == "Backstage passes to a TAFKAL80ETC concert" {
			// At the end of each day our system lowers both values for every item
			item.sellIn = item.sellIn - 1

			// The quality of an item is never more than 50
			if item.quality < 50 {
				item.quality = item.quality + 1
			}

			if item.name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.sellIn < 0 {
					// "Backstage passes" quality drops to 0 after the concert
					item.quality = item.quality - item.quality
				}
				if item.sellIn < 11 && item.quality > 0 && item.quality < 50 {
					item.quality = item.quality + 1
				}
				if item.sellIn < 6 && item.quality > 0 && item.quality < 50 {
					item.quality = item.quality + 1
				}
			}

			if item.sellIn < 0 && item.quality > 0 && item.quality < 50 {
				item.quality = item.quality + 1
			}
		} else {
			// The quality of an item is never negative
			if item.quality > 0 && item.name != "Sulfuras, Hand of Ragnaros" {
				// At the end of each day our system lowers both values for every item
				item.quality = item.quality - 1
			}

			if item.name != "Sulfuras, Hand of Ragnaros" {
				item.sellIn = item.sellIn - 1
			}

			if item.sellIn < 0 && item.quality > 0 && item.name != "Sulfuras, Hand of Ragnaros" {
				item.quality = item.quality - 1
			}
		}

	}

}
