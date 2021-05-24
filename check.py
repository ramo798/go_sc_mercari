import csv


def p_no(p):
    return p[p.find("[")+1:p.find("]")]


inv = []

ya = []

me = []

with open('Inventory_20210521.csv', "r", encoding="utf-8", errors="", newline="") as f:
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


for a in inv:
    for b in ya:
        if p_no(a[1]) == b[0]:
            if int(a[5]) == 0:
                print(b)
