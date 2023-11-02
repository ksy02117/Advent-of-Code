import static org.junit.Assert.assertEquals;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

import puzzle.*;

public class Main {
    private static List<Puzzle> puzzles = Arrays.asList(
            new Puzzle01(),
            new Puzzle02(),
            new Puzzle03(),
            new Puzzle04());

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

    @ParameterizedTest
    @CsvSource({
            "01, 7",
            "02, 150",
            "03, 198",
            "04, 4512"
    })
    void testA(String arg, String expected) {
        int puzzleID = Integer.parseInt(arg);
        String file_path = "resources/" + arg + "/ex.txt";
        String result = puzzles.get(puzzleID - 1).solveA(file_path);

        assertEquals(expected, result);
    }

    @ParameterizedTest
    @CsvSource({
            "01, 5",
            "02, 900",
            "03, 230",
            "04, 1924"
    })
    void testB(String arg, String expected) {
        int puzzleID = Integer.parseInt(arg);
        String file_path = "resources/" + arg + "/ex.txt";
        String result = puzzles.get(puzzleID - 1).solveB(file_path);

        assertEquals(expected, result);
    }
}
