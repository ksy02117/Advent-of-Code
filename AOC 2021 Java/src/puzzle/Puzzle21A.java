package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class Puzzle21A extends Puzzle {

    private static class Dice {
        private static int currentValue = 0;
        public static int rolledCount = 0;

        private static int getNextValue() {
            int out = currentValue + 1;
            currentValue = (currentValue + 1) % 100;
            rolledCount++;
            return out;
        }

        private static int getRolledCount() {
            return rolledCount;
        }
    }

    private class Player {
        private int location;
        private int score;

        public Player(int location) {
            this.location = location;
            score = 0;
        }

        public void move(int amount) {
            location = (location + amount) % 10;
            score += location + 1;
        }

        public int getScore() {
            return score;
        }
    }

    @Override
    public String solve(String file_path) {
        List<String> inputLines = processInput(file_path);

        // Initialize players
        List<Player> players = new ArrayList<>();
        for (String inputLine : inputLines) {
            int location = Integer.parseInt(inputLine.substring(28));
            players.add(new Player(location - 1));
        }

        // game loop
        int currentPlayerID = 1;
        while (true) {
            currentPlayerID = currentPlayerID == 0 ? 1 : 0;
            Player player = players.get(currentPlayerID);

            int roll = 0;
            roll += Dice.getNextValue();
            roll += Dice.getNextValue();
            roll += Dice.getNextValue();

            player.move(roll);
            if (player.getScore() >= 1000) {
                break;
            }
        }

        int loosingPlayerID = currentPlayerID == 0 ? 1 : 0;

        int result = Dice.getRolledCount() * players.get(loosingPlayerID).getScore();

        return Integer.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/21/ex.txt";
        String result = solve(file_path);

        assertEquals("739785", result);
    }
}
