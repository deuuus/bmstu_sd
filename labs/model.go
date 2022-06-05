package labsMed

type Patient struct {
	PatientId  int    `json:"p_id" db:"p_id" binding:"required"`
	Name       string `json:"name" db:"name" binding:"required"`
	SecondName string `json:"second_name" db:"second_name" binding:"required"`
}

type Doctor struct {
	DoctorId       int    `json:"d_id" db:"d_id" binding:"required"`
	Name           string `json:"name" db:"name" binding:"required"`
	SecondName     string `json:"second_name" db:"second_name" binding:"required"`
	MiddleName     string `json:"middle_name" db:"middle_name" binding:"required"`
	Specialization string `json:"specialization" db:"specialization" binding:"required"`
}

type Record struct {
	RecordId  int    `json:"r_id" db:"r_id" binding:"required"`
	Date      string `json:"date" db:"date"`
	DoctorId  int    `json:"doctor_id" db:"doctor_id"`
	PacientId int    `json:"pacient_id" db:"pacient_id"`
	Diagnosis string `json:"diagnosis" db:"diagnosis"`
}
