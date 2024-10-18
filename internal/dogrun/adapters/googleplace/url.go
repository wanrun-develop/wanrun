package googleplace

const (
	MAP_BASE    string = "https://maps.googleapis.com/maps/api/place"
	PLACES_BASE string = "https://places.googleapis.com/v1/places"
)

const (
	SEARCH_NEARBY = "searchNearby"
	SEARCH_TEXT   = "searchText"
)

func urlPlacesWPlaceId(placeeId string) string {
	return PLACES_BASE + "/" + placeeId
}

func urlPlacesWSearchNearBy() string {
	return PLACES_BASE + ":" + SEARCH_NEARBY
}

func urlPlacesWSearchText() string {
	return PLACES_BASE + ":" + SEARCH_TEXT
}
