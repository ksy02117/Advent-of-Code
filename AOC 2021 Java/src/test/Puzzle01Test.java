package test;

import static org.junit.Assert.assertEquals;

import org.junit.Test;

import puzzle.*;

public class Puzzle01Test {
    private final String file_path = "resources/01/ex.txt";

    @Test
    public void testA() {
        Puzzle puzzle = new Puzzle01();
        String result = puzzle.solveA(file_path);

        assertEquals("7", result);
    }

    @Test
    public void testB() {
        Puzzle01 puzzle = new Puzzle01();
        String result = puzzle.solveB(file_path);

        assertEquals("5", result);
    }
}
