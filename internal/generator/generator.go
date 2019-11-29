package generator

type generator struct{}

// ? iota?
const (
	MYSQL_VARCHAR = iota + 1
	MYSQL_INT
)

func dummy(typ string) interface{} {
	switch typ {
	case "a":
		{

		}
	case "b":
		{

		}
	case "c":
		{

		}
	default:
		println("unsupported.")
	}

	return nil
}
