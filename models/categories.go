package models

import Config "money-game-2/config"

func GetAllCategories(categories *[]Categories) (err error) {
	if err = Config.DB.Find(categories).Error; err != nil {
		return err
	}
	return nil
}
