package service

import (
	"advanced-blog-management-system/internal/model"
	"advanced-blog-management-system/internal/repository"
	"context"
)

type CommentService struct {
	repo     repository.CommentRepository
	postRepo repository.PostRepository
}

func NewCommentService(repo repository.CommentRepository, postRepo repository.PostRepository) *CommentService {
	return &CommentService{
		repo:     repo,
		postRepo: postRepo,
	}
}

// TODO: Реализовать Create() в internal/service/comment_service.go.
// Проверьте userID и postID на положительные значения, убедитесь через postRepo,
// что пост существует, провалидируйте content (1-1000 символов) и сохраните
// Comment с переданными ID. userID приходит из middleware, postID — из URL;
// тело запроса использует поле text и не содержит post_id. Верните модель с ID.
func (s *CommentService) Create(ctx context.Context, userID, postID int, content string) (*model.Comment, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByPost(): проверить postID > 0 и наличие поста,
// limit 1-100 и offset >= 0; запросить страницу и count у commentRepo.
// Значение default limit задает handler при отсутствии query-параметра.
// При ошибке любой зависимости верните ошибку вместо частичного результата.
func (s *CommentService) GetByPost(ctx context.Context, postID, limit, offset int) ([]*model.Comment, int, error) {
	// TODO: реализовать
	return nil, 0, nil
}
