fn part1(input: &Vec<(&str, i32)>) -> i32 {
    let (mut depth, mut horizontal) = (0, 0);
    for line in input {
        match line.0 {
            "forward" => horizontal += line.1,
            "up" => depth -= line.1,
            "down" => depth += line.1,
            _ => println!("Invalid command!"),
        }
    }
    depth * horizontal
}

fn part2(input: &Vec<(&str, i32)>) -> i32 {
    let (mut depth, mut horizontal, mut aim) = (0, 0, 0);
    for line in input {
        match line.0 {
            "forward" => {
                horizontal += line.1;
                depth += aim * line.1
            }
            "up" => aim -= line.1,
            "down" => aim += line.1,
            _ => println!("Invalid command!"),
        }
    }
    depth * horizontal
}

fn main() {
    let input: Vec<(&str, i32)> = include_str!("../input")
        .lines()
        .map(|line| {
            let line: Vec<&str> = line.split(" ").collect();
            (line[0], line[1].parse().unwrap())
        })
        .collect();

    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}
