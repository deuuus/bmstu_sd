COPY doctor(name, second_name, middle_name, specialization) 
FROM '/home/polina/sem_06/bmstu_sd/labs/data/doctors.txt' delimiter ',' csv NULL AS 'null';

COPY pacient(name, second_name) 
FROM '/home/polina/sem_06/bmstu_sd/labs/data/patients.txt' delimiter ',' csv NULL AS 'null';

COPY record(date, doctor_id, pacient_id, diagnosis) 
FROM '/home/polina/sem_06/bmstu_sd/labs/data/records.txt' delimiter ',' csv NULL AS 'null';