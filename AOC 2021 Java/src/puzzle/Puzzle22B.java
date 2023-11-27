package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.junit.Test;

public class Puzzle22B extends Puzzle {

    private class Cube {
        int minX, maxX, minY, maxY, minZ, maxZ;

        public Cube(int minX, int maxX, int minY, int maxY, int minZ, int maxZ) {
            this.minX = minX;
            this.maxX = maxX;
            this.minY = minY;
            this.maxY = maxY;
            this.minZ = minZ;
            this.maxZ = maxZ;
        }

        public boolean contains(int x, int y, int z) {
            return x >= minX && x <= maxX && y >= minY && y <= maxY && z >= minZ && z <= maxZ;
        }
    }

    private class Command {
        private enum Type {
            On, Off
        }

        public Type type;
        public Cube cube;

        public Command(boolean turnOn, Cube cube) {
            type = turnOn ? Type.On : Type.Off;
            this.cube = cube;
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Integer> xSplits = new ArrayList<>();
        List<Integer> ySplits = new ArrayList<>();
        List<Integer> zSplits = new ArrayList<>();

        List<Command> commands = new ArrayList<>();
        for (String inputLine : inputLines) {
            String[] inputSplit = inputLine.split("( x=)|(\\,y=)|(\\,z=)|(\\.\\.)");

            boolean turnOn = (inputSplit[0].equals("on")) ? true : false;
            int minX = Integer.parseInt(inputSplit[1]);
            int maxX = Integer.parseInt(inputSplit[2]);
            int minY = Integer.parseInt(inputSplit[3]);
            int maxY = Integer.parseInt(inputSplit[4]);
            int minZ = Integer.parseInt(inputSplit[5]);
            int maxZ = Integer.parseInt(inputSplit[6]);

            commands.add(new Command(turnOn, new Cube(minX, maxX, minY, maxY, minZ, maxZ)));

            xSplits.add(minX);
            xSplits.add(maxX + 1);
            ySplits.add(minY);
            ySplits.add(maxY + 1);
            zSplits.add(minZ);
            zSplits.add(maxZ + 1);
        }

        Collections.sort(xSplits);
        Collections.sort(ySplits);
        Collections.sort(zSplits);
        Collections.reverse(commands);

        long cnt = 0;
        for (int i = 0; i < xSplits.size() - 1; i++) {
            for (int j = 0; j < ySplits.size() - 1; j++) {
                for (int k = 0; k < zSplits.size() - 1; k++) {
                    int x = xSplits.get(i);
                    int y = ySplits.get(j);
                    int z = zSplits.get(k);

                    long dx = xSplits.get(i + 1) - xSplits.get(i);
                    long dy = ySplits.get(j + 1) - ySplits.get(j);
                    long dz = zSplits.get(k + 1) - zSplits.get(k);

                    long volume = dx * dy * dz;

                    if (volume < 0)
                        return "ERROR";

                    if (cnt < 0)
                        return "ERROR 2";

                    for (Command command : commands) {
                        if (command.cube.contains(x, y, z)) {
                            if (command.type == Command.Type.On)
                                cnt += volume;
                            break;
                        }
                    }
                }
            }
        }

        return Long.toString(cnt);
    }

    @Test
    public void test() {
        String file_path = "resources/22/ex2.txt";
        String result = solve(file_path);

        assertEquals("2758514936282235", result);
    }
}
