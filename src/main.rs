use std::io;

const STATE_OFF: &str = "Off";
const STATE_COOKING: &str = "Cooking in progress";
const STATE_KEEP_WARM: &str = "Keep warm mode";

static mut RICE_COOKER_STATE: &'static str = STATE_OFF;
static mut RICE_AND_WATER_ADDED: bool = false;

fn display_state() {
    println!("\nCurrent state of the rice cooker: {}", unsafe { RICE_COOKER_STATE });
}

fn plug_in() {
    unsafe {
        if RICE_COOKER_STATE == STATE_OFF && RICE_AND_WATER_ADDED {
            RICE_COOKER_STATE = STATE_COOKING;
            println!("The rice cooker is plugged in, and cooking begins.");
        } else if !RICE_AND_WATER_ADDED {
            println!("Error: Add rice and water before plugging in and starting cooking.");
        } else {
            println!("Error: The rice cooker is already cooking or in keep warm mode.");
        }
    }
}

fn finish_cooking() {
    unsafe {
        if RICE_COOKER_STATE == STATE_COOKING {
            RICE_COOKER_STATE = STATE_KEEP_WARM;
            println!("Cooking is finished. The rice cooker is in keep warm mode.");
        } else {
            println!("Error: No cooking in progress.");
        }
    }
}

fn quit_program() {
    println!("\n=======================================================================");
    println!("Goodbye! Thanks for using the rice cooker program.\n");
}

fn set_rice_and_water_added(value: bool) {
    unsafe {
        RICE_AND_WATER_ADDED = value;
    }
}

fn main() {
    println!("\nWelcome to the rice cooker program.");

    loop {
        println!(
            "\n=======================================================================\n        Menu:\n"
        );
        println!("1. Add rice and water");
        println!("2. Plug in the rice cooker");
        println!("3. Cook rice");
        println!("4. Keep warm");
        println!("5. Rice cooker state");
        println!("6. End of cooking notification");
        println!("7. Quit the program");

        let mut choice = String::new();
        println!(
            "\n=======================================================================\nEnter your choice number: "
        );
        io::stdin().read_line(&mut choice).expect("Failed to read line");

        if !choice.trim().parse::<u32>().is_ok() {
            println!("Error: Please enter a number.");
            continue;
        }

        let choice_number: u32 = match choice.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Error: Invalid input. Please enter a valid number.");
                continue;
            }
        };

        match choice_number {
            1 => {
                println!("You added rice and water.");
                set_rice_and_water_added(true);
            }
            2 => plug_in(),
            3 => unsafe {
                if RICE_COOKER_STATE == STATE_OFF {
                    println!("Error: Add rice and water before starting cooking.");
                } else {
                    println!("Cooking rice is in progress.");
                }
            },
            4 => unsafe {
                if RICE_COOKER_STATE == STATE_COOKING {
                    finish_cooking();
                } else {
                    println!("Error: No cooking in progress.");
                }
            },
            5 => display_state(),
            6 => unsafe {
                if RICE_COOKER_STATE == STATE_KEEP_WARM {
                    println!("Cooking is finished. The rice cooker is in keep warm mode.");
                } else {
                    println!("Error: No finished cooking.");
                }
            },
            7 => {
                quit_program();
                break;
            }
            _ => println!("Error: Invalid choice. Please enter a valid number."),
        }
    }
}
