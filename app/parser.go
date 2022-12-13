package dicparser

import (
	"fmt"
	"os"
	"strconv"

	// "log"
	"strings"

	// db "dic-morphy/database"
	types1 "dic-morphy/types/interface"
	types2 "dic-morphy/types/struct"

	_ "github.com/go-sql-driver/mysql"
)

type IParser types1.ParserInterface
type Parser types2.ParserStruct

// Constructor
func NewParser(id int, id_normal int, line string, normal string) IParser {
	if len(normal) > 1 {
		normal = strings.Split(strings.TrimSpace(normal), "|")[0]
	}

	var obj IParser = &Parser{
		Id:        id,
		Id_normal: id_normal,
		Line:      line,
		Normal:    normal,
	}

	return obj
}

func (ps *Parser) Execute(ids *map[string]int) {
	expl := strings.Split(ps.Line, "|")

	fmt.Println(expl[0])

	var filename string = "./assets/output/csv/words.csv"

	file, err := os.OpenFile(filename, os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}

	var joins = []string{strconv.Itoa(ps.Id), strconv.Itoa(ps.Id_normal), strings.TrimSpace(expl[0]), strings.TrimSpace(expl[2])}

	file.Write([]byte(strings.Join(joins, ";") + "\n"))
	file.Close()

	morphies := ps.MorphyBuild(expl[1])

	for k1 := range morphies {
		// fmt.Println(k1)
		var filename string = "./assets/output/csv/" + k1 + ".csv"

		if morphies[k1] != nil {
			for k2 := range morphies[k1] {
				(*ids)[k1]++
				mk := (*ids)[k1]
				// fmt.Println(morphies[k1][k2]["ru"])

				file, _ := os.OpenFile(filename, os.O_APPEND, 0644)

				var joins = []string{strconv.Itoa(mk), strconv.Itoa(ps.Id), morphies[k1][k2]["ru"]}

				file.Write([]byte(strings.Join(joins, ";") + "\n"))
				file.Close()
			}
		}

		// fmt.Println("")
	}

	// fmt.Println(morphies)
	// fmt.Println("")
}

func (ps *Parser) ExecuteOld() {
	// expl := strings.Split(ps.Line, "|")

	// fmt.Println(expl[0])

	// morphies := ps.MorphyBuild(expl[1])

	// Подготовленный запрос
	// var sql string = "SELECT `id` FROM `words` WHERE word = ? ORDER BY `id` DESC"

	// // Проверяем была ли добавлена нормальная форма словоформы
	// rows := db.Select(*ps.DB, sql, []any{ps.Normal}...)

	// defer rows.Close()

	// var dbnormal types2.WordStruct

	// // Пытаемся получить Id нормальной формы
	// for rows.Next() {
	// 	err := rows.Scan(&dbnormal.Id)

	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}

	// 	rows.Close()
	// }

	// Начало транзакции
	// tx, err := db.Begin(*ps.DB)

	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatalln("db.Begin", err)
	// }

	/*/ Добавление слова/словоформы
	stmt, err := tx.Prepare("INSERT INTO words (id_normal, word, detail) VALUES (?,?,?)")

	if err != nil {
		tx.Rollback()
		log.Fatalln("tx.Prepare", err)
	}

	defer stmt.Close() // danger!

	ins, err := stmt.Exec(dbnormal.Id, strings.TrimSpace(expl[0]), strings.TrimSpace(expl[2]))

	if err != nil {
		tx.Rollback()
		log.Fatalln("stmt.Exec", err)
	}

	lastInserId, err := ins.LastInsertId()

	if err != nil {
		tx.Rollback()
		log.Fatalln("stmt.LastInsertId", err)
	}

	fmt.Println(lastInserId)

	if lastInserId > 0 {
		for k := range morphies {
			for k2 := range morphies[k] {
				stmt2, _ := tx.Prepare("INSERT INTO " + k + " (id_word, name) VALUES (?,?,?)")

				_, err2 := stmt2.Exec(lastInserId, morphies[k][k2]["ru"])
				if err2 != nil {
					log.Fatalln("stmt.LastInsertId", err2)
				}
			}
		}
	}
	*/

	// Комит транзакции
	// err = tx.Commit()

	// if err != nil {
	// 	log.Fatalln("tx.Commit", err)
	// }
}

