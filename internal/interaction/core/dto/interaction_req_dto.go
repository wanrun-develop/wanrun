package dto

// bookmark 登録用
type BookmarkAddReq struct {
	DogrunIDs []int64 `json:"bookmarkDogrunId" validate:"required,notEmpty"`
}

type BookmarkDeleteReq struct {
	DogrunIDs []int64 `json:"bookmarkDogrunId" validate:"required,notEmpty"`
}

type CheckinReq struct {
	DogrunID int64   `json:"dogrunId" validate:"required"`
	DogIDs   []int64 `json:"dogId" validate:"required"`
}

type CheckoutReq struct {
	DogrunID int64   `json:"dogrunId" validate:"required"`
	DogIDs   []int64 `json:"dogId" validate:"required"`
}
