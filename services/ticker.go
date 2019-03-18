/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE services
*/
package services

import (
	"bysj/models/mgodb"
	"bysj/services/hotel_api_services"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func syncHotel() {
	toolbox.NewTask("tk1", "0 43 21 * * *", func() error {
		//rslt := make([]map[string]interface{}, 0)
		//if err := mgodb.FindPage("bysj", "city", 0,
		//	10, nil, nil, &rslt); err != nil {
		//	return err
		//} else {
		//	for index, _ := range rslt {
		//
		//		fmt.Println(rslt[index])
		//	}
		//}
			return nil
	})
}

func SyncCity() {
	tk := toolbox.NewTask("tk", "0/10 * * * * *", func() error {
		if res, err := hotel_api_services.ApiCity(); err != nil {
			return err
		} else {
			m := make(map[string]interface{})
			v := res.ShowapiResBody.CityNameList
			m["_id"] = bson.NewObjectId().Hex()
			m["list"] = v
			mgodb.Insert("bysj", "hotel", m)
		}
		fmt.Println(time.Now().String())
		return nil
	})

	tk.Run()
	toolbox.AddTask("tk", tk)
	toolbox.StartTask()
}
