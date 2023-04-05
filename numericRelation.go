package MovingOn

type NumericRelation int64

const (
	GreaterThan NumericRelation = iota
	LesserThan
	GreaterThanOrEqualTo
	LesserThanOrEqualTo
	EqualTo
	UndefinedComparison
	NotNumericComparison
)

func (n NumericRelation) String() string {
	switch n {
	case GreaterThan:
		return "GreaterThan"
	case LesserThan:
		return "LesserThan"
	case GreaterThanOrEqualTo:
		return "GreaterThanOrEqualTo"
	case LesserThanOrEqualTo:
		return "LesserThanOrEqualTo"
	case EqualTo:
		return "EqualTo"
	case UndefinedComparison:
		return "UndefinedComparison"
	case NotNumericComparison:
		return "NotNumericComparison"
	default:
		return "ERROR"
	}
}
