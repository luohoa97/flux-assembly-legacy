#include <iostream>
#include <fstream>
#include <streambuf>
#include "antlr4-runtime.h"

#include "../build/lua_grammar_antlr4Lexer.h"
#include "../build/lua_grammar_antlr4Parser.h"

using namespace std;
using namespace antlr4;

int main(int argc, const char *args[]) {
    if (argc < 2) {
        cerr << "Usage: " << args[0] << " <input_file>" << endl;
        return 1;
    }

    // Read the input file
    ifstream stream;
    stream.open(args[1]);
    if (!stream.is_open()) {
        cerr << "Error opening file: " << args[1] << endl;
        return 1;
    }

    string file_content((istreambuf_iterator<char>(stream)), istreambuf_iterator<char>());
    ANTLRInputStream input(file_content);

    lua_grammar_antlr4Lexer lexer(&input);
    CommonTokenStream tokens(&lexer);
    lua_grammar_antlr4Parser parser(&tokens);

    lexer.removeErrorListeners();
    parser.removeErrorListeners();

    tree::ParseTree *tree = parser.program();

    cout << "Parse Tree: " << endl;
    cout << tree->toStringTree(&parser) << endl;

    return 0;
}
