package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kulmenevaa/go-start/app/models"
)

type PostRepository struct {
	dbClient *sql.DB
}

func NewPostRepository(dbClient *sql.DB) *PostRepository {
	return &PostRepository{
		dbClient: dbClient,
	}
}

type PostRepositoryInterface interface {
	GetPostById(ID int) (*models.Post, error)
	GetAllPosts() (*[]models.Post, error)
	SavePost(*models.Post) (bool, error)
	DeletePost(ID int) (bool, error)
	UpdatePost(post *models.Post, PostID int) (bool, error)
}

func (i *PostRepository) GetPostById(ID int) (*models.Post, error) {
	var post models.Post
	err := i.dbClient.QueryRow(`SELECT * FROM posts WHERE id=$1`, ID).Scan(&post.ID, &post.Title, &post.Description,
		&post.CreatedAt, &post.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows were returned!")
		return nil, nil
	case nil:
		return &post, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return &post, nil
}

func (i *PostRepository) GetAllPosts() (*[]models.Post, error) {
	rows, err := i.dbClient.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Printf("ERROR SELECT QUERY - %s", err)
		return nil, err
	}
	var postList []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.ID, &post.Title, &post.Description,
			&post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Printf("ERROR QUERY SCAN - %s", err)
			return nil, err
		}
		postList = append(postList, post)
	}
	return &postList, nil
}

func (i *PostRepository) SavePost(post *models.Post) (bool, error) {
	sqlStatement := `INSERT into posts (name, description, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := i.dbClient.Exec(sqlStatement, post.Title, post.Description, time.Now().Local(), time.Now().Local())
	if err != nil {
		log.Printf("ERROR EXEC INSERT QUERY - %s", err)
		return false, err
	}
	return true, nil
}

func (i *PostRepository) DeletePost(ID int) (bool, error) {
	_, err := i.dbClient.Exec(`DELETE FROM posts WHERE id=$1`, ID)
	if err != nil {
		log.Printf("ERROR EXEC DELETE QUERY - %s", err)
		return false, err
	}
	return true, nil
}

func (i *PostRepository) UpdatePost(post *models.Post, PostID int) (bool, error) {
	sqlStatement := `UPDATE posts SET name=$1, description=$2, updated_at=$3 WHERE id=$4`
	_, err := i.dbClient.Exec(sqlStatement, post.Title, post.Description, time.Now().Local(), PostID)
	if err != nil {
		fmt.Printf("ERROR EXEC UPDATE QUERY - %s", err)
	}
	return true, nil
}
