use std::fs;

pub fn solve() {
    part_1();
}

fn part_1() {
    let mut content = fs::read_to_string("src/inputs/day02.txt").unwrap();
    content = content.trim_end().to_string();

    let mut part1 = 0;
    let mut part2 = 0;

    for range in content.split(",") {
        let (from, to) = range.split_once("-").unwrap();

        let from = from.parse::<u64>().unwrap();
        let to = to.parse::<u64>().unwrap();

        for i in from..=to {
            let id = &i.to_string();

            if is_invalid(id) {
                part1 += i;
            }
            if is_invalid2(id) {
                part2 += i;
            }
        }
    }

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn is_invalid(id: &str) -> bool {
    if id.len() & 1 != 0 {
        return false;
    }

    let mid = id.len() / 2;
    return id[0..mid] == id[mid..];
}

fn is_invalid2(id: &str) -> bool {
    (1..=id.len() / 2)
        .filter(|i| id.len() % i == 0)
        .any(|size| {
            (0..id.len())
                .step_by(size)
                .all(|i| id[i..i + size] == id[0..size])
        })
}
