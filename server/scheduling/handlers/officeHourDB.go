package handlers

import (
	"encoding/hex"
	"log"

	"github.com/alabama/final-project-alabama/server/scheduling/models"
	"gopkg.in/mgo.v2/bson"
)

func (ctx *Context) OfficeHoursInsert(oh *models.NewOfficeHourSession, username string) error {
	if err := officeHoursIsClean(oh); err != nil {
		return err
	}
	oh.TAs = append(oh.TAs, username)
	if err := ctx.OfficeHourCollection.Collection.Insert(oh); err != nil {
		return err
	}
	return nil
}

func (ctx *Context) GetOfficeHours() ([]models.OfficeHourSession, error) {
	var results []models.OfficeHourSession
	if err := ctx.OfficeHourCollection.Collection.Find(bson.M{}).All(&results); err != nil {
		return nil, err
	}
	for _, oh := range results {
		log.Println(oh.ID.Hex())
		decodedID, err := hex.DecodeString(oh.ID.Hex())
		if err != nil {
			log.Println(err)
			return nil, err
		}
		oh.ID = bson.ObjectId(decodedID)
	}
	return results, nil
}

func officeHoursIsClean(oh *models.NewOfficeHourSession) error {
	return nil
}
