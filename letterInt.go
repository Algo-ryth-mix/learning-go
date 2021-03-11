package main

func intToLetters(v int) string {
	var result string

	for x := v; x-1 >= 0; x-- {
		result = string(rune(int('A')+x%26)) + result
		x /= 26
	}
	return result
}

/*
public static string IntToLetters(int value)
{
    string result = string.Empty;
    while (--value >= 0)
    {
        result = (char)('A' + value % 26 ) + result;
        value /= 26;
    }
    return result;
}
*/
