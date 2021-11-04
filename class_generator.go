package main

import (
	"context"
	"fmt"
	. "github.com/dave/jennifer/jen"
	"github.com/machinebox/graphql"
)

const BaseURL = "https://www.warcraftlogs.com/api/v2/client/"

func generateClass(pkgName string) (*File, error) {
	graphqlClient := graphql.NewClient(BaseURL)

	graphqlRequest := graphql.NewRequest(`
{
 gameData {
     classes {
		 id
         name
         specs {
             name
         }
     }
 }
}
	`)

	var bearer = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI5NDQ3MGJiMi1jNTYwLTRkN2ItYjUzMC0wNjgyYzhkNjc3OWYiLCJqdGkiOiI1NGMwOTE5ZmVjYWI4ZWUzOWM1MDY1NmZmY2M3MTVhMGU0ZjU3YzEwNTIwMzczNjRhZGZlMmE1ZjUxODgwMzkwZTA2NmFiOGZkZTQ5ZDA2ZCIsImlhdCI6MTYzMDQ5MTcxMiwibmJmIjoxNjMwNDkxNzEyLCJleHAiOjE2NDA4NTk3MTIsInN1YiI6IiIsInNjb3BlcyI6WyJ2aWV3LXVzZXItcHJvZmlsZSIsInZpZXctcHJpdmF0ZS1yZXBvcnRzIl19.EUwgchJjCVlrea530TWLPvTcAQeLkCcafOmpcn6iSjdOHUjA9sSt-ldmQ7Hspk6Nz8WlMirGaXckXoy5u_gLL1Y2D2KQ1PFe5JHEvK5OeWoKXoBk30_jy5rsaQJJdWZaqqwTse86hmUGVIlNA6Zt3oQbhr2_hPGjQiRz-dCJ9ty7209tC0d5lU8qIbIjvFm3OiYxbc6NLqliEt9kNU4g3Zy2MBeCCPDzGYXQ8cIhEFe_CCLfA03VYGzIrftXPVq0xED73XYsoQVsOkY0_I_wNXlMwhoyJAjvGwFHt5uXT5P-qGI62MvnoMWOvE8fcRC3doEorMlTSOK5gWRLHC_04fuU78fjKyzQrajhTxroBBEN5TvagRvJamXBSy8VkwHoiNS9a0b-o352VDvcNfTQRqkY_5RMrI-WbHIpJQwdk906WDjy59WGibYcyXGJvP4mGAtENFyf10uo5dXwmz6dDaW6HbRTYn7mkEoIG_S-J9imMG_JCveCialQ1tL1ztE8K8EpeNKACyq7-p1wnPvAb-Y1BPaqlLO_HP11YR1hY8BdeVGtHHMgBFGNQ3Dk8ak_kWfVeve_8yIwJPT8UDKWYDA4Q6fNDTORLCp_R9eO1e-F8DmGsJq5PhqoLhB4cAqSDwW9LWtCa2PMY1Ac7RQ5gXKN6vmSi8UhHvqj11oo39I"

	// set header fields
	graphqlRequest.Header.Set("Authorization", bearer)

	var data Data
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &data); err != nil {
		return nil, err
	}

	var DPSSpecs []string
	var healSpecs []string
	for _, class := range data.GameData.Classes {
		for _, spec := range class.Specs {
			classSpecName := fmt.Sprintf("%s-%s", class.Name, spec.Name)
			if Contains(HealAssociation, classSpecName) {
				healSpecs = append(healSpecs, classSpecName)
			} else {
				DPSSpecs = append(DPSSpecs, classSpecName)
			}
		}
	}

	return generateCode(pkgName, data.GameData, DPSSpecs, healSpecs), nil
}

func generateCode(packageName string, gameData GameData, DPSspecs []string, healSpecs []string) *File {
	f := NewFile(packageName)

	f.Comment("Generated types to construct dictionnaries")
	f.Type().Id("ClassID").Id("int64")
	f.Type().Id("PointCounter").Id("uint32")
	f.Type().Id("ClassName").Id("string")
	f.Type().Id("SpecName").Id("string")

	f.Comment("Class associated with its id")
	f.
		Var().Id("Class").
		Op("=").Map(Qual("", "ClassID")).Qual("", "ClassName").
		Values(DictFunc(func(d Dict) {
			for _, class := range gameData.Classes {
				d[Lit(class.ID)] = Lit(class.Name)
			}
		}))

	f.Comment("DPSPointCounter: Dps Class (per spec) associated with its stat numbers")
	f.
		Var().Id("DPSPointCounter").
		Op("=").Map(Qual("", "SpecName")).Qual("", "PointCounter").
		Values(DictFunc(func(d Dict) {
			for _, spec := range DPSspecs {
				d[Lit(spec)] = Lit(0)
			}
		}))

	f.Comment("HealPointCounter: Healing Class (per spec) associated with its stat numbers")
	f.
		Var().Id("HealPointCounter").
		Op("=").Map(Qual("", "SpecName")).Qual("", "PointCounter").
		Values(DictFunc(func(d Dict) {
			for _, spec := range healSpecs {
				d[Lit(spec)] = Lit(0)
			}
		}))

	return f
}
