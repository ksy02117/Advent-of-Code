package puzzle;

import static org.junit.Assert.assertEquals;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.Test;

public class Puzzle21B extends Puzzle {

    private static Map<Integer, Integer> rollsToUniverseCount;
    static {
        rollsToUniverseCount = new HashMap<>();
        rollsToUniverseCount.put(3, 1);
        rollsToUniverseCount.put(4, 3);
        rollsToUniverseCount.put(5, 6);
        rollsToUniverseCount.put(6, 7);
        rollsToUniverseCount.put(7, 6);
        rollsToUniverseCount.put(8, 3);
        rollsToUniverseCount.put(9, 1);
    }

    private class Player {
        private static long player1WinCount = 0;
        private static long player2WinCount = 0;

        private int location;
        private int score;

        public Player(int location) {
            this.location = location;
            score = 0;
        }

        public Player(Player p) {
            location = p.location;
            score = p.score;
        }

        public void move(int amount) {
            location = (location + amount) % 10;
            score += location + 1;
        }

        public int getScore() {
            return score;
        }
    }

    private void player1Turn(Player player1, Player player2, long universeCount) {
        for (var entry : rollsToUniverseCount.entrySet()) {
            Player movedPlayer = new Player(player1);
            movedPlayer.move(entry.getKey());
            if (movedPlayer.getScore() >= 21) {
                Player.player1WinCount += universeCount * entry.getValue();
            } else {
                player2Turn(movedPlayer, player2, universeCount * entry.getValue());
            }
        }
    }

    private void player2Turn(Player player1, Player player2, long universeCount) {
        for (var entry : rollsToUniverseCount.entrySet()) {
            Player movedPlayer = new Player(player2);
            movedPlayer.move(entry.getKey());
            if (movedPlayer.getScore() >= 21) {
                Player.player2WinCount += universeCount * entry.getValue();
            } else {
                player1Turn(player1, movedPlayer, universeCount * entry.getValue());
            }
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
        player1Turn(players.get(0), players.get(1), 1);

        long result = Player.player1WinCount > Player.player2WinCount ? Player.player1WinCount : Player.player2WinCount;

        return Long.toString(result);
    }

    @Test
    public void test() {
        String file_path = "resources/21/ex.txt";
        String result = solve(file_path);

        assertEquals("444356092776315", result);
    }
}
