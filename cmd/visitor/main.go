package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/company"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/visitor"
)

func main() {
	auditor := visitor.NewVisitor()
	for _, auditedCompany := range [...]company.Company{
		company.NewSteelMill(),
		company.NewChemicalFactory(),
		company.NewCarFactory(),
	} {
		auditedCompany.Accept(auditor)
	}
}
