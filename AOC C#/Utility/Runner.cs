namespace AdventOfCode;

public class Runner
{
    public static void Run(int year, int day, Part part, InputType input = InputType.Input)
    {
        try
        {
            string solverTypeString = $"AdventOfCode.Y{year}.D{day:00}.Solver{part}";
            Solver? solver = (Solver?)Activator.CreateInstance("AdventOfCode", solverTypeString)?.Unwrap();
            if (solver is null)
            {
                Console.WriteLine("Solver Does Not Exists Yet!");
                return;
            }

            Console.WriteLine($"Running Year {year}, Day {day}, Part{part}");
            string inputFilePath = $"Year{year}/Day{day:00}/{input.ToString().ToLower()}.txt";
            string result = solver.Solve(inputFilePath);
            Console.WriteLine(result);
        }
        catch (Exception)
        {
            Console.WriteLine("Solver Does Not Exists Yet!");
            return;
        }
    }

    public static void RunLatest(InputType input = InputType.Input)
    {
        int year, day;
        for (year = 2030; year >= 2015; year--) 
        {
            if (Directory.Exists($"Year{year}")) 
            {
                break;
            }
        }

        for (day = 25; day >= 1; day--) 
        {
            if (Directory.Exists($"Year{year}/Day{day:00}")) 
            {
                break;
            }
        }

        if (File.Exists($"Year{year}/Day{day:00}/SolverB.cs"))
        {
            Run(year, day, Part.B, input);
        }
        else
        {
            Run(year,day,Part.A, input);
        }
    }
}