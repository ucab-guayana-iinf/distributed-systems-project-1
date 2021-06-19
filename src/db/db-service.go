package dbService

import (
	"io/ioutil"

	utils "proyecto1.com/main/src/utils"
)

var db_filename = "countdb.txt"

func Initialize() int {
	content, err := ioutil.ReadFile(db_filename)

	if (err != nil) {
		val := "0"
    data := []byte(val)
		ioutil.WriteFile(db_filename, data, 0644)
		return 0
	}

	return utils.StringToInt(string(content))
}

func UpdateCount(value int) {
	val := utils.IntToString(value)
	data := []byte(val)
	ioutil.WriteFile(db_filename, data, 0644)
}