package googleplace

const mapBase string = "https://maps.googleapis.com/maps/api/place/"
const placeBase string = "https://places.googleapis.com/v1/places/"

func urlApiPlaceWPlaceId(placeeId string) string {
	return placeBase + placeeId
}
