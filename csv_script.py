import csv


def p_no(p):
    return p[p.find("[")+1:p.find("]")]


inv = []

ya = []

me = []

with open('Inventory_20210524.csv', "r", encoding="utf-8", errors="", newline="") as f:
    reader = csv.reader(f)
    next(reader)
    for row in reader:
        inv.append(row)
with open('mercari.csv', "r", encoding="utf-8", errors="", newline="") as f:
    reader = csv.reader(f)
    next(reader)
    for row in reader:
        me.append(row)
with open('yahuoku.csv', "r", encoding="utf-8", errors="", newline="") as f:
    reader = csv.reader(f)
    next(reader)
    for row in reader:
        ya.append(row)

res = []
fl = ["在庫ID", "品番", "young", "tomokimi", "mercari"]

print(len(inv))

result = []
row = ["在庫ID", "品番", "younghoho_1121", "tomokimi_777",
       "maron", "younghoho_1121_URL", "tomokimi_777_URL", "maron_URL", "価格"]
result.append(row)

print(inv[0])
print(ya[0])

for a in inv:
    # print(p_no(a[1]))
    rowin = ["", "", "", "", "", "", "", "", ""]

    rowin[0] = a[0]
    rowin[1] = p_no(a[1])

    for i in ya:
        # print(i[0])
        if p_no(a[1]) == i[0]:
            # print(i[4])
            if i[4] == "younghoho_1121":
                rowin[2] = "1"
                rowin[5] = i[2]
                rowin[8] = i[6]
            if i[4] == "tomokimi_777":
                rowin[3] = "1"
                rowin[6] = i[2]
                rowin[8] = i[6]

    for j in me:
        if p_no(a[1]) == j[0]:
            rowin[4] = "1"
            rowin[7] = j[2]
            rowin[8] = i[6]

            print(j[2])

    # print(rowin)

    for k in range(len(rowin)):
        if rowin[k] == "":
            rowin[k] = "NaN"

    result.append(rowin)


with open('result.csv', "w", encoding="utf-8", errors="", newline="") as f:
    writer = csv.writer(f)
    for a in result:
        # print(a)
        writer.writerow(a)
