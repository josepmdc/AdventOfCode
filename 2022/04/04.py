data = open("input.test").read().splitlines()
data = [[[int(value) for value in elf.split('-')] for elf in pair.split(',')] for pair in data]

# part 1
def is_subset(pair):
    return ((pair[0][0] <= pair[1][0] and pair[0][1] >= pair[1][1]) or
            (pair[1][0] <= pair[0][0] and pair[1][1] >= pair[0][1]))

result = sum(1 if is_subset(pair) else 0 for pair in data)
print("part 1: ", result)

# part 2
result = sum(1 if pair[0][0] <= pair[1][1] and pair[1][0] <= pair[0][1] else 0 for pair in data)
print("part 1: ", result)
