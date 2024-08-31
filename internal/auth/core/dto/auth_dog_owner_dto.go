package dto

type ReqAuthDogOwnerDto struct {
	Password     string `json:"password"`
	DogOwnerName string `json:"dogOwnerName"`
	Email        string `json:"email"`
}

type ResAuthDogOwner struct {
	AuthDogOwnerID uint   `json:"authDogOwnerId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Sex            string `json:"sex"`
}
