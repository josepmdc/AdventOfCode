rucksacks = open("input").read().splitlines()

def get_priority(char: str):
    if char.islower():
        return ord(char) % 96
    else:
        return ord(char) % 64 + 26

# part 1 #
priority_sum = sum(
    get_priority("".join(set(rucksack[:len(rucksack)//2]).intersection(rucksack[len(rucksack)//2:]))) 
    for rucksack in rucksacks
)

print("part 1: ", priority_sum)

# part 2 #
groups = [rucksacks[i:i+3] for i in range(0, len(rucksacks), 3)]

priority_sum = sum(
    get_priority("".join(set(group[0]).intersection(group[1]).intersection(group[2]))) 
    for group in groups
)

print("part 2: ", priority_sum)
