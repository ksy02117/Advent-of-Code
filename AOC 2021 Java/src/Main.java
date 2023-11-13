import java.util.Arrays;
import java.util.List;

import puzzle.*;

public class Main {
    private static List<Puzzle> puzzles = Arrays.asList(
            new Puzzle01A(), new Puzzle01B(),
            new Puzzle02A(), new Puzzle02B(),
            new Puzzle03A(), new Puzzle03B(),
            new Puzzle04A(), new Puzzle04B(),
            new Puzzle05A(), new Puzzle05B(),
            new Puzzle06A(), new Puzzle06B(),
            new Puzzle07A(), new Puzzle07B(),
            new Puzzle08A(), new Puzzle08B(),
            new Puzzle09A(), new Puzzle09B(),
            new Puzzle10A(), new Puzzle10B(),
            new Puzzle11A(), new Puzzle11B(),
            new Puzzle12A(), new Puzzle12B());

    public static void main(String... args) {
        int puzzleID = Integer.parseInt(args[0]);
        String file_path = "resources/" + args[0] + "/input.txt";

        int IDX = (puzzleID - 1) * 2;
        if (args[1].equals("B"))
            IDX++;

        String result = puzzles.get(IDX).solve(file_path);

        System.out.println(result);
    }
}
