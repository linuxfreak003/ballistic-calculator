package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	ports "github.com/linuxfreak003/ballistic-calculator/ports/repository"
	_ "github.com/mattn/go-sqlite3" // sqlite3
)

type repo struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(filename string) (ports.Repository, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	if err := InitializeDatabase(db); err != nil {
		return nil, err
	}

	return &repo{
		db: db,
	}, nil
}

// CreateLoad ...
func (r *repo) CreateLoad(ctx context.Context, load *ports.Load) (int64, error) {
	query := `INSERT INTO loads (
		bullet_caliber,
		bullet_weight,
		bullet_bc_value,
		bullet_bc_drag_func,
		bullet_length,
		muzzle_velocity
	) VALUES (?,?,?,?,?,?);`
	res, err := r.db.ExecContext(ctx, query,
		load.Bullet.Caliber,
		load.Bullet.Weight,
		load.Bullet.BC.Value,
		load.Bullet.BC.DragFunc,
		load.Bullet.Length,
		load.MuzzleVelocity,
	)
	if err != nil {
		return 0, fmt.Errorf("could not execute query: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not get last insert id: %v", err)
	}
	return id, nil
}

// ListLoads ...
func (r *repo) ListLoads(ctx context.Context) (loads []*ports.Load, err error) {
	query := `SELECT
		id,
		bullet_caliber,
		bullet_weight,
		bullet_bc_value,
		bullet_bc_drag_func,
		bullet_length,
		muzzle_velocity
		FROM loads;`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		load := &ports.Load{
			Bullet: &ports.Bullet{
				BC: &ports.Coefficient{},
			},
		}
		if err := rows.Scan(
			&load.LoadId,
			&load.Bullet.Caliber,
			&load.Bullet.Weight,
			&load.Bullet.BC.Value,
			&load.Bullet.BC.DragFunc,
			&load.Bullet.Length,
			&load.MuzzleVelocity,
		); err != nil {
			return nil, err
		}
		loads = append(loads, load)
	}
	return
}

// CreateRifle ...
func (r *repo) CreateRifle(ctx context.Context, rifle *ports.Rifle) (int64, error) {
	return 0, fmt.Errorf("DEATH")
}

// ListRifles ...
func (r *repo) ListRifles(ctx context.Context) ([]*ports.Rifle, error) {
	return nil, fmt.Errorf("DEATH")
}
