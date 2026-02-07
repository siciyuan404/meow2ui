package policy

type Decision struct {
	Action           string
	RuleID           string
	Reason           string
	RecommendedModel string
}

func Evaluate(currentSpent float64, budget float64, predictedCost float64) Decision {
	if budget <= 0 {
		return Decision{Action: "allow", RuleID: "COST-ALLOW-NO-BUDGET", Reason: "budget disabled"}
	}
	if currentSpent+predictedCost > budget {
		return Decision{Action: "block", RuleID: "COST-BLOCK-OVER-BUDGET", Reason: "budget exceeded"}
	}
	if currentSpent+predictedCost > budget*0.9 {
		return Decision{Action: "degrade_model", RuleID: "COST-DEGRADE-NEAR-BUDGET", Reason: "near budget", RecommendedModel: "low-cost"}
	}
	return Decision{Action: "allow", RuleID: "COST-ALLOW", Reason: "within budget"}
}
