use std::fs;

#[derive(Debug, Clone, Copy)]
enum Tile {
    RollOfPaper,
    Empty,
}

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day04.txt").unwrap();
    let mut grid = parse_input(content);

    let part1 = rolls_to_remove(&grid).len();
    let part2 = rolls_to_remove_until_none(&mut grid);

    println!("part1: {}", part1);
    println!("part2: {}", part2);
}

fn parse_input(input: String) -> Vec<Vec<Tile>> {
    return input
        .lines()
        .map(|line| {
            line.chars()
                .map(|char| match char {
                    '@' => Tile::RollOfPaper,
                    '.' => Tile::Empty,
                    chr => unreachable!("Unknown char: {chr}"),
                })
                .collect()
        })
        .collect();
}

fn rolls_to_remove_until_none(grid: &mut Vec<Vec<Tile>>) -> usize {
    let to_remove = rolls_to_remove(grid);
    let remove_count = to_remove.len();

    if remove_count == 0 {
        return 0;
    }

    for roll in to_remove {
        grid[roll.0][roll.1] = Tile::Empty;
    }

    return remove_count + rolls_to_remove_until_none(grid);
}

fn rolls_to_remove(grid: &Vec<Vec<Tile>>) -> Vec<(usize, usize)> {
    let mut rolls_of_paper: Vec<(usize, usize)> = vec![];

    for (row, cols) in grid.iter().enumerate() {
        for (col, tile) in cols.iter().enumerate() {
            match tile {
                Tile::Empty => continue,
                Tile::RollOfPaper => {}
            };

            let mut rolls_of_paper_count = 0;
            for tile in get_adjacent_tiles(&grid, row, col) {
                match tile {
                    Tile::RollOfPaper => rolls_of_paper_count += 1,
                    Tile::Empty => {}
                };
            }

            if rolls_of_paper_count < 4 {
                rolls_of_paper.push((row, col));
            }
        }
    }

    return rolls_of_paper;
}

fn get_adjacent_tiles(grid: &Vec<Vec<Tile>>, row: usize, col: usize) -> Vec<Tile> {
    let rows = grid.len();
    let cols = grid[0].len();

    let mut adjacent_tiles = vec![];

    if row >= 1 && col >= 1 {
        adjacent_tiles.push(grid[row - 1][col - 1]);
    }
    if row >= 1 && col < cols - 1 {
        adjacent_tiles.push(grid[row - 1][col + 1]);
    }
    if col >= 1 {
        adjacent_tiles.push(grid[row][col - 1]);
    }
    if col < cols - 1 {
        adjacent_tiles.push(grid[row][col + 1]);
    }
    if row >= 1 {
        adjacent_tiles.push(grid[row - 1][col]);
    }
    if row < rows - 1 {
        adjacent_tiles.push(grid[row + 1][col]);
    }
    if row < rows - 1 && col >= 1 {
        adjacent_tiles.push(grid[row + 1][col - 1]);
    }
    if row < rows - 1 && col < cols - 1 {
        adjacent_tiles.push(grid[row + 1][col + 1]);
    }

    return adjacent_tiles;
}
