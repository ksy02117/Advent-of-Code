package test;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

import puzzle.*;

public class Puzzle02Test {
    private final String file_path = "resources/02/ex.txt";

    @Test
    public void testA() {
        Puzzle02 puzzle = new Puzzle02();
        String result = puzzle.solveA(file_path);

        assertEquals("150", result);
    }

    @Test
    public void testB() {
        Puzzle02 puzzle = new Puzzle02();
        String result = puzzle.solveB(file_path);

        assertEquals("900", result);
    }
}
