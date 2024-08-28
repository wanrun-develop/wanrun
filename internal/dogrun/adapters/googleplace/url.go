package googleplace

const MAP_BASE string = "https://maps.googleapis.com/maps/api/place/"
const PLACE_BASE string = "https://places.googleapis.com/v1/places/"

func urlApiPlaceWPlaceId(placeeId string) string {
	return PLACE_BASE + placeeId
}
