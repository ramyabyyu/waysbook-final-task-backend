package bookpurchaseddto

type BookPurchasedRequest struct {
	BookID        string `json:"book_id"`
	BookTitle     string `json:"book_title"`
	BookThumbnail string `json:"book_thumbnail"`
}