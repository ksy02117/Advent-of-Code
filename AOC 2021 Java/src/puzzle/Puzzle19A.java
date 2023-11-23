package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle19A extends Puzzle {

    private class Coordinate {
        final int x, y, z;

        public Coordinate(int x, int y, int z) {
            this.x = x;
            this.y = y;
            this.z = z;
        }

        public Coordinate translate(final int dx, final int dy, final int dz) {
            return new Coordinate(x + dx, y + dy, z + dz);
        }

        public Coordinate rotate(final int id) {
            switch (id) {
                case 0:
                    return new Coordinate(x, y, z);
                case 1:
                    return new Coordinate(x, z, -y);
                case 2:
                    return new Coordinate(x, -y, -z);
                case 3:
                    return new Coordinate(x, -z, y);
                case 4:
                    return new Coordinate(y, -x, z);
                case 5:
                    return new Coordinate(y, z, x);
                case 6:
                    return new Coordinate(y, x, -z);
                case 7:
                    return new Coordinate(y, -z, -x);
                case 8:
                    return new Coordinate(-x, -y, z);
                case 9:
                    return new Coordinate(-x, z, y);
                case 10:
                    return new Coordinate(-x, y, -z);
                case 11:
                    return new Coordinate(-x, -z, -y);
                case 12:
                    return new Coordinate(-y, x, z);
                case 13:
                    return new Coordinate(-y, z, -x);
                case 14:
                    return new Coordinate(-y, -x, -z);
                case 15:
                    return new Coordinate(-y, -z, x);
                case 16:
                    return new Coordinate(z, y, -x);
                case 17:
                    return new Coordinate(z, -x, -y);
                case 18:
                    return new Coordinate(z, -y, x);
                case 19:
                    return new Coordinate(z, x, y);
                case 20:
                    return new Coordinate(-z, y, x);
                case 21:
                    return new Coordinate(-z, x, -y);
                case 22:
                    return new Coordinate(-z, -y, -x);
                case 23:
                    return new Coordinate(-z, -x, y);
            }

            return null;
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
            if (x == o.x && y == o.y && z == o.z)
                return true;

            return false;
        }
    }

    private class Scanner {
        private List<Coordinate> beacons = new ArrayList<>();
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        List<Scanner> scanners = new ArrayList<>();

        Scanner scanner = new Scanner();
        for (String inputLine : inputLines) {
            if (inputLine.contains("scanner")) {
                scanner = new Scanner();
                continue;
            }

            if (inputLine.equals("")) {
                scanners.add(scanner);
                continue;
            }

            String[] coordinatesInString = inputLine.split(",");
            int x = Integer.parseInt(coordinatesInString[0]);
            int y = Integer.parseInt(coordinatesInString[1]);
            int z = Integer.parseInt(coordinatesInString[2]);
            Coordinate beaconCoordinate = new Coordinate(x, y, z);

            scanner.beacons.add(beaconCoordinate);
        }
        scanners.add(scanner);

        // adding beacons in the first scanner to know beacon locations
        List<Coordinate> beacons = new ArrayList<>();
        scanner = scanners.remove(0);
        for (Coordinate coordinate : scanner.beacons) {
            beacons.add(coordinate);
        }

        scannerLoop: while (!scanners.isEmpty()) {
            scanner = scanners.remove(0);

            for (int rotation = 0; rotation < 24; rotation++) {
                final int rotationID = rotation;
                List<Coordinate> rotatedBeaconCoordinates = scanner.beacons.stream()
                        .map((Coordinate c) -> c.rotate(rotationID)).toList();

                for (int i = 0; i < rotatedBeaconCoordinates.size(); i++) {
                    for (int j = 0; j < beacons.size(); j++) {
                        int dx = beacons.get(j).x - rotatedBeaconCoordinates.get(i).x;
                        int dy = beacons.get(j).y - rotatedBeaconCoordinates.get(i).y;
                        int dz = beacons.get(j).z - rotatedBeaconCoordinates.get(i).z;

                        List<Coordinate> translatedCoordinates = rotatedBeaconCoordinates.stream()
                                .map((Coordinate c) -> c.translate(dx, dy, dz)).toList();

                        int cnt = 0;
                        for (Coordinate translatedCoordinate : translatedCoordinates) {
                            if (beacons.contains(translatedCoordinate)) {
                                cnt++;
                            }
                        }

                        if (cnt >= 12) {
                            for (Coordinate translatedCoordinate : translatedCoordinates) {
                                if (!beacons.contains(translatedCoordinate)) {
                                    beacons.add(translatedCoordinate);
                                }
                            }
                            continue scannerLoop;
                        }
                    }
                }
            }

            scanners.add(scanner);
        }

        return Integer.toString(beacons.size());
    }

    @Test
    public void test() {
        String file_path = "resources/19/ex.txt";
        String result = solve(file_path);

        assertEquals("79", result);
    }
}
