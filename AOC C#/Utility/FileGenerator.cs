

namespace AdventOfCode;

public static class FileGenerator
{
    public static void GenerateNextSolver()
    {
        Console.WriteLine("Creating New Line");
        for (int year = 2015; year < 2030; year++)
        {
            for (int day = 1; day <= 25; day++)
            {
                string directoryPath = $"Year{year}/Day{day:00}";
                string solverAFile = $"Year{year}/Day{day:00}/SolverA.cs";
                if (!File.Exists(solverAFile))
                {
                    // Create SolutionA From Template File
                    Console.WriteLine($"Creating {solverAFile}");
                    System.IO.Directory.CreateDirectory(directoryPath);
                    File.WriteAllText(solverAFile, GenerateSolutionTemplate(year, day));

                    return;
                }

                string solverBFile = $"Year{year}/Day{day:00}/SolverB.cs";
                if (!File.Exists(solverBFile))
                {
                    // Create SolutionB.cs by copying SolutionA.cs;
                    Console.WriteLine($"Creating {solverBFile}, by copying {solverAFile}");
                    File.Copy(solverAFile, solverBFile);

                    return;
                }

            }
        }
    }

    private static string GenerateSolutionTemplate(int year, int day)
    {
        return $@"
namespace AdventOfCode.Y{year}.D{day:00};
            
class SolverA : Solver
{{
    public string Solve(string inputFilePath)
    {{
        IEnumerable<string> inputLines = InputFileReader.readInputFile(inputFilePath);
        
        throw new NotImplementedException();
    }}
}}";
    }
}