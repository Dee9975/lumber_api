package stores

import (
	"context"
	"github.com/jackc/pgx/v5"
	"lumber/data"
)

type WorkdayStore struct {
	db *pgx.Conn
}

const (
	getAllWorkdays = "select id, team_id, warehouse_id, created_at from workday"
	//Timestamp needs to be in a format YYYY-M-D
	getWorkdayFromTimestamp = "select id, team_id, warehouse_id, created_at from workday where DATE(date_trunc('day', created_at)) = $1"
)

func NewWorkdayStore(db *pgx.Conn) (*WorkdayStore, error) {
	return &WorkdayStore{
		db: db,
	}, nil
}

func (s *WorkdayStore) GetAllWorkdays() ([]data.WorkdayRaw, error) {
	rows, err := s.db.Query(context.Background(), getAllWorkdays)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workdays := []data.WorkdayRaw{}
	for rows.Next() {
		var workday data.WorkdayRaw
		err := rows.Scan(&workday.ID, &workday.TeamID, &workday.WarehouseID, &workday.CreatedAt)
		if err != nil {
			return nil, err
		}
		workdays = append(workdays, workday)
	}

	return workdays, nil
}

func (s *WorkdayStore) GetWorkdaysFromDate(date string) ([]data.WorkdayRaw, error) {
	rows, err := s.db.Query(context.Background(), getWorkdayFromTimestamp, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workdays := []data.WorkdayRaw{}
	for rows.Next() {
		var workday data.WorkdayRaw
		err := rows.Scan(&workday.ID, &workday.TeamID, &workday.WarehouseID, &workday.CreatedAt)
		if err != nil {
			return nil, err
		}
		workdays = append(workdays, workday)
	}

	return workdays, nil
}
