pub fn solve() {
    let content = std::fs::read_to_string("src/inputs/day08.txt").unwrap();

    let mut screen = Screen::new();

    for cmd in Command::parse(&content) {
        match cmd {
            Command::Rect { width, height } => {
                screen.rect(width, height);
            }
            Command::RotateRow { y, by } => {
                screen.rotate_row_by(y, by);
            }
            Command::RotateColum { x, by } => {
                screen.rotate_column_by(x, by);
            }
        }
    }

    println!("part1: {}", screen.count_lit_pixels());
    println!("part2:\n{}", screen);
}

const WIDTH: usize = 50;
const HEIGHT: usize = 6;

enum Command {
    RotateColum { x: usize, by: usize },
    RotateRow { y: usize, by: usize },
    Rect { width: usize, height: usize },
}

impl Command {
    fn parse(input: &str) -> Vec<Self> {
        let mut commands = Vec::new();

        for line in input.lines() {
            let parts: Vec<&str> = line
                .split(|c: char| c.is_whitespace() || c == '=')
                .collect();

            match parts.as_slice() {
                ["rotate", row_or_column, _, index_str, "by", by_str] => {
                    let index = index_str.parse::<usize>().unwrap();
                    let by = by_str.parse::<usize>().unwrap();
                    if *row_or_column == "row" {
                        commands.push(Self::RotateRow { y: index, by: by });
                    } else {
                        commands.push(Self::RotateColum { x: index, by: by });
                    }
                }
                ["rect", size_str] => {
                    let size: Vec<&str> = size_str.split("x").collect();

                    let width = size[0].parse::<usize>().unwrap();
                    let height = size[1].parse::<usize>().unwrap();
                    commands.push(Self::Rect { width, height });
                }
                _ => unreachable!(),
            }
        }

        return commands;
    }
}

struct Screen {
    pixels: [[bool; WIDTH]; HEIGHT],
}

impl Screen {
    fn new() -> Self {
        return Screen {
            pixels: [[false; WIDTH]; HEIGHT],
        };
    }

    fn width(&self) -> usize {
        return WIDTH;
    }

    fn height(&self) -> usize {
        return HEIGHT;
    }

    fn rect(&mut self, width: usize, height: usize) {
        for row in 0..height {
            for col in 0..width {
                self.pixels[row][col] = true;
            }
        }
    }

    fn rotate_column_by(&mut self, column: usize, by: usize) {
        let height = self.height();

        let mut column_copy = [false; HEIGHT];
        for row in 0..height {
            column_copy[row] = self.pixels[row][column]
        }

        for row in 0..height {
            let new_row = (row + height - by).rem_euclid(height);
            self.pixels[row][column] = column_copy[new_row]
        }
    }

    fn rotate_row_by(&mut self, row: usize, by: usize) {
        let width = self.width();

        let mut row_copy = [false; WIDTH];
        for col in 0..width {
            row_copy[col] = self.pixels[row][col]
        }

        for col in 0..width {
            let new_col = (col + width - by).rem_euclid(width);
            self.pixels[row][col] = row_copy[new_col]
        }
    }

    fn count_lit_pixels(&self) -> usize {
        let mut count = 0;

        for row in self.pixels {
            for col in row {
                if col {
                    count += 1;
                }
            }
        }

        return count;
    }
}

impl std::fmt::Display for Screen {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut result: Vec<String> = vec![];

        for i in 0..self.pixels.len() {
            let data: String = self.pixels[i]
                .map(|pixel| if pixel { return '#' } else { return '.' })
                .iter()
                .collect();

            result.push(data);
        }

        write!(f, "{}", result.join("\n"))
    }
}
