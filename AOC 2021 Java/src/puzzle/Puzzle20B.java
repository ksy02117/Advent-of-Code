package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.junit.Test;

public class Puzzle20B extends Puzzle {
    private String enhancer;
    private boolean invert;

    private class Coordinate {
        private int x, y;

        public Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public int hashCode() {
            return x * 100000 + y;
        }

        @Override
        public boolean equals(Object obj) {
            if (!(obj instanceof Coordinate))
                return false;

            Coordinate o = (Coordinate) obj;
            if (x == o.x && y == o.y)
                return true;
            return false;
        }

    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        // Enhancer
        enhancer = inputLines.get(0);
        invert = enhancer.charAt(0) == '#';

        // Get original input
        Set<Coordinate> litCoordinates = new HashSet<>();
        for (int i = 2; i < inputLines.size(); i++) {
            String inputLine = inputLines.get(i);
            int y = i - 2;

            for (int x = 0; x < inputLine.length(); x++) {
                if (inputLine.charAt(x) == '#') {
                    litCoordinates.add(new Coordinate(x, y));
                }
            }
        }

        for (int i = 0; i < 50; i++) {
            // Get Board Boundary
            int minX = Integer.MAX_VALUE, minY = Integer.MAX_VALUE;
            int maxX = Integer.MIN_VALUE, maxY = Integer.MIN_VALUE;
            for (Coordinate coordinate : litCoordinates) {
                if (minX > coordinate.x - 1)
                    minX = coordinate.x - 1;

                if (minY > coordinate.y - 1)
                    minY = coordinate.y - 1;

                if (maxX < coordinate.x + 1)
                    maxX = coordinate.x + 1;

                if (maxY < coordinate.y + 1)
                    maxY = coordinate.y + 1;
            }

            // Get Next State
            Set<Coordinate> previousLitCoordinates = litCoordinates;
            litCoordinates = new HashSet<>();
            for (int x = minX; x <= maxX; x++) {
                for (int y = minY; y <= maxY; y++) {
                    // get binary
                    int binary = 0;
                    if (previousLitCoordinates.contains(new Coordinate(x - 1, y - 1)))
                        binary += 256;
                    if (previousLitCoordinates.contains(new Coordinate(x, y - 1)))
                        binary += 128;
                    if (previousLitCoordinates.contains(new Coordinate(x + 1, y - 1)))
                        binary += 64;
                    if (previousLitCoordinates.contains(new Coordinate(x - 1, y)))
                        binary += 32;
                    if (previousLitCoordinates.contains(new Coordinate(x, y)))
                        binary += 16;
                    if (previousLitCoordinates.contains(new Coordinate(x + 1, y)))
                        binary += 8;
                    if (previousLitCoordinates.contains(new Coordinate(x - 1, y + 1)))
                        binary += 4;
                    if (previousLitCoordinates.contains(new Coordinate(x, y + 1)))
                        binary += 2;
                    if (previousLitCoordinates.contains(new Coordinate(x + 1, y + 1)))
                        binary += 1;

                    if (!invert) {
                        if (enhancer.charAt(binary) == '#') {
                            litCoordinates.add(new Coordinate(x, y));
                        }
                    } else {
                        if (i % 2 == 0) {
                            if (enhancer.charAt(binary) != '#') {
                                litCoordinates.add(new Coordinate(x, y));
                            }
                        } else {
                            binary = 511 - binary;
                            if (enhancer.charAt(binary) == '#') {
                                litCoordinates.add(new Coordinate(x, y));
                            }
                        }
                    }
                }
            }

            // Print map
            // System.out.println("\n-- Map " + (i + 1) + " --");
            // for (int y = minY; y <= maxY; y++) {
            // for (int x = minX; x <= maxX; x++) {
            // if (litCoordinates.contains(new Coordinate(x, y)))
            // System.out.print("#");
            // else
            // System.out.print(".");
            // }
            // System.out.println();
            // }
        }

        return Integer.toString(litCoordinates.size());
    }

    @Test
    public void test() {
        String file_path = "resources/20/ex.txt";
        String result = solve(file_path);

        assertEquals("3351", result);
    }
}
