package fields

type Column struct {
	Name string
	Type string
}

type FieldType byte

const (
	FieldTypeDecimal FieldType = iota
	FieldTypeTiny
	FieldTypeShort
	FieldTypeLong
	FieldTypeFloat
	FieldTypeDouble
	FieldTypeNULL
	FieldTypeTimestamp
	FieldTypeLongLong
	FieldTypeInt24
	FieldTypeDate
	FieldTypeTime
	FieldTypeDateTime
	FieldTypeYear
	FieldTypeNewDate
	FieldTypeVarChar
	FieldTypeBit
	FieldUnsupported
)

func (c *Column) convertFieldType() FieldType {
	switch c.Type {
	case "date":
		return FieldTypeDate
	case "decimal":
		return FieldTypeDecimal
	case "tinyint":
		return FieldTypeTiny
	case "varchar":
		return FieldTypeVarChar
	case "varbinary":
		return FieldTypeVarChar
	//  case "bit":
	//    return FieldTypeBit
	// 	case FieldTypeDateTime:
	// 		return "DATETIME"
	// 	case FieldTypeDouble:
	// 		return "DOUBLE"
	// 	case FieldTypeEnum:
	// 		return "ENUM"
	// 	case FieldTypeFloat:
	// 		return "FLOAT"
	// 	case FieldTypeGeometry:
	// 		return "GEOMETRY"
	// 	case FieldTypeInt24:
	// 		return "MEDIUMINT"
	// 	case FieldTypeJSON:
	// 		return "JSON"
	// 	case FieldTypeLong:
	// 		return "INT"
	// 	case FieldTypeLongBLOB:
	// 		if mf.charSet != collations[binaryCollation] {
	// 			return "LONGTEXT"
	// 		}
	// 		return "LONGBLOB"
	// 	case FieldTypeLongLong:
	// 		return "BIGINT"
	// 	case FieldTypeMediumBLOB:
	// 		if mf.charSet != collations[binaryCollation] {
	// 			return "MEDIUMTEXT"
	// 		}
	// 		return "MEDIUMBLOB"
	// 	case FieldTypeNewDate:
	// 		return "DATE"
	// 	case FieldTypeNewDecimal:
	// 		return "DECIMAL"
	// 	case FieldTypeNULL:
	// 		return "NULL"
	// 	case FieldTypeSet:
	// 		return "SET"
	// 	case FieldTypeShort:
	// 		return "SMALLINT"
	// 	case FieldTypeString:
	// 		if mf.charSet == collations[binaryCollation] {
	// 			return "BINARY"
	// 		}
	// 		return "CHAR"
	// 	case FieldTypeTime:
	// 		return "TIME"
	// 	case FieldTypeTimestamp:
	// 		return "TIMESTAMP"
	// 	case FieldTypeTinyBLOB:
	// 		if mf.charSet != collations[binaryCollation] {
	// 			return "TINYTEXT"
	// 		}
	// 		return "TINYBLOB"
	// 	case FieldTypeVarString:
	// 		if mf.charSet == collations[binaryCollation] {
	// 			return "VARBINARY"
	// 		}
	// 		return "VARCHAR"
	// 	case FieldTypeYear:
	// 		return "YEAR"
	default:
		return FieldUnsupported
	}
}
