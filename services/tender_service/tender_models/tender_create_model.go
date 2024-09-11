package tender_models

import "avitoTest/shared/constants"

type TenderCreateModel struct {
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	ServiceType    string                 `json:"service_type"`
	OrganizationID int                    `json:"organization_id"`
	CreatorID      int                    `json:"creator_id"`
	Status         constants.TenderStatus `json:"status"`
}
