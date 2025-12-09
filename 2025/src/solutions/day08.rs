use std::{fs, hash::Hash};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Vec3 {
    x: usize,
    y: usize,
    z: usize,
}

impl Vec3 {
    fn new(x: usize, y: usize, z: usize) -> Self {
        return Self { x, y, z };
    }

    fn distance(&self, other: &Self) -> f64 {
        let dx = self.x.abs_diff(other.x);
        let dy = self.y.abs_diff(other.y);
        let dz = self.z.abs_diff(other.z);
        ((dx * dx + dy * dy + dz * dz) as f64).sqrt()
    }
}

struct UnionFind {
    id: Vec<usize>,
}

impl UnionFind {
    fn new(len: usize) -> Self {
        let id = (0..len).enumerate().map(|(id, _)| id).collect();
        return Self { id };
    }

    fn root(&self, mut id: usize) -> usize {
        while self.id[id] != id {
            id = self.id[id];
        }
        return id;
    }

    fn unite(&mut self, p: usize, q: usize) {
        let i = self.root(p);
        let j = self.root(q);
        self.id[i] = j;
    }

    fn sizes(&self) -> Vec<usize> {
        let mut sz: Vec<usize> = (0..self.id.len()).map(|_| 0).collect();
        for &id in &self.id {
            let i = self.root(id);
            sz[i] += 1;
        }
        return sz;
    }
}

pub fn solve() {
    let content = fs::read_to_string("src/inputs/day08.txt").unwrap();

    let junction_boxes = parse_input(content);

    println!("part1: {}", part1(&junction_boxes));
    println!("part2: {}", part2(&junction_boxes));
}

fn part2(junction_boxes: &Vec<Vec3>) -> usize {
    let size = junction_boxes.len();
    let mut uf = UnionFind::new(size);

    for (v1, v2) in find_pairs(&junction_boxes) {
        uf.unite(v1, v2);
        if uf.sizes().iter().any(|&sz| sz == size) {
            let x1 = junction_boxes[v1].x;
            let x2 = junction_boxes[v2].x;
            return x1 * x2;
        }
    }

    unreachable!()
}

fn part1(junction_boxes: &Vec<Vec3>) -> usize {
    let size = junction_boxes.len();
    let mut uf = UnionFind::new(size);

    for &(v1, v2) in find_pairs(&junction_boxes).iter().take(1000) {
        uf.unite(v1, v2);
    }

    let mut sz = uf.sizes();
    sz.sort_by(|a, b| b.cmp(a));

    return sz[0] * sz[1] * sz[2];
}

fn find_pairs(junction_boxes: &Vec<Vec3>) -> Vec<(usize, usize)> {
    let mut pairs: Vec<(usize, usize)> = vec![];

    for i in 0..junction_boxes.len() - 1 {
        for j in i + 1..junction_boxes.len() {
            pairs.push((i, j));
        }
    }

    pairs.sort_by(|&(a1, b1), &(a2, b2)| {
        let v1 = junction_boxes[a1].distance(&junction_boxes[b1]);
        let v2 = junction_boxes[a2].distance(&junction_boxes[b2]);
        return v1.total_cmp(&v2);
    });

    return pairs;
}

fn parse_input(input: String) -> Vec<Vec3> {
    return input
        .lines()
        .map(|line| {
            let coords: Vec<usize> = line.split(",").map(|s| s.parse().unwrap()).collect();
            assert_eq!(coords.len(), 3);
            return Vec3::new(coords[0], coords[1], coords[2]);
        })
        .collect();
}
