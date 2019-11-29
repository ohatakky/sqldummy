package generator

import (
	"log"

	"github.com/ohatakky/sqldummy/internal/fields"
)

type generator struct{}

func dummy(f fields.FieldType) interface{} {
	switch f {
	case fields.FieldTypeDecimal:
		{

		}
	case fields.FieldTypeDate:
		{

		}
	case fields.FieldTypeTiny:
		{

		}
	case fields.FieldTypeVarChar:
		{

		}
	default:
		log.Println("unsupported.")
	}

	return nil
}
