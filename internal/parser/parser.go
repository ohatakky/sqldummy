package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/ohatakky/sqldummy/adapter/mysql"
	"github.com/ohatakky/sqldummy/internal/fields"
)

type parser struct {
	Mysql *mysql.Mysql
}

type Parser interface {
	DumpTableDDL(table string) (string, error)
	ParseDDL(string) ([]*fields.Column, error)
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

func (p *parser) ParseDDL(ddl string) ([]*fields.Column, error) {
	columns := make([]*fields.Column, 0)
	lines := strings.Split(ddl, "\n")
	for idx, line := range lines {
		if idx == 0 {
			if line[0:6] != "CREATE" {
				return nil, errors.New("parse DDL failed")
			}
		}
		if idx == len(lines)-1 {
			continue
		}
		if line[0:3] != "  `" {
			continue
		}
		columns = append(columns, parseColumn(line))
	}

	return columns, nil
}

func parseClosure(str string) string {
	r := regexp.MustCompile("`.*`")
	find := r.FindStringSubmatch(str)
	return strings.Trim(find[0], "`")
}

func parseColumn(str string) *fields.Column {
	return &fields.Column{
		Name: parseClosure(str),
		Type: parseType(str),
	}
}

func parseType(str string) string {
	r := regexp.MustCompile("` [^A-Z]+ ")
	find := r.FindStringSubmatch(str)
	return strings.TrimSpace(find[0][2:])
}
