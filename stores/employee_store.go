package stores

import (
	"context"
	"github.com/jackc/pgx/v5"
	"lumber/data"
)

type EmployeeStore struct {
	db *pgx.Conn
}

const (
	createTeamsTableSQL         = "create table if not exists teams (id serial not null primary key, name text not null, created_at timestamp default current_timestamp)"
	createEmployeesTableSQL     = "create table if not exists employees (id serial not null primary key, first_name text not null, last_name text not null, email text not null)"
	createTeamEmployeesTableSQL = "create table if not exists team_employees (id serial not null primary key, team_id int not null, employee_id int not null, hours_worked int default 0, hours_worked_forklift int default 0, hours_worked_heating int default 0)"
	getAllEmployeesQuery        = "select * from employees"
	getEmployeesByIdQuery       = "select * from employees where id = $1"
	createEmployeeSQL           = "insert into employees (first_name, last_name, email) values ($1, $2, $3)"
	updateEmployeesSQL          = "update employees set first_name = $1, last_name = $2, email = $3 where id = $4"
	deleteEmployeeSQL           = "delete from employees where id = $1"
	getAllTeamEmployees         = `select 
       	employees.id,
       	employees.first_name,
       	employees.last_name,
       	employees.email,
       	te.hours_worked,
       	te.hours_worked_heating,
       	te.hours_worked_forklift
	from 
	    employees
    join 
	    team_employees te
    on 
        employees.id = te.employee_id and team_id = $1`
)

func NewEmployeeStore(db *pgx.Conn) (*EmployeeStore, error) {
	if _, err := db.Exec(context.Background(), createEmployeesTableSQL); err != nil {
		return nil, err
	}
	if _, err := db.Exec(context.Background(), createTeamsTableSQL); err != nil {
		return nil, err
	}
	if _, err := db.Exec(context.Background(), createTeamEmployeesTableSQL); err != nil {
		return nil, err
	}
	return &EmployeeStore{
		db: db,
	}, nil
}

func (s *EmployeeStore) CreateEmployee(e data.Employee) error {
	_, err := s.db.Exec(context.Background(), createEmployeeSQL, e.FirstName, e.LastName, e.Email)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeStore) UpdateEmployee(id int, e data.Employee) error {
	_, err := s.db.Exec(context.Background(), updateEmployeesSQL, e.FirstName, e.LastName, e.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeStore) DeleteEmployee(id int) error {
	_, err := s.db.Exec(context.Background(), deleteEmployeeSQL, id)

	if err != nil {
		return err
	}
	return nil
}

func (s *EmployeeStore) GetAllEmployees() ([]data.EmployeeResponse, error) {
	rows, err := s.db.Query(context.Background(), getAllEmployeesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []data.EmployeeResponse
	for rows.Next() {
		var employee data.EmployeeResponse

		err := rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Email)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *EmployeeStore) GetEmployeeById(id int) (data.EmployeeResponse, error) {
	row := s.db.QueryRow(context.Background(), getEmployeesByIdQuery, id)
	var employee data.EmployeeResponse
	err := row.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Email)
	if err != nil {
		return data.EmployeeResponse{}, err
	}
	return employee, nil
}

func (s *EmployeeStore) GetEmployeesByTeamId(id int) ([]data.TeamEmployeeResponse, error) {
	rows, err := s.db.Query(context.Background(), getAllTeamEmployees, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workers []data.TeamEmployeeResponse
	for rows.Next() {
		var response data.TeamEmployeeResponse

		err := rows.Scan(&response.ID, &response.FirstName, &response.LastName, &response.Email, &response.HoursWorked, &response.HoursWorkedHeating, &response.HoursWorkedForklift)
		if err != nil {
			return nil, err
		}
		workers = append(workers, response)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return workers, nil
}
