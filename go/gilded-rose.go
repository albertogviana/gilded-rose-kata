package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
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
		if items[i].name == "Aged Brie" || items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
			// At the end of each day our system lowers both values for every item
			items[i].sellIn = items[i].sellIn - 1

			// The quality of an item is never more than 50
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
			}

			if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
				if items[i].sellIn < 0 {
					// "Backstage passes" quality drops to 0 after the concert
					items[i].quality = items[i].quality - items[i].quality
				}
				if items[i].sellIn < 11 && items[i].quality > 0 && items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
				if items[i].sellIn < 6 && items[i].quality > 0 && items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}

			if items[i].sellIn < 0 && items[i].quality > 0 && items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
			}
		} else {
			// The quality of an item is never negative
			if items[i].quality > 0 && items[i].name != "Sulfuras, Hand of Ragnaros" {
				// At the end of each day our system lowers both values for every item
				items[i].quality = items[i].quality - 1
			}

			if items[i].name != "Sulfuras, Hand of Ragnaros" {
				items[i].sellIn = items[i].sellIn - 1
			}

			if items[i].sellIn < 0 && items[i].quality > 0 && items[i].name != "Sulfuras, Hand of Ragnaros" {
				items[i].quality = items[i].quality - 1
			}
		}

	}

}
