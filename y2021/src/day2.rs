use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn day2_step1() {
    let reader = BufReader::new(File::open("inputs/d2").expect("cannot open 'input' file"));
    let mut x = 0;
    let mut d = 0;
    for line in reader.lines() {
        let token = line.unwrap();
        let split: Vec<&str> = token.split(" ").collect();
        let distance = split[1].parse::<i32>().unwrap();
        match split[0].to_string().as_ref() {
            "up" => d -= distance,
            "down" => d += distance,
            "forward" => x += distance,
            _ => println!("command not found"),
        }
    }
    println!("d2s1 | x{} X d{} => {}", x, d, x * d);
}

pub fn day2_step2() {
    let reader = BufReader::new(File::open("inputs/d2").expect("cannot open 'input' file"));
    let mut d = 0;
    let mut a = 0;
    let mut x = 0;
    for line in reader.lines() {
        let token = line.unwrap();
        let split: Vec<&str> = token.split(" ").collect();
        let distance = split[1].parse::<i32>().unwrap();
        match split[0].to_string().as_ref() {
            "up" => a -= distance,
            "down" => a += distance,
            "forward" => {
                x += distance;
                d += a * distance;
            }
            _ => println!("command not found"),
        }
    }
    println!("d2s2 | x{} X d{} => {}", x, d, x * d);
}
