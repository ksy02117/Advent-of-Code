package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;

import org.junit.Test;

public class Puzzle12B extends Puzzle {

    private Map<String, List<String>> nodes = new HashMap<>();
    private List<List<String>> allPaths = new ArrayList<>();

    private int traverse(List<String> visited, String nodeName) {
        if (nodeName.equals("end")) {
            allPaths.add(visited);
            return 1;
        }

        int output = 0;
        for (String neighbor : nodes.get(nodeName)) {
            if (neighbor.equals("start"))
                continue;

            if (neighbor.charAt(0) >= 'a' && neighbor.charAt(0) <= 'z') {
                Stream<String> stream1 = visited.stream().filter(n -> n.charAt(0) >= 'a' && n.charAt(0) <= 'z');
                Stream<String> stream2 = visited.stream().filter(n -> n.charAt(0) >= 'a' && n.charAt(0) <= 'z');

                // revisited already visited already
                if (stream1.count() != stream2.distinct().count() && visited.contains(neighbor))
                    continue;
            }

            List<String> newPath = new ArrayList<>(visited);
            newPath.add(neighbor);
            output += traverse(newPath, neighbor);
        }

        return output;
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        for (String inputLine : inputLines) {
            String[] nodeNames = inputLine.split("-");

            if (!nodes.containsKey(nodeNames[0]))
                nodes.put(nodeNames[0], new ArrayList<>());
            nodes.get(nodeNames[0]).add(nodeNames[1]);

            if (!nodes.containsKey(nodeNames[1]))
                nodes.put(nodeNames[1], new ArrayList<>());
            nodes.get(nodeNames[1]).add(nodeNames[0]);
        }

        List<String> initialPath = new ArrayList<>();
        initialPath.add("start");
        int numberOfPath = traverse(initialPath, "start");

        // System.out.println(allPaths);

        return Integer.toString(numberOfPath);
    }

    @Test
    public void test() {
        String file_path = "resources/12/ex1.txt";
        String result = solve(file_path);

        assertEquals("36", result);
    }

    @Test
    public void test2() {
        String filePath = "resources/12/ex2.txt";
        String result = solve(filePath);

        assertEquals("103", result);
    }

    @Test
    public void test3() {
        String filePath = "resources/12/ex3.txt";
        String result = solve(filePath);

        assertEquals("3509", result);
    }

}
