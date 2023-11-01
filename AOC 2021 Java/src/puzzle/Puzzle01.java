package puzzle;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class Puzzle01 extends Puzzle {

    public List<Integer> processInput(String file_path) {
        List<Integer> out = new ArrayList<Integer>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file_path));) {
            for (String line = reader.readLine(); line != null; line = reader.readLine())
                out.add(Integer.parseInt(line));
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(-1);
        }

        return out;
    }

    public String solveA(String file_path) {
        List<Integer> depths = processInput(file_path);

        int cnt = 0;
        for (int i = 0; i < depths.size() - 1; i++)
            if (depths.get(i + 1) > depths.get(i))
                cnt++;

        return Integer.toString(cnt);
    }

    public String solveB(String file_path) {
        List<Integer> depths = processInput(file_path);

        int cnt = 0;
        for (int i = 0; i < depths.size() - 3; i++)
            if (depths.get(i + 3) > depths.get(i))
                cnt++;

        return Integer.toString(cnt);
    }

}