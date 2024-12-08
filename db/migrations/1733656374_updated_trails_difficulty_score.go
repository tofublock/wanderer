package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("e864strfxo14pm4")
		if err != nil {
			return err
		}

		// add
		new_difficulty_score := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "qt3hlfjk",
			"name": "difficulty_score",
			"type": "number",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"noDecimal": false
			}
		}`), new_difficulty_score); err != nil {
			return err
		}
		collection.Schema.AddField(new_difficulty_score)

		if err := dao.SaveCollection(collection); err != nil {
			return err
		}

		// Initialize the new field to 0 for all existing records
		_, err = db.NewQuery("UPDATE trails SET difficulty_score = 0 WHERE difficulty_score IS NULL").
			Bind(dbx.Params{
				"Collection": collection.Name,
			}).
			Execute()

		return err
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("e864strfxo14pm4")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("qt3hlfjk")

		return dao.SaveCollection(collection)
	})
}
