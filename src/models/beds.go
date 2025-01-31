package models

type Beds struct {
	BedTypeId  string `json:"bed_type_id" bson:"bed_type_id" binding:"required"`
	BedType    string `json:"bed_type" bson:"bed_type" binding:"required"`
	T_Capacity string `json:"t_capacity" bson:"t_capacity" binding:"required"`
	Available  string `json:"available" bson:"available" binding:"required"`
	Occupied   string `json:"occupied" bson:"occupied" binding:"required"`
}
