package googleplace

const (
	// MAP_BASE    string = "https://maps.googleapis.com/maps/api/place"
	PLACES_BASE        string = "https://places.googleapis.com/v1"
	PLACES_BASE_PLACES string = PLACES_BASE + "/" + PLACES
)

const (
	PLACES        string = "places"
	SEARCH_NEARBY string = "searchNearby"
	SEARCH_TEXT   string = "searchText"
	MEDIA         string = "media"
)

func urlPlacesWPlaceId(placeeId string) string {
	return PLACES_BASE_PLACES + "/" + placeeId
}

func urlPlacesWSearchNearBy() string {
	return PLACES_BASE_PLACES + ":" + SEARCH_NEARBY
}

func urlPlacesWSearchText() string {
	return PLACES_BASE_PLACES + ":" + SEARCH_TEXT
}

func urlPlacesPhotoWName(name string) string {
	return PLACES_BASE + "/" + name + "/" + MEDIA
}
