#!/bin/bash

# I think the code is already pretty self explainatory with all these step comments. Don't beg me for an explanation, user, the responsibility's YOURS

# Directories
grammar="src/lua_grammar_antlr4.g4"
grammarpure="lua_grammar_antlr4.g4"
output="build"
src="src"
pwd=$(pwd)

# Stage 1: Resolve filename mismatch
GRAMMAR_NAME=$(basename "$grammar" .g4)

if ! grep -q "grammar $GRAMMAR_NAME" "$grammar"; then
    echo "Resolver: Mismatch detected"
    echo "Resolver: Resolving mismatch..."
    sed -i "s#^grammar .*#grammar $GRAMMAR_NAME#" "$grammar"
fi

# Stage 2: Resolve semicolons
if ! grep -q "grammar $GRAMMAR_NAME;" "$grammar"; then
    echo "Resolver: Adding semicolon to grammar declaration..."
    sed -i "s#^grammar $GRAMMAR_NAME#grammar $GRAMMAR_NAME;#" "$grammar"
fi

# Stage 3: Move to correct dir

if [ ! -d "$output" ]; then
    echo "Build folder doesn't exist, creating it..."
    mkdir "$output"
fi

cp "$grammar" "$output"
cd "$output"

# Stage 4: Generate files
antlr4 -Dlanguage=Cpp "$grammarpure"

# Stage 5: Back out
cd "$pwd"

# Stage 6: Compile
compile() {
    g++ -std=c++11 -I$output -I/usr/include/antlr4-runtime -L/usr/lib -o $src/parser $src/parser.cpp $output/lua_grammar_antlr4Lexer.cpp $output/lua_grammar_antlr4Parser.cpp -lantlr4-runtime
}

build() {
    echo "Building..."
    compile
}

clean() {
    echo "Cleaning up the '$output' directory..."
    find "$output" -type f -exec rm -f {} \;
    echo "Cleaned out the '$output' directory."
}

rebuild() {
    clean
    build
}

case "$1" in
    clean)
        clean
        ;;
    rebuild)
        rebuild
        ;;
    build)
        build
        ;;
    *)
        echo "Usage: $0 {clean|rebuild|build}"
        exit 1
        ;;
esac
