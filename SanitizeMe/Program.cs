using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace SanitizeMe
{
    class Program
    {
        static void Main(string[] args)
        {
            string fileName = "CON./sqdsfdsds() ?*f *ù*ù \"com1 .mp3";
            string sanitizedFileName = SanitizeFileName(fileName);
            Console.WriteLine($"{fileName} --> {sanitizedFileName} ");
            Console.ReadLine();
        }

        public static string SanitizeFileName(string filename)
        {
            string replaceChar = "_";

            //Remove InvalidFileNameChars
            var invalidChars = Regex.Escape(new string(Path.GetInvalidFileNameChars()));
            var invalidReStr = string.Format(@"([{0}]*\.+$)|([{0}]+)", invalidChars);
            var sanitisedNamePart = Regex.Replace(filename, invalidReStr, replaceChar);

            //
            var illegalRe = @"[\/\?<>\\:\*\|""():]";
            sanitisedNamePart = Regex.Replace(sanitisedNamePart, illegalRe, replaceChar);

            //Remove spaces around and between
            sanitisedNamePart = Regex.Replace(sanitisedNamePart, @"\s+|^\s+|\s+$", "");

            //Remove . at the begining
            sanitisedNamePart = Regex.Replace(sanitisedNamePart, @"^\.+", "");

            //Remove trailing replace char  
            sanitisedNamePart = Regex.Replace(sanitisedNamePart, @"_+|^_+|_+$", replaceChar);

            //Remove reserved words
            //sanitisedNamePart = Regex.Replace(sanitisedNamePart, @"^(con|prn|aux|nul|com[0-9]|lpt[0-9])(\..*)?$", m => m.Groups[2].Value , RegexOptions.IgnoreCase);


            var reservedWords = new[]
            {
                "CON", "PRN", "AUX", "CLOCK$", "NUL", "COM0", "COM1", "COM2", "COM3", "COM4",
                "COM5", "COM6", "COM7", "COM8", "COM9", "LPT0", "LPT1", "LPT2", "LPT3", "LPT4",
                "LPT5", "LPT6", "LPT7", "LPT8", "LPT9"
            };

            foreach (var reservedWord in reservedWords)
            {
                var reservedWordPattern = $"^{reservedWord}\\.";
                sanitisedNamePart = Regex.Replace(sanitisedNamePart, reservedWordPattern, "_reservedWord_.", RegexOptions.IgnoreCase);
            }

            return sanitisedNamePart.ToUpper();
        }
    }
}
