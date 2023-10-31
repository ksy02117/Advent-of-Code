import java.util.Arrays;
import java.util.List;

import puzzle.*;

public class Main {
    private static List<Puzzle> puzzles = Arrays.asList(new Puzzle01("resources/01/input.txt"));

    public static void main(String[] args) {
        puzzles.get(0).solveA();
    }
}