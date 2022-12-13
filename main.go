package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"regexp"
	"strings"
	"time"

	dicparser "dic-morphy/app"
	db "dic-morphy/database"
)

func init() {
	for k, val := range db.NameTables {
		var filename string = "./assets/output/csv/" + k + ".csv"

		if _, err := os.Stat(filename); err == nil {
			err := os.Remove(filename)

			if err != nil {
				panic(err)
			} else {
				fmt.Println("Dont remove file '" + k + ".csv'!")
			}
		}

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)

		if err != nil {
			panic(err)
		} else {
			fmt.Println("File '" + k + ".csv' created!")
		}

		file.Write([]byte(strings.Join(val, ";") + "\n"))

		file.Close()
	}

	// // Подключение к базе данных
	// dbn := db.Connect(types.DBStruct{
	// 	Driver:   "mysql",
	// 	Host:     "127.0.0.1",
	// 	Port:     3306,
	// 	Username: "root",
	// 	Password: "",
	// 	DBname:   "morphy_ru",
	// })

	// defer dbn.Close()

	// fmt.Println("Connect to MySQL Server!")
}

func main() {
	startTime := time.Now()

	// Читаем файл словаря русского языка
	file, err := os.Open("./assets/ru/morphy-words.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Сканируем файл
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	// Обрабатываем файл для построчного чтения
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	// Объявляем переменные
	var normal string
	var idnormal int

	// Обявляем id'шники для таблиц
	tables := db.TableIds

	// Читаем каждую строку и обратаываем
	for i, eachline := range txtlines {
		tables["words"] = i

		if len(strings.TrimSpace(eachline)) > 0 {
			re, _ := regexp.Compile(`^ `)
			res := re.MatchString(eachline)

			if !res {
				normal = eachline
				idnormal = tables["words"]
			}

			word := dicparser.NewParser(tables["words"], idnormal, eachline, normal)

			word.Execute(&tables)

			fmt.Println("")
		}
	}

	endTime := time.Now()

	fmt.Println("Start time: ", startTime.Format("2006-01-02 15:04:05"))
	fmt.Println("End time: ", endTime.Format("2006-01-02 15:04:05"))
}
