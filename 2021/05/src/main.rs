use std::cmp;
use std::collections::HashMap;

#[derive(Debug, Clone, Copy)]
struct Coordinate {
    x: i32,
    y: i32,
}

#[derive(Debug, Clone, Copy)]
struct Vent {
    start: Coordinate,
    end: Coordinate,
}

fn part1(input: &Vec<Vent>) -> i32 {
    let mut map = HashMap::new();
    let mut result = 0;

    for vent in input {
        if !(vent.start.x == vent.end.x || vent.start.y == vent.end.y) {
            continue; // Skip the diagonals
        }

        let start_x = cmp::min(vent.start.x, vent.end.x);
        let start_y = cmp::min(vent.start.y, vent.end.y);
        let end_x = cmp::max(vent.start.x, vent.end.x);
        let end_y = cmp::max(vent.start.y, vent.end.y);

        for i in start_x..end_x + 1 {
            for j in start_y..end_y + 1 {
                *map.entry((i, j)).or_insert(0) += 1;

                if map[&(i, j)] == 2 {
                    result += 1;
                }
            }
        }
    }
    result
}

fn part2(input: &Vec<Vent>) -> i32 {
    let mut map = HashMap::new();
    let mut result = 0;

    for vent in input {
        let x_step = if vent.start.x > vent.end.x {
            -1
        } else if vent.start.x < vent.end.x {
            1
        } else {
            0
        };

        let y_step = if vent.start.y > vent.end.y {
            -1
        } else if vent.start.y < vent.end.y {
            1
        } else {
            0
        };

        let (mut x, mut y) = (vent.start.x, vent.start.y);
        while x != vent.end.x || y != vent.end.y {
            *map.entry((x, y)).or_insert(0) += 1;
            if map[&(x, y)] == 2 {
                result += 1;
            }
            x += x_step;
            y += y_step;
        }
        *map.entry((x, y)).or_insert(0) += 1;
        if map[&(x, y)] == 2 {
            result += 1;
        }
    }
    result
}

fn main() {
    let input: Vec<Vent> = include_str!("../input")
        .lines()
        .map(|line| {
            line.split(" -> ")
                .map(|pair| {
                    let num = pair.split_once(',').unwrap();
                    Coordinate {
                        x: num.0.parse().unwrap(),
                        y: num.1.parse().unwrap(),
                    }
                })
                .collect::<Vec<Coordinate>>()
        })
        .map(|pair| Vent {
            start: pair[0],
            end: pair[1],
        })
        .collect();

    println!("Part 1: {:?}", part1(&input));
    println!("Part 2: {:?}", part2(&input));
}
