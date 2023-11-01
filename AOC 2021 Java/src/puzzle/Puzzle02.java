package puzzle;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Puzzle02 extends Puzzle {

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

    public List<Command> processInput(String file_path) {
        List<Command> out = new ArrayList<>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file_path));) {
            for (String line = reader.readLine(); line != null; line = reader.readLine())
                out.add(new Command(line));
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(-1);
        }

        return out;
    }

    public String solveA(String file_path) {
        List<Command> commands = processInput(file_path);

        int length = 0, depth = 0;
        for (Command command : commands) {
            switch (command.direction) {
                case Forward:
                    length += command.amount;
                    break;
                case Down:
                    depth += command.amount;
                    break;
                case Up:
                    depth -= command.amount;
                    break;
            }
        }

        return Integer.toString(length * depth);
    }

    public String solveB(String file_path) {
        List<Command> commands = processInput(file_path);

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

}