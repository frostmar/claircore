// Code generated by "stringer -type=V3Metric,v3Valid -linecomment"; DO NOT EDIT.

package cvss

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V3AttackVector-0]
	_ = x[V3AttackComplexity-1]
	_ = x[V3PrivilegesRequired-2]
	_ = x[V3UserInteraction-3]
	_ = x[V3Scope-4]
	_ = x[V3Confidentiality-5]
	_ = x[V3Integrity-6]
	_ = x[V3Availability-7]
	_ = x[V3ExploitMaturity-8]
	_ = x[V3RemediationLevel-9]
	_ = x[V3ReportConfidence-10]
	_ = x[V3ConfidentialityRequirement-11]
	_ = x[V3IntegrityRequirement-12]
	_ = x[V3AvailabilityRequirement-13]
	_ = x[V3ModifiedAttackVector-14]
	_ = x[V3ModifiedAttackComplexity-15]
	_ = x[V3ModifiedPrivilegesRequired-16]
	_ = x[V3ModifiedUserInteraction-17]
	_ = x[V3ModifiedScope-18]
	_ = x[V3ModifiedConfidentiality-19]
	_ = x[V3ModifiedIntegrity-20]
	_ = x[V3ModifiedAvailability-21]
}

const _V3Metric_name = "AVACPRUISCIAERLRCCRIRARMAVMACMPRMUIMSMCMIMA"

var _V3Metric_index = [...]uint8{0, 2, 4, 6, 8, 9, 10, 11, 12, 13, 15, 17, 19, 21, 23, 26, 29, 32, 35, 37, 39, 41, 43}

func (i V3Metric) String() string {
	if i < 0 || i >= V3Metric(len(_V3Metric_index)-1) {
		return "V3Metric(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _V3Metric_name[_V3Metric_index[i]:_V3Metric_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[v3AttackVectorValid-0]
	_ = x[v3AttackComplexityValid-1]
	_ = x[v3PrivilegesRequiredValid-2]
	_ = x[v3UserInteractionValid-3]
	_ = x[v3ScopeValid-4]
	_ = x[v3ConfidentialityValid-5]
	_ = x[v3IntegrityValid-6]
	_ = x[v3AvailabilityValid-7]
	_ = x[v3ExploitMaturityValid-8]
	_ = x[v3RemediationLevelValid-9]
	_ = x[v3ReportConfidenceValid-10]
	_ = x[v3ConfidentialityRequirementValid-11]
	_ = x[v3IntegrityRequirementValid-12]
	_ = x[v3AvailabilityRequirementValid-13]
	_ = x[v3ModifiedAttackVectorValid-14]
	_ = x[v3ModifiedAttackComplexityValid-15]
	_ = x[v3ModifiedPrivilegesRequiredValid-16]
	_ = x[v3ModifiedUserInteractionValid-17]
	_ = x[v3ModifiedScopeValid-18]
	_ = x[v3ModifiedConfidentialityValid-19]
	_ = x[v3ModifiedIntegrityValid-20]
	_ = x[v3ModifiedAvailabilityValid-21]
}

const _v3Valid_name = "NALPLHNLHNRUCHLNHLNHLNXHFPUXUWTOXCRUXHMLXHMLXHMLXNALPXLHXNLHXNRXUCXHLNXHLNXHLN"

var _v3Valid_index = [...]uint8{0, 4, 6, 9, 11, 13, 16, 19, 22, 27, 32, 36, 40, 44, 48, 53, 56, 60, 63, 66, 70, 74, 78}

func (i v3Valid) String() string {
	if i < 0 || i >= v3Valid(len(_v3Valid_index)-1) {
		return "v3Valid(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _v3Valid_name[_v3Valid_index[i]:_v3Valid_index[i+1]]
}
