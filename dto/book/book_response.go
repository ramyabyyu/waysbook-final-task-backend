package bookdto

import (
	"time"
)

type BookResponse struct {
	ID int `json:"id"`
	Title           string    `json:"title"`
	Slug string `json:"slug"`
	PublicationDate time.Time `json:"publication_date"`
	Pages           int       `json:"pages"`
	ISBN            int       `json:"ISBN"`
	Price           int       `json:"price"`
	IsPromo bool `json:"is_promo"`
	Discount int `json:"discount"`
	PriceAfterDiscount int `json:"price_after_discount"`
	Description     string    `json:"description"`
	BookAttachment  string    `json:"book_attachment"`
	Thumbnail       string    `json:"thumbnail"`
	UserID int `json:"user_id"`
	Author string `json:"author"`
}

type UpdateBookAttachmentResponse struct {
	BookAttachment string `json:"book_attachment"`
}

type UpdateBookThumbnailResponse struct {
	Thumbnail string `json:"thumbnail"`
}