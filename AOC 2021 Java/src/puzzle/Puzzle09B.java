package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.junit.Test;

public class Puzzle09B extends Puzzle {

    private NodeState[][] grid;
    private List<Set<Coordinate>> basins;

    private enum NodeState {
        NotVisited,
        Visited,
        Wall
    }

    private class Coordinate {
        int x, y;

        public Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    private Set<Coordinate> getBasinCoordinates(Coordinate coordinate) {
        Set<Coordinate> basinCoordinates = new HashSet<>();

        if (grid[coordinate.y][coordinate.x] != NodeState.NotVisited) {
            return basinCoordinates;
        }

        grid[coordinate.y][coordinate.x] = NodeState.Visited;
        basinCoordinates.add(coordinate);

        // Top
        if (coordinate.y != 0) {
            basinCoordinates.addAll(getBasinCoordinates(new Coordinate(coordinate.x, coordinate.y - 1)));
        }

        // Bottom
        if (coordinate.y != grid.length - 1) {
            basinCoordinates.addAll(getBasinCoordinates(new Coordinate(coordinate.x, coordinate.y + 1)));
        }

        // Left
        if (coordinate.x != 0) {
            basinCoordinates.addAll(getBasinCoordinates(new Coordinate(coordinate.x - 1, coordinate.y)));
        }

        // Right
        if (coordinate.x != grid[coordinate.y].length - 1) {
            basinCoordinates.addAll(getBasinCoordinates(new Coordinate(coordinate.x + 1, coordinate.y)));
        }

        return basinCoordinates;
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);
        int x, y;

        grid = new NodeState[inputLines.size()][inputLines.get(0).length()];

        y = 0;
        for (String inputLine : inputLines) {
            x = 0;
            for (char inputLetter : inputLine.toCharArray()) {
                if (inputLetter == '9') {
                    grid[y][x] = NodeState.Wall;
                } else {
                    grid[y][x] = NodeState.NotVisited;
                }
                x++;
            }
            y++;
        }

        basins = new ArrayList<>();
        for (y = 0; y < grid.length; y++) {
            for (x = 0; x < grid[y].length; x++) {
                if (grid[y][x] == NodeState.NotVisited) {
                    basins.add(getBasinCoordinates(new Coordinate(x, y)));
                }
            }
        }

        basins.sort(new Comparator<Set<Coordinate>>() {
            @Override
            public int compare(Set<Coordinate> o1, Set<Coordinate> o2) {
                return -o1.size() + o2.size();
            }
        });

        int answer = 1;
        for (int i = 0; i < 3; i++) {
            answer *= basins.get(i).size();
        }

        return Integer.toString(answer);
    }

    @Test
    public void test() {
        String file_path = "resources/09/ex.txt";
        String result = solve(file_path);

        assertEquals("1134", result);
    }
}
