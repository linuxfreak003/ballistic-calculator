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

// InitializeDatabase ...
func InitializeDatabase(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := `CREATE TABLE IF NOT EXISTS rifles (
		id INTEGER PRIMARY KEY,
		name TEXT,
		sight_height FLOAT,
		barrel_twist FLOAT,
		twist_direction_left BOOL,
		zero_range FLOAT
	)`
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS loads (
		id INTEGER PRIMARY KEY,
		bullet_caliber FLOAT,
		bullet_weight FLOAT,
		bullet_bc_value FLOAT,
		bullet_bc_drag_func INTEGER,
		bullet_length FLOAT,
		muzzle_velocity FLOAT
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS environments (
		id INTEGER PRIMARY KEY,
		temperature FLOAT,
		altitude INTEGER,
		pressure FLOAT,
		humidity FLOAT,
		wind_angle FLOAT,
		wind_speed FLOAT,
		pressure_is_absolute BOOL,
		latitude FLOAT,
		azimuth FLOAT
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS scenarios (
		id INTEGER PRIMARY KEY,
		name TEXT,
		environment_id INTEGER,
		rifle_id INTEGER,
		load_id INTEGER
	)`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	return nil
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