func (ps *Parser) MorphyBuild(str string) map[string][]map[string]string {
	data := map[string][]map[string]string{
		"parts_speech": {},
		"grammar":      {},
		"cases":        {},
	}

	expl := strings.Split(str, " ")

	// Части речи
	if ps.inAarrayStr("сущ", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "noun", "ru": "сущ"})
	}
	if ps.inAarrayStr("гл", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "verb", "ru": "глагол"})
	}
	if ps.inAarrayStr("прл", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "adjective", "ru": "прилагательное"})
	}
	if ps.inAarrayStr("прл", expl) && ps.inAarrayStr("крат", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "adjective short", "ru": "прилагательное краткое"})
	}
	if ps.inAarrayStr("мест", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "pronoun", "ru": "местоимение"})
	}
	if ps.inAarrayStr("числ", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "numeral", "ru": "числительное"})
	}
	if ps.inAarrayStr("нар", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "adverb", "ru": "наречие"})
	}
	if ps.inAarrayStr("предл", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "preposition", "ru": "предлог"})
	}
	if ps.inAarrayStr("союз", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "conjunction", "ru": "союз"})
	}
	if ps.inAarrayStr("част", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "particle", "ru": "частица"})
	}
	if ps.inAarrayStr("межд", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "interjection", "ru": "междометие"})
	}
	if ps.inAarrayStr("инф", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "infinitive", "ru": "инфинитив"})
	}
	if ps.inAarrayStr("прч", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "participle", "ru": "причастие"})
	}
	if ps.inAarrayStr("прч", expl) && ps.inAarrayStr("крат", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "participle short", "ru": "причастие краткое"})
	}
	if ps.inAarrayStr("дееп", expl) {
		data["parts_speech"] = append(data["parts_speech"], map[string]string{"en": "ad. participle", "ru": "деепричастие"})
	}

	// ============ Грамматика ============

	//Род имен существительных
	if ps.inAarrayStr("муж", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "male gender", "ru": "мужской род"})
	}
	if ps.inAarrayStr("жен", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "famale gender", "ru": "женский род"})
	}
	if ps.inAarrayStr("ср", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "neuter gender", "ru": "средний род"})
	}
	if ps.inAarrayStr("общ", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "all gender", "ru": "общий род"})
	}

	if ps.inAarrayStr("ед", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "singular", "ru": "единсвтенное число"})
	}
	if ps.inAarrayStr("мн", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "plural", "ru": "множественное число"})
	}
	if ps.inAarrayStr("прош", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "past tense", "ru": "прошедшее время"})
	}
	if ps.inAarrayStr("наст", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "present simple", "ru": "настоящее время"})
	}
	if ps.inAarrayStr("буд", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "future tense", "ru": "будущее время"})
	}

	// Другое
	if ps.inAarrayStr("одуш", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "animate", "ru": "одушевленный"})
	}
	if ps.inAarrayStr("неод", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "inanimate", "ru": "неодушевленный"})
	}
	if ps.inAarrayStr("непер", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "intransitive", "ru": "непереходный"})
	}
	if ps.inAarrayStr("перех", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "transitive", "ru": "переходный"})
	}
	if ps.inAarrayStr("сов", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "perfect", "ru": "совершенный"})
	}
	if ps.inAarrayStr("несов", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "imperfective", "ru": "несовершенный"})
	}

	if ps.inAarrayStr("1-е", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "1st declension", "ru": "1-е склонение"})
	}
	if ps.inAarrayStr("2-е", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "2st declension", "ru": "2-е склонение"})
	}
	if ps.inAarrayStr("3-е", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "3st declension", "ru": "3-е склонение"})
	}

	if ps.inAarrayStr("обст", expl) && ps.inAarrayStr("места", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "location", "ru": "обст. места"})
	}
	if ps.inAarrayStr("обст", expl) && ps.inAarrayStr("врем", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "time", "ru": "обст. времени"})
	}
	if ps.inAarrayStr("обст", expl) && ps.inAarrayStr("цель", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "goal", "ru": "обст. цели"})
	}
	if ps.inAarrayStr("обст", expl) && ps.inAarrayStr("причин", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "cause", "ru": "обст. причины"})
	}
	if ps.inAarrayStr("обст", expl) && ps.inAarrayStr("напр", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "direction", "ru": "обст. направления"})
	}

	if ps.inAarrayStr("страд", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "passive voice", "ru": "страдательный залог"})
	}
	if ps.inAarrayStr("кач", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "good", "ru": "качественное"})
	}
	if ps.inAarrayStr("опред", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "define", "ru": "определяемое"})
	}
	if ps.inAarrayStr("нескл", expl) {
		data["grammar"] = append(data["grammar"], map[string]string{"en": "unbending", "ru": "несклоняемое"})
	}

	// Падежи
	if ps.inAarrayStr("им", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "nominative", "ru": "именительный"})
	}
	if ps.inAarrayStr("род", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "genitive", "ru": "родительный"})
	}
	if ps.inAarrayStr("дат", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "dative", "ru": "дательный"})
	}
	if ps.inAarrayStr("вин", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "accusative", "ru": "винительный"})
	}
	if ps.inAarrayStr("тв", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "creative", "ru": "творительный"})
	}
	if ps.inAarrayStr("пр", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "prepositional", "ru": "предложный"})
	}
	if ps.inAarrayStr("парт", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "partitive", "ru": "партитивный"})
	}
	if ps.inAarrayStr("счет", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "countable", "ru": "счетный"})
	}
	if ps.inAarrayStr("мест", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "locative", "ru": "местный"})
	}
	if ps.inAarrayStr("зват", expl) {
		data["cases"] = append(data["cases"], map[string]string{"en": "vocative", "ru": "звательный"})
	}

	/* Не понятные сокращения */
	// пер/не
	// предик
	// прл

	return data
}

func (ps *Parser) inAarrayStr(needle string, arr []string) bool {
	for k := range arr {
		if arr[k] == needle {
			return true
		}
	}

	return false
}
