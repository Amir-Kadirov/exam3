package postgres

import (
	"context"
	"database/sql"
	"exam/api/models"
	"exam/pkg"
	smtp "exam/pkg/helper"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type customersRepo struct {
	db *pgxpool.Pool
}

func NewCustomers(db *pgxpool.Pool) customersRepo {
	return customersRepo{
		db: db,
	}
}

func (s *customersRepo) Create(ctx context.Context, customers models.Customers) (string, error) {

	id := uuid.New()
	ExternalId, err := smtp.ExternalIdGenrerator(s.db)
	if err != nil {
		return "", err
	}

	query := `
		INSERT INTO customers (id, external_id, first_name, last_name, age, phone, mail, birthday, sex, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW()) `

	_, err = s.db.Exec(ctx, query, id, ExternalId, customers.FirstName, customers.LastName, customers.Age,
		pq.Array(customers.Phone), customers.Mail, customers.Birthday, customers.Sex)

	if err != nil {
		fmt.Println("in customers insert")
		return "", err
	}

	return id.String(), nil
}

func (s *customersRepo) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	resp := models.GetAllCustomersResponse{}
	filter := ""
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `
        SELECT id,
		       external_id,
			   first_name,
			   last_name,
			   age,
			   COALESCE(phone, '{}'),
			   mail,
			   sex,
			   birthday,
			   created_at,
			   updated_at
        FROM customers
        WHERE TRUE ` + filter + `
        OFFSET $1 LIMIT $2
    `

	rows, err := s.db.Query(ctx, query, offset, req.Limit)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			customers models.GetCustomers
			UpdatedAt sql.NullString
			Phone     []string
		)

		if err := rows.Scan(
			&customers.Id,
			&customers.ExternalId,
			&customers.FirstName,
			&customers.LastName,
			&customers.Age,
			&Phone,
			&customers.Mail,
			&customers.Sex,
			&customers.Birthday,
			&customers.CreatedAt,
			&UpdatedAt,
		); err != nil {
			return resp, err
		}

		customers.Phone = Phone
		customers.UpdatedAt = pkg.NullStringToString(UpdatedAt)
		resp.Customers = append(resp.Customers, customers)
	}

	if err := rows.Err(); err != nil {
		return resp, err
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from customers WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}



func (s *customersRepo) Update(ctx context.Context, ExternalId string, customers models.Customers) error {

	query := ` UPDATE customers set first_name=$1,
								   last_name=$2,
								   age=$3,
								   phone=$4,
								   mail=$5,
								   birthday=$6,
								   sex=$7,
								   updated_at = NOW()
								       WHERE id=$8 `

	_, err := s.db.Exec(ctx, query,
		customers.FirstName,
		customers.LastName,
		customers.Age,
		pq.Array(customers.Phone),
		customers.Mail,
		customers.Birthday,
		customers.Sex,
		ExternalId)
	if err != nil {
		return err
	}

	return nil
}

func (s *customersRepo) GetById(ctx context.Context, id string) (models.GetCustomers, error) {
	resp := models.GetCustomers{}

	query := `SELECT id,
		external_id,
		first_name,
		last_name,
		age,
		COALESCE(phone, '{}'),
		mail,
		birthday,
		sex,
		created_at,
		updated_at
			  FROM customers
			  WHERE id=$1`

	row := s.db.QueryRow(ctx, query, id)

	var updatedAt sql.NullString
	var phone []string

	if err := row.Scan(&resp.Id,
		&resp.ExternalId,
		&resp.FirstName,
		&resp.LastName,
		&resp.Age,
		&phone,
		&resp.Mail,
		&resp.Birthday,
		&resp.Sex,
		&resp.CreatedAt,
		&updatedAt,
	); err != nil {
		return resp, err
	}

	resp.Phone = phone
	resp.UpdatedAt = pkg.NullStringToString(updatedAt)
	return resp, nil
}

func (s *customersRepo) Delete(ctx context.Context, id string) error {

	query := `DELETE FROM customers WHERE id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
