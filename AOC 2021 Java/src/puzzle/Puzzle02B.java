package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

import org.junit.Test;

public class Puzzle02B extends Puzzle {

    private enum Direction {
        Forward,
        Down,
        Up
    }

    private class Command {
        public Direction direction;
        public int amount;

        public Command(String input) {
            Scanner inputScanner = new Scanner(input);

            String direction = inputScanner.next();
            if (direction.equals("forward"))
                this.direction = Direction.Forward;
            else if (direction.equals("down"))
                this.direction = Direction.Down;
            else if (direction.equals("up"))
                this.direction = Direction.Up;

            this.amount = inputScanner.nextInt();
            inputScanner.close();
        }
    }

    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Command> commands = new ArrayList<>();
        for (String inputLine : inputLines) {
            commands.add(new Command(inputLine));
        }

        int length = 0, depth = 0, aim = 0;
        for (Command command : commands) {
            switch (command.direction) {
                case Forward:
                    length += command.amount;
                    depth += command.amount * aim;
                    break;
                case Down:
                    aim += command.amount;
                    break;
                case Up:
                    aim -= command.amount;
                    break;
            }
        }

        return Integer.toString(length * depth);
    }

    @Test
    public void test() {
        String file_path = "resources/02/ex.txt";
        String result = solve(file_path);

        assertEquals("900", result);
    }
}