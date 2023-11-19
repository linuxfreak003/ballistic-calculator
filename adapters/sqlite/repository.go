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

// GetLoad ...
func (r *repo) GetLoad(ctx context.Context, id int64) (*ports.Load, error) {
	return nil, fmt.Errorf("DEATH")
}

// CreateRifle ...
func (r *repo) CreateRifle(ctx context.Context, in *ports.Rifle) (int64, error) {
	return 0, fmt.Errorf("DEATH")
}

// ListRifles ...
func (r *repo) ListRifles(ctx context.Context) ([]*ports.Rifle, error) {
	return nil, fmt.Errorf("DEATH")
}

// GetRifle ...
func (r *repo) GetRifle(ctx context.Context, id int64) (*ports.Rifle, error) {
	return nil, fmt.Errorf("DEATH")
}

// CreateEnvironment ...
func (r *repo) CreateEnvironment(ctx context.Context, in *ports.Environment) (int64, error) {
	return 0, fmt.Errorf("DEATH")
}

// ListEnvironments ...
func (r *repo) ListEnvironments(ctx context.Context) ([]*ports.Environment, error) {
	return nil, fmt.Errorf("DEATH")
}

// GetEnvironment ...
func (r *repo) GetEnvironment(ctx context.Context, id int64) (*ports.Environment, error) {
	return nil, fmt.Errorf("DEATH")
}

// CreateScenario ...
func (r *repo) CreateScenario(ctx context.Context, in *ports.Scenario) (int64, error) {
	return 0, fmt.Errorf("DEATH")
}

// ListScenarios ...
func (r *repo) ListScenarios(ctx context.Context) ([]*ports.Scenario, error) {
	return nil, fmt.Errorf("DEATH")
}

// GetScenario ...
func (r *repo) GetScenario(ctx context.Context, id int64) (*ports.Scenario, error) {
	return nil, fmt.Errorf("DEATH")
}
