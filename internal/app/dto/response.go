package dto

type Response struct {
	Page int					`json:"page"`
	Per_page int				`json:"per_page"`
	Total int					`json:"total"`
	Total_pages int				`json:"total_pages"`
	Data []Data					`json:"data"`
	Support Support				`json:"support"`
}