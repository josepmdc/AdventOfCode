data = open("input").read().split("\n\n")
data = list(map(lambda x: x.strip("\n").split("\n"), data))

sums = [
    sum([int(calorie) for calorie in elf_calories])
    for elf_calories in data
]

print("part 1: ", max(sums))
print("part 2: ", sum(sorted(sums, reverse=True)[:3]))
