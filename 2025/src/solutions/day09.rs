use std::fs;

#[derive(Clone, Copy)]
struct Vec2 {
    x: usize,
    y: usize,
}

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day09.txt").unwrap();
    let vecs = parse_input(content);

    println!("part1: {}", part1(vecs));
}

fn parse_input(input: String) -> Vec<Vec2> {
    return input
        .lines()
        .map(|line| {
            let (x, y) = line.split_once(",").unwrap();
            return Vec2 {
                x: x.parse().unwrap(),
                y: y.parse().unwrap(),
            };
        })
        .collect();
}

fn part1(input: Vec<Vec2>) -> usize {
    let mut pairs: Vec<(Vec2, Vec2)> = vec![];

    for i in 0..input.len() - 1 {
        for j in i + 1..input.len() {
            let v1 = input[i];
            let v2 = input[j];
            pairs.push((v1, v2));
        }
    }

    return pairs
        .iter()
        .map(|(v1, v2)| {
            let width = 1 + v1.x.abs_diff(v2.x);
            let height = 1 + v1.y.abs_diff(v2.y);
            return width * height;
        })
        .max()
        .unwrap();
}
