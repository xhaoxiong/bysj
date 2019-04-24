/**
*@Author: haoxiongxiao
*@Date: 2019/3/18
*@Description: CREATE GO FILE services
 */
package services

import (
	"bysj/models"
	"github.com/astaxie/beego/toolbox"
	"time"
)

func SyncDashBoard() {
	tk := toolbox.NewTask("tk1", "0 0 3 * * *", func() error {
		s, _ := time.ParseDuration("-6h")
		tx := models.GetMysqlDB().Begin()
		yesterDay := time.Now().Add(s).Format("2006-01-02")
		orderTotal := 0
		models.GetMysqlDB().Model(&models.Order{}).Where("created_at >= ?", yesterDay).
			Where("created_at < ?", time.Now().Format("2006-01-02")).Count(&orderTotal)
		var orderVolume models.OrderVolume

		orderVolume.Date = yesterDay
		orderVolume.Volume = orderTotal
		if err := models.GetMysqlDB().Where("date = ?", time.Now().Format("2006-01-02")).First(&models.OrderVolume{}).Error; err == nil {
			return err
		}

		if err := tx.Create(&orderVolume).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})

	tk.Run()
	toolbox.AddTask("tk", tk)
	toolbox.StartTask()
}
