package todo_app

type User struct {
	Id       int    `json:"-" db:"id"` // Уникальный идентификатор пользователя (скрывается при JSON-сериализации)
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
