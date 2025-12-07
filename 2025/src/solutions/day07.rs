use std::collections::{HashMap, HashSet};
use std::fs;

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day07.txt").unwrap();
    let lines: Vec<Vec<char>> = content.lines().map(|line| line.chars().collect()).collect();

    let start_col = lines[0].iter().position(|&c| c == 'S').unwrap();
    let start = (0, start_col);

    println!("part1: {}", part1_impl(&lines, start));
    println!("part2: {}", part2_memo_impl(&lines, start));
}

fn part1(
    seen: &mut HashSet<(usize, usize)>,
    grid: &Vec<Vec<char>>,
    current: (usize, usize),
) -> usize {
    let (row, col) = current;

    if grid.len() <= row {
        return 0;
    }

    if seen.contains(&current) {
        return 0;
    }

    match grid[row][col] {
        'S' | '.' => part1(seen, grid, (row + 1, col)),
        '^' => {
            seen.insert(current);
            return 1 + part1(seen, grid, (row, col - 1)) + part1(seen, grid, (row, col + 1));
        }
        _ => unreachable!(),
    }
}

fn part1_impl(grid: &Vec<Vec<char>>, current: (usize, usize)) -> usize {
    let mut seen = HashSet::new();
    part1(&mut seen, grid, current)
}

fn part2(
    memo: &mut HashMap<(usize, usize), usize>,
    grid: &Vec<Vec<char>>,
    current: (usize, usize),
) -> usize {
    let (row, col) = current;

    if grid.len() <= row {
        return 1;
    }

    match grid[row][col] {
        'S' | '.' => part2_memo(memo, grid, (row + 1, col)),
        '^' => part2_memo(memo, grid, (row, col + 1)) + part2_memo(memo, grid, (row, col - 1)),
        _ => unreachable!(),
    }
}

fn part2_memo(
    memo: &mut HashMap<(usize, usize), usize>,
    grid: &Vec<Vec<char>>,
    current: (usize, usize),
) -> usize {
    if let Some(&value) = memo.get(&current) {
        return value;
    }

    let result = part2(memo, grid, current);

    memo.insert(current, result);
    return result;
}

fn part2_memo_impl(grid: &Vec<Vec<char>>, current: (usize, usize)) -> usize {
    let mut cache = HashMap::new();
    return part2_memo(&mut cache, grid, current);
}
