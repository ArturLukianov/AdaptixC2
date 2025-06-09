use std::{thread, time};

mod config;

fn main() {
    let is_active = true;

    // Initialize agent - generate ID
    let agent_id: u32 = 1; // Generate random ID

    // Send initial packet to the listener (for now only HTTP, later be more)
    loop {
        if let Ok(body) = reqwest::blocking::get("http://127.0.0.1:8000/initial") {
            break 
        }
        thread::sleep(time::Duration::from_secs(30)); // Wait for reconnect. TODO: make incremental, like in chisel with 5m limit
    } 

    // Wait for task
    while is_active {
        thread::sleep(time::Duration::from_secs(config::SLEEP_TIME)); 
    }
}
