use std::fs::File;
use std::io::{BufRead, BufReader};

fn main() {
    step1();
    step2();
}

fn step1() {
    let reader = BufReader::new(File::open("input").expect("cannot open 'input' file"));
    let mut previous: i32 = 0;
    let mut count = 0;
    for line in reader.lines() {
        let current: i32 = line.unwrap().parse().unwrap();
        if previous != 0 {
            if current > previous {
                count += 1;
            }
        }
        previous = current;
    }
    println!("step 1 > {}", count)
}

fn step2() {
    let reader = BufReader::new(File::open("input").expect("cannot open 'input' file"));
    let mut values: Vec<i32> = Vec::new();
    for line in reader.lines() {
        values.push(line.unwrap().parse().unwrap());
    }
    let l = values.len();
    let mut windows: Vec<i32> = Vec::new();

    for (i, val) in values.iter().enumerate() {
        if i + 2 < l {
            windows.push(val + values[i + 1] + values[i + 2]);
        }
    }
    let mut previous :i32 = 0;
    let mut count = 0;
    for (i, w) in windows.iter().enumerate() {
        if i > 0 && w > &previous {
            count += 1;
        }
        previous = *w
    }
    println!("step 2 > {}", count)
}
