/**
*@Author: haoxiongxiao
*@Date: 2019/1/29
*@Description: CREATE GO FILE mgo
 */
package mgodb

import (
	"bysj/models"

	"gopkg.in/mgo.v2"
)

func connect() *mgo.Session {
	ms := models.DB.Mgo.Copy()
	ms.SetMode(mgo.Monotonic, true)
	return ms
}

func Insert(db, collection string, doc interface{}) error {
	ms := connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Insert(doc)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms := connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms := connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Find(query).Select(selector).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms := connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms := connect()
	defer ms.Close()

	_, err := ms.DB(db).C(collection).Upsert(selector, update)
	return err
}
