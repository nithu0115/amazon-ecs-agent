package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Dockerversion struct {
	major int
	minor int
}

// Matches returns whether or not a version matches a given selector.
// The selector can be any of the following:
//
// * x.y -- Matches a version exactly the same as the selector version
// * >=x.y -- Matches a version greater than or equal to the selector version
// * >x.y -- Matches a version greater than the selector version
// * <=x.y -- Matches a version less than or equal to the selector version
// * <x.y -- Matches a version less than the selector version
// * x.y,a.b -- Matches if the version matches either of the two selector versions
func (lhs Version) matches(selector string) (bool, error) {
	lhsVersion, err := parseVersion(string(lhs))
	if err != nil {
		return false, err
	}

	if strings.Contains(selector, ",") {
		orElements := strings.Split(selector, ",")
		for _, el := range orElements {
			if elMatches, err := lhs.matches(el); err != nil {
				return false, err
			} else if elMatches {
				return true, nil
			}
		}
		// No elements matched
		return false, nil
	}

	if strings.HasPrefix(selector, ">=") {
		rhsVersion, err := parseVersion(selector[2:])
		if err != nil {
			return false, err
		}
		return compareVersion(lhsVersion, rhsVersion) >= 0, nil
	} else if strings.HasPrefix(selector, ">") {
		rhsVersion, err := parseVersion(selector[1:])
		if err != nil {
			return false, err
		}
		return compareVersion(lhsVersion, rhsVersion) > 0, nil
	} else if strings.HasPrefix(selector, "<=") {
		rhsVersion, err := parseVersion(selector[2:])
		if err != nil {
			return false, err
		}
		return compareVersion(lhsVersion, rhsVersion) <= 0, nil
	} else if strings.HasPrefix(selector, "<") {
		rhsVersion, err := parseVersion(selector[1:])
		if err != nil {
			return false, err
		}
		return compareVersion(lhsVersion, rhsVersion) < 0, nil
	}

	rhsVersion, err := parseVersion(selector)
	if err != nil {
		return false, err
	}
	return compareVersion(lhsVersion, rhsVersion) == 0, nil
}

func parseVersion(version string) (Dockerversion, error) {
	var result Dockerversion
	versionParts := strings.Split(version, ".")
	// [0 0]
	if len(versionParts) != 2 {
		return result, fmt.Errorf("Not enough '.' characters in the version part")
	}
	major, err := strconv.Atoi(versionParts[0])
	if err != nil {
		return result, fmt.Errorf("Cannot parse major version as int: %v", err)
	}
	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return result, fmt.Errorf("Cannot parse minor version as int: %v", err)
	}
	result.major = major
	result.minor = minor
	return result, nil
}

// compareVersion compares two versions, 'lhs' and 'rhs', and returns -1 if lhs is less
// than rhs, 0 if they are equal, and 1 lhs is greater than rhs
func compareVersion(lhs, rhs Dockerversion) int {
	if lhs.major < rhs.major {
		return -1
	}
	if lhs.major > rhs.major {
		return 1
	}
	if lhs.minor < rhs.minor {
		return -1
	}
	if lhs.minor > rhs.minor {
		return 1
	}
	return 0
}
