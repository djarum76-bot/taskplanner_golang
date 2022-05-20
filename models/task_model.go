package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/djarum76-bot/taskplanner_golang/db"
)

func AddTask(userId int, title string, tanggal string, waktu string, date string, deskripsi string) (Response, error) {
	var task Task
	var err error
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT into tasks (user_id, title, tanggal, waktu, date, deskripsi) values ($1, $2, $3, $4, $5, $6)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(userId, title, tanggal, waktu, date, deskripsi)
	if err != nil {
		return res, err
	}

	getID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	task.Id = int(getID)
	task.User_Id = userId
	task.Title = title
	task.Tanggal = tanggal
	task.Waktu = waktu
	task.Date = date
	task.Deskripsi = deskripsi

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Tambah Data"
	res.Data = task

	return res, err
}

func GetAllTask(user_id int) (Response, error) {
	var task Task
	var arrTask []Task = []Task{}
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tasks WHERE user_id = ($1) AND date > ($2) ORDER BY date"

	rows, err := con.Query(sqlStatement, user_id, time.Now().String())
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&task.Id, &task.User_Id, &task.Title, &task.Tanggal, &task.Waktu, &task.Date, &task.Deskripsi)
		if err != nil {
			return res, err
		}

		arrTask = append(arrTask, task)
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = arrTask

	return res, nil
}

func GetAllTaskDate(user_id int) (Response, error) {
	var task Task
	var arrTask []Task = []Task{}
	var taskDate TaskDate
	var arrTaskDate []TaskDate = []TaskDate{}
	var tanggal string
	var res Response

	con := db.CreateCon()

	sqlStatement1 := "SELECT DISTINCT tanggal FROM tasks WHERE date > ($1) ORDER BY date"

	rows1, err := con.Query(sqlStatement1, time.Now().String())
	if err != nil {
		return res, err
	}
	defer rows1.Close()

	for rows1.Next() {
		err = rows1.Scan(&tanggal)
		if err != nil {
			return res, err
		}

		sqlStatement2 := "SELECT * FROM tasks WHERE user_id = ($1) AND tanggal = ($2) AND date > ($3) ORDER BY date"
		rows2, err := con.Query(sqlStatement2, user_id, tanggal, time.Now().String())
		if err != nil {
			return res, err
		}
		defer rows2.Close()

		for rows2.Next() {
			err = rows2.Scan(&task.Id, &task.User_Id, &task.Title, &task.Tanggal, &task.Waktu, &task.Date, &task.Deskripsi)
			if err != nil {
				return res, err
			}

			arrTask = append(arrTask, task)
		}

		taskDate.Tanggal = tanggal
		taskDate.Task = arrTask

		arrTaskDate = append(arrTaskDate, taskDate)

		arrTask = []Task{}
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = arrTaskDate

	return res, nil
}

func GetTask(id string) (Response, error) {
	var res Response
	var task Task

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tasks WHERE id = ($1)"

	err := con.QueryRow(sqlStatement, id).Scan(&task.Id, &task.User_Id, &task.Title, &task.Tanggal, &task.Waktu, &task.Date, &task.Deskripsi)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Pesan = "Sukses"
	res.Data = task

	return res, nil
}

func DeleteTask(id string, userId int) (Response, error) {
	var res Response
	var task Task

	con := db.CreateCon()

	sqlStatement := "DELETE FROM tasks WHERE id = ($1)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return res, err
	}

	task.Id, _ = strconv.Atoi(id)
	task.User_Id = userId

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Hapus Data"
	res.Data = task

	return res, err
}

func UpdateTask(id string, userId int, title string, tanggal string, waktu string, date string, deskripsi string) (Response, error) {
	var task Task
	var err error
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE tasks set title = ($1), tanggal = ($2), waktu = ($3), date = ($4), deskripsi = ($5) WHERE id = ($6)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(title, tanggal, waktu, date, deskripsi, id)
	if err != nil {
		return res, err
	}

	task.Id, _ = strconv.Atoi(id)
	task.User_Id = userId
	task.Title = title
	task.Tanggal = tanggal
	task.Waktu = waktu
	task.Date = date
	task.Deskripsi = deskripsi

	res.Status = http.StatusOK
	res.Pesan = "Berhasil Ubah Data"
	res.Data = task

	return res, err
}
