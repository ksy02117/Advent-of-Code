package test;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

import puzzle.Puzzle01;

public class Puzzle01Test {
    private Puzzle01 puzzle = new Puzzle01("resources/01/input.txt");

    @Test
    public void solveA() {
        String result = puzzle.solveA();

        assertEquals("", result);
    }

    @Test
    public void solveB() {
        String result = puzzle.solveB();

        assertEquals("", result);
    }
}
