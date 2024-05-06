package handler

import (
	"github.com/google/uuid"
	"github.com/rnikrozoft/go-graphql/model"
	"github.com/samber/lo"
)

var players = []*model.Player{
	{
		ID:    "1",
		Name:  "Ronaldo",
		Class: &model.AllPlayerClass[0],
		Items: []*model.Item{
			{ID: "1", Name: "Ball"},
			{ID: "2", Name: "Shoes"},
		},
	},
	{
		ID:    "2",
		Name:  "Messi",
		Class: &model.AllPlayerClass[1],
		Items: []*model.Item{
			{ID: "1", Name: "Ball"},
			{ID: "2", Name: "Shoes"},
		},
	},
}

func GetPlayers() []*model.Player {
	return players
}

func CreatePlayer(req *model.NewPlayer) *model.Player {
	new := &model.Player{
		ID:    lo.Must1(uuid.NewV7()).String(),
		Name:  req.Name,
		Class: req.Class,
		Items: []*model.Item{
			{ID: "3", Name: "Posion"},
		},
		Level: 1,
	}
	players = append(players, new)
	return new
}
