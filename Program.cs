using System;
using SanitizeMe;
using System.IO;


namespace Sanitizeme
{
    class Program
    {
        static void Main(string[] args)
        {
            try
            {
                if (args.Length == 0 || args.Length > 1)
                {
                    ShowUsage();
                    return;
                }

                Sanitizer sanitizer = new Sanitizer();
                string fileName = args[0].ToString();

                if (!File.Exists(fileName))
                    throw new Exception("File not found");

                string sanitizedFileName = sanitizer.Sanitize(fileName);

                Console.WriteLine($"Renaming {fileName} to {sanitizedFileName} ");
            }
            catch (Exception e)
            {
                Console.WriteLine($"Error : {e.Message}");
            }
        }

        static void ShowUsage()
        {
            Console.WriteLine("Usage : sanitizeme filename");
        }
    }
}
