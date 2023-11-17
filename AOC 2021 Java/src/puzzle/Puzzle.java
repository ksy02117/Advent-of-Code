package puzzle;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public abstract class Puzzle {

    public List<String> processInput(String file_path) {
        List<String> out = new ArrayList<String>();

        try (BufferedReader reader = new BufferedReader(new FileReader(file_path));) {
            for (String line = reader.readLine(); line != null; line = reader.readLine())
                out.add(line);
        } catch (IOException e) {
            e.printStackTrace();
            System.exit(-1);
        }

        return out;
    }

    public abstract String solve(String file_path);
}
