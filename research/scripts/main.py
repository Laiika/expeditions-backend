import random

LOCATIONS = 100
EXPEDITIONS = 10000

file = open(str(EXPEDITIONS) + '.sql', 'w')

# locations

file.write("insert into locations(name, country, nearest_town) values \n")

countries = ["Russia", "Kyrgyzstan", "Iceland", "Latvia", "Norway"]

for i in range(1, LOCATIONS + 1):
    file.write('(')

    file.write('\'')
    file.write("location" + str(i))
    file.write('\',')

    file.write('\'')
    file.write(countries[i % len(countries)])
    file.write('\',')

    file.write('\'')
    file.write("town" + str(i))
    file.write('\'')

    if i != LOCATIONS:
        file.write("),\n")
    else:
        file.write(");\n\n")


#expeditions

file.write("insert into expeditions(location_id, start_date, end_date) values \n")

for i in range(1, EXPEDITIONS + 1):
    file.write('(')

    file.write('\'')
    file.write(str(random.randint(1, LOCATIONS)))
    file.write('\',')

    file.write('\'')
    file.write("2024-07-07")
    file.write('\',')

    file.write('\'')
    file.write("2024-08-07")
    file.write('\'')

    if i != EXPEDITIONS:
        file.write("),\n")
    else:
        file.write(");\n\n")
