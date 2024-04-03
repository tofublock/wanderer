package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4wbv9tz5zjdrjh1")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id != \"\"")

		// remove
		collection.Schema.RemoveField("3neuurse")

		// remove
		collection.Schema.RemoveField("oynpksx3")

		// remove
		collection.Schema.RemoveField("tkoppkjw")

		// remove
		collection.Schema.RemoveField("uxc0eimu")

		// remove
		collection.Schema.RemoveField("qcktbht6")

		// remove
		collection.Schema.RemoveField("qpkhyujt")

		// add
		new_max_distance := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "54a1hvuq",
			"name": "max_distance",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_max_distance)
		collection.Schema.AddField(new_max_distance)

		// add
		new_max_elevation_gain := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "stcnunkf",
			"name": "max_elevation_gain",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_max_elevation_gain)
		collection.Schema.AddField(new_max_elevation_gain)

		// add
		new_max_duration := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "bholagx7",
			"name": "max_duration",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_max_duration)
		collection.Schema.AddField(new_max_duration)

		// add
		new_min_distance := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qjrtggwi",
			"name": "min_distance",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_min_distance)
		collection.Schema.AddField(new_min_distance)

		// add
		new_min_elevation_gain := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "wtlhqwd2",
			"name": "min_elevation_gain",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_min_elevation_gain)
		collection.Schema.AddField(new_min_elevation_gain)

		// add
		new_min_duration := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "sputfuyq",
			"name": "min_duration",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), new_min_duration)
		collection.Schema.AddField(new_min_duration)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("4wbv9tz5zjdrjh1")
		if err != nil {
			return err
		}

		collection.ListRule = nil

		// add
		del_max_distance := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "3neuurse",
			"name": "max_distance",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_max_distance)
		collection.Schema.AddField(del_max_distance)

		// add
		del_max_elevation_gain := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "oynpksx3",
			"name": "max_elevation_gain",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_max_elevation_gain)
		collection.Schema.AddField(del_max_elevation_gain)

		// add
		del_max_duration := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "tkoppkjw",
			"name": "max_duration",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_max_duration)
		collection.Schema.AddField(del_max_duration)

		// add
		del_min_distance := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "uxc0eimu",
			"name": "min_distance",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_min_distance)
		collection.Schema.AddField(del_min_distance)

		// add
		del_min_elevation_gain := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qcktbht6",
			"name": "min_elevation_gain",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_min_elevation_gain)
		collection.Schema.AddField(del_min_elevation_gain)

		// add
		del_min_duration := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "qpkhyujt",
			"name": "min_duration",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"maxSize": 1
			}
		}`), del_min_duration)
		collection.Schema.AddField(del_min_duration)

		// remove
		collection.Schema.RemoveField("54a1hvuq")

		// remove
		collection.Schema.RemoveField("stcnunkf")

		// remove
		collection.Schema.RemoveField("bholagx7")

		// remove
		collection.Schema.RemoveField("qjrtggwi")

		// remove
		collection.Schema.RemoveField("wtlhqwd2")

		// remove
		collection.Schema.RemoveField("sputfuyq")

		return dao.SaveCollection(collection)
	})
}
