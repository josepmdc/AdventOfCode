
fn main() {
    let input: Vec<&str> = include_str!("../input").lines().collect();
    println!("Part 1: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &Vec<&str>) -> i32 {
    let mut gamma_binary = String::new();
    let mut epsilon_binary = String::new();

    for column_index in 0..input[0].len() {
        let (zeros, ones) = input.iter().fold((0, 0), |(zeros, ones), row| {
            if row.as_bytes()[column_index] as char == '0' {
                (zeros + 1, ones)
            } else {
                (zeros, ones + 1)
            }
        });

        if zeros > ones {
            gamma_binary.push('0');
            epsilon_binary.push('1');
        } else {
            gamma_binary.push('1');
            epsilon_binary.push('0');
        }
    }

    let gamma = i32::from_str_radix(&gamma_binary, 2).unwrap();
    let epsilon = i32::from_str_radix(&epsilon_binary, 2).unwrap();
    gamma * epsilon
}

fn part2(input: &Vec<&str>) -> i32 {
    let oxygen = select_number(
        &input,
        |zeros, ones| if zeros > ones { true } else { false },
    );
    let co2 = select_number(
        &input,
        |zeros, ones| if ones < zeros { false } else { true },
    );
    oxygen * co2
}

fn select_number(input: &Vec<&str>, take_zeros: fn(usize, usize) -> bool) -> i32 {
    let mut numbers = input.clone();
    for column_index in 0..input[0].len() {
        let split_numbers =
            numbers
                .iter()
                .fold((Vec::new(), Vec::new()), |(mut zeros, mut ones), row| {
                    if row.as_bytes()[column_index] as char == '0' {
                        zeros.push(*row);
                    } else {
                        ones.push(*row);
                    }
                    (zeros, ones)
                });

        if take_zeros(split_numbers.0.len(), split_numbers.1.len()) {
            numbers = split_numbers.0;
        } else {
            numbers = split_numbers.1;
        }

        if numbers.len() == 1 {
            break;
        }
    }
    return i32::from_str_radix(&numbers[0], 2).unwrap();
}
