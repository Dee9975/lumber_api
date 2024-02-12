package stores

import (
	"context"
	"github.com/jackc/pgx/v5"
	"lumber/data"
	"lumber/util"
)

type LumberStore struct {
	db *pgx.Conn
}

const (
	createLumberTableSQL = `create table if not exists 
		lumber(
		id serial primary key not null, 
		width integer not null,
		height integer not null,
		len integer not null,
		amount integer not null,
		team_id integer)`

	createWarehouseTableSQL = "create table if not exists warehouses (id serial not null primary key, name text not null, rate int not null, custom_rate text not null)"

	getAllLumberQuery  = "select width, height, len, amount, team_id from lumber"
	createLumberSQL    = "insert into lumber (width, height, len, amount, team_id) values ($1, $2, $3, $4, $5)"
	getLumberByIDQuery = "select width, height, len, amount, team_id from lumber where id = $1"

	updateLumberSQL = `update lumber set 
		width = $1,
		len = $2,
		height = $3,
		amount=$4,
		team_id=$5
		where id = $6`

	getAllTeamLumber = `select id, len, width, height, amount from lumber where team_id = $1`
	getWarehouseById = "select name, rate, custom_rate from warehouses where id = $1"
)

func NewLumberStore(db *pgx.Conn) (*LumberStore, error) {
	// Create the table if it doesn't exist

	_, err := db.Exec(context.Background(), createLumberTableSQL)
	_, err = db.Exec(context.Background(), createWarehouseTableSQL)

	if err != nil {
		return nil, err
	}

	return &LumberStore{
		db: db,
	}, nil
}

func (s *LumberStore) GetLumber() ([]data.LumberResponse, error) {
	var lumber []data.LumberResponse

	rows, err := s.db.Query(context.Background(), getAllLumberQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var l data.Lumber

		if err := rows.Scan(&l.Width, &l.Height, &l.Length, &l.Amount, &l.TeamID); err != nil {
			return nil, err
		}

		volume := l.Height * l.Width * l.Length * l.Amount

		response := data.LumberResponse{
			Lumber: l,
			Volume: volume,
		}

		lumber = append(lumber, response)
	}

	if lumber == nil {
		lumber = []data.LumberResponse{}
	}

	return lumber, nil
}

func (s *LumberStore) CreateLumber(lumber data.Lumber) error {
	_, err := s.db.Exec(
		context.Background(),
		createLumberSQL,
		lumber.Width,
		lumber.Height,
		lumber.Length,
		lumber.Amount,
		lumber.TeamID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *LumberStore) GetLumberById(id int) (*data.LumberResponse, error) {
	var l data.Lumber

	err := s.db.QueryRow(
		context.Background(),
		getLumberByIDQuery,
		id,
	).Scan(
		&l.Width,
		&l.Height,
		&l.Length,
		&l.Amount,
		&l.TeamID,
	)

	if err != nil {
		return nil, err
	}

	return util.CreateLumberResponse(l), err
}

func (s *LumberStore) DeleteLumber(id int) error {
	_, err := s.db.Exec(context.Background(), "delete from lumber where id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (s *LumberStore) UpdateLumber(id int, lumber data.Lumber) (*data.LumberResponse, error) {
	_, err := s.db.Exec(context.Background(),
		updateLumberSQL,
		lumber.Width,
		lumber.Length,
		lumber.Height,
		lumber.Amount,
		lumber.TeamID,
		id,
	)

	if err != nil {
		return nil, err
	}

	return util.CreateLumberResponse(lumber), nil
}

func (s *LumberStore) GetAllTeamLumber(teamId int) ([]data.TeamLumberResponse, error) {
	rows, err := s.db.Query(context.Background(), getAllTeamLumber, teamId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lumber []data.TeamLumberResponse

	for rows.Next() {
		var l data.TeamLumberResponse

		if err := rows.Scan(&l.ID, &l.Length, &l.Width, &l.Height, &l.Amount); err != nil {
			return nil, err
		}

		lumber = append(lumber, l)
	}

	if lumber == nil {
		lumber = []data.TeamLumberResponse{}
	}

	return lumber, nil
}

func (s *LumberStore) GetWarehouseById(id int) (*data.Warehouse, error) {
	row := s.db.QueryRow(context.Background(), getWarehouseById, id)
	var warehouse data.Warehouse
	err := row.Scan(&warehouse.Name, &warehouse.Rate, &warehouse.CustomRate)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}
