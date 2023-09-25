package main

import "github.com/Ammyy9908/code-assisstant/pkg/assistant/factory/impl"

func main() {

	assesstant := impl.NewAssistant("<OPEN AI KEY>", "<MODEL NAME>")

	response, err := assesstant.GetChatResponse("public class BinarySearch {\n"+
		"\n"+
		"    public;\n"+
		"        while (l <= r) {\n"+
		"            int m = l + (r - l) / 2; // calculate mid of the array\n"+
		"\n"+
		"            // Check if x is present at mid\n"+
		"            if (arr[m] == x) {\n"+
		"                return m;\n"+
		"            }\n"+
		"\n"+
		"            // If x is greater, ignore left half\n"+
		"            if (arr[m] < x) {\n"+
		"                l = m + 1;\n"+
		"            }\n"+
		"            // If x is smaller, ignore right half\n"+
		"            else {\n"+
		"                r = m - 1;\n"+
		"            }\n"+
		"        }\n"+
		"\n"+
		"        // x was not present in the array\n"+
		"        return -1;\n"+
		"    }\n"+
		"\n"+
		"    public static void main(String[] args) {\n"+
		"        int[] arr = {2, 5, 8, 12, 16, 23, 38, 45, 56, 72, 91};\n"+
		"        int x = 23;\n"+
		"\n"+
		"        int result = binarySearch(arr, x);\n"+
		"        if (result == -1) {\n"+
		"            System.out.println(\"Element not present in the array\");\n"+
		"        } else {\n"+
		"            System.out.println(\"Element found at index: \" + result);\n"+
		"        }\n"+
		"    }\n"+
		"}", "Java")
	if err != nil {
		panic(err)
	}

	println(response)

}
