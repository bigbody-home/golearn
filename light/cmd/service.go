package cmd

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func getDb() *sql.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", Username, Secret, Host, Port, Db)
	db1, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db1
}

func ParseCsvToDb(path string, table string) {
	var wg sync.WaitGroup
	p, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range p {
		abspath := path + "/" + f.Name()
		wg.Add(1)
		go parse(abspath, &wg, table)
	}
	wg.Wait()
}

func parse(name string, wg *sync.WaitGroup, table string) {
	f, err := os.Open(name)
	defer func() {
		wg.Done()
		f.Close()
	}()
	if err != nil {
		log.Fatalln(err)
		return
	}
	reader := csv.NewReader(bufio.NewReader(f))
	line0, err := reader.Read()
	
	sql := getSql(line0, table)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
			return
		}
		tline := getExec(line)
		sql += tline
	}
	db := getDb()
	sql = sql[0 : len(sql)-1]
	res, err := db.Exec(sql)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.LastInsertId())
}

func getSql(title []string, table string) string {
	str := strings.Join(title, ",")
	res := "insert into " + table + "(" + str + ") values"
	return res
}

func getExec(csv []string) string {
	var newcsv []string
	for i := 0; i < len(csv); i++ {
		tmp := "\"" + csv[i] + "\""
		newcsv = append(newcsv, tmp)
	}
	str := strings.Join(newcsv, ",")
	return "(" + str + "),"
}
