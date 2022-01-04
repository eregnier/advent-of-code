use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn day3_step1() {
  let reader = BufReader::new(File::open("inputs/d3").expect("cannot open 'input' file"));
    let mut ones = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    let mut zeroes = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    for line in reader.lines() {
        let chars = line.unwrap().chars().collect::<Vec<char>>();
        for i in 0..12 {
            match chars[i] {
                '0' => zeroes[i] += 1,
                '1' => ones[i] += 1,
                _ => print!("Err{}#", chars[i]),
            }
        }
    }
    let mut gamma = Vec::new();
    let mut epsilon = Vec::new();
    for i in 0..12 {
        if ones[i] > zeroes[i] {
            gamma.push('1');
            epsilon.push('0');
        } else {
            gamma.push('0');
            epsilon.push('1');
        }
    }
    let e = &epsilon.iter().cloned().collect::<String>();
    let eval = isize::from_str_radix(e, 2).unwrap();
    let g = &gamma.iter().cloned().collect::<String>();
    let gval = isize::from_str_radix(g, 2).unwrap();
    println!("d3s1 | {} X {} = {}", eval, gval, eval * gval);
}

pub fn day3_step2() {
}
