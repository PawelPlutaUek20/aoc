pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day20.txt").unwrap();

    println!("part1: {}", part_1(&input));
    println!("part2: {}", part_2(&input));
}

#[derive(Debug)]
struct Range {
    from: u32,
    to: u32,
}

impl Range {
    fn parse(input: &str) -> Self {
        let (from, to) = input.split_once('-').unwrap();
        return Range {
            from: from.parse().unwrap(),
            to: to.parse().unwrap(),
        };
    }
}

fn part_1(input: &str) -> u32 {
    let mut ranges: Vec<Range> = Vec::new();

    for line in input.lines() {
        let range = Range::parse(line);
        let pos = ranges
            .binary_search_by(|e| e.from.cmp(&range.from))
            .unwrap_or_else(|e| e);

        ranges.insert(pos, range);
    }

    let reduced_ranges = reduce_ranges(ranges);
    return reduced_ranges[0].to + 1;
}

fn part_2(input: &str) -> u32 {
    let mut ranges: Vec<Range> = Vec::new();

    for line in input.lines() {
        let range = Range::parse(line);
        let pos = ranges
            .binary_search_by(|e| e.from.cmp(&range.from))
            .unwrap_or_else(|e| e);

        ranges.insert(pos, range);
    }

    let reduced_ranges = reduce_ranges(ranges);
    return count_whitelisted(&reduced_ranges);
}

fn reduce_ranges(ranges: Vec<Range>) -> Vec<Range> {
    ranges.into_iter().fold(Vec::new(), |mut acc, e| {
        match acc.last_mut() {
            Some(last) if last.to == u32::MAX || e.from <= last.to + 1 => {
                last.to = last.to.max(e.to);
            }
            _ => {
                acc.push(e);
            }
        }
        acc
    })
}

fn count_whitelisted(ranges: &Vec<Range>) -> u32 {
    let mut result = ranges
        .into_iter()
        .enumerate()
        .fold(0, |mut acc, (idx, range)| {
            if idx == 0 {
                return acc;
            }

            let prev = ranges[idx - 1].to;
            let curr = range.from;

            acc += curr - prev - 1;
            return acc;
        });

    result += ranges.first().unwrap().from;
    result += u32::MAX - ranges.last().unwrap().to;
    return result;
}
