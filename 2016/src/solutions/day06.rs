use std::{collections::HashMap, fs};

pub fn solve() {
    part_1();
    part_2();
}

fn part_1() {
    let content = fs::read_to_string("src/inputs/day06.txt").unwrap();

    let grid: Vec<Vec<char>> = content.lines().map(|line| line.chars().collect()).collect();

    let row_len = grid.len();
    let column_len = grid.first().map(|col| col.len()).unwrap();

    let mut result = String::new();

    for col in 0..column_len {
        let mut occurrences = HashMap::new();

        for row in 0..row_len {
            let char = grid[row][col];
            *occurrences.entry(char).or_insert(0) += 1
        }

        let mut chars: Vec<char> = occurrences.keys().copied().collect();
        chars.sort_by(|&a, &b| {
            let count_a = occurrences[&a];
            let count_b = occurrences[&b];
            return count_b.cmp(&count_a);
        });

        let most_common = chars.first().unwrap().to_owned();
        result.push(most_common);
    }

    println!("part1: {result}");
}

fn part_2() {
    let content = fs::read_to_string("src/inputs/day06.txt").unwrap();

    let grid: Vec<Vec<char>> = content.lines().map(|line| line.chars().collect()).collect();

    let row_len = grid.len();
    let column_len = grid.first().map(|col| col.len()).unwrap();

    let mut result = String::new();

    for col in 0..column_len {
        let mut occurrences = HashMap::new();

        for row in 0..row_len {
            let char = grid[row][col];
            *occurrences.entry(char).or_insert(0) += 1
        }

        let mut chars: Vec<char> = occurrences.keys().copied().collect();
        chars.sort_by(|&a, &b| {
            let count_a = occurrences[&a];
            let count_b = occurrences[&b];
            return count_a.cmp(&count_b);
        });

        let lest_common = chars.first().unwrap().to_owned();
        result.push(lest_common);
    }

    println!("part2: {result}");
}
