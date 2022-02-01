package client

import "github.com/NETWAYS/go-check"

type StateOverride map[string]int

func LeveltoState(level string) (state int) {
	// https://docs.microsoft.com/en-us/graph/api/resources/servicehealthissue?view=graph-rest-1.0#servicehealthstatus-values
	switch level {
	case "SERVICEOPERATIONAL":
		return check.OK
	case "INVESTIGATING":
		return check.Critical
	case "RESTORINGSERVICE":
		return check.Warning
	case "VERIFYINGSERVICE":
		return check.Warning
	case "SERVICERESTORED":
		return check.OK
	case "POSTINCIDENTREVIEWPUBLISHED":
		return check.Critical
	case "SERVICEDEGRADATION":
		return check.Critical
	case "SERVICEINTERRUPTION":
		return check.Critical
	case "EXTENDEDRECOVERY":
		return check.Critical
	case "FALSEPOSITIVE":
		return check.OK
	case "INVESTIGATIONSUSPENDED":
		return check.Critical
	case "RESOLVED":
		return check.OK
	case "MITIGATEDEXTERNAL":
		return check.OK
	case "MITIGATED":
		return check.OK
	case "RESOLVEDEXTERNAL":
		return check.OK
	case "CONFIRMED":
		return check.OK
	case "REPORTED":
		return check.OK
	case "UNKNOWNFUTUREVALUE":
		return check.Critical
	default:
		return check.Unknown
	}
}
