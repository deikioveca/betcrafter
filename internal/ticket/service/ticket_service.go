package service

import (
	"gorm.io/gorm"
)

type TicketService interface {
	Creator
	Updater
	Fetcher
	Analyzer
}

type ticketService struct {
	db *gorm.DB
}

func NewTicketService(db *gorm.DB) TicketService {
	return &ticketService{db: db}
}