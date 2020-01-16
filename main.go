package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kobtea/go-todoist/todoist"
)

var ItemsByID map[todoist.ID]todoist.Item
var ItemsByLabelID map[todoist.ID][]todoist.Item
var ItemsByParentID map[todoist.ID][]todoist.Item

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TodoistKey := os.Getenv("TODOIST_KEY")

	cli, _ := todoist.NewClient("", TodoistKey, "*", "", nil)
	ctx := context.Background()

	// sync contents
	cli.FullSync(ctx, []todoist.Command{})

	// add item
	// item := todoist.Item{Content: "hello go-todoist hogehoge"}
	Items := cli.Item.GetAll()
	Labels := cli.Label.GetAll()
	var RoutineLabel todoist.Label
	var GLabel todoist.Label

	LabelsFound := 0

	for _, Label := range Labels {
		if Label.Name == "routine" {
			RoutineLabel = Label
			LabelsFound++
		}

		if Label.Name == "g" {
			GLabel = Label
			LabelsFound++
		}

		if LabelsFound == 2 {
			break
		}
	}

	for _, Item := range Items {
		ItemsByID[Item.ID] = Item

		for _, LabelID := range Item.Labels {
			if LabelID == RoutineLabel.ID {
				ItemsByLabelID[RoutineLabel.ID] = append(ItemsByLabelID[RoutineLabel.ID], Item)
			}
		}
	}
}
