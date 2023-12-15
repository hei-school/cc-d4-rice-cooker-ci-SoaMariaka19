import * as readlineSync from 'readline-sync';

const STATE_OFF: string = "Off";
const STATE_COOKING: string = "Cooking in progress";
const STATE_KEEP_WARM: string = "Keep warm mode";

let riceCookerState: string = STATE_OFF;
let riceAndWaterAdded: boolean = false;

function displayState(): void {
    console.log("\nCurrent state of the rice cooker:", riceCookerState);
}

function plugIn(): void {
    if (riceCookerState === STATE_OFF && riceAndWaterAdded) {
        riceCookerState = STATE_COOKING;
        console.log("The rice cooker is plugged in, and cooking begins.");
    } else if (!riceAndWaterAdded) {
        console.log("Error: Add rice and water before plugging in and starting cooking.");
    } else {
        console.log("Error: The rice cooker is already cooking or in keep warm mode.");
    }
}

function finishCooking(): void {
    if (riceCookerState === STATE_COOKING) {
        riceCookerState = STATE_KEEP_WARM;
        console.log("Cooking is finished. The rice cooker is in keep warm mode.");
    } else {
        console.log("Error: No cooking in progress.");
    }
}

function quitProgram(): void {
    console.log("\n=======================================================================");
    console.log("Goodbye! Thanks for using the rice cooker program.\n");
}

function setRiceAndWaterAdded(value: boolean): void {
    riceAndWaterAdded = value;
}




if(require.main === module){
    console.log("\nWelcome to the rice cooker program.");

    const isProgramRunning: boolean = true;

    while (isProgramRunning) {
        console.log("\n=======================================================================\n        Menu:\n");
        console.log("1. Add rice and water");
        console.log("2. Plug in the rice cooker");
        console.log("3. Cook rice");
        console.log("4. Keep warm");
        console.log("5. Rice cooker state");
        console.log("6. End of cooking notification");
        console.log("7. Quit the program");

        const choice: string = readlineSync.question("\n=======================================================================\nEnter your choice number: ");

        if (!choice.match(/^\d+$/)) {
            console.log("Error: Please enter a number.");
            continue;
        }

        const choiceNumber: number = parseInt(choice);

        switch (choiceNumber) {
            case 1:
                console.log("You added rice and water.");
                riceAndWaterAdded = true;
                break;
            case 2:
                plugIn();
                break;
            case 3:
                if (riceCookerState === STATE_OFF) {
                    console.log("Error: Add rice and water before starting cooking.");
                } else {
                    console.log("Cooking rice is in progress.");
                }
                break;
            case 4:
                if (riceCookerState === STATE_COOKING) {
                    finishCooking();
                } else {
                    console.log("Error: No cooking in progress.");
                }
                break;
            case 5:
                displayState();
                break;
            case 6:
                if (riceCookerState === STATE_KEEP_WARM) {
                    console.log("Cooking is finished. The rice cooker is in keep warm mode.");
                } else {
                    console.log("Error: No finished cooking.");
                }
                break;
            case 7:
                quitProgram();
                process.exit(0);
                break;
            default:
                console.log("Error: Invalid choice. Please enter a valid number.");
                break;
        }
    }

}

export {STATE_COOKING, STATE_KEEP_WARM, STATE_OFF, riceAndWaterAdded, riceCookerState, displayState, plugIn, finishCooking, quitProgram, setRiceAndWaterAdded}