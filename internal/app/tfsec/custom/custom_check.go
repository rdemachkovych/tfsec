package custom

import (
	"github.com/tfsec/tfsec/internal/app/tfsec/scanner"
)

type MatchType string
type CheckAction string

var ValidCheckActions = []CheckAction{
	InModule, IsPresent, NotPresent, StartsWith, EndsWith, Contains, Equals, RegexMatches, IsAny, IsNone,
}

// InModule checks that the block is part of a module
const InModule CheckAction = "inModule"

// IsPresent checks that the named child is present in the block
const IsPresent CheckAction = "isPresent"

// NotPresent checks that the named child is absent in the block
const NotPresent CheckAction = "notPresent"

// StartsWith checks that the named child attribute has a value that starts with the check value
const StartsWith CheckAction = "startsWith"

// EndsWith checks that the named child attribute has a value that ends with the check value
const EndsWith CheckAction = "endsWith"

// Contains checks that the named child attribute has a value in the map, list or attribute
const Contains CheckAction = "contains"

// Contains checks that the named child attribute has a value equal to the check value
const Equals CheckAction = "equals"

// RegexMatches checks that the named attribute has a value that matches the regex
const RegexMatches CheckAction = "regexMatches"

// IsAny checks that the named attribute value can be found in the provided slice
const IsAny CheckAction = "isAny"

// IsNone checks that the named attribute value cannot be found in the provided slice
const IsNone CheckAction = "isNone"

// MatchSpec specifies the checks that should be performed
type MatchSpec struct {
	Name       string      `json:"name,omitempty"`
	MatchValue interface{} `json:"value,omitempty"`
	Action     CheckAction `json:"action,omitempty"`
	SubMatch   *MatchSpec  `json:"subMatch,omitempty"`
}

// CustomCheck specifies the check definition represented in json
type Check struct {
	Code           scanner.RuleCode    `json:"code"`
	Description    scanner.RuleSummary `json:"description"`
	RequiredTypes  []string            `json:"requiredTypes"`
	RequiredLabels []string            `json:"requiredLabels"`
	Severity       scanner.Severity    `json:"severity"`
	ErrorMessage   string              `json:"errorMessage,omitempty"`
	MatchSpec      *MatchSpec          `json:"matchSpec"`
	RelatedLinks   []string            `json:"relatedLinks,omitempty"`
}

func (action *CheckAction) isValid() bool {
	for _, checkAction := range ValidCheckActions {
		if checkAction == *action {
			return true
		}
	}
	return false
}
