package validator

import (
	"regexp"
	"s/dto"
)

func Check(tokens []string) []dto.IOut {
	Patterns :=
		map[string]string{
			"^[a-zA-Z_][a-zA-Z_0-9]*$":               "variable",
			"^[0-9]+$":                               "integer",
			`^[0-9]+\.[0-9]+$`:                       "float",
			`^".*"$`:                                 "string",
			"^[><=!][=]{0,1}$":                       "operator",
			"(^if$|^else$|^for$)":                    "keyword",
			"(^int$|^float$|^string$|^bool$|^char$)": "type",
			"^[(){}]{1}$":                            "bracket",
			"(^return$|^break$)":                     "control",
			"^[;,]{1}$":                              "separator",
		}

	var outty []dto.IOut

	for _, e := range tokens {
		for pat, typ := range Patterns {
			match, _ := regexp.MatchString(pat, e)

			if match {
				// fmt.Println(pat, typ, e)

				outty = append(outty, dto.IOut{Token: e, Typee: typ, Status: true})
				break
			}
			// fmt.Println(pat, typ, e)
			// outty = append(outty, IOut{token: e, typee: "", status: false})

		}
	}

	return outty
}
