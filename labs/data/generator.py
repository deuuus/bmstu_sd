from mimesis import Person
from mimesis.enums import Gender
from random import randint

import datetime

dt = datetime.datetime(2010, 12, 1)
end = datetime.datetime(2010, 12, 10)
step = datetime.timedelta(days=1)

result = []

while dt < end:
    result.append(dt.strftime('%Y-%m-%d') + "," + str(randint(0, 9)) + "," + str(randint(0, 4)) + "\n")
    dt += step

f = open("records.txt", "w")
for r in result:
    f.write(r)
f.close()

gender = [Gender.FEMALE, Gender.MALE, "f", "m"]

person = Person('ru')

#patients = []
#doctors = []

#for i in range(10):
#    doctors.append(", ".join((person.name(gender[randint(0, 1)]) + " " + person.full_name(gender[randint(0, 1)])).split()) + "\n")

#for i in range(5):
#    patients.append(", ".join(person.full_name(gender[randint(0, 1)]).split()) + "\n")

#f = open("patients.txt", "w")
#for p in patients:
#    f.write(p)
#f.close()

#f = open("doctors.txt", "w")
#for d in doctors:
#    f.write(d)
#f.close()