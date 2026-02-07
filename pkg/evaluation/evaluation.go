package evaluation

type EvalScore struct {
	SchemaValid        bool   `json:"schema_valid"`
	ComponentValidRate int    `json:"component_valid_rate"`
	PropValidRate      int    `json:"prop_valid_rate"`
	RepairCount        int    `json:"repair_count"`
	Success            bool   `json:"success"`
	FailureType        string `json:"failure_type,omitempty"`
}

func Score(schemaValid bool, compRate int, propRate int, repairCount int, failureType string) EvalScore {
	success := schemaValid && compRate >= 90 && propRate >= 90
	return EvalScore{
		SchemaValid:        schemaValid,
		ComponentValidRate: compRate,
		PropValidRate:      propRate,
		RepairCount:        repairCount,
		Success:            success,
		FailureType:        failureType,
	}
}

func Regressed(current EvalScore, baseline EvalScore) bool {
	if baseline.Success && !current.Success {
		return true
	}
	if current.ComponentValidRate+3 < baseline.ComponentValidRate {
		return true
	}
	if current.PropValidRate+3 < baseline.PropValidRate {
		return true
	}
	return false
}
