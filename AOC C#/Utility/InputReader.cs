using System.Linq.Expressions;

namespace AdventOfCode;

public static class InputFileReader
{
    public static IEnumerable<string> readInputFile(string filePath)
    {
        IEnumerable<string> output = new List<string>();

        try
        {
            using (StreamReader reader = File.OpenText(filePath))
            {
                string? line;
                while ((line = reader.ReadLine()) is not null)
                {
                    output = output.Append(line);
                }
            };
        }
        catch (Exception)
        {
            Console.WriteLine($"Input File [{filePath}] Does Not Exists!");
            Environment.Exit(-1);
        }

        return output;
    }
}