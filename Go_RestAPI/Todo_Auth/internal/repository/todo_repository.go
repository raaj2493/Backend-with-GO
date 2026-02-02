package repository

import (
	"Todo_Auth/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error){

	var ctx context.Context
	var cancel context.CancelFunc
	ctx , cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Query string = `
	
	        INSERT INTO Todo (title , completed)
			VALUES(&1, &2)
			RETURING id,title,completed,created_at,updated_at
	 

	`
}
