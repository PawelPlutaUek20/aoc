use std::{fs, vec};

#[derive(Clone, Copy)]
enum Operator {
    Add,
    Multiply,
}

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day06.txt").unwrap();

    let part1 = grand_total(parse_horizontal(&content));
    let part2 = grand_total(parse_verical(&content));

    println!("part1: {part1}");
    println!("part2: {part2}");
}

fn grand_total(input: (Vec<Vec<u64>>, Vec<Operator>)) -> u64 {
    let (numbers, operators) = input;

    return numbers
        .iter()
        .zip(operators.iter())
        .map(|(nums, op)| match op {
            Operator::Add => nums.iter().fold(0, |acc, x| acc + x),
            Operator::Multiply => nums.iter().fold(1, |acc, x| acc * x),
        })
        .sum::<u64>();
}

fn parse_horizontal(input: &String) -> (Vec<Vec<u64>>, Vec<Operator>) {
    let mut numbers: Vec<Vec<u64>> = vec![];
    let mut operators: Vec<Operator> = vec![];

    let lines: Vec<&str> = input.lines().collect();
    for i in 0..lines.len() - 1 {
        for (j, num) in lines[i]
            .split_ascii_whitespace()
            .map(|x| x.parse().unwrap())
            .enumerate()
        {
            if i == 0 {
                numbers.push(vec![num])
            } else {
                numbers[j].push(num)
            }
        }
    }

    let ops = lines[lines.len() - 1];
    for op in ops.split_ascii_whitespace() {
        match op {
            "*" => operators.push(Operator::Multiply),
            "+" => operators.push(Operator::Add),
            _ => unreachable!(),
        }
    }

    return (numbers, operators);
}

fn parse_verical(input: &String) -> (Vec<Vec<u64>>, Vec<Operator>) {
    let mut numbers: Vec<Vec<char>> = vec![];
    let mut operators: Vec<Operator> = vec![];

    let lines: Vec<&str> = input.lines().collect();
    for i in 0..lines.len() - 1 {
        for (j, num) in lines[i].chars().enumerate() {
            if i == 0 {
                numbers.push(vec![num])
            } else {
                numbers[j].push(num)
            }
        }
    }

    let numbers: Vec<String> = numbers.iter().map(|v| v.iter().collect()).collect();
    let numbers: Vec<Vec<u64>> = numbers
        .split(|s| s.trim().is_empty())
        .filter(|chunk| !chunk.is_empty())
        .map(|v| v.iter().map(|s| s.trim().parse().unwrap()).collect())
        .collect();

    let ops = lines[lines.len() - 1];
    for op in ops.split_ascii_whitespace() {
        match op {
            "*" => operators.push(Operator::Multiply),
            "+" => operators.push(Operator::Add),
            _ => unreachable!(),
        }
    }

    return (numbers, operators);
}
