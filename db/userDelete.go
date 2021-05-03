package db

import "log"

func DelUser(ced_id int) error {
	stmt, err := DB.Prepare("DELETE from user where id = ? or cedula = ? AND role = 1")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ced_id)

	return err
}
