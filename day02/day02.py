import re
import sys

print('Loading file:', str(sys.argv[1]))

# Read integers from the file
lines = []
with open(str(sys.argv[1])) as f:
    lines = f.read().splitlines()

items = []
for line in lines:
    # group 1: min
    # group 2: max
    # group 3: character
    # group 4: password
    m = re.match(r'(\d+)-(\d+)\s(.*):\s(.*)', line)
    items.append((int(m.group(1)), int(m.group(2)), m.group(3), m.group(4)))

pass_count = 0
for item in items:
    min_val = item[0]
    max_val = item[1]
    char = item[2]
    password = item[3]

    first = False
    try:
        first = password[min_val-1] == char
    except:
        pass

    second = False
    try:
        second = password[max_val-1] == char
    except:
        pass

    if first ^ second:
        pass_count += 1

print(pass_count)
