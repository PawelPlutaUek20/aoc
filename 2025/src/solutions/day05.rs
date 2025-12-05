use std::{fs, vec};

enum Mode {
    Ranges,
    Ids,
}

#[derive(Debug, Clone, Copy)]
struct Range {
    from: u64,
    to: u64,
}

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day05.txt").unwrap();

    let mut ranges: Vec<Range> = vec![];
    let mut ids: Vec<u64> = vec![];
    let mut mode = Mode::Ranges;

    for line in content.lines() {
        if line.is_empty() {
            mode = Mode::Ids;
            continue;
        }

        match mode {
            Mode::Ids => {
                let id = line.parse::<u64>().unwrap();
                ids.push(id);
            }
            Mode::Ranges => {
                let (from, to) = line.split_once('-').unwrap();
                let range = Range {
                    from: from.parse::<u64>().unwrap(),
                    to: to.parse::<u64>().unwrap(),
                };
                ranges.push(range);
            }
        }
    }

    let part1 = get_fresh_count(ids, &ranges);
    let part2 = get_uniq_ranges(&mut ranges);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn get_fresh_count(ids: Vec<u64>, ranges: &Vec<Range>) -> u64 {
    let mut fresh_count = 0;

    for id in ids {
        for range in ranges {
            if id >= range.from && id <= range.to {
                fresh_count += 1;
                break;
            }
        }
    }

    return fresh_count;
}

fn get_uniq_ranges(ranges: &mut Vec<Range>) -> u64 {
    ranges.sort_by_key(|range| range.from);

    let mut merged_ranges: Vec<Range> = vec![ranges[0]];

    for range in ranges {
        let last = merged_ranges.last_mut().unwrap();
        if range.from <= last.to {
            last.to = last.to.max(range.to);
        } else if range.from == last.from {
            last.to = last.to.max(range.to);
        } else {
            merged_ranges.push(range.clone());
        }
    }

    let mut uniq_ranges = 0;

    for range in merged_ranges {
        uniq_ranges += range.to - range.from + 1;
    }

    return uniq_ranges;
}
