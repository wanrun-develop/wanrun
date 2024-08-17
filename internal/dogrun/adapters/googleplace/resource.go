package googleplace

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BaseResource struct {
	ID                    string   `json:"id"`
	Location              Location `json:"location"`
	ShortFormattedAddress string   `json:"shortFormattedAddress"`
}
