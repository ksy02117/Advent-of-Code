package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.HashSet;
import java.util.List;
import java.util.Objects;
import java.util.Set;

import org.junit.Test;

public class Puzzle22A extends Puzzle {

    private Set<Coordinate> onCoordinates = new HashSet<>();

    private class Coordinate {
        int x, y, z;

        public Coordinate(int x, int y, int z) {
            this.x = x;
            this.y = y;
            this.z = z;
        }

        @Override
        public boolean equals(Object obj) {
            if (!(obj instanceof Coordinate))
                return false;
            Coordinate o = (Coordinate) obj;
            if (x == o.x && y == o.y && z == o.z)
                return true;
            return false;
        }

        @Override
        public int hashCode() {
            return Objects.hash(x, y, z);
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        for (String inputLine : inputLines) {
            String[] inputSplit = inputLine.split("( x=)|(\\,y=)|(\\,z=)|(\\.\\.)");

            boolean turnOn = (inputSplit[0].equals("on")) ? true : false;
            int minX = Integer.parseInt(inputSplit[1]);
            if (minX < -50)
                minX = -50;
            int maxX = Integer.parseInt(inputSplit[2]);
            if (maxX > 50)
                maxX = 50;
            int minY = Integer.parseInt(inputSplit[3]);
            if (minY < -50)
                minY = -50;
            int maxY = Integer.parseInt(inputSplit[4]);
            if (maxY > 50)
                maxY = 50;
            int minZ = Integer.parseInt(inputSplit[5]);
            if (minZ < -50)
                minZ = -50;
            int maxZ = Integer.parseInt(inputSplit[6]);
            if (maxZ > 50)
                maxZ = 50;

            for (int x = minX; x <= maxX; x++) {
                for (int y = minY; y <= maxY; y++) {
                    for (int z = minZ; z <= maxZ; z++) {
                        if (turnOn) {
                            onCoordinates.add(new Coordinate(x, y, z));
                        } else {
                            onCoordinates.remove(new Coordinate(x, y, z));
                        }
                    }
                }
            }
        }

        int result = onCoordinates.size();
        return Integer.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/22/ex.txt";
        String result = solve(file_path);

        assertEquals("590784", result);
    }
}
