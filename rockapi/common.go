package rockapi


type AttributeValueDetail struct {
	AttributeId    int    `json:"AttributeId"`
	EntityId       int    `json:"EntityId"`
	Value          string `json:"Value"`
	ValueFormatted string `json:"ValueFormatted"`
}

const MemberRoleId = 118
const LeaderRoleId = 119
const locationId = "14"
const scheduleId = "2835"
const authToken = ""
const rockURL = ""



