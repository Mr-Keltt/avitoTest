package validators

import "avitoTest/shared/constants"

var AllowedServiceTypes = []constants.ServiceType{
	constants.ServiceTypeConstruction,
	constants.ServiceTypeIT,
	constants.ServiceTypeConsulting,
}

func IsValidServiceType(serviceType string) bool {
	for _, st := range AllowedServiceTypes {
		if string(st) == serviceType {
			return true
		}
	}
	return false
}
