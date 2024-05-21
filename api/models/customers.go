package models

import "time"

type Customers struct {
	Id         string   `json:"id"`
	ExternalId string   `json:"external_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Age        int      `json:"age"`
	Phone      []string `json:"phone"`
	Mail       string   `json:"mail"`
	Birthday   string   `json:"birthday"`
	Sex        string   `json:"sex"`
}

type CreateOrUpdate struct {
	Id         string   `json:"id"`
	ExternalId string   `json:"external_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Age        int      `json:"age"`
	Phone      []string `json:"phone"`
	Mail       string   `json:"mail"`
	Birthday   string   `json:"birthday"`
	Sex        string   `json:"sex"`
}

type GetCustomers struct {
	Id         string   `json:"id"`
	ExternalId string   `json:"external_id"`
	FirstName  string   `json:"first_name"`
	LastName   string   `json:"last_name"`
	Age        int      `json:"age"`
	Phone      []string `json:"phone"`
	Mail       string   `json:"mail"`
	Birthday   time.Time   `json:"birthday"`
	Sex        string   `json:"sex"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type GetAllCustomersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCustomersResponse struct {
	Customers []GetCustomers `json:"Customerss"`
	Count      int            `json:"count"`
}
