package facades

import (
	"context"

	"studentgit.kata.academy/Zhodaran/go-kata/internal/domain"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/entities"
	"studentgit.kata.academy/Zhodaran/go-kata/internal/infrastructure/postgres"
)

type LibraryFacade struct {
	authRepo   domain.AuthRepository
	bookRepo   domain.BookRepository
	authorRepo domain.AuthorRepository
	userRepo   domain.UserRepository
}

func NewLibraryFacade(
	auth domain.AuthRepository,
	book domain.BookRepository,
	author domain.AuthorRepository,
	user domain.UserRepository,
) *LibraryFacade {
	return &LibraryFacade{
		authRepo:   auth,
		bookRepo:   book,
		authorRepo: author,
		userRepo:   user,
	}
}

func (f *LibraryFacade) GetAllBooks() ([]entities.Book, error) {
	return f.bookRepo.GetAllBooks()
}

func (f *LibraryFacade) UpdateBook(index int, updatedBook entities.Book) error {
	return f.bookRepo.UpdateBook(index, updatedBook)
}

func (f *LibraryFacade) AddBook(addaderBook postgres.AddaderBookRequest) error {
	return f.bookRepo.AddBook(addaderBook)
}

func (f *LibraryFacade) TakeBook(id int) (entities.Book, error) {
	return f.bookRepo.TakeBook(id)
}

func (f *LibraryFacade) ReturnBook(id int) error {
	return f.bookRepo.ReturnBook(id)
}

func (f *LibraryFacade) ListAuthors() ([]entities.Author, error) {
	return f.authorRepo.ListAuthors()
}

func (f *LibraryFacade) GetAuthors() ([]entities.Author, error) {
	return f.authorRepo.GetAuthors()
}

func (f *LibraryFacade) AddAuthor(author entities.Author) error {
	return f.authorRepo.AddAuthor(author)
}

func (f *LibraryFacade) Create(ctx context.Context, user entities.User) error {
	return f.userRepo.Create(ctx, user)
}

func (f *LibraryFacade) GetByID(ctx context.Context, id string) (entities.User, error) {
	return f.userRepo.GetByID(ctx, id)
}

func (f *LibraryFacade) Update(ctx context.Context, user entities.User) error {
	return f.userRepo.Update(ctx, user)
}

func (f *LibraryFacade) Delete(ctx context.Context, id string) error {
	return f.userRepo.Delete(ctx, id)
}

func (f *LibraryFacade) List(ctx context.Context, limit, offset int) ([]entities.User, error) {
	return f.userRepo.List(ctx, limit, offset)
}

func (f *LibraryFacade) Login(username, password string) (entities.UserAuth, error) {
	return f.authRepo.Login(username, password)
}

func (f *LibraryFacade) Register(user entities.UserAuth) error {
	return f.authRepo.Register(user)
}
