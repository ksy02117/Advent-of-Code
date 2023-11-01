import java.util.Arrays;
import java.util.List;

import puzzle.*;

public class Main {
    private static List<Puzzle> puzzles = Arrays.asList(
            new Puzzle01(),
            new Puzzle02());

    public static void main(String... args) {
        int puzzleID = Integer.parseInt(args[0]);
        String file_path = "resources/" + args[0] + "/input.txt";
        String result;

        if (args[1].equals("A"))
            result = puzzles.get(puzzleID - 1).solveA(file_path);
        else
            result = puzzles.get(puzzleID - 1).solveB(file_path);

        System.out.println(result);
    }
}
