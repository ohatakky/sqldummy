package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/ohatakky/sqldummy/adapter/mysql"
)

type parser struct {
	Mysql *mysql.Mysql
}

type TypeMapper struct {
	table   string
	columns []map[string]string
}

type Parser interface {
	DumpTableDDL(table string) (string, error)
	ParseDDL(string) (*TypeMapper, error)
}

func New(mysql *mysql.Mysql) Parser {
	return &parser{
		Mysql: mysql,
	}
}

func (p *parser) DumpTableDDL(table string) (string, error) {
	var ddl string
	sql := fmt.Sprintf("SHOW CREATE TABLE `%s`;", table)

	err := p.Mysql.DB.QueryRow(sql).Scan(&table, &ddl)
	if err != nil {
		return "", err
	}

	return ddl, nil
}

func (p *parser) ParseDDL(ddl string) (*TypeMapper, error) {
	res := &TypeMapper{}
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
	find := r.FindStringSubmatch(str)
	return strings.Trim(find[0], "`")
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
