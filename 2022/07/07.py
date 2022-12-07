from __future__ import annotations


class File:
    def __init__(self, name: str, size: int) -> None:
        self.name = name
        self.size = size


class Dir:
    def __init__(self, name: str, parent: Dir | None):
        self.name = name
        self.parent = parent
        self.children: dict[str, Dir] = {}
        self.files: list[File] = []

    def size(self):
        return sum(file.size for file in self.files) + sum(
            child.size() for _, child in self.children.items()
        )


class Shell:
    def __init__(self, data: list):
        self.cwd: Dir = Dir("/", None)
        self.input = data
        self.current_line = 1

    def parse(self):
        while self.current_line < len(self.input):
            cmd, args = self.parse_command(self.input[self.current_line][1:])
            if cmd == "cd":
                self.change_dir(args[0])
            self.advance()
            self.parse_output()

    def parse_command(self, line: str) -> tuple[str, list]:
        cmd, *args = line.split()
        return cmd, args

    def parse_output(self):
        while self.current_line < len(self.input) and self.current()[0] != "$":
            line = self.current().split()
            if line[0] == "dir":
                self.add_child(line[1])
            else:
                self.add_file(int(line[0]), line[1])
            self.advance()

    def change_dir(self, name: str):
        if name == "..":
            self.cwd = self.cwd.parent if self.cwd.parent is not None else self.cwd
        else:
            if name in self.cwd.children:
                self.cwd = self.cwd.children[name]
            else:
                self.cwd = Dir(name, self.cwd)

    def add_child(self, name: str):
        if name not in self.cwd.children:
            self.cwd.children[name] = Dir(name, self.cwd)

    def add_file(self, size: int, name: str):
        self.cwd.files.append(File(name, size))

    def advance(self):
        self.current_line += 1

    def current(self):
        return self.input[self.current_line]

    def get_root(self) -> Dir:
        root = self.cwd
        while root.parent is not None:
            root = root.parent
        return root


def part1(dir, total):
    size = dir.size()

    if size <= 100000:
        total += size

    for _, child in dir.children.items():
        total = part1(child, total)

    return total


data = open("input").read().splitlines()
shell = Shell(data)
shell.parse()

print("part 1: ", part1(shell.get_root(), 0))


def part2(dir, space_needed, possible_solutions):
    size = dir.size()

    avaialable_space = space_needed + size

    if avaialable_space >= 30000000:
        possible_solutions.append(size)

    for _, child in dir.children.items():
        possible_solutions = part2(child, space_needed, possible_solutions)

    return possible_solutions


root = shell.get_root()
print("part 2: ", min(part2(root, 70000000 - root.size(), [])))
