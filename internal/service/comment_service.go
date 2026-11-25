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

// TODO: Реализовать Create()
// Валидировать postID, проверить что пост существует,
// валидировать content (не пусто, не более 1000 символов),
// создать комментарий, сохранить в БД, вернуть Comment
// userID должен приходить из middleware, а postID — из параметра URL; тело
// запроса использует поле text и не содержит post_id.
func (s *CommentService) Create(ctx context.Context, userID, postID int, content string) (*model.Comment, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByPost()
// Валидировать postID, проверить что пост существует,
// валидировать limit/offset (limit 1-100, default 10, offset >= 0),
// получить комментарии поста, вернуть слайс и общее количество
func (s *CommentService) GetByPost(ctx context.Context, postID, limit, offset int) ([]*model.Comment, int, error) {
	// TODO: реализовать
	return nil, 0, nil
}
