package main

import "testing"

// Regular Item tests
func TestSellInDropsOne(t *testing.T) {
	testSellIn(t, "aRegularItem", 6, 9, 5)
}

func TestQualityReducesOne(t *testing.T) {
	testQuality(t, "aRegularItem", 6, 9, 8)
}

func TestQualityReducesTwoAfterSellDate(t *testing.T) {
	testQuality(t, "aRegularItem", 0, 9, 7)
}

func TestQualityIsNeverNegative(t *testing.T) {
	testQuality(t, "aRegularItem", 5, 0, 0)
}

// Aged Brie Item tests
func TestAgedBrieSellInDropsOne(t *testing.T) {
	testSellIn(t, "Aged Brie", 5, 10, 4)
}

func TestAgedBrieQualityImprovesOne(t *testing.T) {
	testQuality(t, "Aged Brie", 5, 10, 11)
}

func TestAgedBrieQualityImprovesTwoAfterSellDate(t *testing.T) {
	testQuality(t, "Aged Brie", 0, 10, 12)
}

func TestAgedBrieQualityIsMaxFifty(t *testing.T) {
	testQuality(t, "Aged Brie", 5, 50, 50)
}

// Sulfuras Item tests
func TestSulfurasSellInNeverDrops(t *testing.T) {
	testSellIn(t, "Sulfuras, Hand of Ragnaros", 0, 80, 0)
}

func TestSulfurasQualityNeverReduces(t *testing.T) {
	testQuality(t, "Sulfuras, Hand of Ragnaros", 0, 80, 80)
}

// Backstage Item tests
func TestBackstageSellInDropsOne(t *testing.T) {
	testSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 6, 10, 5)
}

func TestBackstageQualityImprovesOneIfSellInIsElevenOrMore(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 12, 20, 21)
}

func TestBackstageQualityImprovesTwoIfSellInIsTenOrLess(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 11, 20, 22)
}

func TestBackstageQualityImprovesTwoIfSellInIsFiveOrLess(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 6, 20, 23)
}

func TestBackstageQualityIsZeroAfterSellDate(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 0, 20, 0)
}

func TestBackstageQualityIsMaxFiftyWhenIncreasingOne(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 11, 50, 50)
}

func TestBackstageQualityIsMaxFiftyWhenIncreasingTwo(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 7, 49, 50)
}

func TestBackstageQualityIsMaxFiftyWhenIncreasingThree(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 1, 49, 50)
}

// Conjured Item tests
func TestConjuredSellInDropsOne(t *testing.T) {
	testSellIn(t, "Conjured Mana Cake", 8, 12, 7)
}

func TestConjuredQualityReducesTwo(t *testing.T) {
	testQuality(t, "Conjured Mana Cake", 8, 12, 10)
}

func TestConjuredQualityReducesFourAfterSellDate(t *testing.T) {
	testQuality(t, "Conjured Mana Cake", 0, 12, 8)
}

func TestConjuredQualityIsNeverNegative(t *testing.T) {
	testQuality(t, "Conjured Mana Cake", 8, 0, 0)
}

// Two reusable test functions
// testSellIn: Tests if the expected SellIn after an update corresponds with expected SellIn
// testQuality: Tests if the expected Quality after an update corresponds with expected Quality
func testSellIn(t *testing.T, itemName string, currentSellIn, currentQuality, expectedSellIn int) {
	items := []*Item{
		{itemName, currentSellIn, currentQuality},
	}

	UpdateQuality(items)

	if items[0].sellIn != expectedSellIn {
		t.Errorf("SellIn (expected: %d, actual: %d).", expectedSellIn, items[0].sellIn)
	}
}

func testQuality(t *testing.T, itemName string, currentSellIn, currentQuality, expectedQuality int) {
	items := []*Item{
		{itemName, currentSellIn, currentQuality},
	}

	UpdateQuality(items)

	if items[0].quality != expectedQuality {
		t.Errorf("Quality (expected: %d, actual: %d).", expectedQuality, items[0].quality)
	}
}
