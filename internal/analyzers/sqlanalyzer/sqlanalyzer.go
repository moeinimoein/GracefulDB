package sqlanalyzer

import (
	"strings"

	"github.com/Kwynto/GracefulDB/internal/engine/core"
)

type tQuery struct {
	Ticket      string
	Instruction string
	Placeholder []string
}

// TODO: Request
func Request(ticket string, instruction string, placeholder []string) string {
	// -
	// Prep
	instruction = strings.TrimSpace(instruction)
	instruction = strings.TrimRight(instruction, "; ")
	// instruction = strings.TrimSpace(instruction)

	var query tQuery = tQuery{
		Ticket:      ticket,
		Instruction: instruction,
		Placeholder: placeholder,
	}

	for _, expName := range core.ParsingOrder {
		if core.RegExpCollection[expName].MatchString(query.Instruction) {
			switch expName {
			case "SearchCreate":
				res, _ := query.DDLCreate()
				return res
			case "SearchAlter":
				res, _ := query.DDLAlter()
				return res
			case "SearchDrop":
				res, _ := query.DDLDrop()
				return res
			case "SearchSelect":
				res, _ := query.DMLSelect()
				return res
			case "SearchInsert":
				res, _ := query.DMLInsert()
				return res
			case "SearchUpdate":
				res, _ := query.DMLUpdate()
				return res
			case "SearchDelete":
				res, _ := query.DMLDelete()
				return res
			case "SearchTruncate":
				res, _ := query.DMLTruncate()
				return res
			case "SearchCommit":
				res, _ := query.DMLCommit()
				return res
			case "SearchRollback":
				res, _ := query.DMLRollback()
				return res
			case "SearchUse":
				res, _ := query.DCLUse()
				return res
			case "SearchGrant":
				res, _ := query.DCLGrant()
				return res
			case "SearchRevoke":
				res, _ := query.DCLRevoke()
				return res
			case "SearchAuth":
				res, _ := query.DCLAuth()
				return res
			}
		}
	}

	res := "Unknown command"
	return res
}
