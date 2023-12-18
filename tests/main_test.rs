mod rust;
use rust::{set_rice_and_water_added, RICE_COOKER_STATE, RICE_AND_WATER_ADDED, STATE_OFF};

#[cfg(test)]
fn test_initial_state() {
    unsafe {
        assert_eq!(RICE_COOKER_STATE, STATE_OFF);
    }

    unsafe {
        assert_eq!(RICE_AND_WATER_ADDED, false);
    }
}

#[cfg(test)]
fn test_add_rice_and_water() {
    set_rice_and_water_added(true);

    unsafe {
        assert_eq!(RICE_AND_WATER_ADDED, true);
    }
}

