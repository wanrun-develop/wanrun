package facade

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/common"
	"github.com/wanrun-develop/wanrun/internal/interaction/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/wrcontext"
)

type IBookmarkFacade interface {
	GetAllUserBookmarks(echo.Context) ([]int64, error)
	GetAllUserBookmarksByPage(echo.Context, common.PaginationReq) ([]int64, error)
	GetBookmarkByDogrunID(echo.Context, int64) (bool, error)
}

type bookmarkFacade struct {
	r repository.IBookmarkRepository
}

func NewBookmarkFacade(br repository.IBookmarkRepository) IBookmarkFacade {
	return &bookmarkFacade{br}
}

// GetAllUserBookmarks: ログインユーザーのブックマークを取得
//
// args:
//   - echo.Context:	コンテキスト
//
// return:
//   - []int64:	bookmarkIDs
//   - error:	エラー
func (f *bookmarkFacade) GetAllUserBookmarks(c echo.Context) ([]int64, error) {
	// ログインユーザーIDの取得
	userID, err := wrcontext.GetLoginUserID(c)
	if err != nil {
		return nil, err
	}

	bookmarks, err := f.r.GetBookmarks(c, userID)
	if err != nil {
		return nil, err
	}

	bookmarkedDogrunIDs := []int64{}
	for _, bookmark := range bookmarks {
		bookmarkedDogrunIDs = append(bookmarkedDogrunIDs, bookmark.DogrunID.Int64)
	}

	return bookmarkedDogrunIDs, nil
}

// GetAllUserBookmarksByPage: ログインユーザーのブックマークを取得
// paginationあり
// args:
//   - echo.Context:	コンテキスト
//
// return:
//   - []int64:	bookmarkIDs
//   - error:	エラー
func (f *bookmarkFacade) GetAllUserBookmarksByPage(c echo.Context, page common.PaginationReq) ([]int64, error) {
	// ログインユーザーIDの取得
	userID, err := wrcontext.GetLoginUserID(c)
	if err != nil {
		return nil, err
	}

	limit := page.Count
	offset := page.Count * (page.Page - 1)
	bookmarks, err := f.r.GetBookmarksByPage(c, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	bookmarkedDogrunIDs := []int64{}
	for _, bookmark := range bookmarks {
		bookmarkedDogrunIDs = append(bookmarkedDogrunIDs, bookmark.DogrunID.Int64)
	}

	return bookmarkedDogrunIDs, nil
}

// GetBookmarkByDogrunID: 特定のdogrunIdに対するブックマーク状態を取得
//
// args:
//   - echo.Context:	コンテキスト
//   - int64:	dogrunID 検索対象のドッグランID
//
// return:
//   - bool:	ブックマーク済みかどうか
//   - error:	エラー
func (f *bookmarkFacade) GetBookmarkByDogrunID(c echo.Context, dogrunID int64) (bool, error) {
	// ログインユーザーIDの取得
	userID, err := wrcontext.GetLoginUserID(c)
	if err != nil {
		return false, err
	}

	// 指定されたdogrunIDのブックマーク情報を検索
	bookmark, err := f.r.FindDogrunBookmark(c, dogrunID, userID)
	if err != nil {
		return false, err
	}

	// ブックマークが存在するかどうかを返す
	return bookmark.IsNotEmpty(), nil
}
