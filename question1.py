import itertools

def get_2020_multi(items):
    for x in items:
        for y in items:
            for z in items:
                if x + y + z == 2020:
                    return x * y * z
    return 0

# Read integers from the file
lines = []
with open('input.txt') as f:
    lines = f.read().splitlines()
ints = [int(i) for i in lines]

print(get_2020_multi(ints))
