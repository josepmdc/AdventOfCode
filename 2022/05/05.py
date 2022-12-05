import re
import copy

data = open("input").read().split("\n\n")

layout = data[0].splitlines()
commands = data[1].splitlines()

number_of_stacks = len(layout[-1].strip().split('   '))

original_stacks = {i: [] for i in range(1, number_of_stacks+1)}
for line in reversed(layout[:-1]):
    stack = 1
    for char in range(0, len(line), 4):
        if line[char] == '[':
            original_stacks[stack].append(line[char+1])
        stack += 1

commands = [[int(i) for i in re.split(" from | to ", command[5:])] for command in commands]

# part 1
stacks = copy.deepcopy(original_stacks)
for command in commands:
    stacks[command[2]] += reversed(stacks[command[1]][-command[0]:])
    del stacks[command[1]][-command[0]:]

print("part 1:", "".join(stack[-1] for _, stack in stacks.items()))

# part 2
stacks = copy.deepcopy(original_stacks)
for command in commands:
    stacks[command[2]] += stacks[command[1]][-command[0]:]
    del stacks[command[1]][-command[0]:]

print("part 2:", "".join(stack[-1] for _, stack in stacks.items()))
