use std::{collections::HashMap, fs};

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day11.txt").unwrap();

    println!("part1: {}", part1_memo_impl(&parse_input(&content)));
    println!("part2: {}", part2_memo_impl(&parse_input(&content)));
}

const fn key(s: &str) -> usize {
    let b = s.as_bytes();
    return (b[0] as usize) << 16 | (b[1] as usize) << 8 | (b[2] as usize);
}

const OUT: usize = key("out");
const DAC: usize = key("dac");
const FFT: usize = key("fft");

fn dfs(
    memo: &mut HashMap<(usize, bool, bool), usize>,
    adj_list: &HashMap<usize, Vec<usize>>,
    curr: usize,
    seen_dac: bool,
    seen_fft: bool,
) -> usize {
    if curr == OUT {
        return if seen_dac && seen_fft { 1 } else { 0 };
    }

    let mut acc = 0;

    for &next in adj_list.get(&curr).unwrap() {
        let seen_dac = seen_dac || next == DAC;
        let seen_fft = seen_fft || next == FFT;
        acc += dfs_memo(memo, adj_list, next, seen_dac, seen_fft);
    }

    return acc;
}

fn part1_memo_impl(adj_list: &HashMap<usize, Vec<usize>>) -> usize {
    let mut cache = HashMap::new();
    return dfs_memo(&mut cache, adj_list, key("you"), true, true);
}

fn part2_memo_impl(adj_list: &HashMap<usize, Vec<usize>>) -> usize {
    let mut cache = HashMap::new();
    return dfs_memo(&mut cache, adj_list, key("svr"), false, false);
}

fn dfs_memo(
    memo: &mut HashMap<(usize, bool, bool), usize>,
    adj_list: &HashMap<usize, Vec<usize>>,
    curr: usize,
    seen_dac: bool,
    seen_fft: bool,
) -> usize {
    if let Some(&value) = memo.get(&(curr, seen_dac, seen_fft)) {
        return value;
    }

    let result = dfs(memo, adj_list, curr, seen_dac, seen_fft);

    memo.insert((curr, seen_dac, seen_fft), result);
    return result;
}

fn parse_input(input: &String) -> HashMap<usize, Vec<usize>> {
    let mut adjacency_list: HashMap<usize, Vec<usize>> = HashMap::new();

    for line in input.lines() {
        let (k, v) = line.split_once(":").unwrap();
        let values: Vec<usize> = v.split_whitespace().map(|s| key(s)).collect();
        adjacency_list.insert(key(k), values);
    }

    return adjacency_list;
}
