package handler

import "github.com/rnikrozoft/go-graphql/model"

var monsters = []*model.Monster{
	{
		ID:    "1",
		Name:  "Goblin",
		Class: &model.AllMonsterClass[0],
		Items: []*model.Item{
			{ID: "1", Name: "Ball"},
			{ID: "2", Name: "Shoes"},
		},
	},
	{
		ID:    "2",
		Name:  "Orc",
		Class: &model.AllMonsterClass[1],
		Items: []*model.Item{
			{ID: "1", Name: "Ball"},
			{ID: "2", Name: "Shoes"},
		},
	},
}

func GetMonster() []*model.Monster {
	return monsters
}
