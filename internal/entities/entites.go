package entities

type UserAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	DeletedAt *string      `json:"deleted_at"` // Для логического удаления
	Books     map[int]Book `json:"books"`
}

type Book struct {
	Index     int    `json:"index"`
	Title     string `json:"book"`
	Author    string `json:"author"`
	Block     *bool  `json:"block"`
	TakeCount int    `json:"take_count"`
}

type Author struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Bio   *string `json:"bio"`   // Биография автора, может быть пустой
	Books []Book  `json:"books"` // Список книг, написанных автором
}
