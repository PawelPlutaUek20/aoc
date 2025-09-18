pub fn solve() {
    let input = std::fs::read_to_string("src/inputs/day09.txt").unwrap();
    let parsed_input = ExperimentalFormat::parse(&input);

    println!("part1: {}", part_1(&parsed_input));
    println!("part2: {}", part_2(&parsed_input));
}

fn part_1(list: &Vec<ExperimentalFormat>) -> usize {
    let mut result = 0;

    for format in list {
        match format {
            ExperimentalFormat::Marker {
                parsed: _,
                raw,
                repeat_times,
            } => result += raw.len() * repeat_times,
            ExperimentalFormat::Character => {
                result += 1;
            }
        }
    }

    return result;
}

fn part_2(list: &Vec<ExperimentalFormat>) -> usize {
    let mut result = 0;

    for format in list {
        match format {
            ExperimentalFormat::Marker {
                raw: _,
                parsed,
                repeat_times,
            } => {
                result += part_2(parsed) * repeat_times;
            }
            ExperimentalFormat::Character => {
                result += 1;
            }
        }
    }

    return result;
}

enum ExperimentalFormat {
    Marker {
        raw: String,
        parsed: Vec<ExperimentalFormat>,
        repeat_times: usize,
    },
    Character,
}

impl ExperimentalFormat {
    fn parse(input: &str) -> Vec<Self> {
        let mut result = Vec::new();

        let chars: Vec<char> = input.chars().collect();

        let mut i = 0;
        while i < chars.len() {
            let char = chars[i];

            if char.is_whitespace() {
                i += 1;
            } else if char == '(' {
                let mut j = i + 1;
                while chars[j] != ')' {
                    j += 1;
                }
                let marker_str: String = chars[i + 1..j].iter().collect();
                let parts: Vec<&str> = marker_str.split("x").collect();

                let characters_count = parts[0].parse::<usize>().unwrap();
                let repeat_times = parts[1].parse::<usize>().unwrap();

                let characters_start = j + 1;
                let characters_end = characters_start + characters_count;

                let raw: String = chars[characters_start..characters_end].iter().collect();
                let parsed = ExperimentalFormat::parse(&raw);

                result.push(ExperimentalFormat::Marker {
                    raw,
                    parsed,
                    repeat_times,
                });

                i = characters_end
            } else {
                result.push(ExperimentalFormat::Character);
                i += 1;
            }
        }

        return result;
    }
}
