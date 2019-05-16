package main

import (
	"testing"
)

// At the end of each day our system lowers both values for every item
func TestNormalItemDecreaseSellInByOne(t *testing.T) {
	testUpdateSellIn(t, "normal item", 10, 20, 9)
}

// At the end of each day our system lowers both values for every item
func TestNormalItemDecreaseQualityByOne(t *testing.T) {
	testUpdateQuality(t, "normal item", 10, 20, 19)
}

// Once the sell by date has passed, quality degrades twice as fast
func TestNormalItemQualityDegradesTwiceAsFastOnceTheSellDatePassed(t *testing.T) {
	testUpdateQuality(t, "normal item", 0, 20, 18)
}

// The quality of an item is never negative
func TestNormalItemQualityIsNeverNegative(t *testing.T) {
	testUpdateQuality(t, "normal item", 10, 0, 0)
}

// "Aged Brie" actually increases in quality the older it gets
func TestAgedBrieIncreasesQualityTheOlderItGetsByOne(t *testing.T) {
	testUpdateQuality(t, "Aged Brie", 2, 0, 1)
}

// "Aged Brie" actually increases in quality the older it gets
func TestAgedBrieIncreasesQualityByTwoAfterSellDatePassed(t *testing.T) {
	testUpdateQuality(t, "Aged Brie", 0, 30, 32)
}

// The quality of an item is never more than 50
func TestAgedBrieQualityIsNeverMoreThanFifty(t *testing.T) {
	testUpdateQuality(t, "Aged Brie", 2, 50, 50)
}

// "Sulfuras", being a legendary item, never has to be sold or decreases in quality
func TestSufurasNeverDecreasesQuality(t *testing.T) {
	testUpdateQuality(t, "Sulfuras, Hand of Ragnaros", 0, 80, 80)
}

// "Sulfuras", being a legendary item, never has to be sold or decreases in quality
func TestSufurasNeverDecreasesSellIn(t *testing.T) {
	testUpdateSellIn(t, "Sulfuras, Hand of Ragnaros", 0, 80, 0)
}

func TestBackstageQualityIncreasesByOne(t *testing.T) {
	testUpdateQuality(t, "Backstage passes to a TAFKAL80ETC concert", 15, 20, 21)
}

// "Backstage passes", like aged brie, increases in quality as its sell-in
// value approaches; quality increases by 2 when there are 10 days or less
func TestBackstageQualityIncreasesByTwoWhen10DaysOrLess(t *testing.T) {
	testUpdateQuality(t, "Backstage passes to a TAFKAL80ETC concert", 10, 20, 22)
}

// "Backstage passes", like aged brie, increases in quality as its sell-in
// value approaches; quality increases and by 3 when there are 5 days or less
func TestBackstageQualityIncreasesByThreeWhen5DaysOrLess(t *testing.T) {
	testUpdateQuality(t, "Backstage passes to a TAFKAL80ETC concert", 5, 20, 23)
}

// "Backstage passes" quality drops to 0 after the concert
func TestBackstageQualityDropsAfterConcert(t *testing.T) {
	testUpdateQuality(t, "Backstage passes to a TAFKAL80ETC concert", 0, 20, 0)
}

// At the end of each day our system lowers both values for every item
func TestConjuredDecreasesQualityByOne(t *testing.T) {
	testUpdateQuality(t, "Conjured Mana Cake", 3, 6, 5)
}

// At the end of each day our system lowers both values for every item
func TestConjuredDecreasesSellInByOne(t *testing.T) {
	testUpdateSellIn(t, "Conjured Mana Cake", 3, 6, 2)
}

// "Conjured" items degrade in quality twice as fast as normal items
func TestConjuredDecreasesQualityByTwo(t *testing.T) {
	testUpdateQuality(t, "Conjured Mana Cake", 0, 6, 4)
}

// helpers
func testUpdateQuality(t *testing.T, itemName string, itemSellIn int, itemQuality int, expectedQuality int) {
	var items = []*Item{
		&Item{itemName, itemSellIn, itemQuality},
	}

	GildedRose(items)

	if items[0].quality != expectedQuality {
		t.Errorf("Quality (expected: %d, actual: %d).", expectedQuality, items[0].quality)
	}
}

func testUpdateSellIn(t *testing.T, itemName string, itemSellIn int, itemQuality int, expectedSellIn int) {
	var items = []*Item{
		&Item{itemName, itemSellIn, itemQuality},
	}

	GildedRose(items)

	if items[0].sellIn != expectedSellIn {
		t.Errorf("SellIn (expected: %d, actual: %d).", expectedSellIn, items[0].sellIn)
	}
}
