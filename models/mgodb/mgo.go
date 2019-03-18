/**
*@Author: haoxiongxiao
*@Date: 2019/1/29
*@Description: CREATE GO FILE mgo
 */
package mgodb

import (
	"bysj/models"

	"gopkg.in/mgo.v2"
	"log"
)

func Connect() *mgo.Session {
	ms := models.DB.Mgo.Copy()
	ms.SetMode(mgo.Monotonic, true)
	return ms
}

func Insert(db, collection string, doc interface{}) error {
	ms := Connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Insert(doc)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	ms := Connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms := Connect()
	defer ms.Close()
	return ms.DB(db).C(collection).Find(query).Select(selector).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	ms := Connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Update(selector, update)
}

func Upsert(db, collection string, selector, update interface{}) error {
	ms := Connect()
	defer ms.Close()

	_, err := ms.DB(db).C(collection).Upsert(selector, update)
	return err
}

func Remove(db, collection string, selector interface{}) error {
	ms := Connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Remove(selector)
}

func RemoveAll(db, collection string, selector interface{}) error {
	ms := Connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Remove(selector)
}

func FindPage(db, collection string, page, limit int, query, selector, result interface{}) error {
	ms := Connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Find(query).Select(selector).Skip(page * limit).Limit(limit).All(result)
}

func Count(db, collection string, query interface{}) (int, error) {
	ms := Connect()
	defer ms.Close()

	return ms.DB(db).C(collection).Find(query).Count()
}

func IsEmpty(db, collection string) bool  {
	ms := Connect()
	defer ms.Close()
	count, err := ms.DB(db).C(collection).Count()
	if err != nil {
		log.Fatal(err)
	}
	return count == 0

}