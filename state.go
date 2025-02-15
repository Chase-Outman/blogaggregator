package main

import (
	"github.com/Chase-Outman/blogaggregator/internal/config"
	"github.com/Chase-Outman/blogaggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
