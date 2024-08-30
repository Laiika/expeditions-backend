import matplotlib.pyplot as plot

file = open("../../result.txt", "r")


size = []
timeInd = []
errorInd = []
timeNo = []
errorNo = []

while True:
    line = file.readline()

    if not line:
        break

    numbers = [int(x) for x in line.split()]
    size.append(numbers[0])
    timeInd.append(numbers[1])
    errorInd.append(numbers[2])
    timeNo.append(numbers[3])
    errorNo.append(numbers[4])

for i in range(len(timeInd)):
    # timeTr[i] = timeTr[i] / (500 - errorTr[i])
    timeInd[i] = (timeInd[i] / 100) / 1000000

for i in range(len(timeNo)):
    # timeApp[i] = timeApp[i] / (500 - errorApp[i])
    timeNo[i] = (timeNo[i] / 100) / 1000000


plot.ylabel("Время, мс")
plot.xlabel("Количество записей в таблице")
plot.grid(True)

plot.plot(size, timeInd, color="green", label='Есть индекс', marker='*')
plot.plot(size, timeNo, color="blue", label='Нет индекса', marker='*')
plot.legend(["Есть индекс", "Нет индекса"])

plot.savefig('resultGraph.pdf')