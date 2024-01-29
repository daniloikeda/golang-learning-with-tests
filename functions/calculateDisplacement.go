package main

import ("fmt")

func main() {
	fmt.Printf("Enter the values for: Acceleration, Initial Velocity and Initial Displament: \n");
	acc, vel, displ := 0.0, 0.0, 0.0

	_, err := fmt.Scan(&acc)
	if err != nil {
		fmt.Errorf("Invalid acceleration value")
	}

	_, err = fmt.Scan(&vel)
	if err != nil {
		fmt.Errorf("Invalid inital velocity value")
	}

	_, err = fmt.Scan(&displ)
	if err != nil {
		fmt.Errorf("Invalid initial displacement value")
	}

	displacement := GenDisplaceFn(acc, vel, displ)

	fmt.Println("Enter the time you which to calculate the displacement: ")
	time := 0.0
	_, err = fmt.Scan(&time)

	if err != nil {
		fmt.Errorf("Invalid time value")
	} else {
		fmt.Printf("Calculated Displacement: %v", displacement(time))
	}
	
}

func GenDisplaceFn(acc, initVel , initDisp float64) func (float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acc * time * time) + (initVel * time) + initDisp;
	}
}
