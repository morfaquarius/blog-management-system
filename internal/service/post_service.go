package service

import (
	"advanced-blog-management-system/internal/model"
	"advanced-blog-management-system/internal/model/repository"
	"context"
)

type PostService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

// Не принимайте authorID от клиента при создании: handler передает сюда ID,
// извлеченный middleware из проверенных JWT claims.
func NewPostService(postRepo repository.PostRepository, userRepo repository.UserRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

// TODO: Реализовать Create() в internal/service/post_service.go.
// Провалидируйте req, проверьте что userID положительный и пользователь
// существует (userRepo), затем сформируйте model.Post с AuthorID=userID и
// сохраните через postRepo.Create. Верните созданную модель с заполненным ID;
// authorID не берите из клиентского DTO.
func (s *PostService) Create(ctx context.Context, userID int, req *model.PostCreateRequest) (*model.Post, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByID(): отклонить неположительный ID, запросить
// postRepo.GetByID и вернуть ErrPostNotFound при согласованном nil-результате.
func (s *PostService) GetByID(ctx context.Context, id int) (*model.Post, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetAll(): проверить limit 1-100 и offset >= 0, получить
// страницу из postRepo.GetAll и общий total из GetTotalCount. Если любой запрос
// завершается ошибкой, верните ее; не возвращайте частичный успех.
func (s *PostService) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, int, error) {
	// TODO: реализовать
	return nil, 0, nil
}

// TODO: Реализовать GetByAuthor(): проверить authorID > 0, limit 1-100 и
// offset >= 0; получить страницу и total через методы репозитория по автору.
// Этот сценарий нужен только если добавлен необязательный маршрут автора.
func (s *PostService) GetByAuthor(ctx context.Context, authorID int, limit, offset int) ([]*model.Post, int, error) {
	// TODO: реализовать
	return nil, 0, nil
}
