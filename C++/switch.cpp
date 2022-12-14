#include <cstring>
#include <iostream>

// A simple integer switch statement
void integerSwitch() {
    int age = 25;
    switch (age) {
        case 25:
            printf("25\n");
            break;
        case 26:
            printf("26\n");
            break;
        case 27:
            printf("27\n");
            break;
        default:
            printf("default\n");
    }
}

// A simple string switch statement
void stringSwitch() {
    char grade = 'B';
    switch (grade) {
        case 'A':
            printf("Excellent!\n");
            break;
        case 'B':
            printf("Good Job!\n");
            break;
        case 'C':
            printf("Well done\n");
            break;
        case 'D':
            printf("You passed\n");
            break;
        case 'F':
            printf("Better try again\n");
            break;
        default:
            printf("Invalid grade\n");
    }
}

// Another example of a switch statement based on a value provided by the function receiver.
void reciverSwitchExample(int content) {
    switch (content) {
        case 1:
            printf("1");
            break;
        case 2:
            printf("2");
            break;
        case 3:
            printf("3");
            break;
        default:
            printf("The default value.");
    }
}

int main() {
    // A simple switch statement with ints
    integerSwitch();
    // A simple switch statement with strings
    stringSwitch();
    // Simple example of a switch statement based on function receiver.
    reciverSwitchExample(2);
    return 0;
}
