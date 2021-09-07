package models

import (
	Config "money-game-2/config"
)

func GetAllActivities(activities *[]Activities) (err error) {
	if err = Config.DB.Find(activities).Error; err != nil {
		return err
	}
	return nil
}
