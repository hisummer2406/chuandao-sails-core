package state

import "chuandao-sails-core/apps/order-service/pkg/constants"

// StateValidator 状态转换规则
type StateValidator struct {
	transitions map[string][]string
}

func NewStateValidator() *StateValidator {
	return &StateValidator{
		transitions: constants.StateTransitions,
	}
}

func (v *StateValidator) IsValidTransition(from, to string) bool {
	allowed, exists := v.transitions[from]
	if !exists {
		return false
	}

	for _, next := range allowed {
		if next == to {
			return true
		}
	}
	return false
}
