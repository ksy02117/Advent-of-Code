package puzzle;

import java.io.BufferedReader;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;

import util.BufferedReaderGetter;

public class Puzzle01 extends Puzzle {
    private String file_path;
    private BufferedReader reader;

    public Puzzle01(String file_path) {
        this.file_path = file_path;
        reader = BufferedReaderGetter.getReader(file_path);
    }

    public String solveA() {
        for (String line = reader.readLine(); line != null; line = reader.readLine()) {
            System.out.print(line);
        }

        return "";
    }

    public String solveB() {
        try (FileReader reader = new FileReader(file_path);) {
            System.out.println("Hello Day 01B");

        } catch (IOException e) {
            e.printStackTrace();
        }

        return "";
    }

}