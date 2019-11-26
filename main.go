package main

import (
	"errors"
	"log"
	"regexp"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/k0kubun/sqldef/adapter"
	"github.com/k0kubun/sqldef/adapter/mysql"
)

func main() {
	config := adapter.Config{
		DbName:   "",
		User:     "root",
		Password: "",
		Host:     "localhost",
		Port:     3306,
	}
	db, err := mysql.NewDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	currentDDLs, err := adapter.DumpDDLs(db)
	if err != nil {
		log.Fatalf("rror on DumpDDLs: %v", err)
	}

	ddls, err := parseDDLs(currentDDLs)
	if err != nil {
		log.Fatal(err)
	}

	pp.Println(ddls)
}

func parseDDLs(str string) ([]*DDL, error) {
	ddls := strings.Split(str, ";")
	res := []*DDL{}
	for _, ddl := range ddls {
		ddl = strings.TrimSpace(ddl)
		if len(ddl) == 0 {
			continue
		}
		println(ddl)

		parsed, err := parseDDL(ddl)
		if err != nil {
			return nil, err
		}
		res = append(res, parsed)
	}

	return res, nil
}

type DDL struct {
	table   string
	columns []map[string]string
}

func parseDDL(ddl string) (*DDL, error) {
	res := &DDL{}
	lines := strings.Split(ddl, "\n")
	for idx, line := range lines {
		if idx == 0 {
			if line[0:6] != "CREATE" {
				return nil, errors.New("parse DDL failed")
			}
			res.table = parseClosure(line)
		}
		if idx == len(lines)-1 {
			continue
		}
		if line[0:3] != "  `" {
			continue
		}
		column := parseColumn(line)
		res.columns = append(res.columns, column)
	}

	return res, nil
}

func parseClosure(str string) string {
	r := regexp.MustCompile("`.*`")
	find := r.FindAllStringSubmatch(str, -1)
	return find[0][0]
}

func parseColumn(str string) map[string]string {
	column := make(map[string]string, 1)
	column[parseClosure(str)] = parseType(str)
	return column
}

func parseType(str string) string {
	r := regexp.MustCompile("` [^A-Z]+ ")
	find := r.FindStringSubmatch(str)
	return strings.TrimSpace(find[0][2:])
}
