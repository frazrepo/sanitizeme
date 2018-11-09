using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Linq.Expressions;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace SanitizeMe
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
