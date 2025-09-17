pub fn solve() {
    part_1();
    part_2();
}

fn part_1() {
    let content = std::fs::read_to_string("src/inputs/day07.txt").unwrap();
    let result = content
        .lines()
        .map(|line| {
            let mut has_abba = false;
            let mut has_abba_within_hypernet = false;

            line.split(|c| c == '[' || c == ']')
                .enumerate()
                .for_each(|(idx, arr)| {
                    let ip: Vec<char> = arr.chars().collect();
                    let is_within_hypernet = idx & 1 == 1;

                    if is_within_hypernet {
                        if check_abba(ip) {
                            has_abba_within_hypernet = true;
                        }
                    } else if check_abba(ip) {
                        has_abba = true;
                    }
                });

            return !has_abba_within_hypernet && has_abba;
        })
        .fold(0, |acc, e| acc + e as i32);

    println!("part1: {}", result);
}

fn part_2() {
    let content = std::fs::read_to_string("src/inputs/day07.txt").unwrap();
    let result = content
        .lines()
        .map(|line| {
            let mut aba_list = vec![];
            let mut aba_list_within_hypernet = vec![];

            line.split(|c| c == '[' || c == ']')
                .enumerate()
                .for_each(|(idx, arr)| {
                    let ip: Vec<char> = arr.chars().collect();
                    let is_within_hypernet = idx & 1 == 1;

                    if is_within_hypernet {
                        for aba in find_aba(ip) {
                            aba_list_within_hypernet.push(aba);
                        }
                    } else {
                        for aba in find_aba(ip) {
                            aba_list.push(aba);
                        }
                    }
                });

            return supports_ssl(&aba_list, &aba_list_within_hypernet);
        })
        .fold(0, |acc, e| acc + e as i32);

    println!("part2: {}", result);
}

fn check_abba(ip: Vec<char>) -> bool {
    for i in 0..ip.len() - 3 {
        let a = ip[i];
        let b = ip[i + 1];
        let c = ip[i + 2];
        let d = ip[i + 3];
        let is_abba = a == d && b == c && a != b;

        if is_abba {
            return true;
        };
    }
    return false;
}

fn find_aba(ip: Vec<char>) -> Vec<(char, char, char)> {
    let mut aba_list = vec![];

    for i in 0..ip.len() - 2 {
        let a = ip[i];
        let b = ip[i + 1];
        let c = ip[i + 2];
        let is_aba = a == c && a != b;
        if is_aba {
            aba_list.push((a, b, c));
        };
    }
    return aba_list;
}

fn supports_ssl(
    aba_list: &Vec<(char, char, char)>,
    aba_list_within_hypernet: &Vec<(char, char, char)>,
) -> bool {
    for aba in aba_list {
        let (a, b, _) = aba;

        for bab in aba_list_within_hypernet {
            let (d, e, _) = bab;

            if a == e && b == d {
                return true;
            }
        }
    }

    return false;
}
