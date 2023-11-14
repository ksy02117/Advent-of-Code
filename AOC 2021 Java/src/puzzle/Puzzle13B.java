package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.junit.Test;

public class Puzzle13B extends Puzzle {

    private class Coordinate {
        int x;
        int y;

        public Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public int hashCode() {
            return x * 10000 + y;
        }

        @Override
        public boolean equals(Object obj) {
            if (this == obj)
                return true;
            if (obj == null)
                return false;
            if (getClass() != obj.getClass())
                return false;
            Coordinate o = (Coordinate) obj;
            if (this.hashCode() != o.hashCode())
                return false;
            return true;
        }
    }

    private class Command {
        char direction;
        int location;

        public Command(char direction, int location) {
            this.direction = direction;
            this.location = location;
        }

        private Coordinate Execute(Coordinate input) {
            if (direction == 'x' && input.x > location) {
                int newX = input.x - 2 * (input.x - location);
                Coordinate newCoordinate = new Coordinate(newX, input.y);
                return newCoordinate;
            }
            if (direction == 'y' && input.y > location) {
                int newY = input.y - 2 * (input.y - location);
                Coordinate newCoordinate = new Coordinate(input.x, newY);
                return newCoordinate;
            }
            return new Coordinate(input.x, input.y);
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        Set<Coordinate> coordinates = new HashSet<>();
        int lineIdx = 0;
        for (lineIdx = 0; lineIdx < inputLines.size(); lineIdx++) {
            String inputLine = inputLines.get(lineIdx);
            if (inputLine.equals(""))
                break;

            String[] coordinateString = inputLine.split(",");

            int x = Integer.parseInt(coordinateString[0]);
            int y = Integer.parseInt(coordinateString[1]);

            coordinates.add(new Coordinate(x, y));
        }

        List<Command> commands = new ArrayList<>();
        lineIdx++;
        for (; lineIdx < inputLines.size(); lineIdx++) {
            String inputLine = inputLines.get(lineIdx);

            String[] inputLineDivided = inputLine.split("=");
            commands.add(new Command(inputLineDivided[0].charAt(11), Integer.parseInt(inputLineDivided[1])));
        }

        Set<Coordinate> foldedCoordinates = new HashSet<>();
        for (Coordinate coordinate : coordinates) {
            for (Command command : commands) {
                coordinate = command.Execute(coordinate);
            }
            foldedCoordinates.add(coordinate);
        }

        int maxX = 0, maxY = 0;
        for (Coordinate coordinate : foldedCoordinates) {
            if (maxX < coordinate.x)
                maxX = coordinate.x;
            if (maxY < coordinate.y)
                maxY = coordinate.y;
        }

        String out = "";
        for (int y = 0; y <= maxY; y++) {
            for (int x = 0; x <= maxX; x++) {
                if (foldedCoordinates.contains(new Coordinate(x, y)))
                    out += '#';
                else
                    out += '.';
            }
            out += "\n";
        }

        return out;
    }

    @Test
    public void test() {
        String file_path = "resources/13/ex.txt";
        String result = solve(file_path);

        assertEquals("#####\n#...#\n#...#\n#...#\n#####\n", result);
    }

}
