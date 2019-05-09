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
		// if item is not Aged Brie and Backstage passes to a TAFKAL80ETC concert, quality
		// is greater than 0, and item is not Sulfuras, Hand of Ragnaros, then quality
		// decreases 1
		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			// The quality of an item is never negative
			if items[i].quality > 0 {
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					// At the end of each day our system lowers both values for every item
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			// The quality of an item is never more than 50
			if items[i].quality < 50 {
				// "Aged Brie" actually increases in quality the older it gets
				items[i].quality = items[i].quality + 1
				// "Backstage passes", like aged brie, increases in quality as its sell-in
				// value approaches; quality increases by 2 when there are 10 days or less
				// and by 3 when there are 5 days or less but quality drops to 0 after the
				// concert
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						// The quality of an item is never more than 50
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						// The quality of an item is never more than 50
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			// At the end of each day our system lowers both [sell in and quality] values for every item
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					// The quality of an item is never negative
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							// At the end of each day our system lowers both values for every item
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					// "Backstage passes" quality drops to 0 after the concert
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				// The quality of an item is never more than 50
				if items[i].quality < 50 {
					// "Aged Brie" actually increases in quality the older it gets
					items[i].quality = items[i].quality + 1
				}
			}
		}
	}

}
