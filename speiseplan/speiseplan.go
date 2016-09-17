package speiseplan

// FeedURL is the URL of the Menu RSS Feed
const FeedURL = "http://www.studentenwerk-dresden.de/feeds/speiseplan-ifsr.rss"

// Mensa names to be used with GetCurrentForCanteen
const (
	AlteMensa           = "Alte Mensa"
	Bruehl              = "Mensa Brühl"
	GrillCube           = "Grill Cube"
	Goerlitz            = "Mensa Görlitz"
	Johannstadt         = "Mensa Johannstadt"
	Kindertagesstaetten = "Kindertagesstätten"
	Kreuzgymnasium      = "Mensa Kreuzgymnasium"
	Mensologie          = "Mensologie"
	PaluccaHochschule   = "Mensa Palucca Hochschule"
	Reichenbachstrasse  = "Mensa Reichenbachstraße"
	Siedepunkt          = "Mensa Siedepunkt"
	Sport               = "Mensa Sport"
	StimmGabel          = "Mensa Stimm-Gabel"
	TellerRandt         = "Mensa TellerRandt"
	UBoot               = "BioMensa U-Boot"
	WUEins              = "Mensa WUeins"
	Zeltschloesschen    = "Zeltschlösschen"
	Zittau              = "Mensa Zittau"
)

// GetCurrent returns current menu data
func GetCurrent() (meals []*Meal, err error) {
	items, err := parseURL(FeedURL)
	if err != nil {
		return nil, err
	}

	for _, mealItem := range items {
		meal := mealFromFeedItem(mealItem)
		meals = append(meals, &meal)
	}

	return meals, nil
}

// GetCurrentForCanteen returns all current meals for a given canteen
func GetCurrentForCanteen(canteen string) (meals []*Meal, err error) {
	all, err := GetCurrent()
	for _, meal := range all {
		if meal.Canteen == canteen {
			meals = append(meals, meal)
		}
	}
	return
}
