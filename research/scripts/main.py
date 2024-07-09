import random

LOCATIONS = 10
EXPEDITIONS = 1000
MEMBERS = 10

file = open(str(MEMBERS) + '.sql', 'w')

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


#members

file.write("insert into members(name, phone_number, login, password) values \n")

for i in range(1, MEMBERS + 1):
    file.write('(')

    file.write('\'')
    file.write("name" + str(i))
    file.write('\',')

    file.write('\'')
    file.write("phone_number" + str(i))
    file.write('\',')

    file.write('\'')
    file.write("login" + str(i))
    file.write('\',')

    file.write('\'')
    file.write("password" + str(i))
    file.write('\'')

    if i != MEMBERS:
        file.write("),\n")
    else:
        file.write(");\n\n")


# expeditions_members

file.write("insert into expeditions_members(expedition_id, member_id) values \n")

for i in range(1, EXPEDITIONS + 1):
    for j in range(0, 5):
        file.write('(')

        file.write('\'')
        file.write(str(i))
        file.write('\',')

        file.write('\'')
        file.write(str(random.randint(1, MEMBERS)))
        file.write('\'')

        if i != EXPEDITIONS or j != 4:
            file.write("),\n")
        else:
            file.write(");\n\n")
