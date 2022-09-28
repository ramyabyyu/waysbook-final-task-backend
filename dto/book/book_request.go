package bookdto

type CreateBookRequest struct {
	Title           string `json:"title" form:"title" validate:"required"`
	PublicationDate string `json:"publication_date" form:"publication_date" validate:"required"`
	Pages           string `json:"pages" form:"pages" validate:"required"`
	ISBN            string `json:"ISBN" form:"ISBN" validate:"required"`
	Price           string `json:"price" form:"price" validate:"required"`
	Description     string `json:"description" form:"description" validate:"required"`
}

type UpdateBookRequest struct {
	Title           string `json:"title" form:"title"`
	PublicationDate string `json:"publication_date" form:"publication_date"`
	Pages           string `json:"pages" form:"pages"`
	ISBN            string `json:"ISBN" form:"ISBN"`
	Price           string `json:"price" form:"price"`
	Description     string `json:"description" form:"description"`
}

type UpdateBookAttachmentRequest struct {
	BookAttachment string `json:"book_attachment" form:"book_attachment"`
}

type UpdateBookThumbnailRequest struct {
	Thumbnail string `json:"thumbnail" form:"thumbnail"`
}